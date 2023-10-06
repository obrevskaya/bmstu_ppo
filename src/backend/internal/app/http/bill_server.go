package http

import (
	"context"
	"net/http"

	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	openapi "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/server"
	"github.com/google/uuid"
)

func (s *Server) PayBill(ctx context.Context, s2 string) (openapi.ImplResponse, error) {
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

	err = s.billLogic.UpdateBillStatus(ctx, id, models.PaidBill)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't update bill.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.PayBillResponse{Payed: true},
	}, nil
}
