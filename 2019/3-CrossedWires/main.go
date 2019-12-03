package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	grid1 := make(map[point]bool)
	currentPoint := point{x: 0, y: 0}

	// Create a map of points that we have moved accross
	for _, move := range strings.Split(input[0], ",") {

		// first character is move direction
		// second is distance
		moveDir := string(move[0])
		moveDist, _ := strconv.Atoi(move[1:])

		// Move
		moveGrid(moveDir, moveDist, grid1, &currentPoint)
	}

	grid2 := make(map[point]bool)
	currentPoint = point{x: 0, y: 0}

	// Create a map of points that we have moved accross
	for _, move := range strings.Split(input[1], ",") {

		// first character is move direction
		// second is distance
		moveDir := string(move[0])
		moveDist, _ := strconv.Atoi(move[1:])

		// Move
		moveGrid(moveDir, moveDist, grid2, &currentPoint)
	}

	// check both lists and find any same points
	matchGrid := make(map[point]bool)
	for i, _ := range grid1 {
		if grid2[i] == true {
			matchGrid[i] = true
		}
	}

	// find closest point to origin
	closest := point{x: 1000000, y: 1000000}
	origin := point{x: 0, y: 0}
	for p, _ := range matchGrid {
		if p.distanceBetween(origin) <= closest.distanceBetween(origin) {
			closest = p
		}
	}

	fmt.Println("The closest point is", closest)

}

func moveGrid(moveDir string, moveDist int, grid map[point]bool, currentPoint *point) {

	xMove := 0
	yMove := 0

	switch moveDir {
	case "R":
		xMove = 1
		break
	case "U":
		yMove = 1
		break
	case "L":
		xMove = -1
		break
	case "D":
		yMove = -1
		break
	}

	for i := 0; i < moveDist; i++ {
		// move first
		currentPoint.x += xMove
		currentPoint.y += yMove

		grid[*currentPoint] = true
	}
}

func PartTwo() {
	input := readInput()

	grid1 := make(map[point]int)
	currentPoint := point{x: 0, y: 0}

	// Create a map of points that we have moved accross
	totalSteps := 0
	for _, move := range strings.Split(input[0], ",") {

		// first character is move direction
		// second is distance
		moveDir := string(move[0])
		moveDist, _ := strconv.Atoi(move[1:])

		// Move
		moveGridCountingSteps(moveDir, moveDist, grid1, &currentPoint, &totalSteps)
	}

	grid2 := make(map[point]int)
	currentPoint = point{x: 0, y: 0}

	totalSteps = 0
	// Create a map of points that we have moved accross
	for _, move := range strings.Split(input[1], ",") {

		// first character is move direction
		// second is distance
		moveDir := string(move[0])
		moveDist, _ := strconv.Atoi(move[1:])

		// Move
		moveGridCountingSteps(moveDir, moveDist, grid2, &currentPoint, &totalSteps)
	}

	// check both lists and find any same points
	matchGrid := make(map[point]int)
	for i, _ := range grid1 {
		if grid2[i] > 0 {
			matchGrid[i] = grid1[i] + grid2[i]
		}
	}

	// find closest point to origin
	fmt.Println(matchGrid)
	closest := 100000
	var closestPoint point
	var steps int
	for p, v := range matchGrid {
		if v <= closest {
			closest = v
			closestPoint = p
			steps = v
		}
	}

	fmt.Println("The closest point is", closestPoint, " with steps", steps)

}

func moveGridCountingSteps(moveDir string, moveDist int, grid map[point]int, currentPoint *point, totalSteps *int) {

	xMove := 0
	yMove := 0

	switch moveDir {
	case "R":
		xMove = 1
		break
	case "U":
		yMove = 1
		break
	case "L":
		xMove = -1
		break
	case "D":
		yMove = -1
		break
	}

	for i := 0; i < moveDist; i++ {
		*totalSteps++
		// move first
		currentPoint.x += xMove
		currentPoint.y += yMove

		grid[*currentPoint] = *totalSteps
	}
}

type point struct {
	x int
	y int
}

func (p point) toString() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

func (p point) distanceBetween(p2 point) int {
	return p.x - p2.x + p.y - p2.y
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
