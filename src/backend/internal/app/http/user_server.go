package http

import (
	"context"
	"net/http"
	"strconv"

	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	openapi "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/server"
	"github.com/google/uuid"
)

func (s *Server) GetUser(ctx context.Context, userID string) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(userID)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	u, err := s.userLogic.Get(ctx, id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't get user.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.User{
			Id:            u.ID.String(),
			Login:         u.Login,
			Password:      u.Password,
			Fio:           u.Fio,
			Email:         u.Fio,
			Points:        strconv.Itoa(u.Points),
			Status:        strconv.Itoa(u.Status),
			CntFavourites: strconv.Itoa(u.CountFavourites),
		},
	}, nil
}

func (s *Server) Register(ctx context.Context, user openapi.RegisterRequest) (openapi.ImplResponse, error) {
	status, err := strconv.Atoi(user.Status)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect status." + user.Status,
				SystemMessage: err.Error(),
			},
		}, nil
	}
	userModel := &models.User{
		Login:    user.Login,
		Password: user.Password,
		Fio:      user.Fio,
		Email:    user.Email,
		Points:   int(user.Points),
		Status:   status,
	}

	err = s.userLogic.Create(ctx, userModel)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusConflict,
			Body: openapi.ErrorResponse{
				Message:       "can't create user.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusCreated,
		Body: openapi.RegisterResponse{Registered: true},
	}, nil
}

func (s *Server) UpdatePoints(ctx context.Context, request openapi.UpdatePointsRequest) (openapi.ImplResponse, error) {
	userID, err := uuid.Parse(request.Id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}
	err = s.userLogic.UpdateUserPoints(ctx, userID, int(request.Points))
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't update points.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.UpdateWineResponse{Updated: true},
	}, nil
}

func (s *Server) Authorize(ctx context.Context, request openapi.AuthRequest) (openapi.ImplResponse, error) {
	user, err := s.userLogic.Authorize(ctx, request.Login, request.Password)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusUnauthorized,
			Body: openapi.ErrorResponse{
				Message:       "Authorization failed.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.User{
			Id:            user.ID.String(),
			Login:         user.Login,
			Password:      user.Password,
			Fio:           user.Fio,
			Email:         user.Email,
			Points:        strconv.Itoa(user.Points),
			Status:        strconv.Itoa(user.Status),
			CntFavourites: strconv.Itoa(user.CountFavourites),
		},
	}, nil
}
