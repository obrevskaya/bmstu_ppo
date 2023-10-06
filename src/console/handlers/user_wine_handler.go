package handlers

import (
	"bytes"
	myErrors "console/errors"
	"fmt"
	"net/http"
)

func CreateUserWine(client *http.Client, userID string, wineID string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/favourite"
	params := fmt.Sprintf("{\"idUser\": \"%s\", \"idWine\": \"%s\"}", userID, wineID)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func DeleteUserWine(client *http.Client, userID string, wineID string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/favourite"
	params := fmt.Sprintf("{\"idUser\": \"%s\", \"idWine\": \"%s\"}", userID, wineID)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func GetUserWines(client *http.Client, userID string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/favourite/" + userID

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}
