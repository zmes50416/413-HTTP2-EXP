package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8888", http.HandlerFunc(limitHandler)))
}

func limitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	}
	
	w.WriteHeader(http.StatusRequestEntityTooLarge)
	w.Write([]byte(http.StatusText(http.StatusRequestEntityTooLarge)))
}