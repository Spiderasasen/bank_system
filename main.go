package main

func main() {
	//boolean for loop
	var login bool

	for !login {
		login = mainLogin()
	}

	//getting the username and password
	var username string = getUsername()
	var password string = getPassword()

	var bank bool
	for !bank {
		bank = mainBank(username, password)
	}
}
