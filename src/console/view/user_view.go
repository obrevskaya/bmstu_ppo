package view

import (
	openapi "console/internal/client"
	"console/internal/consts"
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

func InputUser() (*openapi.RegisterRequest, error) {
	user := &openapi.RegisterRequest{}

	fmt.Print("Enter login: ")
	if _, err := fmt.Scan(&user.Login); err != nil {
		return nil, fmt.Errorf("input login: %w", err)
	}

	fmt.Print("Enter password: ")
	if _, err := fmt.Scan(&user.Password); err != nil {
		return nil, fmt.Errorf("input password: %w", err)
	}

	fmt.Print("Enter name: ")
	if _, err := fmt.Scan(&user.Fio); err != nil {
		return nil, fmt.Errorf("input name: %w", err)
	}

	fmt.Print("Enter email: ")
	if _, err := fmt.Scan(&user.Email); err != nil {
		return nil, fmt.Errorf("input email: %w", err)
	}

	user.Points = nil
	user.Status = strconv.Itoa(consts.Customer)

	return user, nil
}

func InputAuth() (*openapi.AuthRequest, error) {
	auth := &openapi.AuthRequest{}

	fmt.Print("Enter login: ")
	if _, err := fmt.Scan(&auth.Login); err != nil {
		return nil, fmt.Errorf("input login: %w", err)
	}

	fmt.Print("Enter password: ")
	if _, err := fmt.Scan(&auth.Password); err != nil {
		return nil, fmt.Errorf("input password: %w", err)
	}

	return auth, nil
}

func PrintErrInputAuthAdmin(err error) {
	c := color.New(color.FgRed)
	c.Printf("Incorrect input field for authorize admin.\n\t%s\n", err)
}

func PrintErrInputAuth(err error) {
	c := color.New(color.FgRed)
	c.Printf("Incorrect input field for authorize.\n\t%s\n", err)
}

func PrintErrAuth(err error) {
	c := color.New(color.FgRed)
	c.Printf("Error of authorize (maybe invalid login or password).\n\t%s\n", err)
}

func PrintUser(user *openapi.User) {
	c := color.New(color.FgCyan)
	c.Printf("User:\n")

	fmt.Printf("\tLogin: %s\n", user.Login)
	fmt.Printf("\tName: %s\n", user.Fio)
	fmt.Printf("\tEmail: %s\n", user.Email)
	fmt.Printf("\tPoints: %s\n", user.Points)
	fmt.Printf("\tCount favourites wines: %s\n", user.CntFavourites)
	if user.Status == strconv.Itoa(consts.Admin) {
		fmt.Printf("\tStatus: Admin\n")
	}

}

func PrintErrInputUser(err error) {
	c := color.New(color.FgRed)
	c.Printf("Incorrect input field for register user.\n\t%s\n", err)
}

func PrintErrParseUser(err error) {
	c := color.New(color.FgRed)
	c.Printf("Error of get user from response.\n\t%s\n", err)
}

func PrintErrAdminRight() {
	c := color.New(color.FgRed)
	c.Printf("Error: this user isn't admin.\n")
}

func PrintErrRegister(err error) {
	c := color.New(color.FgRed)
	c.Printf("Error of register user.\n\t%s\n", err)
}
