package main

import "fmt"

func main() {
	var username string = "Caesar"
	fmt.Println(username)
	fmt.Printf("Variable username is of type: %T \n", username)

	var isloggedin bool = true
	fmt.Println(isloggedin)
	fmt.Printf("Variable isloggedin is of type: %T \n", isloggedin)

	var smallval uint8 = 254
	fmt.Println(smallval)
	fmt.Printf("Variable username is of type: %T \n", smallval)

	var smallfloat float32 = 23.891111111
	fmt.Println(smallfloat)
	fmt.Printf("Variable username is of type: %T \n", smallfloat)

	//default values and aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Variable username is of type: %T \n", anothervariable)

	//implicit type
	var website = "cobratate.com"
	fmt.Println(website)

	// no var style
	number := 42.098
	fmt.Println(number)
}
