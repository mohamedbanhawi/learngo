package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var inputString string

	fmt.Printf("Enter String: ")

	in := bufio.NewReader(os.Stdin)
	inputString, err := in.ReadString('\n')

	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	var lowerCaseString string = strings.ToLower(inputString)
	var trimmedString string = strings.TrimSpace(lowerCaseString)

	var startsWithI bool = strings.HasPrefix(trimmedString, "i")
	var endsWithN bool = strings.HasSuffix(trimmedString, "n")
	var containsA bool = strings.Contains(trimmedString, "a")

	if startsWithI && endsWithN && containsA {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
