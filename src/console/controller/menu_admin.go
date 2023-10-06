package controller

import (
	"errors"
	"fmt"
	"log"

	"github.com/dixonwille/wmenu"
	"github.com/fatih/color"
)

func (m *Menu) AddOptionsAdmin(ctx ClientEntity) {
	m.aMenu.Option("View wines", ctx, false, m.getWines)
	m.aMenu.Option("View wine", ctx, false, m.getWine)
	m.aMenu.Option("Register", ctx, false, m.registerAdmin)
	m.aMenu.Option("Add wine", ctx, false, m.addWine)
	m.aMenu.Option("Delete wine", ctx, false, m.deleteWine)
	m.aMenu.Option("Update wine", ctx, false, m.updateWine)
	m.aMenu.Option("Exit", ctx, false, func(_ wmenu.Opt) error {
		return errExit
	})
}

func (m *Menu) adminMenu(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}

	m.aMenu = wmenu.NewMenu("Choose an option.")
	m.AddOptionsAdmin(clientEntity)

	for {
		err := m.aMenu.Run()
		fmt.Println()
		if err != nil {
			if errors.Is(err, errExit) {
				break
			}

			c := color.New(color.FgRed)
			c.Printf("ERROR: %s\n\n", err)
		}
	}

	c := color.New(color.FgMagenta)
	c.Printf("Exited menu.\n")
	return nil
}

func (m *Menu) adminAuthorizeMenu(opt wmenu.Opt) error {

	err := m.authorizeAdmin(opt)
	if err != nil {
		return err
	}
	err = m.adminMenu(opt)
	if err != nil {
		return err
	}
	return nil
}
