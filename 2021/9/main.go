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

func init() {
	// Use Flags to run a part
	methodP = flag.String("method", "p2", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()
}

func main() {
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

	// Get data
	points := make(map[Point]int)
	for i, v := range input {
		for j, c := range v {
			pointStr := Point{i, j}
			num, _ := strconv.Atoi(string(c))
			points[pointStr] = num
		}
	}

	// Find each low point
	// add up all risk levels
	sum := 0
	for p := range points {
		if IsLowPoint(points, p) {
			sum += GetRiskLevel(points, p)
		}
	}

	fmt.Println(sum)

}

func PartTwo() {
	input := readInput()

	// Get data
	points := make(map[Point]int)
	for i, v := range input {
		for j, c := range v {
			pointStr := Point{i, j}
			num, _ := strconv.Atoi(string(c))
			points[pointStr] = num
		}
	}

	// Find each low point
	// get each basin
	var basins []map[Point]int
	for p := range points {
		if IsLowPoint(points, p) {
			b := make(map[Point]int)
			basins = append(basins, GetBasinFromPoint(points, p, b))
		}
	}

	// Get three largest basins
	var biggestThreeBasins []map[Point]int

	for i := 0; i < 3; i++ {
		biggest := 0
		var biggestBasin map[Point]int
		numInArray := 0
		for i, v := range basins {
			if len(v) > biggest {
				biggest = len(v)
				biggestBasin = v
				numInArray = i
			}
		}
		biggestThreeBasins = append(biggestThreeBasins, biggestBasin)
		basins = remove(basins, numInArray)
	}

	sum := len(biggestThreeBasins[0]) * len(biggestThreeBasins[1]) * len(biggestThreeBasins[2])
	fmt.Println(sum)
}

type Point struct {
	X int
	Y int
}

func IsLowPoint(points map[Point]int, p Point) bool {

	up := Point{p.X + 1, p.Y}
	if _, ok := points[up]; ok {
		//do something here
		if points[p] >= points[up] {
			return false
		}
	}

	right := Point{p.X, p.Y + 1}
	if _, ok := points[right]; ok {
		//do something here
		if points[p] >= points[right] {
			return false
		}
	}
	down := Point{p.X - 1, p.Y}
	if _, ok := points[down]; ok {
		//do something here
		if points[p] >= points[down] {
			return false
		}
	}
	left := Point{p.X, p.Y - 1}
	if _, ok := points[left]; ok {
		//do something here
		if points[p] >= points[left] {
			return false
		}
	}

	return true
}

func GetRiskLevel(points map[Point]int, p Point) int {
	return points[p] + 1
}

func GetBasinFromPoint(points map[Point]int, p Point, basin map[Point]int) map[Point]int {

	// If we already have this point, return
	if _, ok := basin[p]; ok {
		return basin
	}

	// otherwise add this point
	basin[p] = points[p]

	// now we try and check a adjacent point
	// check if it exists and is not 9
	// then recursively add to basin
	up := Point{p.X + 1, p.Y}
	if _, ok := points[up]; ok {
		if points[up] != 9 {
			basin = merge(basin,GetBasinFromPoint(points, up, basin))
		}
	}

	right := Point{p.X, p.Y + 1}
	if _, ok := points[right]; ok {
		if points[right] != 9 {
			basin = merge(basin,GetBasinFromPoint(points, right, basin))
		}
	}
	down := Point{p.X - 1, p.Y}
	if _, ok := points[down]; ok {
		if points[down] != 9 {
			basin = merge(basin,GetBasinFromPoint(points, down, basin))
		}
	}
	left := Point{p.X, p.Y - 1}
	if _, ok := points[left]; ok {
		if points[left] != 9 {
			basin = merge(basin,GetBasinFromPoint(points, left, basin))
		}
	}

	return basin
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

// remove element from array
func remove(s []map[Point]int, i int) []map[Point]int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func merge(ms ...map[Point]int) map[Point]int {
	res := map[Point]int{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}
