package main

import (
	"./objects"
	"log"
	"net/http"
)

const (
	LISTEN_ADDRESS = "127.0.0.1:12345"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(LISTEN_ADDRESS, nil))
}
