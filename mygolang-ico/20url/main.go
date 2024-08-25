package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://go.dev"

func main() {
	fmt.Println("Welcome to Go lang")
	fmt.Println(myurl)

	//parsing the url
	result, _ := url.Parse(myurl)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

}
