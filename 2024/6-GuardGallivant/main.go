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

	grid := make(map[Point]string)
	markedGrid := make(map[Point]string)

	for y, line := range input {
		for x, char := range line {
			grid[Point{x, y}] = string(char)
		}
	}

	maxX := 0
	maxY := 0
	for p, _ := range grid {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	for {
		MoveGuard(grid, markedGrid, maxX, maxY)
		if GuardLeftGrid(grid) {
			break
		}
	}

	total := CountSquares(markedGrid)

	return strconv.Itoa(total)
}

// MoveGuard moves the guard a single step, or rotates 90 degrees clockwise if unable to move
func MoveGuard(grid map[Point]string, markedGrid map[Point]string, maxX, maxY int) {
	// Find Guard
	currentX := 0
	currentY := 0
	guardDirection := ""
	for p, v := range grid {
		if v == "^" || v == ">" || v == "v" || v == "<" {
			currentX = p.x
			currentY = p.y
			guardDirection = v
		}
	}

	// Move
	switch guardDirection {
	case "^":
		// OOB check
		for {
			// Mark current position
			markedGrid[Point{currentX, currentY}] = "X"

			if currentY-1 < 0 {
				grid[Point{currentX, currentY}] = "."
				break
			}
			// Obstacle check
			// Then basic movement
			if grid[Point{currentX, currentY - 1}] == "#" {
				grid[Point{currentX, currentY}] = ">"
				break
			} else {
				grid[Point{currentX, currentY}] = "."
				grid[Point{currentX, currentY - 1}] = "^"
				currentY--
			}
		}
	case ">":
		for {
			// Mark current position
			markedGrid[Point{currentX, currentY}] = "X"
			if currentX+1 > maxX {
				grid[Point{currentX, currentY}] = "."
				break
			}

			if grid[Point{currentX + 1, currentY}] == "#" {
				grid[Point{currentX, currentY}] = "v"
				break
			} else {
				grid[Point{currentX, currentY}] = "."
				grid[Point{currentX + 1, currentY}] = ">"
				currentX++
			}
		}
	case "v":
		for {
			// Mark current position
			markedGrid[Point{currentX, currentY}] = "X"
			if currentY+1 > maxY {
				grid[Point{currentX, currentY}] = "."
				break
			}

			if grid[Point{currentX, currentY + 1}] == "#" {
				grid[Point{currentX, currentY}] = "<"
				break
			} else {
				grid[Point{currentX, currentY}] = "."
				grid[Point{currentX, currentY + 1}] = "v"
				currentY++
			}
		}
	case "<":
		for {
			// Mark current position
			markedGrid[Point{currentX, currentY}] = "X"
			if currentX-1 < 0 {
				grid[Point{currentX, currentY}] = "."
				break
			}

			if grid[Point{currentX - 1, currentY}] == "#" {
				grid[Point{currentX, currentY}] = "^"
				break
			} else {
				grid[Point{currentX, currentY}] = "."
				grid[Point{currentX - 1, currentY}] = "<"
				currentX--
			}
		}
	}
}

func GuardLeftGrid(grid map[Point]string) bool {
	for _, v := range grid {
		if v == "^" || v == ">" || v == "v" || v == "<" {
			return false
		}
	}
	return true
}

func CountSquares(grid map[Point]string) int {
	total := 0
	for _, v := range grid {
		if v == "X" {
			total++
		}
	}
	return total
}

type Point struct {
	x int
	y int
}

func PartTwo(filename string) string {
	input := readInput(filename)

	totalLoops := 0
	grid := make(map[Point]string)
	// Setup Grid
	for y, line := range input {
		for x, char := range line {
			grid[Point{x, y}] = string(char)
		}
	}

	maxX := 0
	maxY := 0
	for p, _ := range grid {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	// Do inital run to get movementGrid
	// These are all the possible points we can make an obstacle on
	movementGrid := make(map[Point]string)
	for {
		MoveGuard(grid, movementGrid, maxX, maxY)
		if GuardLeftGrid(grid) {
			break
		}
	}

	pointsToCheck := make([]Point, 0)
	for p, v := range movementGrid {
		if v == "X" {
			pointsToCheck = append(pointsToCheck, p)
		}
	}

	for _, extraObstacle := range pointsToCheck {
		guardSet := make(map[Guard]int)

		for p := range grid {
			delete(grid, p)
		}

		// Setup Grid
		for y, line := range input {
			for x, char := range line {
				grid[Point{x, y}] = string(char)
			}
		}
		grid[extraObstacle] = "#"

		// Run loop until either Guard leaves or we find a loop
		for {
			MoveGuardNumberMark(grid, guardSet, maxX, maxY)
			if GuardLeftGrid(grid) {
				break
			}
			if GuardInLoop(guardSet) {
				totalLoops++
				break
			}
		}
	}

	return strconv.Itoa(totalLoops)
}

func MoveGuardNumberMark(grid map[Point]string, guardSet map[Guard]int, maxX, maxY int) {
	// Find Guard
	currentX := 0
	currentY := 0
	guardDirection := ""
	for p, v := range grid {
		if v == "^" || v == ">" || v == "v" || v == "<" {
			currentX = p.x
			currentY = p.y
			guardDirection = v
		}
	}

	// Mark current position
	_, exists := guardSet[Guard{currentX, currentY, guardDirection}]
	if !exists {
		guardSet[Guard{currentX, currentY, guardDirection}] = 1
	} else {
		guardSet[Guard{currentX, currentY, guardDirection}] = 2
		return // We've already been here
	}

	// Move
	switch guardDirection {
	case "^":
		// OOB check
		for {

			if currentY-1 < 0 {
				grid[Point{currentX, currentY}] = "."
				break
			}
			// Obstacle check
			// Then basic movement
			if grid[Point{currentX, currentY - 1}] == "#" {
				grid[Point{currentX, currentY}] = ">"
				break
			} else {
				grid[Point{currentX, currentY}] = "."
				grid[Point{currentX, currentY - 1}] = "^"
				currentY--
			}
		}
	case ">":
		for {
			if currentX+1 > maxX {
				grid[Point{currentX, currentY}] = "."
				break
			}

			if grid[Point{currentX + 1, currentY}] == "#" {
				grid[Point{currentX, currentY}] = "v"
				break
			} else {
				grid[Point{currentX, currentY}] = "."
				grid[Point{currentX + 1, currentY}] = ">"
				currentX++
			}
		}
	case "v":
		for {
			if currentY+1 > maxY {
				grid[Point{currentX, currentY}] = "."
				break
			}

			if grid[Point{currentX, currentY + 1}] == "#" {
				grid[Point{currentX, currentY}] = "<"
				break
			} else {
				grid[Point{currentX, currentY}] = "."
				grid[Point{currentX, currentY + 1}] = "v"
				currentY++
			}
		}
	case "<":
		for {
			if currentX-1 < 0 {
				grid[Point{currentX, currentY}] = "."
				break
			}

			if grid[Point{currentX - 1, currentY}] == "#" {
				grid[Point{currentX, currentY}] = "^"
				break
			} else {
				grid[Point{currentX, currentY}] = "."
				grid[Point{currentX - 1, currentY}] = "<"
				currentX--
			}
		}
	}
}

func GuardInLoop(guardSet map[Guard]int) bool {
	for _, v := range guardSet {
		if v == 2 {
			return true
		}
	}
	return false
}

type Guard struct {
	x         int
	y         int
	direction string
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
