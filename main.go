package main

import "fmt"

func main() {
	//boolean for loop
	var login bool

	for !login {
		login = mainLogin()
	}

	//getting the username and password
	var username string = getUsername()
	var password string = getPassword()
	fmt.Println(username)
	fmt.Println(password)
}
