package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")
	var fruitlist = []string{"cherry", "Guava"}
	fmt.Printf("Type of fruitlist is: %T\n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Apple")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[:3])
	fmt.Println(fruitlist)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 233
	highScores[2] = 236
	highScores[3] = 223

	highScores = append(highScores, 333, 444, 555)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// Slicing via index

	var courses = []string{"reactjs", "ruby", "rust", "go", "swift", "python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
