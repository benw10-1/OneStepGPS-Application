package main

import (
	. "server/config"
	"server/handlers"
	"net/http"
	"log"
)

func main() {
	port := GetEnv("PORT", "3000")

	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/signup", handlers.SignUp)
	http.HandleFunc("/preferences", handlers.Preferences)

	fs := http.FileServer(http.Dir("../client/dist"))
	http.Handle("/", fs)

	log.Print("Listening on : " + port + "...")
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal(err)
	}
}