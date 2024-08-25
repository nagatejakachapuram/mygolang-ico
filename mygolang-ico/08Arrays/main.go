package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays class: ")

	// In arrays we must explicitly mention the size of the array
	var fruitlist [4]string
	fruitlist[0] = "Apple"
	fruitlist[1] = "Pear"
	fruitlist[2] = "Cherry"
	fruitlist[3] = "Banana"

	fmt.Println("fruitlist is:", fruitlist)
	fmt.Println("fruitlist length is:", len(fruitlist))

	var veglist = [3]string{"potato", "tomato", "cabbage"}
	fmt.Println("veglist is:", veglist)
	fmt.Println("veglist length is:", len(veglist))

}
