/*
Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.
*/

package main

import (
	"fmt"
)

func GenDisplaceFn(a_mpsps, v_mps, s_m float64) func(float64) float64 {
	return func(t_s float64) float64 {
		return 0.5*a_mpsps*t_s*t_s + v_mps*t_s + s_m
	}
}

func main() {
	// constants
	var a_mpsps float64 // constant acceleration
	var v_mps float64   // initial velocity
	var s_m float64     // initial displacement
	var t_s float64     //Elapsed time

	// prompt user for input
	fmt.Print("Enter constant acceleration (m/s2): ")
	fmt.Scan(&a_mpsps)
	fmt.Print("Enter inital velocity (m/s): ")
	fmt.Scan(&v_mps)
	fmt.Print("Enter initial displacement (m): ")
	fmt.Scan(&s_m)

	var fn func(float64) float64 = GenDisplaceFn(a_mpsps, v_mps, s_m)

	for {
		fmt.Print("Enter elapsed time (s): ")
		fmt.Scan(&t_s)
		fmt.Printf("Displacement after %2.2f (s): %2.2f (m)\n", t_s, fn(t_s))
	}
}
