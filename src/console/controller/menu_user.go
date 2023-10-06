package controller

import (
	"errors"
	"fmt"
	"log"

	"github.com/dixonwille/wmenu"
	"github.com/fatih/color"
)

var (
	errExit = errors.New("exit")
)

func (m *Menu) AddOptionsGuest(ctx ClientEntity) {
	m.gMenu.Option("Authorize", ctx, false, m.authorizeMenu)
	m.gMenu.Option("Register", ctx, false, m.registerMenu)
	m.gMenu.Option("View wines", ctx, false, m.getWines)
	m.gMenu.Option("Exit", ctx, false, func(_ wmenu.Opt) error {
		return errExit
	})
}

func (m *Menu) AddOptionsUser(ctx ClientEntity) {
	m.uMenu.Option("View wines", ctx, false, m.getWines)
	m.uMenu.Option("View wine", ctx, false, m.getWine)
	m.uMenu.Option("Add element of order", ctx, false, m.createElem)
	m.uMenu.Option("Increase element of order", ctx, false, m.addElem)
	m.uMenu.Option("Decrease element of order", ctx, false, m.decreaseElem)
	m.uMenu.Option("Delete element of order", ctx, false, m.deleteElem)
	m.uMenu.Option("View current order", ctx, false, m.getOrder)
	m.uMenu.Option("Place an order", ctx, false, m.placeOrder)
	m.uMenu.Option("View current user", ctx, false, m.getUser)
	m.uMenu.Option("Add wine in favourites", ctx, false, m.createUserWine)
	m.uMenu.Option("View favourites wines", ctx, false, m.getUserWines)
	m.uMenu.Option("Delete wine from favorites", ctx, false, m.deleteUserWine)
	m.uMenu.Option("Exit", ctx, false, func(_ wmenu.Opt) error {
		return errExit
	})
}

func (m *Menu) guestMenu(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}

	m.gMenu = wmenu.NewMenu("Choose an option.")
	m.AddOptionsGuest(clientEntity)

	for {
		err := m.gMenu.Run()
		fmt.Println()
		if err != nil {
			if errors.Is(err, errExit) {
				break
			}
		}
	}

	c := color.New(color.FgMagenta)
	c.Printf("Exited menu.\n")
	return nil

}

func (m *Menu) authorizeMenu(opt wmenu.Opt) error {
	err := m.authorize(opt)
	if err != nil {
		return err
	}
	err = m.userMenu(opt)
	if err != nil {
		return err
	}
	return nil
}

func (m *Menu) registerMenu(opt wmenu.Opt) error {
	err := m.register(opt)
	if err != nil {
		return err
	}
	err = m.userMenu(opt)
	if err != nil {
		return err
	}
	return nil
}

func (m *Menu) userMenu(opt wmenu.Opt) error {
	clientEntity, ok := opt.Value.(ClientEntity)
	if !ok {
		log.Fatal("Could not cast option's value to ClientEntity")
	}

	m.uMenu = wmenu.NewMenu("Choose an option.")
	m.AddOptionsUser(clientEntity)

	for {
		err := m.uMenu.Run()
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
