package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fatih/color"

	"github.com/dixonwille/wmenu"
)

type ClientEntity struct {
	Client *http.Client
}

type Menu struct {
	mainMenu *wmenu.Menu
	uMenu    *wmenu.Menu
	gMenu    *wmenu.Menu
	aMenu    *wmenu.Menu

	login    string
	password string
}

func NewMenu() *Menu {
	return &Menu{}
}

func (m *Menu) AddOptionsMain(client *http.Client) {
	m.mainMenu.Option("User", ClientEntity{client}, false, m.guestMenu)
	m.mainMenu.Option("Admin", ClientEntity{client}, false, m.adminAuthorizeMenu)
	m.mainMenu.Option("Exit", ClientEntity{client}, false, func(_ wmenu.Opt) error {
		return errExit
	})
}

func (m *Menu) RunMenu(client *http.Client) error {
	m.mainMenu = wmenu.NewMenu("Who are you?")
	m.AddOptionsMain(client)

	for {
		err := m.mainMenu.Run()
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
