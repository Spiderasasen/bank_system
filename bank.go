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
		fmt.Println("withdraw...")
		onHand = withdraw(&amount, &onHand)
		break
	case "c":
		fmt.Println("Logging out...")
		return true
	}

	return false
}

// depositing
func deposit(balance *int, hand *int) int {
	fmt.Println("\nHow much would you like to deposit?")

	var depositAmount int
	_, err := fmt.Scanf("%d", &depositAmount)

	if err != nil {
		fmt.Println("Invalid input")
		return *balance
	}

	if depositAmount < 0 {
		fmt.Println("Please enter a positive amount")
		depositAmount = int(math.Abs(float64(depositAmount)))
	}

	if *hand < depositAmount {
		fmt.Println("You don't have enough cash on hand")
		return *balance
	}

	*hand -= depositAmount
	*balance += depositAmount

	return *balance
}

// withdrawing
func withdraw(balance *int, hand *int) int {
	fmt.Println("\nHow much would you like to withdraw?")

	var withdrawAmount int
	_, err := fmt.Scanf("%d", &withdrawAmount)

	if err != nil {
		fmt.Println("Invalid input")
		return *hand
	}

	if withdrawAmount < 0 {
		fmt.Println("Please enter a positive amount")
		withdrawAmount = int(math.Abs(float64(withdrawAmount)))
	}

	if *balance < withdrawAmount {
		fmt.Println("You don't have enough in your account")
		return *hand
	}

	*balance -= withdrawAmount
	*hand += withdrawAmount

	return *hand
}

//withdrawing
