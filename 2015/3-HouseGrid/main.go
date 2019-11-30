package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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

	// Traverse while building a map
	grid := make(map[point]int)

	currentPoint := point{x: 0, y: 0}
	grid[currentPoint] = 1

	for _, move := range input {

		switch string(move) {
		case ">":
			currentPoint.x = currentPoint.x + 1
			break
		case "<":
			currentPoint.x = currentPoint.x - 1
			break
		case "^":
			currentPoint.y = currentPoint.y + 1
			break
		case "v":
			currentPoint.y = currentPoint.y - 1
			break
		}
		grid[currentPoint] = 1
	}

	// Finish setting up Grid
	count := len(grid)

	fmt.Println("The number of houses with at least one present is", count)

}

// Robo Santa takes every even move, while normal santa takes the odd moves
func PartTwo() {
	input := readInput()

	// Traverse while building a map
	gridNormalSanta := make(map[point]int)
	gridRoboSanta := make(map[point]int)

	cpNormalSanta := point{x: 0, y: 0}
	cpRoboSanta := point{x: 0, y: 0}
	gridNormalSanta[cpNormalSanta] = 1
	gridRoboSanta[cpRoboSanta] = 1

	for i, move := range input {

		if i%2 == 0 {
			// Normal Santa
			switch string(move) {
			case ">":
				cpNormalSanta.x = cpNormalSanta.x + 1
				break
			case "<":
				cpNormalSanta.x = cpNormalSanta.x - 1
				break
			case "^":
				cpNormalSanta.y = cpNormalSanta.y + 1
				break
			case "v":
				cpNormalSanta.y = cpNormalSanta.y - 1
				break
			}
			gridNormalSanta[cpNormalSanta] = 1

		} else {
			// Robo Santa
			switch string(move) {
			case ">":
				cpRoboSanta.x = cpRoboSanta.x + 1
				break
			case "<":
				cpRoboSanta.x = cpRoboSanta.x - 1
				break
			case "^":
				cpRoboSanta.y = cpRoboSanta.y + 1
				break
			case "v":
				cpRoboSanta.y = cpRoboSanta.y - 1
				break
			}
			gridRoboSanta[cpRoboSanta] = 1
		}
	}

	// Finish setting up Grids
	// Now need to merge
	for i, p := range gridRoboSanta {
		gridNormalSanta[i] = p
	}

	count := len(gridNormalSanta)

	fmt.Println("The number of houses with at least one present is", count)
}

type point struct {
	x int
	y int
}

func (p point) toString() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

// Read data from input.txt
// Return the string, so that we can deal with it however
func readInput() string {

	var input string

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}
	return input
}
