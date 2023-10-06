package handlers

import (
	"bytes"
	myErrors "console/errors"
	openapi "console/internal/client"
	"fmt"
	"net/http"
)

func CreateElem(client *http.Client, r *openapi.CreateElemRequest, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/elems"
	params := fmt.Sprintf("{\"IdWine\": \"%s\", \"Count\": %d}", r.IdWine, r.Count)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	return DoRequestWithAuth(client, request, login, password)
}

func AddElem(client *http.Client, r *openapi.AddElemRequest, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/elems/" + r.Id + "/add"

	request, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	return DoRequestWithAuth(client, request, login, password)
}

func DecreaseElem(client *http.Client, r *openapi.DecreaseElemRequest, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/elems/" + r.Id + "/decrease"

	request, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	return DoRequestWithAuth(client, request, login, password)
}

func DeleteElem(client *http.Client, id string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/elems/" + id

	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func GetByOrder(client *http.Client, id string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/elems"
	params := fmt.Sprintf("{\"id\": \"%s\"}", id)
	jsonStr := []byte(params)
	request, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	return DoRequestWithAuth(client, request, login, password)
}
