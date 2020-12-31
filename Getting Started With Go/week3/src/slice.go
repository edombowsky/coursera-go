package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var intArray = []int{}
	var intInput int

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter an integer, X to quit: ")

		// Scans a line from Stdin(Console)
		scanner.Scan()

		// Holds the string that scanned
		text := scanner.Text()

		if strings.EqualFold(text, "X") {
			break
		}

		if len(text) != 0 {
			_, err := fmt.Sscan(text, &intInput)
			if err == nil {
				// Allows duplicates
				intArray = append(intArray, intInput)
				sort.Ints(intArray)
				fmt.Printf("After insert: %v\n", intArray)
			}
		}
	}

	if len(intArray) > 0 {
		fmt.Println()
		fmt.Printf("Final sorted slice: %v\n", intArray)
	}
}

// InsertIntoSortedSlice insert an integer to a sorted slice
func InsertIntoSortedSlice(ss []int, s int) []int {
	i := sort.SearchInts(ss, s)
	ss = append(ss, 0)
	copy(ss[i+1:], ss[i:])
	ss[i] = s
	return ss
}

// InsertIntoSortedStringSlice insert string into sorted slice
func InsertIntoSortedStringSlice(ss []string, s string) []string {
	i := sort.SearchStrings(ss, s)
	ss = append(ss, "")
	copy(ss[i+1:], ss[i:])
	ss[i] = s
	return ss
}
