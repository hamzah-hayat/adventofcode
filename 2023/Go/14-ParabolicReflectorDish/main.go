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

type Point struct {
	x int
	y int
}

func PartOne(filename string) string {
	input := readInput(filename)

	for x := 0; x < len(input); x++ {
		for y := 0; y < len(input[0]); y++ {
			if input[x][y] == 'O' {
				moveTo := MoveNorth(input, x, y)
				if moveTo.x != x {
					// Delete old rock
					input[x] = input[x][:y] + string('.') + input[x][y+1:]
					// Add new rock
					input[moveTo.x] = input[moveTo.x][:y] + string('O') + input[moveTo.x][y+1:]
				}
			}
		}
	}

	// Score the rocks
	result := 0
	for x := 0; x < len(input); x++ {
		for y := 0; y < len(input[0]); y++ {
			if input[x][y] == 'O' {
				result += len(input) - x
			}
		}
	}

	num := strconv.Itoa(result)

	return num
}

// Move the Rock at x,y as far as possible North
func MoveNorth(input []string, x, y int) Point {

	if x-1 < 0 {
		return Point{x, y}
	}

	if input[x-1][y] == '.' {
		return MoveNorth(input, x-1, y)
	} else {
		return Point{x, y}
	}
}

// Move the Rock at x,y as far as possible East
func MoveEast(input []string, x, y int) Point {

	if y+1 >= len(input[0]) {
		return Point{x, y}
	}

	if input[x][y+1] == '.' {
		return MoveEast(input, x, y+1)
	} else {
		return Point{x, y}
	}
}

// Move the Rock at x,y as far as possible South
func MoveSouth(input []string, x, y int) Point {

	if x+1 >= len(input) {
		return Point{x, y}
	}

	if input[x+1][y] == '.' {
		return MoveSouth(input, x+1, y)
	} else {
		return Point{x, y}
	}
}

// Move the Rock at x,y as far as possible West
func MoveWest(input []string, x, y int) Point {

	if y-1 < 0 {
		return Point{x, y}
	}

	if input[x][y-1] == '.' {
		return MoveWest(input, x, y-1)
	} else {
		return Point{x, y}
	}
}

func findRocks(input []string) []Point {
	// Find E point in input
	var points []Point
	for y, line := range input {
		for x, v := range line {
			if string(v) == "O" {
				points = append(points, Point{x, y})
			}
		}
	}
	return points
}

func PartTwo(filename string) string {
	input := readInput(filename)
	inputCopy := make([]string, len(input))
	copy(inputCopy, input)

	// Find each possible pattern
	var patterns [][]string
	endX := 0
	cycleLegnth := 0
	c := make([]string, len(input))
	copy(c, input)
	patterns = append(patterns, c)

	for x := 0; x < 1000000000; x++ {
		input = MoveAllRocksInDirectin(input, 0)
		input = MoveAllRocksInDirectin(input, 3)
		input = MoveAllRocksInDirectin(input, 2)
		input = MoveAllRocksInDirectin(input, 1)

		for i, v := range patterns {
			if areEqual(v, input) {
				cycleLegnth = x - i + 1
				endX = x + 1
				break
			}
		}
		if cycleLegnth != 0 {
			break
		}

		c2 := make([]string, len(input))
		copy(c2, input)
		patterns = append(patterns, c2)
	}

	// Now we know that the rocks cycle after cycleLegnth times
	// So we can shorten the work we do
	stepsRemaining := (1000000000 - endX) % cycleLegnth
	for x := 0; x < stepsRemaining; x++ {
		input = MoveAllRocksInDirectin(input, 0)
		input = MoveAllRocksInDirectin(input, 3)
		input = MoveAllRocksInDirectin(input, 2)
		input = MoveAllRocksInDirectin(input, 1)
	}

	// Score the rocks
	result := 0
	for x := 0; x < len(input); x++ {
		for y := 0; y < len(input[0]); y++ {
			if input[x][y] == 'O' {
				result += len(input) - x
			}
		}
	}

	num := strconv.Itoa(result)

	return num
}

func areEqual(v []string, input []string) bool {
	equal := true
	for i := 0; i < len(v); i++ {
		if v[i] != input[i] {
			equal = false
		}
	}
	return equal
}

// Move in a specific direction
func MoveAllRocksInDirectin(input []string, direction int) []string {

	// 0 is North
	// 1 is East
	// 2 is South
	// 3 is West

	switch direction {
	case 0:
		for x := 0; x < len(input); x++ {
			for y := 0; y < len(input[0]); y++ {
				if input[x][y] == 'O' {
					moveTo := MoveNorth(input, x, y)
					if moveTo.x != x {
						// Delete old rock
						input[x] = input[x][:y] + string('.') + input[x][y+1:]
						// Add new rock
						input[moveTo.x] = input[moveTo.x][:y] + string('O') + input[moveTo.x][y+1:]
					}
				}
			}
		}
	case 1:
		for x := 0; x < len(input); x++ {
			for y := len(input[x]) - 1; y >= 0; y-- {
				if input[x][y] == 'O' {
					moveTo := MoveEast(input, x, y)
					if moveTo.y != y {
						// Delete old rock
						input[x] = input[x][:y] + string('.') + input[x][y+1:]
						// Add new rock
						input[x] = input[moveTo.x][:moveTo.y] + string('O') + input[x][moveTo.y+1:]
					}
				}
			}
		}
	case 2:
		for x := len(input) - 1; x >= 0; x-- {
			for y := 0; y < len(input[0]); y++ {
				if input[x][y] == 'O' {
					moveTo := MoveSouth(input, x, y)
					if moveTo.x != x {
						// Delete old rock
						input[x] = input[x][:y] + string('.') + input[x][y+1:]
						// Add new rock
						input[moveTo.x] = input[moveTo.x][:y] + string('O') + input[moveTo.x][y+1:]
					}
				}
			}
		}
	case 3:
		for x := 0; x < len(input); x++ {
			for y := 0; y < len(input[0]); y++ {
				if input[x][y] == 'O' {
					moveTo := MoveWest(input, x, y)
					if moveTo.y != y {
						// Delete old rock
						input[x] = input[x][:y] + string('.') + input[x][y+1:]
						// Add new rock
						input[x] = input[moveTo.x][:moveTo.y] + string('O') + input[moveTo.x][moveTo.y+1:]
					}
				}
			}
		}
	}

	return input
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
