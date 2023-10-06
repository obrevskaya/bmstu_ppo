package controller

import (
	"console/handlers"
	openapi "console/internal/client"
	"console/utils"
	"console/view"
	"fmt"
	"log"

	"github.com/dixonwille/wmenu"
	"github.com/fatih/color"
)

func (m *Menu) getOrder(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}

	authRequest := &openapi.AuthRequest{Password: m.password, Login: m.login}
	response, err := handlers.AuthorizeClient(clientEntity.Client, authRequest)
	if err != nil {
		view.PrintErrAuth(err)
		return fmt.Errorf("error of auth: %w", err)
	}

	user, err := utils.ParseUserBody(response)
	if err != nil {
		view.PrintErrParseUser(err)
		return fmt.Errorf("error of parse user: %w", err)
	}

	response, err = handlers.GetOrderByUser(clientEntity.Client, user.Id, m.login, m.password)
	if err != nil {
		view.PrintErrGetOrderByUser(err)
		return fmt.Errorf("error of get order by user: %w", err)
	}

	order, err := utils.ParseOrderBody(response)
	if err != nil {
		view.PrintErrParseOrder(err)
		return fmt.Errorf("error of parse order: %w", err)
	}

	view.PrintOrder(order)

	response, err = handlers.GetByOrder(clientEntity.Client, order.Id, m.login, m.password)

	elems, err := utils.ParseElemsBody(response)
	if err != nil {
		view.PrintErrParseElems(err)
		return fmt.Errorf("error of parse elems: %w", err)
	}
	for _, el := range elems {
		response, err = handlers.GetWine(clientEntity.Client, el.IdWine)
		if err != nil {
			view.PrintErrGetWine(err)
			return fmt.Errorf("error of get wine: %w", err)
		}

		wine, err := utils.ParseWineBody(response)
		if err != nil {
			view.PrintErrParseWine(err)
			return fmt.Errorf("error of parse wine: %w", err)

		}

		view.PrintWineInOrder(wine, &el)
	}
	c := color.New(color.FgGreen)
	c.Printf("\nSuccessful get order\n")
	return nil
}

func (m *Menu) placeOrder(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}
	authRequest := &openapi.AuthRequest{Password: m.password, Login: m.login}
	response, err := handlers.AuthorizeClient(clientEntity.Client, authRequest)
	if err != nil {
		view.PrintErrAuth(err)
		return fmt.Errorf("error of auth: %w", err)
	}

	user, err := utils.ParseUserBody(response)
	if err != nil {
		view.PrintErrParseUser(err)
		return fmt.Errorf("error of parse user: %w", err)

	}

	response, err = handlers.GetOrderByUser(clientEntity.Client, user.Id, m.login, m.password)
	if err != nil {
		view.PrintErrGetOrderByUser(err)
		return fmt.Errorf("error of get order by user: %w", err)

	}

	order, err := utils.ParseOrderBody(response)
	if err != nil {
		view.PrintErrParseOrder(err)
		return fmt.Errorf("error of parse order: %w", err)
	}

	order.IsPoints, err = view.InputIsPoints()
	if err != nil {
		view.PrintErrInputIsPoints(err)
		return fmt.Errorf("error of input is points: %w", err)

	}

	_, err = handlers.PlaceOrder(clientEntity.Client, order, m.login, m.password)
	if err != nil {
		view.PrintErrPlaceOrder(err)
		return fmt.Errorf("error of place order: %w", err)

	}
	c := color.New(color.FgGreen)
	c.Printf("Successful place order\n")
	return nil
}
