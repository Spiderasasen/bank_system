package main

import (
	"fmt"
	"math"
	"strconv" //makes it possible to print out integers in the print statment
)

// import vars here
var options []string = []string{"a. deposit", "b. withdraw", "c. send money to another account", "d. logout"}

// colors
var Red = "\033[31m"
var Cyan = "\033[36m"
var Green = "\033[32m"

func mainBank(user string, pass string) bool {
	//getting the username and password
	var userPlace string = user
	var passPlace string = pass

	var choice string
	//printing out the user and the password
	fmt.Println("\nUsername: " + userPlace)
	fmt.Println("Password: " + passPlace)

	//showing the amount and the options
	fmt.Println("\nAmount: " + strconv.Itoa(accounts[currentAccount].depositAmount))
	fmt.Println("On Hand: " + strconv.Itoa(accounts[currentAccount].withdrawAmount))
	for option := range options {
		fmt.Println(options[option])
	}

	fmt.Scan(&choice)

	//seeing what option the user selected
	switch choice {
	case "a":
		fmt.Println(Green + "deposit..." + Reset)
		accounts[currentAccount].depositAmount = deposit(&accounts[currentAccount].depositAmount, &accounts[currentAccount].withdrawAmount)
		break
	case "b":
		fmt.Println(Green + "withdraw..." + Reset)
		accounts[currentAccount].withdrawAmount = withdraw(&accounts[currentAccount].depositAmount, &accounts[currentAccount].withdrawAmount)
		break
	case "c":
		fmt.Println(Green + "Sending money..." + Reset)
		accounts[currentAccount].withdrawAmount = sendingMoney(&accounts[currentAccount].withdrawAmount, user)
		break
	case "d":
		fmt.Println(Red + "Logging out..." + Reset)
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

// just showing the user that there are more accounts in the sending money section
func sendingMoney(hand *int, username string) int {
	var intOptions []int
	fmt.Println("Who would you like to send money too?")

	for i, accout := range accounts {
		if accout.username == username {
			continue
		}
		fmt.Println(strconv.Itoa(i) + ". " + accout.username)
		intOptions = append(intOptions, i)
	}

	var balance int
	var option int

	//scanning the option
	_, err := fmt.Scanf("%d", &option)
	if err != nil {
		fmt.Println(Red + "Invalid input" + Reset)
		return *hand
	} else if !isAOption(intOptions, option) {
		fmt.Println(Red + "Please enter a valid option." + Reset)
		return *hand
	}

	fmt.Println("How much would you like to send?")
	_, err = fmt.Scanf("%d", &balance)
	if err != nil {
		fmt.Println(Red + "Invalid input" + Reset)
	} else if balance < 0 {
		fmt.Println(Red + "Please enter a positive amount" + Reset)
	} else {
		*hand -= balance
	}

	return *hand
}

// checking if the user selected 1 of the options
func isAOption(options []int, choice int) bool {
	for _, option := range options {
		if choice == option {
			return true
		}
	}

	return false
}
