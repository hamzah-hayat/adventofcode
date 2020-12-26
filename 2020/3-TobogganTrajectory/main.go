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

	currentX := 0

	treesHit := 0

	for _, v := range input {

		if v[currentX] == '#' {
			treesHit++
		}

		if currentX+3 >= len(v) {
			currentX = currentX + 3 - len(v)
		} else {
			currentX = currentX + 3
		}
	}

	fmt.Println("The number of trees we hit is: ", treesHit)
}

func PartTwo() {
	first := checkSlope(1, 1)
	second := checkSlope(3, 1)
	third := checkSlope(5, 1)
	fourth := checkSlope(7, 1)
	fifth := checkSlope(1, 2)

	fmt.Println("The number of trees hit for all slopes multipled together is: ", first*second*third*fourth*fifth)
}

func checkSlope(right, down int) int {

	slope := readInput()

	currentX := 0

	treesHit := 0

	endSkip := 0
	skipping := false

	for currentY, v := range slope {

		// Down movement
		if down != 1 {
			// Check if we need to skip loops
			if skipping {
				if currentY == endSkip {
					skipping = false
				}
				continue
			} else {
				skipping = true
				endSkip = currentY + down - 1
			}
		}

		// Right movement
		if v[currentX] == '#' {
			treesHit++
		}

		if currentX+right >= len(v) {
			currentX = currentX + right - len(v)
		} else {
			currentX = currentX + right
		}

	}

	return treesHit
}

// Read data from input.txt
// Return the string, so that we can deal with it however
func readInput() []string {

	var input []string

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
