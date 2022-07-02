package main

import (
	. "server/config"
	. "server/handlers"
	"net/http"
	"log"
	"os"
)

func main() {
	port := GetEnv("PORT", "3000")
	apiKey := GetEnv("API_KEY", "")

	if apiKey == "" {
		log.Fatal("Please set API_KEY environment variable in .env file")
		os.Exit(1)
	}

	http.HandleFunc("/api/info", GetAPI)

	fs := http.FileServer(http.Dir("../client/dist"))
	http.Handle("/", fs)

	log.Print("Listening on : " + port + "...")
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal(err)
	}
}