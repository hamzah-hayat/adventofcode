package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	PartOne()
	//PartTwo()
}

func PartOne() {
	input := readInput()
	fmt.Println(input)
}

func PartTwo() {
	//input := readInput()
}

type claim struct {
	id     int
	left   int
	top    int
	width  int
	height int
}

// Read data from input.txt
// Load it into claim array
func readInput() []claim {

	var input []claim

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() != "" {
			line := scanner.Text()

			fields := strings.Fields(line)

			var newClaim claim
			newClaim.id, _ = strconv.Atoi(strings.TrimLeft(fields[0], "#"))
			newClaim.left, _ = strconv.Atoi(strings.Split(strings.TrimRight(fields[2], ":"), ",")[0])
			newClaim.top, _ = strconv.Atoi(strings.Split(strings.TrimRight(fields[2], ":"), ",")[1])
			newClaim.width, _ = strconv.Atoi(strings.Split(fields[3], "x")[0])
			newClaim.height, _ = strconv.Atoi(strings.Split(fields[3], "x")[1])

			input = append(input, newClaim)
		}
	}
	return input
}

// Find the amount of overlap between all the claims
func FindOverlap(input []claim) int {
	overlap := 0

	return overlap
}
