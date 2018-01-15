package main

import (
	"bytes"
	"encoding/json"
	"endpoints"
	. "models/session_token"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestSessionTokenEndpoint(test *testing.T) {
	new(SessionToken).Initialize().RemoveAll(nil)

	apiHost := "http://127.0.0.1"
	resource := "/session"
	jsonValue := []byte(`{"email": "some_valid@email.com"}`)
	encodedData := bytes.NewBuffer(jsonValue)

	u, _ := url.ParseRequestURI(apiHost)
	u.Path = resource
	url := u.String()

	req, err := http.NewRequest("POST", url, encodedData)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		test.Fatalf("Can not create the request")
	}

	recorder := httptest.NewRecorder()
	endpoints.SessionResource(recorder, req, nil)

	if recorder.Code != http.StatusCreated {
		test.Errorf("expected status 201, got %d", recorder.Code)
	}

	type SessionResponse struct {
		token string
	}

	var response SessionResponse

	jsonResponseError := json.Unmarshal(recorder.Body.Bytes(), &response)

	if jsonResponseError != nil {
		test.Fatalf("Response does not have `token` key")
	}
}
