package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on loops")
	days := []string{"Monday", "Tuesday", "Wednesday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }
	// for _, day := range days {
	// 	fmt.Printf("index is and day is %v\n", day)
	// }

	value := 5
	for value < 10 {
		if value == 8 {
			goto lco
		}
		if value == 9 {
			break
		}
		fmt.Println("Print value is ", value)
		value++
	}
lco:
	fmt.Println("Your welcome")

}
