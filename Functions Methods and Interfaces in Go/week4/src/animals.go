package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Animal an interface with three main behaviours
type Animal interface {
	Eat() string   // prints the type of food animal eats
	Move() string  // prints the kind kind of locomotion an animal uses
	Speak() string // prints the kind of noise an animal produces
}

// ======================================================
// COW
// ======================================================

// A Cow Animal
type Cow struct {
	name string
}

// Associated Methods

// Eat prints the type of food animal eats
func (c Cow) Eat() string {
	return "grass"
}

// Move prints the kind kind of locomotion an animal uses
func (c Cow) Move() string {
	return "walk"
}

// Speak prints the kind of noise an animal produces
func (c Cow) Speak() string {
	return "moo"
}

// Name prints out the Animal's Name
func (c Cow) Name() string {
	return c.name
}

// ======================================================
// BIRD
// ======================================================

// A Bird Animal
type Bird struct {
	name string
}

// Associated Methods

// Eat prints the type of food animal eats
func (c Bird) Eat() string {
	return "worms"
}

// Move prints the kind kind of locomotion an animal uses
func (c Bird) Move() string {
	return "fly"
}

// Speak prints the kind of noise an animal produces
func (c Bird) Speak() string {
	return "peep"
}

// Name prints out the Animal's Name
func (c Bird) Name() string {
	return c.name
}

// ======================================================
// SNAKE
// ======================================================

// A Snake Animal
type Snake struct {
	name string
}

// Associated Methods

// Eat prints the type of food animal eats
func (c Snake) Eat() string {
	return "mice"
}

// Move prints the kind kind of locomotion an animal uses
func (c Snake) Move() string {
	return "slither"
}

// Speak prints the kind of noise an animal produces
func (c Snake) Speak() string {
	return "hsss"
}

// Name prints out the Animal's Name
func (c Snake) Name() string {
	return c.name
}

func main() {

	animalDatabase := make(map[string]Animal)

	for {
		// requestedAnimal, requestedAction, quit := getUserRequest()
		requestedCommand, requestedName, requestedOption, quit := getUserRequest()

		if quit == true {
			break
		}

		switch requestedCommand {
		case "newanimal":
			// fmt.Println("Make new animal:  " + requestedOption + ", named: " + requestedName)
			err := createNewAnimal(requestedName, requestedOption, animalDatabase)
			if err != nil {
				fmt.Printf("%s. Please try again\n", err)
			} else {
				fmt.Println("Created it!")
			}

		case "query":
			response, err := queryAnimalAction(requestedName, requestedOption, animalDatabase)
			if err != nil {
				fmt.Printf("%s. Please try again\n", err)
			} else {
				fmt.Println(response)
			}
			fmt.Println("Query:  " + requestedName + ", action: " + requestedOption)

		default:
			fmt.Printf("Unknown command '%s', please try again.\n", requestedCommand)
			continue
		}
	}
}

func getUserRequest() (string, string, string, bool) {

	stdin := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter command newanimal or query, X to quit\n> ")
		stdin.Scan()
		scanned := stdin.Text()
		scanned = strings.TrimSpace(scanned)

		if strings.EqualFold(scanned, "X") {
			return "", "", "", true
		}

		fields := strings.Fields(scanned)

		if len(fields) != 3 {
			fmt.Println("Bad request, three strings are required")
			continue
		} else {
			return strings.ToLower(fields[0]), strings.ToLower(fields[1]), strings.ToLower(fields[2]), false
		}

	}
}

func createNewAnimal(theName, theType string, db map[string]Animal) error {

	if _, existing := db[theName]; existing {
		return errors.New("there is already an animal stored with the name [" + theName + "]")
	}

	switch theType {
	case "cow":
		db[theName] = Cow{theName}
	case "bird":
		db[theName] = Bird{theName}
	case "snake":
		db[theName] = Snake{theName}
	default:
		return errors.New("type of animal [" + theType + "] is not in the list of accepted types")
	}

	return nil
}

func queryAnimalAction(theName, theAction string, db map[string]Animal) (string, error) {

	animal, existing := db[theName]

	if !existing {
		return "", errors.New("there is no such animal [" + theName + "] " + "stored")
	}

	switch theAction {
	case "eat":
		return animal.Eat(), nil
	case "move":
		return animal.Move(), nil
	case "speak":
		return animal.Speak(), nil
	default:
		return "", errors.New("action [" + theAction + "] of animal is not in the list of accepted types")
	}
}
