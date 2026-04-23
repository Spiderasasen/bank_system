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

// colors
var Red = "\033[31m"
var Cyan = "\033[36m"

func mainBank(user string, pass string) bool {
	var choice string
	//printing out the user and the password
	fmt.Println("\nUsername: " + user)
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
	default:
		fmt.Println(Red + "Invalid choice" + Reset)
	}

	return false
}

// depositing
func deposit(balance *int, hand *int) int {
	fmt.Println(Cyan + "\nHow much would you like to deposit?" + Reset)

	var depositAmount int
	_, err := fmt.Scanf("%d", &depositAmount)

	if err != nil {
		fmt.Println(Red + "Invalid input" + Reset)
		return *balance
	}

	if depositAmount < 0 {
		fmt.Println(Red + "Please enter a positive amount" + Reset)
		depositAmount = int(math.Abs(float64(depositAmount)))
	}

	if *hand < depositAmount {
		fmt.Println(Red + "You don't have enough cash on hand" + Reset)
		return *balance
	}

	*hand -= depositAmount
	*balance += depositAmount

	return *balance
}

// withdrawing
func withdraw(balance *int, hand *int) int {
	fmt.Println(Cyan + "\nHow much would you like to withdraw?" + Reset)

	var withdrawAmount int
	_, err := fmt.Scanf("%d", &withdrawAmount)

	if err != nil {
		fmt.Println(Red + "Invalid input" + Reset)
		return *hand
	}

	if withdrawAmount < 0 {
		fmt.Println(Red + "Please enter a positive amount" + Reset)
		withdrawAmount = int(math.Abs(float64(withdrawAmount)))
	}

	if *balance < withdrawAmount {
		fmt.Println(Red + "You don't have enough in your account" + Reset)
		return *hand
	}

	*balance -= withdrawAmount
	*hand += withdrawAmount

	return *hand
}

//withdrawing
