package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type TokenFile struct {
	Token string `json:"access_token"`
}

func GetUserToken(user string) (*TokenFile, error) {
	f, err := ioutil.ReadFile(user)
	if err != nil {
		fmt.Printf("User %v has not authorized this bot\n", user)
		return nil, errors.New("User not found")
	}

	var tokenFile TokenFile
	json.NewDecoder(bytes.NewReader(f)).Decode(&tokenFile)

	return &tokenFile, nil
}
