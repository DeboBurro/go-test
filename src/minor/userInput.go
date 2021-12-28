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
	bookingSlice = append(bookingSlice, "Nana is good")
	fmt.Printf("The whole array is %v \n", bookingSlice)
	fmt.Printf("The first value is %v \n", bookingSlice[0])
	fmt.Printf("Array type : %T\n" , bookingSlice)
	fmt.Printf("Array length : %v \n", len(bookingSlice))	
}