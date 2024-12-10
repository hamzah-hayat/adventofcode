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

	topMap := make(map[Point]int)
	// Setup Map
	for y, line := range input {
		for x, char := range line {
			num, _ := strconv.Atoi(string(char))
			topMap[Point{x, y}] = num
		}
	}

	startSpots := FindSpotsWithNum(0, topMap)
	totalTrailHeads := 0
	maxX := 0
	maxY := 0
	for p, _ := range topMap {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	// DFS
	for s := range startSpots {
		reachableEndSpots := make(map[Point]int)
		totalTrailHeadSpots := FindTotalTrailHeads(s, 0, maxX, maxY, reachableEndSpots, topMap)
		totalTrailHeads += len(totalTrailHeadSpots)
	}

	return strconv.Itoa(totalTrailHeads)
}

func FindTotalTrailHeads(s Point, currentNum, maxX, maxY int, reachableEndSpots, topMap map[Point]int) map[Point]int {

	if currentNum == 9 {
		// We've found a trailhead
		reachableEndSpots[s] = currentNum
		return reachableEndSpots
	}

	// Check each direction with recursion
	north := Point{s.x, s.y - 1}
	east := Point{s.x + 1, s.y}
	south := Point{s.x, s.y + 1}
	west := Point{s.x - 1, s.y}

	if topMap[north] == currentNum+1 && CheckBounds(north, maxX, maxY) {
		FindTotalTrailHeads(Point{s.x, s.y - 1}, currentNum+1, maxX, maxY, reachableEndSpots, topMap)
	}
	if topMap[east] == currentNum+1 && CheckBounds(east, maxX, maxY) {
		FindTotalTrailHeads(Point{s.x + 1, s.y}, currentNum+1, maxX, maxY, reachableEndSpots, topMap)
	}
	if topMap[south] == currentNum+1 && CheckBounds(south, maxX, maxY) {
		FindTotalTrailHeads(Point{s.x, s.y + 1}, currentNum+1, maxX, maxY, reachableEndSpots, topMap)
	}
	if topMap[west] == currentNum+1 && CheckBounds(west, maxX, maxY) {
		FindTotalTrailHeads(Point{s.x - 1, s.y}, currentNum+1, maxX, maxY, reachableEndSpots, topMap)
	}

	return reachableEndSpots
}

type Point struct {
	x int
	y int
}

func CheckBounds(point Point, xMax, yMax int) bool {

	if point.x < 0 || point.x > xMax {
		return false
	}
	if point.y < 0 || point.y > yMax {
		return false
	}
	return true
}

func FindSpotsWithNum(num int, topMap map[Point]int) map[Point]int {
	spotMap := make(map[Point]int)

	for s, n := range topMap {
		if num == n {
			spotMap[s] = n
		}
	}

	return spotMap
}

func PartTwo(filename string) string {
	input := readInput(filename)

	topMap := make(map[Point]int)
	// Setup Map
	for y, line := range input {
		for x, char := range line {
			num, _ := strconv.Atoi(string(char))
			topMap[Point{x, y}] = num
		}
	}

	startSpots := FindSpotsWithNum(0, topMap)
	maxX := 0
	maxY := 0
	for p, _ := range topMap {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	// DFS
	total := 0
	for s := range startSpots {
		total += FindTotalTrailHeadsRating(s, 0, maxX, maxY, topMap)
	}

	return strconv.Itoa(total)
}

func FindTotalTrailHeadsRating(s Point, currentNum, maxX, maxY int, topMap map[Point]int) int {
	currentTotal := 0
	if currentNum == 9 {
		// We've found a trailhead
		return 1
	}

	// Check each direction with recursion
	north := Point{s.x, s.y - 1}
	east := Point{s.x + 1, s.y}
	south := Point{s.x, s.y + 1}
	west := Point{s.x - 1, s.y}

	if topMap[north] == currentNum+1 && CheckBounds(north, maxX, maxY) {
		currentTotal += FindTotalTrailHeadsRating(Point{s.x, s.y - 1}, currentNum+1, maxX, maxY, topMap)
	}
	if topMap[east] == currentNum+1 && CheckBounds(east, maxX, maxY) {
		currentTotal += FindTotalTrailHeadsRating(Point{s.x + 1, s.y}, currentNum+1, maxX, maxY, topMap)
	}
	if topMap[south] == currentNum+1 && CheckBounds(south, maxX, maxY) {
		currentTotal += FindTotalTrailHeadsRating(Point{s.x, s.y + 1}, currentNum+1, maxX, maxY, topMap)
	}
	if topMap[west] == currentNum+1 && CheckBounds(west, maxX, maxY) {
		currentTotal += FindTotalTrailHeadsRating(Point{s.x - 1, s.y}, currentNum+1, maxX, maxY, topMap)
	}

	return currentTotal
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
