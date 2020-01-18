package handler

import (
	"encoding/json"
	"net/http"
)

type ResponseOAuth struct {
	AccessToken string `json:"access_token"`
}

func Token(w http.ResponseWriter, r *http.Request) {
	var data ResponseOAuth
	json, err := json.Marshal(data)
	if err != nil {
		panic("Failed to write response")
	}
	w.Write(json)
}
