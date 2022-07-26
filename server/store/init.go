// Load a struct from a file into memory, and write it back to the file whenever it is changed
package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"sync"

	"golang.org/x/crypto/bcrypt"
)

// User data structure
type User struct {
	Name     string
	Password string
	APIKey   string
	// key value pairs for preferences for formless yet structured
	Preferences map[string]interface{}
}

// User without password
type CleanUser struct {
	Name   string
	APIKey string
}

// Structure of the JSON file
type UserStorage struct {
	Users   []User
	UserMap map[string]*User
}

// Declare all methods
type userStorageMethods interface {
	AddUser(name string, password string, apiKey string)
	GetUser(name string) *User
	RemoveUser(name string)
	GetUsers() []User
	IsUnique(name string, apiKey string) bool
	SetPreferences(name string, preferences map[string]string)
	GetPreferences(name string) map[string]string
	CorrectCredentials(name string, password string) bool
	GetDevices(name string) string
	ReverseGeocode(name string, lat string, lon string) string
	UpdateUserMap()
	GetFullUser(name string) *User
}

// Indirectly implement the GetUsers method
func (x *UserStorage) GetUsers() []User {
	return x.Users
}

func (x *UserStorage) UpdateUserMap() {
	userMap := make(map[string]*User)
	for i, user := range x.Users {
		userMap[user.Name] = &x.Users[i]
	}
	x.UserMap = userMap
}

// API proxy (no CORS on the server)
func (x *UserStorage) GetDevices(name string) string {
	// Get the user
	if user, ok := x.UserMap[name]; ok {
		if user.APIKey == "" {
			return "No API key"
		}

		requestURL := "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=" + user.APIKey

		req, err := http.NewRequest(http.MethodGet, requestURL, nil)
		if err != nil {
			return "Error making request"
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return "Error making request"
		}

		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "Could not read response"
		}
		return string(resBody)
	}

	return ""
}

func (x *UserStorage) ReverseGeocode(name string, lat string, lon string) string {
	// Get the user
	if user, ok := x.UserMap[name]; ok {
		if user.APIKey == "" {
			return "No API key"
		}
		requestURL := "https://track.onestepgps.com/v3/api/public/reverse-geocode?lat_lng=" + lat + "%2C" + lon + "&api-key=" + user.APIKey

		req, err := http.NewRequest(http.MethodGet, requestURL, nil)
		if err != nil {
			return "Error making request"
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return "Error making request"
		}

		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "Could not read response"
		}
		return string(resBody)
	}

	return ""
}

// Checks if the user name and API key are unique
func (x *UserStorage) IsUnique(name string, apiKey string) bool {
	if _, ok := x.UserMap[name]; ok {
		return false
	}
	return true
}

// Creates a new user in the store
func (x *UserStorage) AddUser(name string, password string, apiKey string) bool {
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
	map_ := make(map[string]interface{})

	newUser := User{name, string(hashedPassword), apiKey, map_}
	x.Users = append(x.GetUsers(), newUser)

	x.UpdateUserMap()

	return true
}

// Validate the user name and password
func (x *UserStorage) CorrectCredentials(name string, password string) bool {
	if user, ok := x.UserMap[name]; ok {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err == nil {
			return true
		}
	}
	return false
}

// Get a user from the store
func (x *UserStorage) GetUser(name string) *CleanUser {
	if user, ok := x.UserMap[name]; ok {
		return &CleanUser{user.Name, user.APIKey}
	}
	return nil
}

func (x *UserStorage) GetFullUser(name string) *User {
	if user, ok := x.UserMap[name]; ok {
		return user
	}
	return nil
}

// Remove a user from the store
func (x *UserStorage) RemoveUser(name string) bool {
	for i, user := range x.GetUsers() {
		if user.Name == name {
			x.Users = append(x.Users[:i], x.Users[i+1:]...)
			return true
		}
	}
	x.UpdateUserMap()
	return false
}

// Gets a user's preferences
func (x *UserStorage) GetPreferences(name string) map[string]interface{} {
	if user, ok := x.UserMap[name]; ok {
		return user.Preferences
	}
	return nil
}

// Sets a user's preferences
func (x *UserStorage) SetPreferences(name string, preferences map[string]interface{}) {
	if user, ok := x.UserMap[name]; ok {
		user.Preferences = preferences
	}
}

// Sets a user's API key
func (x *UserStorage) SetAPIKey(name string, apiKey string) *CleanUser {
	if user, ok := x.UserMap[name]; ok {
		user.APIKey = apiKey
		return &CleanUser{user.Name, apiKey}
	}

	return nil
}

// Interface for the data store
type StoreMethods interface {
	Get() (*UserStorage, error)
	Save() error
}

// Top-level data store struct
type DataStore struct {
	dataRef    *UserStorage
	filepath   string
	accessLock sync.Mutex
}

// Get the data stored in the store
func (x *DataStore) Get() *UserStorage {
	x.accessLock.Lock()
	defer x.accessLock.Unlock()
	return x.dataRef
}

// Save the data stored in the store
func (x *DataStore) Save() error {
	x.accessLock.Lock()
	defer x.accessLock.Unlock()

	data, err := json.Marshal(x.dataRef.Users)
	if err != nil {
		fmt.Println("Error marshalling data")
		return err
	}
	err = ioutil.WriteFile(x.filepath, data, 0644)
	if err != nil {
		fmt.Println("Error writing data")
		return err
	}
	fmt.Println("Saved data")
	return nil
}

// Constructor for the data store
func initStore(filepath string) *DataStore {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		data = []byte("[]")
	}

	store := &DataStore{}
	store.filepath = filepath
	store.dataRef = &UserStorage{
		Users:   []User{},
		UserMap: make(map[string]*User),
	}
	// load file into store memory
	_ = json.Unmarshal(data, &store.dataRef.Users)

	store.dataRef.UpdateUserMap()

	return store
}

// global variable for single data store across all files
var Store = initStore("data.json")
