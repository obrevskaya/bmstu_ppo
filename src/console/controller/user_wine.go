package controller

import (
	"console/handlers"
	openapi "console/internal/client"
	"console/utils"
	"console/view"
	"fmt"
	"log"

	"github.com/dixonwille/wmenu"
)

func (m *Menu) createUserWine(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}
	idWine, err := view.InputID()
	if err != nil {
		view.PrintErrInputID(err)
		return fmt.Errorf("error of input id: %w", err)
	}

	authRequest := &openapi.AuthRequest{Login: m.login, Password: m.password}
	response, err := handlers.AuthorizeClient(clientEntity.Client, authRequest)
	if err != nil {
		view.PrintErrAuth(err)
		return fmt.Errorf("error of authorize: %w", err)
	}

	user, err := utils.ParseUserBody(response)
	if err != nil {
		view.PrintErrParseUser(err)
		return fmt.Errorf("error of parse user: %w", err)
	}

	_, err = handlers.CreateUserWine(clientEntity.Client, user.Id, idWine, m.login, m.password)
	if err != nil {
		view.PrintErrCreateUserWine(err)
		return fmt.Errorf("error of create user wine: %w", err)
	}

	return nil
}

func (m *Menu) deleteUserWine(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}
	idWine, err := view.InputID()
	if err != nil {
		view.PrintErrInputID(err)
		return fmt.Errorf("error of input id: %w", err)
	}

	authRequest := &openapi.AuthRequest{Login: m.login, Password: m.password}
	response, err := handlers.AuthorizeClient(clientEntity.Client, authRequest)
	if err != nil {
		view.PrintErrAuth(err)
		return fmt.Errorf("error of authorize: %w", err)
	}

	user, err := utils.ParseUserBody(response)
	if err != nil {
		view.PrintErrParseUser(err)
		return fmt.Errorf("error of parse user: %w", err)
	}

	_, err = handlers.DeleteUserWine(clientEntity.Client, user.Id, idWine, m.login, m.password)
	if err != nil {
		view.PrintErrDeleteUserWine(err)
		return fmt.Errorf("error of delete user wine: %w", err)
	}

	return nil
}

func (m *Menu) getUserWines(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}
	authRequest := &openapi.AuthRequest{Login: m.login, Password: m.password}
	response, err := handlers.AuthorizeClient(clientEntity.Client, authRequest)
	if err != nil {
		view.PrintErrAuth(err)
		return fmt.Errorf("error of authorize: %w", err)
	}

	user, err := utils.ParseUserBody(response)
	if err != nil {
		view.PrintErrParseUser(err)
		return fmt.Errorf("error of parse user: %w", err)
	}

	response, err = handlers.GetUserWines(clientEntity.Client, user.Id, m.login, m.password)
	if err != nil {
		view.PrintErrGetUserWine(err)
		return fmt.Errorf("error of get user wine: %w", err)
	}

	userWines, err := utils.ParseUserWinesBody(response)
	if err != nil {
		view.PrintErrParseUserWines(err)
		return fmt.Errorf("error of parse user wines: %w", err)
	}

	var wines []openapi.Wine
	for _, userWine := range userWines {
		response, err = handlers.GetWine(clientEntity.Client, userWine.IdWine)
		if err != nil {
			view.PrintErrGetWine(err)
			return fmt.Errorf("error of get wine: %w", err)
		}
		wine, err := utils.ParseWineBody(response)
		if err != nil {
			view.PrintErrParseWine(err)
			return fmt.Errorf("error of parse wine: %w", err)
		}
		wines = append(wines, *wine)
	}

	view.PrintWines(wines)

	return nil
}
