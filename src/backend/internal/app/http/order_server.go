package http

import (
	"context"
	"net/http"
	"strconv"

	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	openapi "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/server"
	"github.com/google/uuid"
)

func (s *Server) PlaceOrder(ctx context.Context, order openapi.Order) (openapi.ImplResponse, error) {
	orderID, err := uuid.Parse(order.Id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect order uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	userID, err := uuid.Parse(order.IdUser)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect user uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	totalPrice, err := strconv.Atoi(order.TotalPrice)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect total price." + order.TotalPrice,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	var isPoints bool
	if order.IsPoints == yes {
		isPoints = true
	} else if order.IsPoints == no {
		isPoints = false
	} else {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect is points." + order.IsPoints,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	orderModel := &models.Order{
		ID:         orderID,
		IDUser:     userID,
		TotalPrice: totalPrice,
		IsPoints:   isPoints,
		Status:     order.Status,
	}

	err = s.orderLogic.Update(ctx, orderModel)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't update order.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.PlaceOrderResponse{Placed: true},
	}, nil
}

func (s *Server) GetOrderByUser(ctx context.Context, request openapi.GetOrderByUserRequest) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	order, err := s.orderLogic.GetByUserInProcess(ctx, id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't get by user.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	orderAPI := openapi.Order{
		Id:         order.ID.String(),
		IdUser:     order.IDUser.String(),
		TotalPrice: strconv.Itoa(order.TotalPrice),
		IsPoints:   strconv.FormatBool(order.IsPoints),
		Status:     order.Status,
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: orderAPI,
	}, nil
}

func (s *Server) GetOrder(ctx context.Context, s2 string) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(s2)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	order, err := s.orderLogic.GetByID(ctx, id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't get order.",
				SystemMessage: err.Error(),
			},
		}, nil
	}
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.Order{
			Id:         order.ID.String(),
			IdUser:     order.IDUser.String(),
			TotalPrice: strconv.Itoa(order.TotalPrice),
			IsPoints:   strconv.FormatBool(order.IsPoints),
			Status:     order.Status,
		},
	}, nil
}
