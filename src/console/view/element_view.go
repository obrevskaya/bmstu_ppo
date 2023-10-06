package view

import (
	openapi "console/internal/client"
	"fmt"

	"github.com/fatih/color"
)

func InputElem() (*openapi.CreateElemRequest, error) {
	elem := &openapi.CreateElemRequest{}

	fmt.Print("Enter id wine: ")
	if _, err := fmt.Scan(&elem.IdWine); err != nil {
		return nil, fmt.Errorf("input id wine: %w", err)
	}

	fmt.Print("Enter count: ")
	if _, err := fmt.Scan(&elem.Count); err != nil {
		return nil, fmt.Errorf("input count: %w", err)
	}

	return elem, nil
}

func PrintErrInputElem(err error) {
	c := color.New(color.FgRed)
	c.Printf("Incorrect field for order element.\n\t%s\n", err)
}

func PrintErrCreateElem(err error) {
	c := color.New(color.FgRed)
	c.Printf("error of create order element.\n\t%s\n", err)
}

func PrintErrAddElem(err error) {
	c := color.New(color.FgRed)
	c.Printf("error of increase order element.\n\t%s\n", err)
}

func PrintErrDecreaseElem(err error) {
	c := color.New(color.FgRed)
	c.Printf("error of decrease order element.\n\t%s\n", err)
}

func PrintErrDeleteElem(err error) {
	c := color.New(color.FgRed)
	c.Printf("error of delete order element.\n\t%s\n", err)
}

func PrintErrParseElems(err error) {
	c := color.New(color.FgRed)
	c.Printf("error of get order elements from responce.\n\t%s\n", err)
}
