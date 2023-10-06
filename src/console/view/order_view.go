package view

import (
	openapi "console/internal/client"
	"fmt"

	"github.com/fatih/color"
)

func InputIsPoints() (string, error) {
	var s string
	fmt.Print("Payed by points? Enter yes or no: ")
	if _, err := fmt.Scan(&s); err != nil {
		return "", fmt.Errorf("input id order: %w", err)
	}
	if s == "yes" || s == "y" {
		return "true", nil
	}

	return "false", nil
}

func PrintOrder(o *openapi.Order) {
	c := color.New(color.FgCyan)
	c.Printf("Order:\n")
	fmt.Printf("\tID: %s\n", o.Id)
	fmt.Printf("\tTotal Price: %s\n", o.TotalPrice)
	fmt.Printf("\tPayed by points: %s\n", o.IsPoints)
	fmt.Printf("\tStatus: %s\n", o.Status)
}

func PrintErrPlaceOrder(err error) {
	c := color.New(color.FgRed)
	c.Printf("error of place order.\n\t%s\n", err)
}

func PrintErrGetOrderByUser(err error) {
	c := color.New(color.FgRed)
	c.Printf("error of get order (maybe there are no current orders).\n\t%s\n", err)
}

func PrintErrParseOrder(err error) {
	c := color.New(color.FgRed)
	c.Printf("error of get order from response.\n\t%s\n", err)
}

func PrintErrInputIsPoints(err error) {
	c := color.New(color.FgRed)
	c.Printf("Incorrect input field.\n\t%s\n", err)
}
