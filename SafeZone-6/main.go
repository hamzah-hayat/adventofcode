package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

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
	largestArea := FindLargestArea(input)
	fmt.Printf("The largest area is %v\n", largestArea)
}

func PartTwo() {
	//input := readInput()
}

func FindLargestArea(points []Point) int {
	largestArea := 0

	return largestArea
}

// Find out which is the closest neighbour to the gridpoint
func FindClosestNeighbour(gridPoint Point, neighbours []Point) {

}

// Figure out the Manhatten Distance between two points
func ManhattenDistance(firstPoint Point, secondPoint Point) int {

}

// Read data from input.txt
// Load it into points array
func readInput() []Point {

	var input []Point

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() != "" {
			values := strings.Split(scanner.Text(), ",")
			x, _ := strconv.Atoi(values[0])
			y, _ := strconv.Atoi(strings.Trim(values[1], " "))
			input = append(input, Point{x, y})
		}
	}
	return input
}
