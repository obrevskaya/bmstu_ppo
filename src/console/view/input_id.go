package view

import (
	"fmt"

	"github.com/fatih/color"
)

func InputID() (string, error) {
	var idString string
	fmt.Print("Enter ID: ")
	if _, err := fmt.Scan(&idString); err != nil {
		return "", fmt.Errorf("input id order: %w", err)
	}
	return idString, nil
}

func PrintErrInputID(err error) {
	c := color.New(color.FgRed)
	c.Printf("Incorrect input for id.\n\t%s\n", err)
}
