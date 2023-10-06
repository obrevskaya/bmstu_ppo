package handlers

import (
	myErrors "console/errors"
	"fmt"
	"net/http"
)

const port = "8081"
const address = "localhost"

func DoRequest(client *http.Client, request *http.Request) (*http.Response, error) {
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, myErrors.ErrorResponse
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return response, fmt.Errorf("%s : %d", myErrors.ErrorResponseStatus, response.StatusCode)
	}

	return response, nil
}

func DoRequestWithAuth(client *http.Client, request *http.Request, login string, password string) (*http.Response, error) {
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	return DoRequest(client, request)
}
