package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Event struct {
	Text string `json:"text"`
	User string `json:"user"`
}

type EventObject struct {
	Event     Event `json:"event"`
	Timestamp int   `json:"event_time"`
}

type SetStatus struct {
	Profile Profile `json:"profile"`
}

type Profile struct {
	StatusText       string `json:"status_text"`
	StatusExpiration int    `json:"status_expiration"`
	StatusEmoji      string `json:"status_emoji"`
}

func EventHandler(w http.ResponseWriter, req *http.Request) {
	var o EventObject

	json.NewDecoder(req.Body).Decode(&o)
	log.Println(o.Event.Text)
	go SetStatusAway(o)
}

func SetStatusAway(e EventObject) {
	user := e.Event.User
	text := strings.ToLower(e.Event.Text)
	expiration := e.Timestamp
	log.Printf("Current unix time: %v\n", expiration)

	var tokenFile *TokenFile
	var status SetStatus
	body := bytes.NewBuffer([]byte{})
	var resp *http.Response
	var data map[string]interface{}

	tokenFile, err := GetUserToken(user)
	if err != nil {
		log.Println(err)
		return
	}

	switch {
	case strings.EqualFold(text, "lunch"), strings.EqualFold(text, "ofl"):
		status.Profile.StatusText = "Out for lunch"
		status.Profile.StatusEmoji = ":sandwich:"
		status.Profile.StatusExpiration = expiration + 3600
	case strings.EqualFold(text, "coffee"), strings.EqualFold(text, "ofc"):
		status.Profile.StatusEmoji = ":coffee:"
		status.Profile.StatusText = "Out for coffee"
		status.Profile.StatusExpiration = expiration + 1800
	default:
		// no matching text
		// POST with empty strings resets status
		return
	}

	json.NewEncoder(body).Encode(status)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://slack.com/api/users.profile.set", body)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tokenFile.Token))
	req.Header.Set("Content-type", "application/json")
	resp, _ = client.Do(req)

	json.NewDecoder(resp.Body).Decode(&data)
	log.Println(data)
}
