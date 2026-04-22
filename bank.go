package main

import (
	"fmt"
	"strconv" //makes it possible to print out integers in the print statment
)

// import vars here
var amount int = 0
var onHand int = 2000
var options []string = []string{"a. deposit", "b. withdraw", "c. logout"}

func mainBank(user string, pass string) bool {
	var choice string
	//printing out the user and the password
	fmt.Println("Username: " + user)
	fmt.Println("Password: " + pass)

	//showing the amount and the options
	fmt.Println("\nAmount: " + strconv.Itoa(amount) + "\n\nPlease select a option")
	for option := range options {
		fmt.Println(options[option])
	}

	fmt.Scan(&choice)

	switch choice {
	case "a":
		fmt.Println("deposit")
		break
	case "b":
		fmt.Println("withdraw")
		break
	case "c":
		fmt.Println("Logging out...")
		return true
	}

	return false
}
