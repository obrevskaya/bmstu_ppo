package view

import (
	openapi "console/internal/client"
	"fmt"

	"github.com/fatih/color"
)

func PrintWine(wine *openapi.Wine) {
	fmt.Printf("\tID: %s\n", wine.Id)
	fmt.Printf("\tName: %s\n", wine.Name)
	fmt.Printf("\tCount: %s\n", wine.Count)
	fmt.Printf("\tYear: %d\n", wine.Year)
	fmt.Printf("\tStrength: %d\n", wine.Strength)
	fmt.Printf("\tPrice: %s\n", wine.Price)
	fmt.Printf("\tType: %s\n", wine.Type)
}

func PrintWines(wines []openapi.Wine) {
	c := color.New(color.FgCyan)
	for i, wine := range wines {
		c.Printf("Wine №%d:\n", i+1)
		PrintWine(&wine)
	}
}

func PrintWineInOrder(w *openapi.Wine, el *openapi.Elem) {
	c := color.New(color.FgCyan)
	c.Printf("Element of order:\n")
	fmt.Printf("\tID Element of order: %s\n", el.Id)
	fmt.Printf("\tName: %s\n", w.Name)
	fmt.Printf("\tPrice of wine: %s\n", w.Price)
	fmt.Printf("\tCount: %d\n\n", el.Count)
}

func InputWine() (*openapi.AddWineRequest, error) {
	w := &openapi.AddWineRequest{}

	fmt.Print("Enter name wine: ")
	if _, err := fmt.Scan(&w.Name); err != nil {
		return nil, fmt.Errorf("input name wine: %w", err)
	}

	fmt.Print("Enter count: ")
	if _, err := fmt.Scan(&w.Count); err != nil {
		return nil, fmt.Errorf("input count: %w", err)
	}

	fmt.Print("Enter year: ")
	if _, err := fmt.Scan(&w.Year); err != nil {
		return nil, fmt.Errorf("input year: %w", err)
	}

	fmt.Print("Enter strength: ")
	if _, err := fmt.Scan(&w.Strength); err != nil {
		return nil, fmt.Errorf("input strength: %w", err)
	}

	fmt.Print("Enter price: ")
	if _, err := fmt.Scan(&w.Price); err != nil {
		return nil, fmt.Errorf("input price: %w", err)
	}

	fmt.Print("Enter type: ")
	if _, err := fmt.Scan(&w.Type); err != nil {
		return nil, fmt.Errorf("input type: %w", err)
	}

	return w, nil
}

func UpdateWine(wine *openapi.Wine) (*openapi.Wine, error) {
	w := &openapi.Wine{
		Id:       wine.Id,
		Name:     wine.Name,
		Count:    wine.Count,
		Year:     wine.Year,
		Strength: wine.Strength,
		Price:    wine.Price,
		Type:     wine.Type,
	}
	var flag string
	fmt.Printf("Update name (Old: %s) (yes or no): ", wine.Name)

	if _, err := fmt.Scan(&flag); err != nil {
		return nil, fmt.Errorf("input flag name: %w", err)
	}
	if flag == "yes" || flag == "y" {
		fmt.Printf("Enter name wine: ")
		if _, err := fmt.Scan(&w.Name); err != nil {
			return nil, fmt.Errorf("input name wine: %w", err)
		}
	}

	fmt.Printf("Update count (Old: %s) (yes or no): ", wine.Count)
	if _, err := fmt.Scan(&flag); err != nil {
		return nil, fmt.Errorf("input flag count: %w", err)
	}
	if flag == "yes" || flag == "y" {
		fmt.Printf("Enter count wine: ")
		if _, err := fmt.Scan(&w.Count); err != nil {
			return nil, fmt.Errorf("input count wine: %w", err)
		}
	}

	fmt.Printf("Update year (Old: %d) (yes or no): ", wine.Year)
	if _, err := fmt.Scan(&flag); err != nil {
		return nil, fmt.Errorf("input flag year: %w", err)
	}
	if flag == "yes" || flag == "y" {
		fmt.Printf("Enter year wine: ")
		if _, err := fmt.Scan(&w.Year); err != nil {
			return nil, fmt.Errorf("input year wine: %w", err)
		}
	}

	fmt.Printf("Update strength (Old: %d) (yes or no): ", wine.Strength)
	if _, err := fmt.Scan(&flag); err != nil {
		return nil, fmt.Errorf("input flag strength: %w", err)
	}
	if flag == "yes" || flag == "y" {
		fmt.Printf("Enter strength wine: ")
		if _, err := fmt.Scan(&w.Strength); err != nil {
			return nil, fmt.Errorf("input strength wine: %w", err)
		}
	}

	fmt.Printf("Update price (yes or no) (Old: %s): ", wine.Price)
	if _, err := fmt.Scan(&flag); err != nil {
		return nil, fmt.Errorf("input flag price: %w", err)
	}
	if flag == "yes" || flag == "y" {
		fmt.Printf("Enter price wine: ")
		if _, err := fmt.Scan(&w.Price); err != nil {
			return nil, fmt.Errorf("input price wine: %w", err)
		}
	}

	fmt.Printf("Update type (Old: %s) (yes or no): ", wine.Type)
	if _, err := fmt.Scan(&flag); err != nil {
		return nil, fmt.Errorf("input flag type: %w", err)
	}
	if flag == "yes" || flag == "y" {
		fmt.Printf("Enter type wine: ")
		if _, err := fmt.Scan(&w.Type); err != nil {
			return nil, fmt.Errorf("input type wine: %w", err)
		}
	}
	return w, nil
}

func InputLimitSkip() (*openapi.GetWinesRequest, error) {
	getWines := &openapi.GetWinesRequest{}

	fmt.Print("Enter limit: ")
	if _, err := fmt.Scan(&getWines.Limit); err != nil {
		return nil, fmt.Errorf("input limit: %w", err)
	}

	fmt.Print("Enter skip: ")
	if _, err := fmt.Scan(&getWines.Skip); err != nil {
		return nil, fmt.Errorf("input skip: %w", err)
	}

	return getWines, nil
}

func PrintErrInputWine(err error) {
	c := color.New(color.FgRed)
	c.Printf("Incorrect input field for wine.\n\t%s\n", err)
}

func PrintErrCreateWine(err error) {
	c := color.New(color.FgRed)
	c.Printf("Error of add wine.\n\t%s\n", err)
}

func PrintErrDeleteWine(err error) {
	c := color.New(color.FgRed)
	c.Printf("Error of add wine.\n\t%s\n", err)
}

func PrintErrGetWine(err error) {
	c := color.New(color.FgRed)
	c.Printf("Error of get wine from service.\n\t%s\n", err)
}

func PrintErrParseWine(err error) {
	c := color.New(color.FgRed)
	c.Printf("Error of get wine from response.\n\t%s\n", err)
}

func PrintErrInputUpdateWine(err error) {
	c := color.New(color.FgRed)
	c.Printf("Incorrect field for update wine.\n\t%s\n", err)
}

func PrintErrUpdateWine(err error) {
	c := color.New(color.FgRed)
	c.Printf("Error of update wine.\n\t%s\n", err)
}

func PrintErrGetWines(err error) {
	c := color.New(color.FgRed)
	c.Printf("Error of get wines from service.\n\t%s\n", err)
}

func PrintErrParseWines(err error) {
	c := color.New(color.FgRed)
	c.Printf("Error of get wines from response.\n\t%s\n", err)
}

func PrintErrInputLimitSkip(err error) {
	c := color.New(color.FgRed)
	c.Printf("Incorrect field limit or skip.\n\t%s\n", err)
}
