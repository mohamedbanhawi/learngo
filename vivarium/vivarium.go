package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// implement Animal Interface and Concrete Types

type Animal interface {
	Eat()
	Speak()
	Move()
}

type Cow struct{}

func (c *Cow) Eat() {
	fmt.Println("grass")
}

func (c *Cow) Speak() {
	fmt.Println("moo")
}

func (c *Cow) Move() {
	fmt.Println("walk")
}

type Bird struct{}

func (b *Bird) Eat() {
	fmt.Println("worms")
}

func (b *Bird) Speak() {
	fmt.Println("peep")
}

func (b *Bird) Move() {
	fmt.Println("fly")
}

type Snake struct{}

func (s *Snake) Eat() {
	fmt.Println("mice")
}

func (s *Snake) Speak() {
	fmt.Println("hiss")
}

func (s *Snake) Move() {
	fmt.Println("slither")
}

// Helpers

func getUserInput() (string, []string) {

	fmt.Printf(">")

	in := bufio.NewReader(os.Stdin)
	inputString, err := in.ReadString('\n')

	if err != nil {
		fmt.Println("Invalid input")
		return "", make([]string, 0)
	}

	var lowerCaseString string = strings.ToLower(inputString)
	lowerCaseString = strings.TrimSuffix(lowerCaseString, "\n")

	// split string into requesttype and requestParameters
	inputSlice := strings.Split(lowerCaseString, " ")

	if len(inputSlice) != 3 {
		fmt.Println("Failed to process inputs, enter `newanimal` or `query`")
	}

	return inputSlice[0], inputSlice[1:]

}

type AnimalStore struct {
	Records map[string]Animal
	Types   map[string]Animal
	Actions map[string]func(Animal)
}

func main() {
	// place holder for valid animals
	var store AnimalStore = AnimalStore{}

	store.Records = make(map[string]Animal, 0)
	store.Types = map[string]Animal{
		"cow":   &Cow{},
		"bird":  &Bird{},
		"snake": &Snake{},
	}
	store.Actions = map[string]func(Animal){
		"eat":   (Animal).Eat,
		"speak": (Animal).Speak,
		"move":  (Animal).Move,
	}

	for {
		requestType, requestParameters := getUserInput()

		if requestType == "query" {

			var queryAnimal string = requestParameters[0]
			var queryAction string = requestParameters[1]

			selectedAnimal, foundAnimal := store.Records[queryAnimal]
			if !foundAnimal {
				fmt.Printf("Entered animal name %s not valid\n", queryAnimal)
				continue
			}
			selectedActions, foundAction := store.Actions[queryAction]

			if !foundAction {
				fmt.Printf("Entered action %s not valid\n", queryAction)
				continue
			}

			selectedActions(selectedAnimal)

		} else if requestType == "newanimal" {

			var newAnimalName string = requestParameters[0]
			var newAnimalType string = requestParameters[1]

			selectedAnimal, foundAnimal := store.Types[newAnimalType]

			if !foundAnimal {
				fmt.Printf("Entered animal type %s not valid\n", newAnimalType)
				continue
			}

			store.Records[newAnimalName] = selectedAnimal
			fmt.Println("Created it!")
		}
	}
}
