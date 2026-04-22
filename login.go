package main

import (
	"fmt"
	"strings"
)

// private vars
var usernameHelp int
var username string

//var password string

func mainLogin() bool {

	//gets the username
	fmt.Println("Please enter a username: \n" +
		"type \"help\" if you forgot your username")
	fmt.Scan(&username)

	//puts the username input in lowercase
	username = strings.TrimSpace(username)

	//if the username is guessed admin correctly it will stop the loop
	if username == "admin" {
		return true

		//else if the username asked for help, it will return some help
	} else if username == "help" {
		usernameHelp++

		//if the user asked for help once, it will give out a hint
		if usernameHelp == 1 {
			fmt.Println("What is the username in DVWA?")
			//otherwise it will give out the username
		} else {
			fmt.Println("The username is \"admin\"")
		}

		//if the username is wrong, it will say its wrong
	} else {
		fmt.Println("Wrong Username")
	}

	return false
}
