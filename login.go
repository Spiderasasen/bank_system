package main

import (
	"fmt"
	"strings"
)

// private vars
var usernameHelp int //used as a pointer in the function i call in line 24
var username string
var passwordHelp int
var password string

// colors
var Reset = "\033[0m"
var Yellow = "\033[33m"

// getters
func getUsername() string { return username }
func getPassword() string { return password }

func mainLogin() bool {
	//username
	if loginArea("username", &username, &usernameHelp, "What is the username in DVWA?", "The username is \"admin\"", "admin", "username") {
		//password
		var passwordSection bool = false
		for !passwordSection {
			passwordSection = loginArea("password", &password, &passwordHelp, "What is the password in DVWA?", "The password is \"password\"", "password", "password")
		}
		return true
	}
	return false
}

// main login area
func loginArea(prompt string, input *string, helpCounter *int, hint1 string, hint2 string, answer string, label string) bool {
	//gets the prompt to repeat the same system needed
	fmt.Println("Enter a " + prompt + ": \n" +
		Yellow + "type \"help\" if you forgot your " + prompt + Reset)

	//gets the string input
	fmt.Scan(input)

	//places the string to be lowercase
	*input = strings.ToLower(*input)

	//checks if the input is valid or not
	if loginChecker(*input, helpCounter, hint1, hint2, answer, label) {
		return true
	}

	return false
}

// checks if the login was successful or not
func loginChecker(userInput string, userHelp *int, hint1 string, hint2 string, answer string, outputType string) bool {
	//if the user input is the same as the intended answer
	if userInput == answer {
		fmt.Println()
		return true

		//the user ask for help
	} else if userInput == "help" {
		*userHelp++ //gets the orginal input from the class then returns that input

		//if userHelp == 1, give hint1
		if *userHelp == 1 {
			fmt.Println(hint1 + "\n")

			//will print hint 2
		} else {
			fmt.Println(hint2 + "\n")
		}

		//if the user just put in a wrong input
	} else {
		fmt.Println(Red + "Wrong " + outputType + " \n" + Reset)
	}

	return false
}
