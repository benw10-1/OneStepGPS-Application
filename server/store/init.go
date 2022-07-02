package store

import (
	"io/ioutil"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
)

type User struct {
	Name string
	Password string
	APIKey string
	preferences map[string]string
}

type JSONStruct struct {
	Users []User
}

type JSONStructor interface {
	AddUser(name string, password string, apiKey string)
	GetUser(name string) *User
	RemoveUser(name string)
	GetUsers() []User
	IsUnique(name string, apiKey string) bool
	CorrectCredentials(name string, password string) bool
}

func (x *JSONStruct) GetUsers() []User {
	return x.Users
}

func (x *JSONStruct) IsUnique(name string, apiKey string) bool {
	for _, user := range x.Users {
		if user.Name == name && user.APIKey == apiKey {
			return false
		}
	}
	return true
}

func (x *JSONStruct) AddUser(name string, password string, apiKey string) bool {
	unique := x.IsUnique(name, apiKey)
	if !unique {
		return false
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }

	newUser := User{name, string(hashedPassword), apiKey, make(map[string]string)}
	x.Users = append(x.Users, newUser)

	return true
}

func (x *JSONStruct) CorrectCredentials(name string, password string) bool {
	for _, user := range x.GetUsers() {
		if user.Name == name {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err == nil {
				return true
			}
		}
	}
	return false
}

func (x *JSONStruct) GetUser(name string) *User {
	for _, user := range x.GetUsers() {
		if user.Name == name {
			return &user
		}
	}
	return nil
}

func (x *JSONStruct) RemoveUser(name string) bool {
	for i, user := range x.GetUsers() {
		if user.Name == name {
			x.Users = append(x.Users[:i], x.Users[i+1:]...)
			return true
		}
	}
	return false
}

func (x *JSONStruct) APIKeyExists (apiKey string) bool {
	for _, user := range x.GetUsers() {
		if user.APIKey == apiKey {
			return true
		}
	}
	return false
}

type Storer interface {
	Get() (*JSONStruct, error) 
	Save() error
}

type DataStore struct {
	dataRef *JSONStruct
	filepath string
}

func (x *DataStore) Get() *JSONStruct {
	return x.dataRef
}

func (x *DataStore) Save() error {
	data, err := json.Marshal(*x.dataRef)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(x.filepath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func initStore(filepath string) *DataStore {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		data = []byte("{\"Users\":[]}")
	}

	store := &DataStore{}
	store.filepath = filepath
	store.dataRef = &JSONStruct{}
	
	_ = json.Unmarshal(data, store.dataRef)

	return store
}

var Store = initStore("data.json")