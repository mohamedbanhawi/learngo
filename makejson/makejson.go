package main

import (
	"encoding/json" // JSON encoding
	"fmt"           // IO
)

func main() {

	// prompt user for input
	var inputName, inputAddress string
	fmt.Print("Enter Name: ")
	_, err := fmt.Scan(&inputName)
	// error handling
	if err != nil {
		fmt.Println("Failed to process name")
		fmt.Print(err)
		return
	}
	// prompt user for input
	fmt.Print("Enter Address: ")
	// error handling
	n, err := fmt.Scan(&inputAddress)

	if err != nil || n < 0 {
		fmt.Println("Failed to process address")
		fmt.Print(err)
		return
	}

	// create person map
	person := map[string]string{
		"name":    inputName,
		"address": inputAddress,
	}

	// encode as json
	bArray, err := json.Marshal(person)

	if err != nil {
		fmt.Println("Failed to process input")
		fmt.Print(err)
		return
	}

	// display to user
	fmt.Println(string(bArray))

}
