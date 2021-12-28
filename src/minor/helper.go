package main

import "strings"

// Need to specify an array of output datatypes
func validateUserInput(firstName string, lastName string, email string) (bool, bool){
	// This is really consise nice
	isValidName := len(firstName) >= 2 || len(lastName) !=2 && strings.Contains(email, "@")
	isValidEmail := len(email) >= 5
    return isValidName, isValidEmail
}

