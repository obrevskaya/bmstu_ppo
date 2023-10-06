package controller

import (
	"console/handlers"
	"console/utils"
	"console/view"
	"fmt"
	"log"

	"github.com/dixonwille/wmenu"
	"github.com/fatih/color"
)

func (m *Menu) addWine(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}
	wine, err := view.InputWine()
	if err != nil {
		view.PrintErrInputWine(err)
		return fmt.Errorf("error of input wine: %w", err)
	}

	_, err = handlers.AddWine(clientEntity.Client, wine, m.login, m.password)
	if err != nil {
		view.PrintErrCreateWine(err)
		return fmt.Errorf("error of create wine: %w", err)
	}

	return nil
}

func (m *Menu) deleteWine(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}
	id, err := view.InputID()
	if err != nil {
		view.PrintErrInputID(err)
		return fmt.Errorf("error of input id: %w", err)
	}

	_, err = handlers.DeleteWine(clientEntity.Client, id, m.login, m.password)
	if err != nil {
		view.PrintErrDeleteWine(err)
		return fmt.Errorf("error of delete wine: %w", err)
	}

	return nil
}

func (m *Menu) updateWine(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}

	id, err := view.InputID()
	if err != nil {
		view.PrintErrInputID(err)
		return fmt.Errorf("error of input id: %w", err)
	}

	response, err := handlers.GetWine(clientEntity.Client, id)
	if err != nil {
		view.PrintErrGetWine(err)
		return fmt.Errorf("error of get wine: %w", err)
	}
	oldWine, err := utils.ParseWineBody(response)
	if err != nil {
		view.PrintErrParseWine(err)
		return fmt.Errorf("error of parse wine: %w", err)
	}

	wine, err := view.UpdateWine(oldWine)
	if err != nil {
		view.PrintErrInputUpdateWine(err)
		return fmt.Errorf("error of input update wine: %w", err)
	}

	_, err = handlers.UpdateWine(clientEntity.Client, wine, m.login, m.password)
	if err != nil {
		view.PrintErrUpdateWine(err)
		return fmt.Errorf("error of update wine: %w", err)
	}

	return nil
}

func (m *Menu) getWines(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}

	winesRequest, err := view.InputLimitSkip()
	if err != nil {
		view.PrintErrInputLimitSkip(err)
		return fmt.Errorf("error of input limit or skip: %w", err)
	}

	response, err := handlers.GetWines(clientEntity.Client, winesRequest)
	if err != nil {
		view.PrintErrGetWines(err)
		return fmt.Errorf("error of get wines: %w", err)

	}

	wines, err := utils.ParseWinesBody(response)
	if err != nil {
		view.PrintErrParseWines(err)
		return fmt.Errorf("error of parse wines: %w", err)
	}

	view.PrintWines(wines)

	c := color.New(color.FgGreen)
	c.Printf("Successful get wines\n")

	return nil
}

func (m *Menu) getWine(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}

	id, err := view.InputID()
	if err != nil {
		view.PrintErrInputID(err)
		return fmt.Errorf("error of input id: %w", err)
	}

	response, err := handlers.GetWine(clientEntity.Client, id)
	if err != nil {
		view.PrintErrGetWine(err)
		return fmt.Errorf("error of get wine: %w", err)
	}

	wine, err := utils.ParseWineBody(response)
	if err != nil {
		view.PrintErrParseWine(err)
		return fmt.Errorf("error of parse wine: %w", err)
	}

	view.PrintWine(wine)

	c := color.New(color.FgGreen)
	c.Printf("Successful get wine\n")

	return nil
}
