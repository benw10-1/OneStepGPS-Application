// http request handler functions
package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/auth"
	"server/store"
)

// User + token combo
type Auth struct {
	User  store.CleanUser
	Token string
}

// Login unmarshal format
type LoginStruct struct {
	Name     string
	Password string
}

// Function and object aliases
var jsonStruct = store.Store.Get()
var save = store.Store.Save

// generic headers in-case more need to be set later
func Headers(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}
func SignAuth(user *store.CleanUser) []byte {
	token, err := auth.Sign(*user)
	if err != nil {
		fmt.Println(err)
		return []byte("Error: " + err.Error())
	}
	auth := Auth{*user, token}
	marshalled, err := json.Marshal(auth)

	if err != nil {
		fmt.Println(err)
		return []byte("Error: " + err.Error())
	}

	return marshalled
}

// Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	Headers(w)
	// guard clause
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// read body
	resBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	var login LoginStruct
	// load JSON into login variable
	err = json.Unmarshal(resBody, &login)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error: " + err.Error()))
		return
	}
	// verify user
	correct := jsonStruct.CorrectCredentials(login.Name, login.Password)

	if !correct {
		w.Write([]byte("Error: Incorrect username or password"))
		return
	}

	marshalled := SignAuth(jsonStruct.GetUser(login.Name))
	fmt.Println(string(marshalled))
	w.Write(marshalled)
}

// Add user handler (same logic as login)
func SignUp(w http.ResponseWriter, r *http.Request) {
	Headers(w)

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var user LoginStruct

	err = json.Unmarshal(body, &user)

	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusInternalServerError)
		return
	}
	// add user using meothod in store (returns false if user already exists)
	userAdded := jsonStruct.AddUser(user.Name, user.Password, "")

	if !userAdded {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error: " + err.Error()))
		return
	}
	save()
	marshalled := SignAuth(jsonStruct.GetUser(user.Name))

	w.Write(marshalled)
}

func SetAPIKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	Headers(w)
	// guard clause
	if r.Header.Get("Authorization") == "" {
		http.Error(w, "No authorization header", http.StatusUnauthorized)
		return
	}
	// get token from header
	token := r.Header.Get("Authorization")[7:]
	// verify token and get user
	user, err := auth.Validate(token)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}
	// unmarshal body into data
	var data map[string]string
	err = json.Unmarshal(body, &data)

	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusInternalServerError)
		return
	}

	edited := jsonStruct.SetAPIKey(user.Name, data["key"])
	// save user
	save()

	marshalled := SignAuth(edited)

	w.Write(marshalled)
}

// Need to be logged in to view or alter preferences
func Preferences(w http.ResponseWriter, r *http.Request) {
	Headers(w)
	// guard clause
	if r.Header.Get("Authorization") == "" {
		http.Error(w, "No authorization header", http.StatusUnauthorized)
		return
	}
	// get token from header
	token := r.Header.Get("Authorization")[7:]
	// verify token and get user
	user, err := auth.Validate(token)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	if r.Method == "GET" {
		preferences := jsonStruct.GetPreferences(user.Name)

		marshalled, err := json.Marshal(preferences)

		if err != nil {
			fmt.Println(err)
			w.Write([]byte("Error: " + err.Error()))
			return
		}

		w.Write(marshalled)
	} else if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}
		// unmarshal body into preferences
		var preferences map[string]string
		err = json.Unmarshal(body, &preferences)

		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusInternalServerError)
		}
		// update preferences
		jsonStruct.SetPreferences(user.Name, preferences)

		save()
		w.Write([]byte("Success"))
	} else {
		w.Write([]byte("Error: Method not allowed"))
	}
}
