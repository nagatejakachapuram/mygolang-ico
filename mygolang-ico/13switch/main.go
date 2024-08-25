// Building a Dice game

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in GoLang")

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNum)
	switch diceNum {
	case 1:
		fmt.Println("Dice value is 1, You can start")
	case 2:
		fmt.Println("Dice can move 2 places")
	case 3:
		fmt.Println("Dice can move 3 places")
	case 4:
		fmt.Println("Dice can move 4 places")
	case 5:
		fmt.Println("Dice can move 5 places")
	case 6:
		fmt.Println("Dice can move 6 places and youy can roll dice again")
		fallthrough
	default:
		fmt.Println("What was that!")
	}
}
