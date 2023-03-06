package main

import (
	"fmt"     // I/O functions
	"sort"    // sort list
	"strconv" // convert string to int types
)

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop,
the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice,
and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character "X" instead of an integer.
*/

func main() {
	// slice to store user input
	var inputSlice = make([]int, 0, 3)

	// Loop through user prompt
	for {
		fmt.Print("Enter Number or X to quit: ")

		// store user input
		var inputString string
		// prompt user input
		_, err := fmt.Scan(&inputString)

		// error handling
		if err != nil {
			fmt.Println("Error processing user input try again")
			fmt.Println(err)
			continue
		}

		// check termination condition
		if inputString == "X" {
			fmt.Println("Quiting.")
			break
		}

		// convert string to integer
		inputInt, err := strconv.Atoi(inputString)

		// error handling
		if err != nil {
			fmt.Println("Error processing user input try again")
			fmt.Println(err)
			continue
		}

		// add integer to slice
		inputSlice = append(inputSlice, inputInt)

		// sort slice
		sort.Slice(inputSlice, func(i, j int) bool {
			return inputSlice[i] < inputSlice[j]
		})

		// iterate through sorted slice
		for _, element := range inputSlice {
			fmt.Println(element)
		}
	}

}
