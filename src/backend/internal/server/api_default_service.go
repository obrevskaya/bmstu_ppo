/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"net/http"
)

// DefaultAPIService is a service that implements the logic for the DefaultAPIServicer
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type DefaultAPIService struct {
}

// NewDefaultAPIService creates a default api service
func NewDefaultAPIService() DefaultAPIServicer {
	return &DefaultAPIService{}
}

// AddElem -
func (s *DefaultAPIService) AddElem(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update AddElem with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, AddElemResponse{}) or use other options such as http.Ok ...
	// return Response(200, AddElemResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddElem method not implemented")
}

// AddWine -
func (s *DefaultAPIService) AddWine(ctx context.Context, addWineRequest AddWineRequest) (ImplResponse, error) {
	// TODO - update AddWine with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, AddWineResponse{}) or use other options such as http.Ok ...
	// return Response(200, AddWineResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddWine method not implemented")
}

// Authorize -
func (s *DefaultAPIService) Authorize(ctx context.Context, authRequest AuthRequest) (ImplResponse, error) {
	// TODO - update Authorize with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, User{}) or use other options such as http.Ok ...
	// return Response(200, User{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("Authorize method not implemented")
}

// CreateElem -
func (s *DefaultAPIService) CreateElem(ctx context.Context, createElemRequest CreateElemRequest) (ImplResponse, error) {
	// TODO - update CreateElem with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, CreateElemResponse{}) or use other options such as http.Ok ...
	// return Response(200, CreateElemResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateElem method not implemented")
}

// CreateUserWine -
func (s *DefaultAPIService) CreateUserWine(ctx context.Context, userWine UserWine) (ImplResponse, error) {
	// TODO - update CreateUserWine with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, CreateUserWineResponse{}) or use other options such as http.Ok ...
	// return Response(200, CreateUserWineResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateUserWine method not implemented")
}

// DecreaseElem -
func (s *DefaultAPIService) DecreaseElem(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update DecreaseElem with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, DecreaseElemResponse{}) or use other options such as http.Ok ...
	// return Response(200, DecreaseElemResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DecreaseElem method not implemented")
}

// DeleteElem -
func (s *DefaultAPIService) DeleteElem(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update DeleteElem with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, DeleteElemResponse{}) or use other options such as http.Ok ...
	// return Response(200, DeleteElemResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteElem method not implemented")
}

// DeleteUserWine -
func (s *DefaultAPIService) DeleteUserWine(ctx context.Context, userWine UserWine) (ImplResponse, error) {
	// TODO - update DeleteUserWine with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, DeleteUserWineResponse{}) or use other options such as http.Ok ...
	// return Response(200, DeleteUserWineResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteUserWine method not implemented")
}

// DeleteWine -
func (s *DefaultAPIService) DeleteWine(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update DeleteWine with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, DeleteWineResponse{}) or use other options such as http.Ok ...
	// return Response(200, DeleteWineResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteWine method not implemented")
}

// GetByOrder -
func (s *DefaultAPIService) GetByOrder(ctx context.Context, getByOrderRequest GetByOrderRequest) (ImplResponse, error) {
	// TODO - update GetByOrder with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Elems{}) or use other options such as http.Ok ...
	// return Response(200, Elems{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetByOrder method not implemented")
}

// GetOrder -
func (s *DefaultAPIService) GetOrder(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update GetOrder with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Order{}) or use other options such as http.Ok ...
	// return Response(200, Order{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetOrder method not implemented")
}

// GetOrderByUser -
func (s *DefaultAPIService) GetOrderByUser(ctx context.Context, getOrderByUserRequest GetOrderByUserRequest) (ImplResponse, error) {
	// TODO - update GetOrderByUser with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Order{}) or use other options such as http.Ok ...
	// return Response(200, Order{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetOrderByUser method not implemented")
}

// GetUser -
func (s *DefaultAPIService) GetUser(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update GetUser with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, User{}) or use other options such as http.Ok ...
	// return Response(200, User{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUser method not implemented")
}

// GetUserWines -
func (s *DefaultAPIService) GetUserWines(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update GetUserWines with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, UserWines{}) or use other options such as http.Ok ...
	// return Response(200, UserWines{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUserWines method not implemented")
}

// GetWine -
func (s *DefaultAPIService) GetWine(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update GetWine with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Wine{}) or use other options such as http.Ok ...
	// return Response(200, Wine{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetWine method not implemented")
}

// GetWines -
func (s *DefaultAPIService) GetWines(ctx context.Context, getWinesRequest GetWinesRequest) (ImplResponse, error) {
	// TODO - update GetWines with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetWinesResponse{}) or use other options such as http.Ok ...
	// return Response(200, GetWinesResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetWines method not implemented")
}

// PayBill -
func (s *DefaultAPIService) PayBill(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update PayBill with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, PayBillResponse{}) or use other options such as http.Ok ...
	// return Response(200, PayBillResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PayBill method not implemented")
}

// PlaceOrder -
func (s *DefaultAPIService) PlaceOrder(ctx context.Context, order Order) (ImplResponse, error) {
	// TODO - update PlaceOrder with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, PlaceOrderResponse{}) or use other options such as http.Ok ...
	// return Response(200, PlaceOrderResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PlaceOrder method not implemented")
}

// Register -
func (s *DefaultAPIService) Register(ctx context.Context, registerRequest RegisterRequest) (ImplResponse, error) {
	// TODO - update Register with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, RegisterResponse{}) or use other options such as http.Ok ...
	// return Response(201, RegisterResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("Register method not implemented")
}

// UpdatePoints -
func (s *DefaultAPIService) UpdatePoints(ctx context.Context, updatePointsRequest UpdatePointsRequest) (ImplResponse, error) {
	// TODO - update UpdatePoints with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, RegisterResponse{}) or use other options such as http.Ok ...
	// return Response(200, RegisterResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdatePoints method not implemented")
}

// UpdateWine -
func (s *DefaultAPIService) UpdateWine(ctx context.Context, wine Wine) (ImplResponse, error) {
	// TODO - update UpdateWine with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, UpdateWineResponse{}) or use other options such as http.Ok ...
	// return Response(200, UpdateWineResponse{}), nil

	// TODO: Uncomment the next line to return response Response(0, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(0, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateWine method not implemented")
}
