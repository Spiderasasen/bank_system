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

func mainLogin() bool {
	//username
	if usernameArea() {
		//password
		var passwordSection bool = false
		for !passwordSection {
			passwordSection = passwordArea()
		}
		return true
	}
	return false
}

// making the password
func passwordArea() bool {
	//gets the password
	fmt.Println("Enter Password: \n" +
		"type \"help\" if you forgot your password")
	fmt.Scan(&password)

	//puts the password in lowercase
	password = strings.ToLower(password)

	if loginChecker(password, &passwordHelp, "What is the password in DVWA?", "The password is \"password\"", "password", "password") {
		return true
	}
	return false
}

// making the username secion
func usernameArea() bool {
	//gets the username
	fmt.Println("Please enter a username: \n" +
		"type \"help\" if you forgot your username")
	fmt.Scan(&username)

	//puts the username input in lowercase
	username = strings.ToLower(username)

	//checks if the username is correct
	if loginChecker(username, &usernameHelp, "What is the username in DVWA?", "The username is \"admin\"", "admin", "username") {
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
		fmt.Println("Wrong " + outputType + " \n")
	}

	return false
}
