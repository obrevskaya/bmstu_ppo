package handlers

import (
	"bytes"
	myErrors "console/errors"
	openapi "console/internal/client"
	"fmt"
	"net/http"
)

func AuthorizeClient(client *http.Client, authRequest *openapi.AuthRequest) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/authorize"
	params := fmt.Sprintf("{\"login\": \"%s\", \"password\": \"%s\"}", authRequest.Login, authRequest.Password)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	return DoRequest(client, request)
}

func CreateClient(client *http.Client, r *openapi.RegisterRequest) (*http.Response, error) {

	url := "http://" + address + ":" + port + "/register"
	params := fmt.Sprintf("{\"login\": \"%s\", \"password\": \"%s\", \"fio\": \"%s\", "+
		"\"email\": \"%s\", \"points\": %d, \"status\": \"%s\"}", r.Login, r.Password, r.Fio, r.Email,
		r.Points, r.Status)
	jsonStr := []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	return DoRequest(client, request)
}
