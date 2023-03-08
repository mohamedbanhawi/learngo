package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	locomotion string
	food       string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func main() {
	// place holder for valid animals
	animalMap := map[string]Animal{
		"cow":   Animal{food: "grass", noise: "moo", locomotion: "walk"},
		"bird":  Animal{food: "worms", noise: "peep", locomotion: "fly"},
		"snake": Animal{food: "mice", noise: "hss", locomotion: "slither"},
	}

	// place holder for action values
	actionMap := map[string]func(Animal){
		"eat":   (Animal).Eat,
		"speak": (Animal).Speak,
		"move":  (Animal).Move,
	}

	for {
		fmt.Printf(">")

		in := bufio.NewReader(os.Stdin)
		inputString, err := in.ReadString('\n')

		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		var lowerCaseString string = strings.ToLower(inputString)
		lowerCaseString = strings.TrimSuffix(lowerCaseString, "\n")

		// split string into animal and aciton
		inputSlice := strings.Split(lowerCaseString, " ")

		if len(inputSlice) != 2 {
			fmt.Println("Enter request as <Animal> <Action> for example cow speak")
			continue
		}

		var inputAnimal string = inputSlice[0]
		var inputAction string = inputSlice[1]

		selectedAnimal, foundAnimal := animalMap[inputAnimal]
		if !foundAnimal {
			fmt.Printf("Entered animal %s not valid\n", inputAnimal)
			continue
		}
		selectedActions, foundAction := actionMap[inputAction]

		if !foundAction {
			fmt.Printf("Entered action %s not valid\n", inputAction)
			continue
		}

		selectedActions(selectedAnimal)
	}
}
