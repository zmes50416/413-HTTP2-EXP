package main

import (
	"io"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	log.Println("Server starting...")
	h2s := &http2.Server{}
	server := &http.Server{
		Addr:    "0.0.0.0:8888",
		Handler: h2c.NewHandler(http.HandlerFunc(limitHandler), h2s),
	}
	log.Fatal(server.ListenAndServe())
}

func limitHandler(w http.ResponseWriter, r *http.Request) {
	data := make([]byte, 128)
	if _, err := r.Body.Read(data); err == io.EOF {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
		return
	} else if err != nil {
		log.Println("Read request body failed, ", err)
		return
	}
	
	w.WriteHeader(http.StatusRequestEntityTooLarge)
	w.Write([]byte(http.StatusText(http.StatusRequestEntityTooLarge)))
}