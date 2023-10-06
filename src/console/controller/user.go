package controller

import (
	myErrors "console/errors"
	"console/handlers"
	openapi "console/internal/client"
	"console/internal/consts"
	"console/utils"
	"console/view"
	"fmt"
	"log"
	"strconv"

	"github.com/dixonwille/wmenu"
	"github.com/fatih/color"
)

func (m *Menu) authorizeAdmin(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}

	authRequest, err := view.InputAuth()
	if err != nil {
		view.PrintErrInputAuthAdmin(err)
		return fmt.Errorf("error of input auth admin: %w", err)
	}

	response, err := handlers.AuthorizeClient(clientEntity.Client, authRequest)
	if err != nil {
		view.PrintErrAuth(err)
		return fmt.Errorf("error of auth: %w", err)
	}

	user, err := utils.ParseUserBody(response)
	if err != nil {
		view.PrintErrParseUser(err)
		return fmt.Errorf("error of input auth admin: %w", err)
	}
	if user.Status != strconv.Itoa(consts.Admin) {
		view.PrintErrAdminRight()
		return myErrors.ErrorAccess
	}

	m.login = authRequest.Login
	m.password = authRequest.Password
	return nil
}

func (m *Menu) registerAdmin(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}
	registerRequest, err := view.InputUser()
	if err != nil {
		view.PrintErrInputUser(err)
		return fmt.Errorf("error of input admin: %w", err)
	}

	registerRequest.Status = strconv.Itoa(consts.Admin)

	_, err = handlers.CreateClient(clientEntity.Client, registerRequest)
	if err != nil {
		view.PrintErrRegister(err)
		return fmt.Errorf("error of register admin: %w", err)
	}

	m.login = registerRequest.Login
	m.password = registerRequest.Password
	return nil
}

func (m *Menu) getUser(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}
	authReq := &openapi.AuthRequest{Login: m.login, Password: m.password}

	response, err := handlers.AuthorizeClient(clientEntity.Client, authReq)
	if err != nil {
		view.PrintErrAuth(err)
		return fmt.Errorf("error of auth: %w", err)

	}

	user, err := utils.ParseUserBody(response)
	if err != nil {
		view.PrintErrParseUser(err)
		return fmt.Errorf("error of parse user: %w", err)
	}

	view.PrintUser(user)
	c := color.New(color.FgGreen)
	c.Printf("Successful get user\n")
	return nil
}

func (m *Menu) authorize(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}

	authRequest, err := view.InputAuth()
	if err != nil {
		view.PrintErrInputAuth(err)
		return fmt.Errorf("error of input authorize: %w", err)
	}

	_, err = handlers.AuthorizeClient(clientEntity.Client, authRequest)
	if err != nil {
		view.PrintErrAuth(err)
		return fmt.Errorf("error of authorize: %w", err)
	}

	m.login = authRequest.Login
	m.password = authRequest.Password

	c := color.New(color.FgGreen)
	c.Printf("Successful authorize\n")

	return nil
}

func (m *Menu) register(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}

	registerRequest, err := view.InputUser()
	if err != nil {
		view.PrintErrInputUser(err)
		return fmt.Errorf("error of input user register: %w", err)
	}

	_, err = handlers.CreateClient(clientEntity.Client, registerRequest)
	if err != nil {
		view.PrintErrRegister(err)
		return fmt.Errorf("error of create user: %w", err)
	}

	m.login = registerRequest.Login
	m.password = registerRequest.Password
	c := color.New(color.FgGreen)
	c.Printf("Successful register\n")
	return nil
}
