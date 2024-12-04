package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	methodP *string
)

func parseFlags() {
	methodP = flag.String("method", "all", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()
}

func main() {

	parseFlags()

	switch *methodP {
	case "all":
		fmt.Println("Silver:" + PartOne("input"))
		fmt.Println("Gold:" + PartTwo("input"))
	case "p1":
		fmt.Println("Silver:" + PartOne("input"))
	case "p2":
		fmt.Println("Gold:" + PartTwo("input"))
	}
}

func PartOne(filename string) string {
	input := readInput(filename)

	totalMatches := 0

	for startY, line := range input {
		// We always have to start at an X
		for startX, c := range line {
			if string(c) == "X" {
				// Look for rest of XMAS
				for directionY := -1; directionY <= 1; directionY++ {
					for directionX := -1; directionX <= 1; directionX++ {
						totalMatches += SearchForXMAS(startX, startY, directionX, directionY, input, "M")
					}
				}
			}
		}
	}

	return strconv.Itoa(totalMatches)
}

// Recursive search
func SearchForXMAS(startX, startY, directionX, directionY int, grid []string, letter string) int {
	totalMatches := 0

	if startY+directionY < 0 || startY+directionY >= len(grid) {
		return totalMatches
	}
	if startX+directionX < 0 || startX+directionX >= len(grid[0]) {
		return totalMatches
	}

	if directionX == 0 && directionY == 0 {
		return totalMatches
	}

	if string(grid[startY+directionY][startX+directionX]) == letter {
		if letter == "M" {
			totalMatches += SearchForXMAS(startX+directionX, startY+directionY, directionX, directionY, grid, "A")
		}
		if letter == "A" {
			totalMatches += SearchForXMAS(startX+directionX, startY+directionY, directionX, directionY, grid, "S")
		}
		if letter == "S" {
			return 1
		}
	}

	return totalMatches
}

func PartTwo(filename string) string {
	input := readInput(filename)

	totalMatches := 0

	AList := make(map[Point]int)

	for startY, line := range input {
		// We always have to start at an M
		for startX, c := range line {
			if string(c) == "M" {
				// Look for rest of XMAS
				for directionY := -1; directionY <= 1; directionY++ {
					for directionX := -1; directionX <= 1; directionX++ {

						// Can only be diagnonal, so ignore vertical/horizontal
						if directionY == 0 {
							continue
						}
						if directionX == 0 {
							continue
						}

						if SearchForXMAS(startX, startY, directionX, directionY, input, "A") == 1 {
							// Add this A to our list
							AList[Point{startX + directionX, startY + directionY}]++
						}
					}
				}
			}
		}
	}

	for _, v := range AList {
		if v == 2 {
			totalMatches++
		}
	}

	return strconv.Itoa(totalMatches)
}

type Point struct {
	x int
	y int
}

// Read data from input.txt
// Return the string, so that we can deal with it however
func readInput(filename string) []string {

	var input []string

	f, _ := os.Open(filename + ".txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

// Read data from input.txt
// Return the string as int
func readInputInt() []int {

	var input []int

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		input = append(input, num)
	}
	return input
}
