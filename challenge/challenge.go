package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Challenge struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
}

type ChallengeResponse struct {
	Challenge string `json:"challenge"`
}

func main() {
	http.HandleFunc("/event", EventHandler)
	log.Fatal(http.ListenAndServeTLS(":443", "../server.crt", "../server.key", nil))
}

func EventHandler(w http.ResponseWriter, req *http.Request) {
	var c Challenge
	var cr ChallengeResponse

	json.NewDecoder(req.Body).Decode(&c)

	cr.Challenge = c.Challenge

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cr)
}
