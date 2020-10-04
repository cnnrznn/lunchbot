package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Command struct {
	UserID  string `json:"user_id"`
	Command string `json:"command"`
}

func ImbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/imback handler")
	var c Command

	r.ParseForm()
	c.UserID = r.FormValue("user_id")
	c.Command = r.FormValue("command")

	if c.Command != "/imback" {
		log.Printf("Wrong handler, command: %v\n", c.Command)
		return
	}

	tokenFile, err := GetUserToken(c.UserID)
	if err != nil {
		log.Printf("User %v has not authorized lunchbot\n", c.UserID)
		return
	}

	var status SetStatus
	body := bytes.NewBuffer([]byte{})

	status.Profile.StatusText = ""
	status.Profile.StatusEmoji = ""
	status.Profile.StatusExpiration = 0

	json.NewEncoder(body).Encode(status)

	log.Println("Sending 'clear status' request")
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://slack.com/api/users.profile.set", body)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tokenFile.Token))
	req.Header.Set("Content-type", "application/json")
	resp, _ := client.Do(req)

	var respBody map[string]interface{}
	var bs []byte
	json.NewDecoder(resp.Body).Decode(&respBody)

	log.Println(resp.Status)
	bs, _ = json.MarshalIndent(respBody, "", "  ")
	log.Println(string(bs))
}
