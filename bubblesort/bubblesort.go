package main

import (
	"fmt"     // I/O functions
	"strconv" // convert string to int types
)

func Swap(arr []int, index int) {
	var nextValue int = arr[index+1]
	arr[index+1] = arr[index]
	arr[index] = nextValue
}

func RunBubbleSortPass(arr []int) bool {
	var swapOccured bool = false
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			Swap(arr, i)
			swapOccured = true
		}
	}
	return swapOccured
}

func BubbleSort(arr []int) int {
	var passCounter int = 1 // at least one pass is needed
	for RunBubbleSortPass(arr) {
		passCounter++
	}
	return passCounter
}

func main() {
	// slice to store user input
	var inputSlice = make([]int, 0, 3)

	// Loop through user prompt
	for len(inputSlice) < 10 {
		fmt.Print("Enter Numbers or X to Sort: ")

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
			fmt.Println("Sorting")
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

	}

	var count int = BubbleSort(inputSlice)
	fmt.Printf("Array sorted after %d passes\n", count)

	// iterate through sorted slice
	for _, element := range inputSlice {
		fmt.Println(element)
	}

}
