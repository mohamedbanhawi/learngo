/*
Write a program to sort an array of integers.
The program should partition the array into 4 parts,
 each of which is sorted by a different goroutine.
Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
When sorting is complete,
the main goroutine should print the entire sorted list.
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInputNumbers() []int {

	fmt.Printf("Enter a series of integers\n>")

	in := bufio.NewReader(os.Stdin)
	inputString, err := in.ReadString('\n')

	if err != nil {
		fmt.Println("Invalid input")
		return make([]int, 0)
	}

	var lowerCaseString string = strings.ToLower(inputString)
	lowerCaseString = strings.TrimSuffix(lowerCaseString, "\n")

	// split string into requesttype and requestParameters
	inputSlice := strings.Split(lowerCaseString, " ")

	var numbersSlice []int = make([]int, 0)

	for _, numberString := range inputSlice {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			fmt.Println(err)
			return make([]int, 0)
		}
		numbersSlice = append(numbersSlice, number)
	}
	return numbersSlice

}

func sortParition(partition []int, c chan []int) {
	sort.Ints(partition)
	fmt.Println(partition)
	c <- partition
}

func main() {

	numbersSlice := getInputNumbers()
	lenghtSlice := len(numbersSlice)

	if lenghtSlice <= 0 {
		return
	}
	numOfPartions := 4
	paritionSize := int(math.Round(float64(lenghtSlice) / float64(numOfPartions)))

	sortedMergeSlice := make([]int, 0)
	c := make(chan []int) // no buffering needed, main routine can only grab one result at a time

	for i := 0; i < int(numOfPartions); i++ {
		if i < numOfPartions-1 { // avoid slicing beyond the list size
			go sortParition(numbersSlice[i*paritionSize:(i+1)*paritionSize], c)
			continue
		}
		go sortParition(numbersSlice[i*paritionSize:], c)

	}

	for i := 0; i < int(numOfPartions); i++ {
		sortedParition := <-c
		sortedMergeSlice = append(sortedMergeSlice, sortedParition...)
	}
	sort.Ints(sortedMergeSlice)
	fmt.Println(sortedMergeSlice)
}
