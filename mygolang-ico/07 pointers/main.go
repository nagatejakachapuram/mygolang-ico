package main

import "fmt"

func main() {
	fmt.Println("Welcome to today's class on pointers")

	// var ptr *int
	// fmt.Println("value of the pointer is", ptr)
	// pointer points to the memory address of the variable
	Mynumber := 99
	var ptr = &Mynumber
	fmt.Println("Value of the actual pointer", ptr)
	// *ptr -> value inside the address of the pointer
	fmt.Println("Value of the actual pointer", *ptr)

	*ptr = *ptr + 2
	fmt.Println("The new value is:", Mynumber)
}
