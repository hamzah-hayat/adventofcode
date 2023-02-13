package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	sandAtRest := 0

	rockGrid := MakeRockGrid(input)

	fmt.Println(rockGrid)

	return strconv.Itoa(sandAtRest)
}

func PartTwo(filename string) string {
	input := readInput(filename)

	return input[0]
}

func MakeRockGrid(input []string) map[Point]bool {
	rockGrid := make(map[Point]bool)

	// 498,4 -> 498,6 -> 496,6
	// 503,4 -> 502,4 -> 502,9 -> 494,9
	for _, line := range input {
		pointsStr := strings.Split(line, " -> ")

		// the first point is the start of the rock
		firstRockSplit := strings.Split(pointsStr[0], ",")
		x, _ := strconv.Atoi(firstRockSplit[0])
		y, _ := strconv.Atoi(firstRockSplit[1])
		startRockPoint := Point{x: x, y: y}

		rockGrid[startRockPoint] = true

		for i := 1; i < len(pointsStr)-1; i++ {

		}

	}

	return rockGrid
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

type Point struct {
	x, y int
}
