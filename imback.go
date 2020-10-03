package main

import (
	"bytes"
	"fmt"
	"net/http"
)

type Command struct {
	UserID  string `json:"user_id"`
	Command string `json:"command"`
}

func ImbackHandler(w http.ResponseWriter, r *http.Request) {
	var c Command

	r.ParseForm()
	c.UserID = r.FormValue("user_id")
	c.Command = r.FormValue("command")

	if c.Command != "imback" {
		fmt.Printf("Wrong handler, command: %v\n", c.Command)
		return
	}

	tokenFile, err := GetUserToken(c.UserID)
	if err != nil {
		fmt.Printf("User %v has not authorized lunchbot\n", c.UserID)
		return
	}

	var status SetStatus
	body := bytes.NewBuffer([]byte{})

	status.Profile.StatusText = ""
	status.Profile.StatusEmoji = ""
	status.Profile.StatusExpiration = 0

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://slack.com/api/users.profile.set", body)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tokenFile.Token))
	req.Header.Set("Content-type", "application/json")
	client.Do(req)
}
