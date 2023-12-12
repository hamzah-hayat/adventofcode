package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
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
		fmt.Println("Gold:" + PartTwo("input", 1000000))
	case "p1":
		fmt.Println("Silver:" + PartOne("input"))
	case "p2":
		fmt.Println("Gold:" + PartTwo("input", 1000000))
	}
}

type Point struct {
	x int
	y int
}

func PartOne(filename string) string {
	input := readInput(filename)

	// First expand galaxy
	expandedGalaxy := ExpandGalaxy(input)

	// Find each Galaxy
	galaxies := findGoals(expandedGalaxy)

	totalDistance := 0
	for len(galaxies) != 0 {
		checkGalaxy := galaxies[0]
		galaxies = galaxies[1:]

		// Distance between each pair
		for _, g := range galaxies {
			totalDistance += ManhattenDistance(checkGalaxy, g)
		}
	}

	num := strconv.Itoa(totalDistance)

	return num
}

func ExpandGalaxy(input []string) []string {
	var expanded []string
	// Expand each row
	for _, v := range input {
		if strings.Contains(v, "#") {
			expanded = append(expanded, v)
		} else {
			expanded = append(expanded, v, v)
		}
	}

	// Transpose
	var transpose []string
	for i := 0; i < len(expanded[0]); i++ {
		transposeLine := ""
		for _, v := range expanded {
			transposeLine += string(v[i])
		}
		transpose = append(transpose, transposeLine)
	}

	// Then expand each column
	var expanded_2 []string
	for _, v := range transpose {
		if strings.Contains(v, "#") {
			expanded_2 = append(expanded_2, v)
		} else {
			expanded_2 = append(expanded_2, v, v)
		}
	}

	// Transpose again
	var transpose_2 []string
	for i := 0; i < len(expanded_2[0]); i++ {
		transposeLine := ""
		for _, v := range expanded_2 {
			transposeLine += string(v[i])
		}
		transpose_2 = append(transpose_2, transposeLine)
	}

	// Then return
	return transpose_2
}

func findGoals(input []string) []Point {
	// Find E point in input
	var points []Point
	for y, line := range input {
		for x, v := range line {
			if string(v) == "#" {
				points = append(points, Point{x, y})
			}
		}
	}
	return points
}

// Figure out the Manhatten Distance between two points
func ManhattenDistance(firstPoint Point, secondPoint Point) int {
	x := abs(firstPoint.x - secondPoint.x)
	y := abs(firstPoint.y - secondPoint.y)
	return x + y
}

// Absoulute value of Int
func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func PartTwo(filename string, size int) string {
	input := readInput(filename)

	// Find rows that are "expanded"
	expandedRows := ExpandGalaxyRows(input)

	// Find columns that are "expanded"
	expandedColumns := ExpandGalaxyColumns(input)

	// Find each Galaxy
	galaxies := findGoals(input)

	totalDistance := 0
	for len(galaxies) != 0 {
		checkGalaxy := galaxies[0]
		galaxies = galaxies[1:]

		// Distance between each pair
		for _, g := range galaxies {
			totalDistance += GalaxyDistance(checkGalaxy, g, expandedRows, expandedColumns, size)
		}
	}

	num := strconv.Itoa(totalDistance)

	return num
}

func ExpandGalaxyRows(input []string) []int {
	var expanded []int
	// Expand each row
	for i, v := range input {
		if !strings.Contains(v, "#") {
			expanded = append(expanded, i)
		}
	}

	// Then return
	return expanded
}

func ExpandGalaxyColumns(input []string) []int {
	var expanded []int

	// Transpose
	var transpose []string
	for i := 0; i < len(input[0]); i++ {
		transposeLine := ""
		for _, v := range input {
			transposeLine += string(v[i])
		}
		transpose = append(transpose, transposeLine)
	}

	// Expand each row
	for i, v := range transpose {
		if !strings.Contains(v, "#") {
			expanded = append(expanded, i)
		}
	}

	// Then return
	return expanded
}

// Figure out the "Galaxy" Distance between two points
func GalaxyDistance(firstPoint Point, secondPoint Point, expandedRows, expandedColumns []int, size int) int {

	distance := 0

	// Work out X
	// Always go from left to right
	if firstPoint.x < secondPoint.x {
		for i := firstPoint.x; i < secondPoint.x; i++ {
			if slices.Contains(expandedColumns, i) {
				distance += size
			} else {
				distance += 1
			}
		}
	} else {
		for i := secondPoint.x; i < firstPoint.x; i++ {
			if slices.Contains(expandedColumns, i) {
				distance += size
			} else {
				distance += 1
			}
		}
	}

	// Work out Y
	// Always go from top to bottom
	if firstPoint.y < secondPoint.y {
		for i := firstPoint.y; i < secondPoint.y; i++ {
			if slices.Contains(expandedRows, i) {
				distance += size
			} else {
				distance += 1
			}
		}
	} else {
		for i := secondPoint.y; i < firstPoint.y; i++ {
			if slices.Contains(expandedRows, i) {
				distance += size
			} else {
				distance += 1
			}
		}
	}

	return distance
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
