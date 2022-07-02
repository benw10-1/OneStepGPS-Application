package store

import (
	"io/ioutil"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	// "fmt"
)

type User struct {
	Name string
	Password string
	APIKey string
	Preferences map[string]string
}

type CleanUser struct {
	Name string
	APIKey string
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
	SetPreferences(name string, preferences map[string]string)
	GetPreferences(name string) map[string]string
	CorrectCredentials(name string, password string) bool
}

func (x *JSONStruct) GetUsers() []User {
	return x.Users
}

// Checks if the user name and API key are unique
func (x *JSONStruct) IsUnique(name string, apiKey string) bool {
	for _, user := range x.Users {
		if user.Name == name || (user.APIKey != "" && user.APIKey == apiKey) {
			return false
		}
	}
	return true
}
// Creates a new user in the store
func (x *JSONStruct) AddUser(name string, password string, apiKey string) bool {
	// Check if the user name and API key are unique
	unique := x.IsUnique(name, apiKey)
	if !unique {
		return false
	}
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
	// Add the user to the store
	map_ := make(map[string]string)
	map_["test"] = "test"
	newUser := User{name, string(hashedPassword), apiKey, map_}
	x.Users = append(x.GetUsers(), newUser)

	return true
}
// Validate the user name and password
func (x *JSONStruct) CorrectCredentials(name string, password string) bool {
	for _, user := range x.GetUsers() {
		if user.Name == name {
			// Check if the password is correct
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err == nil {
				return true
			}
		}
	}
	return false
}
// Get a user from the store
func (x *JSONStruct) GetUser(name string) *CleanUser {
	for _, user := range x.GetUsers() {
		if user.Name == name {
			// Return a clean user
			return &CleanUser{user.Name, user.APIKey}
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

func (x *JSONStruct) APIKeyExists(apiKey string) bool {
	for _, user := range x.GetUsers() {
		if user.APIKey == apiKey {
			return true
		}
	}
	return false
}

func (x *JSONStruct) GetPreferences(name string) map[string]string {
	for _, user := range x.GetUsers() {
		if user.Name == name {
			return user.Preferences
		}
	}
	return nil
}

func (x *JSONStruct) SetPreferences(name string, preferences map[string]string) {
	for i, user := range x.GetUsers() {
		if user.Name == name {
			x.Users[i].Preferences = preferences
		}
	}
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