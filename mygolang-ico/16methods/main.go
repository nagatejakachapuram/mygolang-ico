package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on Structs")
	// No inheritance in go lang
	Teja := User{"Teja", "e@gmail.com", true, 21}
	fmt.Println(Teja)
	fmt.Printf("Teja details are: %+v\n", Teja)
	Teja.GetStatus()
	Teja.NewEmail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)

}

func (u User) NewEmail() {
	u.Email = "Teja@go.dev"
	fmt.Println("New email of user is:", u.Email)

}
