package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Access struct {
	Token string `json:"access_token"`
}

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Authorization redirect!")
	fmt.Println(r)

	code := r.URL.Query()["code"][0]

	payload := url.Values{}
	payload.Set("code", code)
	payload.Set("client_id", os.Getenv("LUNCHBOT_ID"))
	payload.Set("client_secret", os.Getenv("LUNCHBOT_SECRET"))

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://slack.com/api/oauth.access", strings.NewReader(payload.Encode()))
	req.Header.Set("Content-type", "x-www-form-urlencoded")
	resp, _ := client.Do(req)

}
