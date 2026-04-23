package main

import (
	"fmt"
	"strings"
)

// private vars
var username string
var password string

// colors
var Reset = "\033[0m"
var Yellow = "\033[33m"

// getters
func getUsername() string { return username }
func getPassword() string { return password }

func mainLogin() bool {

	//places all username in an array
	var usernames []string
	for _, account := range accounts {
		usernames = append(usernames, account.username)
	}

	//username
	if loginArea("username", "Type \"new\" to make a new account", &username) {
		//password
		var passwordSection bool = false
		for !passwordSection {
			passwordSection = loginArea("password", "Type \"help\" if you forgot your password", &password)
		}
		return true
	}
	return false
}

// main login area
func loginArea(prompt string, questionArea string, input *string) bool {
	//gets the prompt to repeat the same system needed
	fmt.Println("Enter a " + prompt + ": \n" +
		Yellow + questionArea + Reset)

	//gets the string input
	fmt.Scan(input)

	//places the string to be lowercase
	*input = strings.ToLower(*input)

	//checks if the input is valid or not
	if loginChecker(*input, prompt) {
		return true
	}

	return false
}

// checks if the login was successful or not
func loginChecker(userInput string, section string) bool {

	// USERNAME SECTION
	if section == "username" {

		// check if username exists in accounts
		for i, acc := range accounts {
			if userInput == acc.username {
				currentAccount = i
				fmt.Println()
				return true
			}
		}

		// new account command
		if userInput == "new" {
			fmt.Println("Making a new account...")
			makeAccount()
			return false
		}

		fmt.Println(Red + "Wrong username\n" + Reset)
		return false
	}

	// PASSWORD SECTION
	if section == "password" {

		// help command
		if userInput == "help" {
			fmt.Println("Your password is " + accounts[currentAccount].password)
			return false
		} else if userInput == accounts[currentAccount].password {
			fmt.Println()
			return true
		}

		fmt.Println(Red + "Wrong password\n" + Reset)
	}

	return false
}

// my version of conatins
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
