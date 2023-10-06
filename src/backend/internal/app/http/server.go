package http

import (
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/interfaces"
	openapi "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/server"
)

const no = "false"
const yes = "true"

type Server struct {
	openapi.DefaultAPIService
	userLogic     interfaces.IUserController
	billLogic     interfaces.IBillController
	orderLogic    interfaces.IOrderController
	elemLogic     interfaces.IOrderElementController
	wineLogic     interfaces.IWineController
	userWineLogic interfaces.IUserWinesController
}

func NewServer(u interfaces.IUserController, b interfaces.IBillController,
	o interfaces.IOrderController, el interfaces.IOrderElementController,
	w interfaces.IWineController, uw interfaces.IUserWinesController) *Server {
	return &Server{
		userLogic:     u,
		billLogic:     b,
		orderLogic:    o,
		elemLogic:     el,
		wineLogic:     w,
		userWineLogic: uw,
	}
}
