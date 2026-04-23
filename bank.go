package main

import (
	"fmt"
	"math"
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
	fmt.Println("\nAmount: " + strconv.Itoa(amount) + "\nOn Hand: " + strconv.Itoa(onHand) + "\n\nPlease select a option")
	for option := range options {
		fmt.Println(options[option])
	}

	fmt.Scan(&choice)

	//seeing what option the user selected
	switch choice {
	case "a":
		fmt.Println("deposit...")
		amount = deposit(&amount, &onHand)
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

// depositing
func deposit(in *int, hand *int) int {
	fmt.Println("How much would you like to deposit?")

	//scanning the input
	_, err := fmt.Scanf("%d", in)
	if err != nil {
		fmt.Println("Invalid input")
		//if the user is less then 0, yell at the user
	} else if *in < 0 {
		fmt.Println("Please enter a positive amount")
		*in = int(math.Abs(float64(*in)))
		//otherwise allow the user to deposit the money
	} else {
		*hand -= *in
	}
	return *in
}

//withdrawing
