package http

import (
	"context"
	"net/http"

	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	openapi "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/server"
	"github.com/google/uuid"
)

func (s *Server) AddElem(ctx context.Context, s2 string) (openapi.ImplResponse, error) {
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

	err = s.elemLogic.Add(ctx, id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't add.",
				SystemMessage: err.Error(),
			},
		}, nil
	}
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.AddElemResponse{Added: true},
	}, nil

}

func (s *Server) CreateElem(ctx context.Context, elem openapi.CreateElemRequest) (openapi.ImplResponse, error) {
	wineID, err := uuid.Parse(elem.IdWine)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect wine uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	elemModel := &models.OrderElement{
		IDWine: wineID,
		Count:  int(elem.Count),
	}

	err = s.elemLogic.Create(ctx, elemModel)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't create elem.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.CreateElemResponse{Created: true},
	}, nil
}

func (s *Server) DecreaseElem(ctx context.Context, s2 string) (openapi.ImplResponse, error) {
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

	err = s.elemLogic.Decrease(ctx, id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't decrease.",
				SystemMessage: err.Error(),
			},
		}, nil
	}
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.DecreaseElemResponse{Decreased: true},
	}, nil

}

func (s *Server) DeleteElem(ctx context.Context, s2 string) (openapi.ImplResponse, error) {
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

	err = s.elemLogic.Delete(ctx, id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't delete elem.",
				SystemMessage: err.Error(),
			},
		}, nil
	}
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.DeleteElemResponse{Deleted: true},
	}, nil
}

func (s *Server) GetByOrder(ctx context.Context, request openapi.GetByOrderRequest) (openapi.ImplResponse, error) {
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

	elems, err := s.elemLogic.GetByOrder(ctx, id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't get by order.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	elemsApi := make([]openapi.Elem, len(elems))

	for i, el := range elems {
		elemsApi[i] = openapi.Elem{
			Id:      el.ID.String(),
			IdOrder: el.IDOrder.String(),
			IdWine:  el.IDWine.String(),
			Count:   int32(el.Count),
		}
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.Elems{
			Elems: elemsApi,
		},
	}, nil

}
