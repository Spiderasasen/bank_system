package main

import "fmt"

// Account making a type of array list for this section
type Account struct {
	username       string
	password       string
	depositAmount  int
	withdrawAmount int
}

// main 4 accounts
var accounts []Account = []Account{
	{"admin", "password", 0, 2000},
	{"gordonb", "abc123", 500, 200},
	{"pablo", "latmein", 1200, 50},
	{"smithy", "qwerty", 300, 100},
}

// if the user wants to make another account
func makeAccount() {
	var accoutLoginItems string
	var placeHolder1 string
	var placeHolder2 string

	//asking the user what there username and password will be
	fmt.Println("What would you like to have as your username?")
	fmt.Scanln(&accoutLoginItems)
	placeHolder1 = accoutLoginItems
	fmt.Println("What would you like to have as your password?")
	fmt.Scanln(&placeHolder2)

	//appened it to the map
	accounts = append(accounts, Account{placeHolder1, placeHolder2, 0, 3000})
}

// getting the username of the Accounts table
func getUsernameAcounts(currentAccount int) string {
	return accounts[currentAccount].username
}

func getPasswordAcounts(currentAccount int) string {
	return accounts[currentAccount].password
}
