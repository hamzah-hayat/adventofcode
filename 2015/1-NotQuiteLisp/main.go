package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p1", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()

	switch *methodP {
	case "p1":
		PartOne()
		break
	case "p2":
		PartTwo()
		break
	case "test":
		break
	}
}

func PartOne() {
	input := readInput()
	fmt.Println(input)
}

func PartTwo() {
	//input := readInput()
}

// Read data from input.txt
// Return the string, so that we can deal with it however
func readInput() string {

	var input string

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() != "" {
			input = scanner.Text()
		}
	}
	return input
}
