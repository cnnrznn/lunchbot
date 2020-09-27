package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Event struct {
	Text string `json:"text"`
	User string `json:"user"`
}

type EventObject struct {
	Event Event `json:"event"`
}

type SetStatus struct {
	Profile Profile `json:"profile"`
}

type Profile struct {
	StatusText       string `json:"status_text"`
	StatusExpiration int    `json:"status_expiration"`
	StatusEmoji      string `json:"status_emoji"`
}

type TokenFile struct {
	Token string `json:"access_token"`
}

func EventHandler(w http.ResponseWriter, req *http.Request) {
	var o EventObject

	json.NewDecoder(req.Body).Decode(&o)
	fmt.Println(o.Event.Text)
	go SetStatusAway(o)
}

func SetStatusAway(o EventObject) {
	user := o.Event.User
	text := strings.ToLower(o.Event.Text)
	expiration := int(time.Now().Unix())

	var tokenFile TokenFile
	var status SetStatus
	body := bytes.NewBuffer([]byte{})
	var resp *http.Response
	var data map[string]interface{}

	f, err := ioutil.ReadFile(user)
	if err != nil {
		fmt.Printf("User %v has not authorized this bot\n", user)
		return
	}

	json.NewDecoder(bytes.NewReader(f)).Decode(&tokenFile)

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
	fmt.Println(data)
	fmt.Println(resp)
}
