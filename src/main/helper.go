package notmain
// this file will be considered as notmain package but under main folder so we need to import "go-test/main" to import this pkg


import "strings"

// Need to specify an array of output datatypes
// Capitalize the function means it will exported as a sort of public function that other package can use
func ValidateUserInput(firstName string, lastName string, email string) (bool, bool){
	// This is really consise nice
	isValidName := len(firstName) >= 2 || len(lastName) !=2 && strings.Contains(email, "@")
	isValidEmail := len(email) >= 5
    return isValidName, isValidEmail
}

