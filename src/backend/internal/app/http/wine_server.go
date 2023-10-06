package http

import (
	"context"
	"net/http"
	"strconv"

	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	openapi "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/server"
	"github.com/google/uuid"
)

func (s *Server) DeleteWine(ctx context.Context, s2 string) (openapi.ImplResponse, error) {
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

	err = s.wineLogic.Delete(ctx, id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't delete wine.",
				SystemMessage: err.Error(),
			},
		}, nil
	}
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.DeleteWineResponse{Deleted: true},
	}, nil
}

func (s *Server) GetWines(ctx context.Context, request openapi.GetWinesRequest) (openapi.ImplResponse, error) {
	limit, err := strconv.Atoi(request.Limit)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect limit." + request.Limit,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	skip, err := strconv.Atoi(request.Skip)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect skip." + request.Skip,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	wines, err := s.wineLogic.GetWines(ctx, limit, skip)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't get wines.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	winesApi := make([]openapi.Wine, len(wines))

	for i, w := range wines {
		winesApi[i] = openapi.Wine{
			Id:       w.ID.String(),
			Name:     w.Name,
			Count:    strconv.Itoa(w.Count),
			Year:     int32(w.Year),
			Strength: int32(w.Strength),
			Price:    strconv.Itoa(w.Price),
			Type:     w.Type,
		}
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GetWinesResponse{
			Wines: winesApi,
		},
	}, nil
}

func (s *Server) UpdateWine(ctx context.Context, wine openapi.Wine) (openapi.ImplResponse, error) {
	wineID, err := uuid.Parse(wine.Id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	count, err := strconv.Atoi(wine.Count)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect count." + wine.Count,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	price, err := strconv.Atoi(wine.Price)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect price." + wine.Price,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	err = s.wineLogic.Update(ctx, &models.Wine{
		ID:       wineID,
		Name:     wine.Name,
		Count:    count,
		Year:     int(wine.Year),
		Strength: int(wine.Strength),
		Price:    price,
		Type:     wine.Type,
	})
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't update wine.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.UpdateWineResponse{Updated: true},
	}, nil
}

func (s *Server) GetWine(ctx context.Context, id string) (openapi.ImplResponse, error) {
	wineID, err := uuid.Parse(id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	w, err := s.wineLogic.GetWine(ctx, wineID)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't get wine.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.Wine{
			Id:       w.ID.String(),
			Name:     w.Name,
			Count:    strconv.Itoa(w.Count),
			Year:     int32(w.Year),
			Strength: int32(w.Strength),
			Price:    strconv.Itoa(w.Price),
			Type:     w.Type,
		},
	}, nil
}

func (s *Server) AddWine(ctx context.Context, wine openapi.AddWineRequest) (openapi.ImplResponse, error) {
	count, err := strconv.Atoi(wine.Count)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect count." + wine.Count,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	price, err := strconv.Atoi(wine.Price)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect price." + wine.Count,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	wineModel := &models.Wine{
		Name:     wine.Name,
		Count:    count,
		Year:     int(wine.Year),
		Strength: int(wine.Strength),
		Price:    price,
		Type:     wine.Type,
	}

	err = s.wineLogic.Create(ctx, wineModel)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't create wine.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.AddWineResponse{Added: true},
	}, nil
}
