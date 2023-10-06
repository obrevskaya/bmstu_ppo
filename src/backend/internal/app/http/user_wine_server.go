package http

import (
	"context"
	"net/http"

	openapi "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/server"
	"github.com/google/uuid"
)

func (s *Server) CreateUserWine(ctx context.Context, request openapi.UserWine) (openapi.ImplResponse, error) {
	wineID, err := uuid.Parse(request.IdWine)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect wine uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	userID, err := uuid.Parse(request.IdUser)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect user uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	err = s.userWineLogic.Create(ctx, userID, wineID)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't create user wine.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.CreateUserWineResponse{Created: true},
	}, nil
}

func (s *Server) DeleteUserWine(ctx context.Context, request openapi.UserWine) (openapi.ImplResponse, error) {
	wineID, err := uuid.Parse(request.IdWine)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect wine uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	userID, err := uuid.Parse(request.IdUser)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect user uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	err = s.userWineLogic.DeleteWine(ctx, userID, wineID)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't delete user wine.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.DeleteUserWineResponse{Deleted: true},
	}, nil
}

func (s *Server) GetUserWines(ctx context.Context, id string) (openapi.ImplResponse, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect user uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	wines, err := s.userWineLogic.GetByUser(ctx, userID)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't get user wines.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	userWinesApi := make([]openapi.UserWine, len(wines))

	for i, w := range wines {
		userWinesApi[i] = openapi.UserWine{
			IdUser: w.IDUser.String(),
			IdWine: w.IDWine.String(),
		}
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.UserWines{
			UserWines: userWinesApi,
		},
	}, nil
}
