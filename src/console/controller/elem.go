package controller

import (
	"console/handlers"
	openapi "console/internal/client"
	"console/view"
	"fmt"
	"log"

	"github.com/dixonwille/wmenu"
	"github.com/fatih/color"
)

func (m *Menu) createElem(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}

	request, err := view.InputElem()
	if err != nil {
		view.PrintErrInputElem(err)
		return fmt.Errorf("error of input elem: %w", err)
	}

	_, err = handlers.CreateElem(clientEntity.Client, request, m.login, m.password)
	if err != nil {
		view.PrintErrCreateElem(err)
		return fmt.Errorf("error of create elem: %w", err)
	}
	c := color.New(color.FgGreen)
	c.Printf("Successful create elem\n")
	return nil
}

func (m *Menu) addElem(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}

	id, err := view.InputID()
	if err != nil {
		view.PrintErrInputID(err)
		return fmt.Errorf("error of input id: %w", err)
	}

	request := &openapi.AddElemRequest{}
	request.Id = id

	_, err = handlers.AddElem(clientEntity.Client, request, m.login, m.password)
	if err != nil {
		view.PrintErrAddElem(err)
		return fmt.Errorf("error of add elem: %w", err)
	}
	c := color.New(color.FgGreen)
	c.Printf("Successful add elem\n")
	return nil
}

func (m *Menu) decreaseElem(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}

	id, err := view.InputID()
	if err != nil {
		view.PrintErrInputID(err)
		return fmt.Errorf("error of input id: %w", err)
	}

	request := &openapi.DecreaseElemRequest{}
	request.Id = id

	_, err = handlers.DecreaseElem(clientEntity.Client, request, m.login, m.password)
	if err != nil {
		view.PrintErrDecreaseElem(err)
		return fmt.Errorf("error of decrease elem: %w", err)
	}
	c := color.New(color.FgGreen)
	c.Printf("Successful decrease elem\n")
	return nil
}

func (m *Menu) deleteElem(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}

	id, err := view.InputID()
	if err != nil {
		view.PrintErrInputID(err)
		return fmt.Errorf("error of input id: %w", err)
	}

	_, err = handlers.DeleteElem(clientEntity.Client, id, m.login, m.password)
	if err != nil {
		view.PrintErrDeleteElem(err)
		return fmt.Errorf("error of delete elem: %w", err)
	}
	c := color.New(color.FgGreen)
	c.Printf("Successful delete elem\n")
	return nil
}
