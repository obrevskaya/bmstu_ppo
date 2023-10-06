package view

import (
	"github.com/fatih/color"
)

func PrintErrCreateUserWine(err error) {
	c := color.New(color.FgRed)
	c.Printf("Error of add favourite wine.\n\t%s\n", err)
}

func PrintErrDeleteUserWine(err error) {
	c := color.New(color.FgRed)
	c.Printf("Error of delete favourite wine.\n\t%s\n", err)
}

func PrintErrGetUserWine(err error) {
	c := color.New(color.FgRed)
	c.Printf("Error of get favourite wines from service.\n\t%s\n", err)
}

func PrintErrParseUserWines(err error) {
	c := color.New(color.FgRed)
	c.Printf("Error of get favourite wines from response.\n\t%s\n", err)
}
