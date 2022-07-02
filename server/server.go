package main

import (
	. "server/config"
	// . "server/handlers"
	. "server/store"
	"net/http"
	"log"
	// "os"
)

func main() {
	port := GetEnv("PORT", "3000")

	// http.HandleFunc("/api/", GetAPI)

	fs := http.FileServer(http.Dir("../client/dist"))
	http.Handle("/", fs)

	log.Print("Listening on : " + port + "...")
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal(err)
	}
}