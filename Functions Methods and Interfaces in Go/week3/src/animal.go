package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal with ts three main properties
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat prints out the Animal's kind of food it eats
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move prints out the Animal's kind of locomotion it uses
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak prints out the Animal's kind of noise it produces
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	am := make(map[string]Animal)
	am["cow"] = Animal{"grass", "walk", "moo"}
	am["bird"] = Animal{"worms", "fly", "peep"}
	am["snake"] = Animal{"mice ", "slither", "hsss"}

	for {
		requestedAnimal, requestedAction, quit := getUserRequest()

		if quit == true {
			break
		}

		animal, existing := am[requestedAnimal]
		if !existing {
			fmt.Printf("Animal '%s' not known, please try again.\n", requestedAnimal)
			continue
		}
		// action := v[1]

		switch requestedAction {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("No such action as '%s', please try again\n", requestedAction)
		}
	}
}

func getUserRequest() (string, string, bool) {

	stdin := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter an animal and an action, X to quit\n> ")
		stdin.Scan()
		scanned := stdin.Text()
		scanned = strings.TrimSpace(scanned)

		if strings.EqualFold(scanned, "X") {
			return "", "", true
		}

		resulting := make([]string, 2, 2)
		cursor := 0
		splitted := strings.Split(scanned, " ")

		if len(splitted) < 2 {
			fmt.Println("Must enter both an animal and an action, please try again.")
			continue
		}

		for _, y := range splitted {
			if len(y) > 0 {
				resulting[cursor] = y
				cursor++
				if cursor >= 2 {
					break
				}
			}
		}
		return strings.ToLower(resulting[0]), strings.ToLower(resulting[1]), false
	}
}
