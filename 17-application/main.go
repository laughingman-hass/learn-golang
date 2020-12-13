package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":4000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
