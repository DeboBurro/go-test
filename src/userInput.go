package main
// package minor
// if the package is not a main pkg, you cannot run : go run ***.go
// import multiple modules
import (
	"fmt"
	"strings"
	"go-test/main"
	"strconv"
)

// Found weird stuff from fmt.Scan but #1359 and #1462 mentioned this is the best for developer to use
// also, we can use n, err := fmt.Scan(...) to determine if the input is valid

// This is a Package level varaible
// Package level variable cannot be create by :=
// conferenceName := "JustKiddingConf"
var conferenceName = "JustKiddingConf"

// [TODO] : take a look on how to use `make` in different ways 
//create a slice of map[string]string, but we need to initialize it with an default size
var userDataSliceMap = make([]map[string]string, 0)


func main(){
    greetUsers()
	// var userName string
	// var userTicket uint16
	// Need to pass the pointer of an variable so scan can assignstdin into it
	// fmt.Scan("Please type user name to be : ", &userName)
	// fmt.Scan(&userTicket)
	// fmt.Println("Hi" , userName)
	// fmt.Println(userTicket)

	// 
	remainTickets := 1
	var userTicket int16
	var totalTicket int16
	totalTicket = 50
	fmt.Println("Enter number of tickets you need: ")
	n, err := fmt.Scan(&userTicket)
	if err == nil{
		fmt.Println(userTicket)
		fmt.Println(n)
		fmt.Printf("%T\n", err)
	}else{
		fmt.Println("Error occured")
		fmt.Println(n)
        fmt.Println(err)
	}
    // int and uint are different data types
    totalTicket -= userTicket
    fmt.Println(totalTicket)
    // var x, y int8
    // n, err := fmt.Scan(&x, &y)
	// fmt.Println(x + y)
	// fmt.Println(n)
	// fmt.Println(err)

	// In GO, array is constant size while Slices is dynamical size of array
	// var bookings = [50]string{}
	// the size of an array need to be an constant integer, the line beow would fail
    // var bookings [totalTicket]string
    var bookings [50]string
	bookings[0] = "Nana" + " is good"
	fmt.Printf("The whole array is %v \n", bookings) // this stdout looks weird : The whole array is [Nana is good                                                 ] 
	fmt.Printf("The first value is %v \n", bookings[0])
	fmt.Printf("Array type : %T\n" , bookings)
	fmt.Printf("Array length : %v \n", len(bookings))

    // var bookingSlice = []string{}
	var bookingSlice []string
	bookingSlice = append(bookingSlice, "Nana good")
	bookingSlice = append(bookingSlice, "Debo bad")
	bookingSlice = append(bookingSlice, "Well done")
	fmt.Printf("The whole array is %v \n", bookingSlice)
	fmt.Printf("The first value is %v \n", bookingSlice[0])
	fmt.Printf("Array type : %T\n" , bookingSlice)
	fmt.Printf("Array length : %v \n", len(bookingSlice))

	// infinite loop
	// for {
	//  Put some code here to run 
	// }

	firstNames_pre := []string{}
    firstNames := printFirstNames(remainTickets, bookingSlice, firstNames_pre)

	remainTickets = 3
	// defined end status in the beginning of the loop
	for remainTickets >= 0 && len(firstNames) > 0{
		// just a random loop
		fmt.Println(remainTickets)
		remainTickets -= 1
	}
    
	// Validate user input
	var lastName string
	var firstName string
	var email string
	var userTickets uint
	fmt.Println("Enter your last name : ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your first name : ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your email : ")
	fmt.Scan(&email)
	fmt.Println("Enter your number of tickets you have got : ")
	fmt.Scan(&userTickets)

	// make creating a map
	var userData = make(map[string]string)
	userData["firstName"] = firstName
    userData["lastName"] = lastName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)
    userDataSliceMap = append(userDataSliceMap, userData)
	firstNameSlice := getFirstName()
	fmt.Println(firstNameSlice)

	isValidName,  isValidEmail := notmain.ValidateUserInput(firstName, lastName, email)
	fmt.Println(isValidName, isValidEmail)
	fmt.Println("lastname is ", lastName)
	// Newcomer need to know the difference between string, rune and characters from the link : https://go.dev/blog/strings
	// Otherwise, you would find it weird when trying to access each character from a string
    fmt.Printf("%c\n", []rune(lastName)[0])

	// switch statement
	city := "London"
	switch city{
	case "London", "England":
		// some code here
	case "Taiwan", "Taipei":
		// some code here
	case "US":
		// some code here
	default:
		fmt.Println("No valid city selected")
	}

}

// [TODO] : find out how to make optional function paramenters (or give them default values) 
// The datatype defined in the back is so weird  LOL
func greetUsers(){
	// Since the conferenceName varaiable is a package level variable, all the function here can access it
	fmt.Println("welcome here : ", conferenceName)
}

// I prefer string[] instead of []string                                            V--- the output data type here looks cute for some reason
func printFirstNames(remainTickets int, bookingSlice []string, firstNames []string) []string {
	for _, booking := range bookingSlice{
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
		if remainTickets == 0{
			fmt.Println("All tickets already got booked. Aborted.")
		    break
		}else if booking == "ABC" {
			fmt.Println("this is jt a dummy line")
		}else{
			remainTickets -= 1
			continue
		}
	}
	fmt.Println(firstNames)
	return firstNames
}

func getFirstName() []string{
	firstNames := []string{}
	for _, booking := range userDataSliceMap{
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}
