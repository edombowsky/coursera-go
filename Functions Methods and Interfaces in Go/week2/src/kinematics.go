package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	for {
		var acceleration float64
		var initialVelocity float64
		var initialDisplacement float64
		var time float64

		fmt.Print("Enter acceleration, X to quit: ")
		if _, err := fmt.Scan(&acceleration); err != nil {
			log.Print("Scan for acceleration failed, due to ", err)
			return
		}
		fmt.Print("Enter initial velocity, X to quit: ")
		if _, err := fmt.Scan(&initialVelocity); err != nil {
			log.Print("Scan for initialVelocity failed, due to ", err)
			return
		}
		fmt.Print("Enter initial displacement, X to quit: ")
		if _, err := fmt.Scan(&initialDisplacement); err != nil {
			log.Print("Scan for initialDisplacement failed, due to ", err)
			return
		}
		fmt.Print("Enter time, X to quit: ")
		if _, err := fmt.Scan(&time); err != nil {
			log.Print("Scan for time failed, due to ", err)
			return
		}

		fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
		fmt.Printf("The resulting displacement value s(t) is [%f]\n", fn(time))
	}
}

// GenDisplaceFn returns a function which computes displacement as a function of time
func GenDisplaceFn(acceleration float64,
	initialVelocity float64,
	initialDisplacement float64) func(float64) float64 {

	return func(time float64) float64 {
		return (0.5*acceleration*time*time + initialVelocity*time + initialDisplacement)
	}
}

func readFloatFromStdin(promptText string) float64 {

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(promptText)
		in.Scan()
		aFloat64, err := strconv.ParseFloat(in.Text(), 64)

		if err != nil {
			fmt.Println("Not a parseable float value, please try again.")
		} else {
			return aFloat64
		}
	}
}
