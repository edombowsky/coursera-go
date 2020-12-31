package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Name Struct
type Name struct {
	fname string
	lname string
}

func main() {
	fmt.Print("Enter the name of data file: ")
	var filename string
	fmt.Scanln(&filename)
	filename = strings.Trim(filename, "n")

	buf, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	snl := bufio.NewScanner(buf)
	names := make([]Name, 0, 3)

	for snl.Scan() {
		v := strings.Fields(snl.Text())
		if len(v[0]) > 20 {
			v[0] = firstTwentyChars(v[0])
		}
		if len(v[1]) > 20 {
			v[1] = firstTwentyChars(v[1])
		}

		var aName Name
		aName.fname, aName.lname = v[0], v[1]
		names = append(names, aName)
	}

	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range names {
		fmt.Println(n.fname, n.lname)
	}
}

func firstTwentyChars(s string) string {
	runes := []rune(s)
	return string(runes[0:20])
}
