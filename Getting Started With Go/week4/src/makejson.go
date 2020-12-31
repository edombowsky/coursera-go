package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	name, quit := getName()
	if quit {
		return
	}
	address, quit := getAddress()
	if quit {
		return
	}

	nameAddress := make(map[string]string)

	nameAddress["name"] = name
	nameAddress["address"] = address

	json, err := json.MarshalIndent(nameAddress, "", "   ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(json))
	}
}

func getName() (name string, quit bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a name, X to quit: ")

		// Scans a line from Stdin(Console)
		scanner.Scan()

		// Holds the string that scanned
		text := scanner.Text()

		if strings.EqualFold(text, "X") {
			return "", true
		} else if len(text) > 0 {
			return text, false
		}
	}
}

func getAddress() (address string, quit bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter an address, X to quit: ")

		// Scans a line from Stdin(Console)
		scanner.Scan()

		// Holds the string that scanned
		text := scanner.Text()

		if strings.EqualFold(text, "X") {
			return "", true
		} else if len(text) > 0 {
			return text, false
		}
	}
}
