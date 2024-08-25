package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println("Present time is: ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	//Creating a date
	createdDate := time.Date(2019, time.July, 6, 6, 6, 6, 6, time.UTC)
	fmt.Println("Created date is:", createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
