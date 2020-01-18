package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGrantType(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Token))
	var jsonStr string = `{"title":"Buy cheese and bread for breakfast."}`
	res, err := http.Post(server.URL, "application/json", bytes.NewBufferString(jsonStr))
	if err != nil {
		t.Log("Failed to create request")
		t.Log(err)
		t.FailNow()
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("Failed to create request")
		t.Log(err)
		t.FailNow()
	}
	var payload ResponseOAuth
	json.Unmarshal(body, &payload)
	if payload.AccessToken == "" {
		t.FailNow()
	}
}
