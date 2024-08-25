package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on maps")

	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["RT"] = "rust"

	fmt.Println("list of all languages:", languages)
	fmt.Println("JS stands for:", languages["JS"])
	delete(languages, "RB")
	fmt.Println("list of all languages:", languages)

}
