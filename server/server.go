package main

import (
	"log"
	"net/http"
	. "server/config"
	"server/handlers"
)

func main() {
	port := GetEnv("PORT", "3000")
	// set up handlers
	http.HandleFunc("/api/login", handlers.Login)
	http.HandleFunc("/api/signup", handlers.SignUp)
	http.HandleFunc("/api/preferences", handlers.Preferences)
	http.HandleFunc("/api/setAPIKey", handlers.SetAPIKey)
	http.HandleFunc("/api/getDevices", handlers.GetDevices)
	http.HandleFunc("/api/ReverseGeocode", handlers.GetDevices)

	// start server
	fs := http.FileServer(http.Dir("../client/dist"))
	http.Handle("/", fs)
	log.Print("Listening on : " + port + "...")
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}
}
