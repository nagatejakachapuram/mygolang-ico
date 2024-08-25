package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://go.dev"

func main() {
	fmt.Println("Web requests")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close()
	//Callers responsibilty to close response
	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)

}
