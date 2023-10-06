package view

import (
	"fmt"
)

func PrintRunMenu() {
	fmt.Printf("Set role\n" +
		"0) user\n1) admin\nChoice:\n")
}

func PrintGuestMenu() {
	fmt.Printf("User menu\n")
	fmt.Printf("0) Exit\n")
	fmt.Printf("1) Register\n")
	fmt.Printf("2) Authorize\n")
	fmt.Printf("3) Get wines\n")
	fmt.Printf("Choice:\n")
}

func PrintUserMenu() {
	fmt.Printf("User menu\n")
	fmt.Printf("0) Exit\n")
	fmt.Printf("1) Log out\n")
	fmt.Printf("2) Get wines\n")
	fmt.Printf("3) Create element of order with wine\n")
	fmt.Printf("4) Add element of order with wine\n")
	fmt.Printf("5) Decrease element of order with wine\n")
	fmt.Printf("6) Delete element of order with wine\n")
	fmt.Printf("7) Get order\n")
	fmt.Printf("8) Place an order\n")
	fmt.Printf("9) Get Wine\n")
	fmt.Printf("10) Get User\n")
	fmt.Printf("Choice:\n")
}

func PrintAdminMenu() {
	fmt.Printf("Admin menu\n")
	fmt.Printf("0) Exit\n")
	fmt.Printf("1) Register\n")
	fmt.Printf("2) Get wines\n")
	//fmt.Printf("3) Confirm pay bill\n")
	fmt.Printf("3) Add wine\n")
	fmt.Printf("4) Delete wine\n")
	fmt.Printf("5) Update wine\n")
	fmt.Printf("Choice:\n")
}
