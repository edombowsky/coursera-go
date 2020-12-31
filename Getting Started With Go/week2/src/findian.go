package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a string, X to quit: ")

		// Scans a line from Stdin(Console)
		scanner.Scan()

		// Holds the string that scanned
		text := scanner.Text()

		if strings.EqualFold(text, "X") {
			break
		}

		if len(text) != 0 {
			if findian(text) {
				fmt.Println("Found")
			} else {
				fmt.Println("Not Found")
			}
		}
	}
}

func findian(str string) bool {

	match, _ := regexp.MatchString(`^[iI].*[aA].*[nN]$`, str)

	return match
}
