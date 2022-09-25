package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	body := bytes.NewReader(make([]byte, 5000_000))
	resp, err := http.Post("http://localhost:8888", "text/plain", body)
	if err != nil {
		log.Fatalln("Request failed", err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("read body failed", err)
	}
	log.Println("Read data(String):\n", string(data))
}