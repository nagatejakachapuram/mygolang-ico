package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on Structs")
	// No inheritance in go lang
	Teja := User{"Teja", "e@gmail.com", true, 21}
	fmt.Println(Teja)
	fmt.Printf("Teja details are: %+v\n", Teja)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
