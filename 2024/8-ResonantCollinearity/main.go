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

	// Figure out all unique frequencies
	uniqueFreq := make(map[string]bool)
	for _, v := range grid {
		if v != "." {
			uniqueFreq[v] = true
		}
	}

	// For each unique frequency, get all antinodes
	uniqueAntiNodes := make(map[Point]bool)
	for freq := range uniqueFreq {
		for p := range GetAllAntiNodesOfFreq(freq, grid, maxX, maxY) {
			uniqueAntiNodes[p] = true
		}
	}

	return strconv.Itoa(len(uniqueAntiNodes))
}

func GetAllAntiNodesOfFreq(freq string, grid map[Point]string, maxX, maxY int) map[Point]bool {

	// First find all points with this freq on grid
	totalPoints := make([]Point, 0)
	for p, f := range grid {
		if f == freq {
			totalPoints = append(totalPoints, p)
		}
	}

	antiNodes := make(map[Point]bool)
	for i := 0; i < len(totalPoints)-1; i++ {
		for j := i + 1; j < len(totalPoints); j++ {
			// Check each pair
			XDifi := totalPoints[i].x - totalPoints[j].x
			YDifi := totalPoints[i].y - totalPoints[j].y
			FirstPoint := Point{totalPoints[i].x + XDifi, totalPoints[i].y + YDifi}

			XDifj := totalPoints[j].x - totalPoints[i].x
			YDifj := totalPoints[j].y - totalPoints[i].y
			SecondPoint := Point{totalPoints[j].x + XDifj, totalPoints[j].y + YDifj}

			if CheckBounds(FirstPoint, maxX, maxY) {
				antiNodes[FirstPoint] = true
			}
			if CheckBounds(SecondPoint, maxX, maxY) {
				antiNodes[SecondPoint] = true
			}
		}
	}

	return antiNodes
}

func CheckBounds(FirstPoint Point, xMax, yMax int) bool {

	if FirstPoint.x < 0 || FirstPoint.x > xMax {
		return false
	}
	if FirstPoint.y < 0 || FirstPoint.y > yMax {
		return false
	}
	return true
}

type Point struct {
	x int
	y int
}

func PartTwo(filename string) string {
	input := readInput(filename)

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

	// Figure out all unique frequencies
	uniqueFreq := make(map[string]bool)
	for _, v := range grid {
		if v != "." {
			uniqueFreq[v] = true
		}
	}

	// For each unique frequency, get all antinodes
	uniqueAntiNodes := make(map[Point]bool)
	for freq := range uniqueFreq {
		for p := range GetAllAntiNodesOfFreqHarmonics(freq, grid, maxX, maxY) {
			uniqueAntiNodes[p] = true
		}
	}

	return strconv.Itoa(len(uniqueAntiNodes))
}

func GetAllAntiNodesOfFreqHarmonics(freq string, grid map[Point]string, maxX, maxY int) map[Point]bool {

	// First find all points with this freq on grid
	totalPoints := make([]Point, 0)
	for p, f := range grid {
		if f == freq {
			totalPoints = append(totalPoints, p)
		}
	}

	antiNodes := make(map[Point]bool)
	for i := 0; i < len(totalPoints)-1; i++ {
		for j := i + 1; j < len(totalPoints); j++ {

			// Add both nodes as well
			antiNodes[totalPoints[i]] = true
			antiNodes[totalPoints[j]] = true

			// Check each pair
			XDifi := totalPoints[i].x - totalPoints[j].x
			YDifi := totalPoints[i].y - totalPoints[j].y
			counter := 1
			for {
				FirstPoint := Point{totalPoints[i].x + (XDifi * counter), totalPoints[i].y + (YDifi * counter)}

				if CheckBounds(FirstPoint, maxX, maxY) {
					antiNodes[FirstPoint] = true
					counter++
				} else {
					break
				}
			}

			XDifj := totalPoints[j].x - totalPoints[i].x
			YDifj := totalPoints[j].y - totalPoints[i].y
			counter = 1
			for {
				SecondPoint := Point{totalPoints[j].x + (XDifj * counter), totalPoints[j].y + (YDifj * counter)}
				if CheckBounds(SecondPoint, maxX, maxY) {
					antiNodes[SecondPoint] = true
					counter++
				} else {
					break
				}
			}
		}
	}

	return antiNodes
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
