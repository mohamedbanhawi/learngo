package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type fullname struct {
	Firstname string
	Lastname  string
}

func main() {
	// array to store fullnames
	var fullnames []fullname

	// prompt user to enter file name
	var fileName string

	fmt.Print("Enter Filename: ")
	_, err := fmt.Scan(&fileName)

	if err != nil {
		fmt.Print(err)
		return
	}

	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Print(err)
		readFile.Close()
		return
	}
	//TODO: buffer reads to handle longer lines
	scanner := bufio.NewScanner(readFile)

	// read file line by line
	for scanner.Scan() {
		var lineString string = scanner.Text()
		// split by whitespace
		var namesArray []string = strings.Fields(lineString)

		// first and last name only
		if len(namesArray) != 2 {
			fmt.Printf("Skipping bad line: %s\n", lineString)
		}

		// store new entry
		fullnames = append(fullnames,
			fullname{Firstname: namesArray[0],
				Lastname: namesArray[1]})
	}

	for _, name := range fullnames {

		bArray, err := json.Marshal(name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(bArray))
	}

}
