package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readInput()
	fmt.Println(input)
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
