// Load a struct from a file into memory, and write it back to the file whenever it is changed
package store

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
type JSONStruct struct {
	Users []User
}

// Declare all methods
type JSONStructor interface {
	AddUser(name string, password string, apiKey string)
	GetUser(name string) *User
	RemoveUser(name string)
	GetUsers() []User
	IsUnique(name string, apiKey string) bool
	SetPreferences(name string, preferences map[string]string)
	GetPreferences(name string) map[string]string
	CorrectCredentials(name string, password string) bool
	GetDevices(name string) string
}

// Indirectly implement the GetUsers method
func (x *JSONStruct) GetUsers() []User {
	return x.Users
}

// API proxy (no CORS on the server)
func (x *JSONStruct) GetDevices(name string) string {
	// Get the user
	for _, user := range x.GetUsers() {
		if user.Name == name {
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
	}
	return ""
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
	map_ := make(map[string]interface{})

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

// Remove a user from the store
func (x *JSONStruct) RemoveUser(name string) bool {
	for i, user := range x.GetUsers() {
		if user.Name == name {
			x.Users = append(x.Users[:i], x.Users[i+1:]...)
			return true
		}
	}
	return false
}

// Check if the APIKey is valid
func (x *JSONStruct) APIKeyExists(apiKey string) bool {
	for _, user := range x.GetUsers() {
		if user.APIKey == apiKey {
			return true
		}
	}
	return false
}

// Gets a user's preferences
func (x *JSONStruct) GetPreferences(name string) map[string]interface{} {
	for _, user := range x.GetUsers() {
		if user.Name == name {
			return user.Preferences
		}
	}
	return nil
}

// Sets a user's preferences
func (x *JSONStruct) SetPreferences(name string, preferences map[string]interface{}) {
	for i, user := range x.GetUsers() {
		if user.Name == name {
			// Set the preferences by reference to original user in array
			x.Users[i].Preferences = preferences
		}
	}
}

// Sets a user's API key
func (x *JSONStruct) SetAPIKey(name string, apiKey string) *CleanUser {
	for i, user := range x.GetUsers() {
		if user.Name == name {
			x.Users[i].APIKey = apiKey
			return &CleanUser{user.Name, apiKey}
		}
	}

	return nil
}

// Interface for the data store
type Storer interface {
	Get() (*JSONStruct, error)
	Save() error
}

// Top-level data store struct
type DataStore struct {
	dataRef  *JSONStruct
	filepath string
}

// Get the data stored in the store
func (x *DataStore) Get() *JSONStruct {
	return x.dataRef
}

// Save the data stored in the store
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

// Constructor for the data store
func initStore(filepath string) *DataStore {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		data = []byte("{\"Users\":[]}")
	}

	store := &DataStore{}
	store.filepath = filepath
	store.dataRef = &JSONStruct{}
	// load file into store memory
	_ = json.Unmarshal(data, store.dataRef)

	return store
}

// global variable for single data store across all files
var Store = initStore("data.json")
