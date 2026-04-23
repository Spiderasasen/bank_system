package main

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

//if the user wants to make another account
