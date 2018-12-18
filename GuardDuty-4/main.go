package main

import (
	"bufio"
	"os"
)

func main() {
	//PartOne()
	//PartTwo()
}

func PartOne() {
	//input := readInput()

}

func PartTwo() {
	//input := readInput()
}

// Read data from input.txt
// Load it into string array
func readInput() []string {

	var input []string

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() != "" {
			line := scanner.Text()
			input = append(input, line)
		}
	}
	return input
}
