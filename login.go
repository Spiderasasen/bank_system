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
	//username
	if loginArea("username", "Type \"new\" to make a new account", &username, "new", "admin", "username") {
		//password
		var passwordSection bool = false
		for !passwordSection {
			passwordSection = loginArea("password", "Type \"help\" if you forgot your password", &password, "help", "password", "password")
		}
		return true
	}
	return false
}

// main login area
func loginArea(prompt string, questionArea string, input *string, hintCommand string, answer string, label string) bool {
	//gets the prompt to repeat the same system needed
	fmt.Println("Enter a " + prompt + ": \n" +
		Yellow + questionArea + prompt + Reset)

	//gets the string input
	fmt.Scan(input)

	//places the string to be lowercase
	*input = strings.ToLower(*input)

	//checks if the input is valid or not
	if loginChecker(*input, prompt, hintCommand, answer, label) {
		return true
	}

	return false
}

// checks if the login was successful or not
func loginChecker(userInput string, section string, hintCommand string, answer string, outputType string) bool {
	//if the user input is the same as the intended answer
	if userInput == answer {
		fmt.Println()
		return true

		//the user ask for help
	} else if userInput == hintCommand {
		//if the user is in the login page and types new
		if hintCommand == "new" && section == "username" {
			fmt.Println("Making a new account...")
			makeAccount()
		}

		//if the user is in the password section and types help
		if hintCommand == "help" && section == "password" {
			fmt.Println("help")
		}

		//if the user just put in a wrong input
	} else {
		fmt.Println(Red + "Wrong " + outputType + " \n" + Reset)
	}

	return false
}
