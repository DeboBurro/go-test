package main
// package minor
// if the package is not a main pkg, you cannot run : go run ***.go
// import multiple modules
import (
	"fmt"
)

// Found weird stuff from fmt.Scan but #1359 and #1462 mentioned this is the best for developer to use
// also, we can use n, err := fmt.Scan(...) to determine if the input is valid

func main(){
	// var userName string
	// var userTicket uint16
	// Need to pass the pointer of an variable so scan can assignstdin into it
	// fmt.Scan("Please type user name to be : ", &userName)
	// fmt.Scan(&userTicket)
	// fmt.Println("Hi" , userName)
	// fmt.Println(userTicket)

	// 
	var userTicket int16
	var totalTicket int16
	totalTicket = 50
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

}