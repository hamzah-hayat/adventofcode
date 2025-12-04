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

	validRolls := 0
	grid := make(map[Point]string)
	// Read grid in
	for y, row := range input {
		for x, char := range row {
			if char == '@' {
				grid[Point{x: x, y: y}] = string(char)
			}
		}
	}

	// Check each roll of paper
	for point, _ := range grid {
		if checkAroundRoll(grid,point) {
			validRolls++
		}
	}

	return strconv.Itoa(validRolls)
}

func checkAroundRoll(grid map[Point]string, point Point) bool {
	numberOfNeighbors := 0
	// Count how many neighbors are '@'
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue // Skip the center point
			}
			neighbor := Point{x: point.x + dx, y: point.y + dy}
			if val, exists := grid[neighbor]; exists && val == "@" {
				numberOfNeighbors++
			}
		}
	}
	if numberOfNeighbors < 4 {
		return true // Less than 4 neighbors are '@'

	}
	return false // More than 4 neighbors are '@', so it's not a valid roll
}

func PartTwo(filename string) string {
	input := readInput(filename)
	removedRolls := 0

	grid := make(map[Point]string)
	// Read grid in
	for y, row := range input {
		for x, char := range row {
			if char == '@' {
				grid[Point{x: x, y: y}] = string(char)
			}
		}
	}

	removedRolls=loopAndRemoveRolls(grid)

	return strconv.Itoa(removedRolls)
}

func loopAndRemoveRolls(grid map[Point]string) int {
	totalRemoved := 0
	finished := false

	for !finished {
		pointsToRemove := make([]Point, 0)
		finished = true
		for point, _ := range grid {
			if checkAroundRoll(grid, point) {
				pointsToRemove = append(pointsToRemove, point)
				totalRemoved++
				finished = false
			}
		}

		for _, p := range pointsToRemove {
			delete(grid,p)
		}

	}

	return totalRemoved
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
