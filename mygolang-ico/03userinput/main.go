package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of the pizza")

	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating the pizza: ", rating)
	fmt.Printf("Type of the rating is %T", rating)
}
