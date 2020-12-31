package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	var floatInput float64

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a floating point number, X to quit: ")

		// Scans a line from Stdin(Console)
		scanner.Scan()

		// Holds the string that scanned
		text := scanner.Text()

		if strings.EqualFold(text, "X") {
			break
		}

		if len(text) != 0 {
			_, e := fmt.Sscan(text, &floatInput)
			if e == nil {
				fmt.Printf("%f truncated = %.0f\n", floatInput, toFixed(floatInput, 0))
			}
		}
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
