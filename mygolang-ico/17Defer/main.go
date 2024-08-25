package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	Mydefer()
}

func Mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
