package main

import (
	"fmt"
	"time"
)

var a int

func main() {
	a = 5
	go inc()
	go dec()
	time.Sleep(8 * time.Second)
}

func inc() {
	for i := 0; i < 10; i++ {
		a = a + 1
		fmt.Println("inc: ", a, "index: ", i)
		time.Sleep(2 * time.Second)
	}
}

func dec() {
	for i := 0; i < 10; i++ {
		a = a - 1
		fmt.Println("dec: ", a, "index: ", i)
		time.Sleep(1 * time.Second)
	}
}
