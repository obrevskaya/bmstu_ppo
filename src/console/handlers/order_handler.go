package handlers

import (
	"bytes"
	myErrors "console/errors"
	openapi "console/internal/client"
	"console/internal/consts"
	"fmt"
	"net/http"
)

func PlaceOrder(client *http.Client, order *openapi.Order, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/orders"

	params := fmt.Sprintf("{\"id\": \"%s\", \"idUser\": \"%s\", \"totalPrice\": \"%s\", \"isPoints\": \"%s\", \"status\": \"%s\"}", order.Id,
		order.IdUser, order.TotalPrice, order.IsPoints, consts.PlacedOrder)
	jsonStr := []byte(params)

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	return DoRequestWithAuth(client, request, login, password)
}

func GetOrderByUser(client *http.Client, id string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/orders"
	params := fmt.Sprintf("{\"id\": \"%s\"}", id)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	return DoRequestWithAuth(client, request, login, password)
}
