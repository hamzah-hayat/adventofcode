package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type shift struct {
	guardID string // The ID of the guard
	sleep   []bool // For each minute of the shift, was the guard asleep or not
}

func main() {
	PartOne()
	//PartTwo()
}

func PartOne() {
	input := readInput()
	for _, line := range input {
		fmt.Println(line)
	}

}

func PartTwo() {
	//input := readInput()
}

// Read data from input.txt
// Load it into string array
func readInput() []shift {

	var input []string

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() != "" {
			line := scanner.Text()
			input = append(input, line)
		}
	}

	// Now have all the input, sort it by date/time
	sort.Strings(input)

	// Now we have to build a "shift" for each day
	var shifts []shift
	i := 0
	for i < len(input) {
		newShift, nexti := BuildShift(i, input)
		i = nexti
		shifts = append(shifts, newShift)
	}

	return shifts
}

func BuildShift(i int, input []string) (shift, int) {
	// Build a shift using the input
	// Return the new shift and the next i
	var newShift shift

	newShift.guardID = strings.TrimLeft(strings.Fields(input[i])[3], "#")
	i++

	for {
		// We always assume we start with a new Guard Line

		break
	}

	return newShift, i
}
