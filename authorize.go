package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Token struct {
	AuthedUser AuthedUser `json:"authed_user"`
}

type AuthedUser struct {
	AccessToken string `json:"access_token"`
	ID          string `json:"id"`
}

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query()["code"][0]

	payload := url.Values{}
	payload.Set("code", code)
	payload.Set("client_id", os.Getenv("LUNCHBOT_ID"))
	payload.Set("client_secret", os.Getenv("LUNCHBOT_SECRET"))

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://slack.com/api/oauth.v2.access", strings.NewReader(payload.Encode()))
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)

	var token Token
	json.NewDecoder(resp.Body).Decode(&token)

	f, _ := os.Create(token.AuthedUser.ID)
	defer f.Close()

	json.NewEncoder(f).Encode(token.AuthedUser)

	w.Write([]byte("Success!"))
}
