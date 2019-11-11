package main

import (
	"./heartbeat"
	"log"
	"net/http"
)

func main() {
	go heartbeat.ListenHeartbeat()

	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)

	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
