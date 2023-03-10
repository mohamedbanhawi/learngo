/*
A Program with two goroutines which have a race condition when executed
concurrently. Explain what the race condition is and how it can occur.
*/

package main

import (
	"fmt"  // Print statements
	"sync" // Wait groups
)

// WaitGroup is used to wait for the "main" program to finish goroutines.
var wg sync.WaitGroup

// global variable
var number int = 0

func incrementNumber() {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()
	println("incrementNumber")
	number++
}

func printNumber() {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()
	println("printNumber")
	fmt.Println(number)
}

func main() {
	/*
		The race condition in this Go program occurs because there are two concurrent goroutines accessing
		and modifying the global variable "number" without proper synchronization.
	*/

	wg.Add(2) // Two threads are running.

	/*
		Specifically, the "incrementNumber" goroutine increments the value of "number",
		while the "printNumber" goroutine prints its value. Since there is no synchronization mechanism in place,
		it is possible that the "printNumber" goroutine may read the value of "number" before it has been incremented by the "incrementNumber" goroutine,
		resulting in unpredictable and inconsistent behavior.
	*/
	go incrementNumber()
	go printNumber()
	/*
		In some cases the value of "number" is printed before its incremented.
		In some cases the "incrementNumber" print statement run,
		then the value of "number" is printed to stdout before incrementing "number" which is not considered determinstic behavior.
		Each time the program is run it will result in different result with the same conditions.
	*/

	wg.Wait() // Wait for the goroutines to finish.
}
