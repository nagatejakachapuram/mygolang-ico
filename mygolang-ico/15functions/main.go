package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in go lang")
	greeter()
	result := adder(3, 5)
	fmt.Println("Result is ", result)
	prores := proAdder(2, 3, 4, 5, 6)
	fmt.Println("Pro result is ", prores)

}
func greeter() {
	fmt.Println("Good morning")

}
func adder(x int, y int) int {
	return x + y

}
func proAdder(values ...int) int {
	total := 0
	for _, values := range values {
		total += values
	}
	return total

}
