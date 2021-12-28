package notmain

import "fmt"

// Notes: 
// 
// const , var
// int, 
// declaration :=  , assignment =
// declared but not used  : need to make it a rvalue somewhere
// cannot use const when only do declaration :=, ex : const c := 2 would generate error


func mainly(){
  fmt.Println("Hello")
  const conferenceName = "Go Conference"
  var conferenceTicket uint8 = 50
  fmt.Println("Welcome to ", conferenceName)
  // The line below would have a warning becasue %s and conferenceTicket int datatype doens't match, I expected this should be an error
  fmt.Printf("Welcome to here %v again %s times \n", conferenceName, conferenceTicket) 
  _ = conferenceTicket

  // when you only do declaration, you need to specify datatype
  // When you do declaration with assignment, you don't have to specify datatype
  var userName string
  var userTicket int
  userName = "Tom"
  userTicket = 2
  fmt.Println(userName)

  // get the datatype
  // Println doesn't do formatting, it only add spaces between them, you need to use Printf like the next line
  fmt.Println("datatype of userTicket is %T, userName is %T \n", userTicket, userName)
  fmt.Printf("datatype of userTicket is %T, userName is %T \n", userTicket, userName)


}
