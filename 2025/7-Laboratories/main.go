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
	numberOfSplits := 0

	grid := make(map[Point]string)

	y := 0
	for _, lines := range input {
		for x := 0; x < len(lines); x++ {
			grid[Point{x, y}] = string(lines[x])
		}
		y++
	}

	// Now we have grid, process each line and find splits
	// Ignore the first row, look above at last row
	y = 1
	for i := 0; i < len(input); i++ {
		for x := 0; x < len(input[i]); x++ {
			switch grid[Point{x, y - 1}] {
			case "S":
				grid[Point{x, y}] = "|"
			case "|":
				switch grid[Point{x, y}] {
				case ".":
					grid[Point{x, y}] = "|"
				case "^":
					numberOfSplits++
					grid[Point{x + 1, y}] = "|"
					grid[Point{x - 1, y}] = "|"
				}
			}
		}
		y++
	}

	return strconv.Itoa(numberOfSplits)
}

func PartTwo(filename string) string {
	input := readInput(filename)

	grid := make(map[Point]string)

	y := 0
	for _, lines := range input {
		for x := 0; x < len(lines); x++ {
			grid[Point{x, y}] = string(lines[x])
		}
		y++
	}

	// Now we have grid, process each line and find splits
	// Ignore the first row, look above at last row
	totalChoicesSum := 0
	alreadySeen := make(map[Point]int)
	for p, v := range grid {
		if v == "S" {
			totalChoicesSum = recursiveSplit(grid, alreadySeen, Point{p.x, p.y + 1}, len(input))
		}
	}

	return strconv.Itoa(totalChoicesSum)
}

func recursiveSplit(grid map[Point]string, alreadySeen map[Point]int, p Point, maxY int) int {

	if alreadySeenRes, exists := alreadySeen[p]; exists {
		return alreadySeenRes
	}

	if p.y == maxY {
		return 1
	}

	if grid[p] == "^" {
		return recursiveSplit(grid, alreadySeen, Point{p.x + 1, p.y + 1}, maxY) + recursiveSplit(grid, alreadySeen, Point{p.x - 1, p.y + 1}, maxY)
	}

	result := recursiveSplit(grid, alreadySeen, Point{p.x, p.y + 1}, maxY)
	alreadySeen[Point{p.x, p.y + 1}] = result

	return result

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
