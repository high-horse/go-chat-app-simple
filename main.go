package main

import (
	"log"
	"net/http"
)

func main() {
	println("Server listening at 8080!")
	SetupApi()

	// log.Fatal(http.ListenAndServe(":8000", nil))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func SetupApi() {

	manager := NewManager()

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.ServeWS)
}
