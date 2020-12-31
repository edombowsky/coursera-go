package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a sequence of up to 10 integers, X to quit: ")

		// Scans a line from Stdin(Console)
		scanner.Scan()

		// Holds the string that scanned
		text := scanner.Text()

		if strings.EqualFold(text, "X") {
			break
		}

		if len(text) != 0 {
			v := strings.Fields(text)
			if len(v) > 10 {
				v = v[0:10]
			}

			values := []int{}
			for i := range v {
				number := v[i]
				num, _ := strconv.Atoi(number)
				values = append(values, num)
			}

			fmt.Println("\n--- Unsorted --- \n\n", values)
			BubbleSort(values)
			fmt.Println("\n--- Sorted ---\n\n", values)
		}
	}
}

// BubbleSort will sort an integer slice
func BubbleSort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				Swap(items, i)
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}

// Swap swaps the position of two adjacent elements in a slice
func Swap(slice []int, index int) {
	temp := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = temp
}
