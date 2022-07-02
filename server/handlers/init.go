package handlers

import (
	"server/auth"
	"server/store"
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Auth struct {
	User store.CleanUser
	Token string
}

type LoginStruct struct {
	Name string
	Password string
}

var jsonStruct = store.Store.Get()
var save = store.Store.Save

func Headers(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	var login LoginStruct

	err = json.Unmarshal(resBody, &login)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	correct := jsonStruct.CorrectCredentials(login.Name, login.Password)

	if !correct {
		w.Write([]byte("Error: Incorrect username or password"))
		return
	}

	user := jsonStruct.GetUser(login.Name)

	token, err := auth.Sign(*user)
	auth := Auth{*user, token}
	marshalled, err := json.Marshal(auth)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	Headers(w)
	w.Write(marshalled)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
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

	userAdded := jsonStruct.AddUser(user.Name, user.Password, "")
	save()

	if !userAdded {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}
	
	newUser := jsonStruct.GetUser(user.Name)

	token, err := auth.Sign(*newUser)
	auth := Auth{*newUser, token}
	marshalled, err := json.Marshal(auth)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	Headers(w)
	w.Write(marshalled)
}

func Preferences(w http.ResponseWriter, r *http.Request) {
	Headers(w)
	if r.Header.Get("Authorization") == "" {
		http.Error(w, "No authorization header", http.StatusUnauthorized)
		return
	}

	token := r.Header.Get("Authorization")[7:]

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

		var preferences map[string]string

		err = json.Unmarshal(body, &preferences)

		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusInternalServerError)
		}
		fmt.Println(preferences)
		jsonStruct.SetPreferences(user.Name, preferences)
		save()

		w.Write([]byte("Success"))
	} else {
		w.Write([]byte("Error: Method not allowed"))
	}
}
