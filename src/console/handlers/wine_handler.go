package handlers

import (
	"bytes"
	myErrors "console/errors"
	openapi "console/internal/client"
	"fmt"
	"net/http"
)

func GetWines(client *http.Client, r *openapi.GetWinesRequest) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/wines"
	params := fmt.Sprintf("{\"limit\": \"%s\", \"skip\": \"%s\"}", r.Limit, r.Skip)
	jsonStr := []byte(params)

	request, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	return DoRequest(client, request)
}

func GetWine(client *http.Client, id string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/wines/" + id

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	return DoRequest(client, request)
}

func AddWine(client *http.Client, r *openapi.AddWineRequest, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/wines"
	params := fmt.Sprintf("{\"name\": \"%s\", \"count\": \"%s\", \"year\": %d, \"strength\": %d, \"price\": \"%s\", \"type\": \"%s\"}",
		r.Name, r.Count, r.Year, r.Strength, r.Price, r.Type)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	return DoRequestWithAuth(client, request, login, password)
}

func DeleteWine(client *http.Client, id string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/wines/" + id

	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	return DoRequestWithAuth(client, request, login, password)
}

func UpdateWine(client *http.Client, r *openapi.Wine, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/wines"
	params := fmt.Sprintf("{\"id\": \"%s\", \"name\": \"%s\", \"count\": \"%s\", \"year\": %d, \"strength\": %d, \"price\": \"%s\", \"type\": \"%s\"}",
		r.Id, r.Name, r.Count, r.Year, r.Strength, r.Price, r.Type)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	return DoRequestWithAuth(client, request, login, password)
}
