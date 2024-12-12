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

	cropMap := make(map[Point]string)
	// Setup Map
	for y, line := range input {
		for x, char := range line {
			cropMap[Point{x, y}] = string(char)
		}
	}
	maxX := 0
	maxY := 0
	for p, _ := range cropMap {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	cropGroupList := make([]map[Point]string, 0)
	// Find each connected group of crops
	for point, crop := range cropMap {
		if !InCropGroupList(point, crop, cropGroupList) {
			// Lets find all crops connected to this
			newCropGroup := make(map[Point]string)
			FindAllConnectedCrops(point, crop, cropMap, newCropGroup, maxX, maxY)
			cropGroupList = append(cropGroupList, newCropGroup)
		}
	}

	// Work out area and perimeter
	total := 0
	for _, cropGroup := range cropGroupList {
		area := len(cropGroup)
		perimeter := FindPerimeter(cropGroup, maxX, maxY)

		total += area * perimeter
	}

	return strconv.Itoa(total)
}

func FindPerimeter(cropGroup map[Point]string, maxX, maxY int) int {

	totalPerm := 0
	for point, crop := range cropGroup {
		fencePerm := 4

		// Check each direction
		north := Point{point.x, point.y - 1}
		east := Point{point.x + 1, point.y}
		south := Point{point.x, point.y + 1}
		west := Point{point.x - 1, point.y}

		if cropGroup[north] == crop && CheckBounds(north, maxX, maxY) {
			fencePerm--
		}
		if cropGroup[east] == crop && CheckBounds(east, maxX, maxY) {
			fencePerm--
		}
		if cropGroup[south] == crop && CheckBounds(south, maxX, maxY) {
			fencePerm--
		}
		if cropGroup[west] == crop && CheckBounds(west, maxX, maxY) {
			fencePerm--
		}
		totalPerm += fencePerm
	}
	return totalPerm
}

func FindAllConnectedCrops(point Point, crop string, cropMap, newCropGroup map[Point]string, maxX, maxY int) map[Point]string {

	if newCropGroup[point] == crop {
		return newCropGroup
	} else {
		// Add current to Crop Group
		newCropGroup[point] = crop
	}

	// Check each direction with recursion
	north := Point{point.x, point.y - 1}
	east := Point{point.x + 1, point.y}
	south := Point{point.x, point.y + 1}
	west := Point{point.x - 1, point.y}

	if cropMap[north] == crop && CheckBounds(north, maxX, maxY) {
		FindAllConnectedCrops(Point{point.x, point.y - 1}, crop, cropMap, newCropGroup, maxX, maxY)
	}
	if cropMap[east] == crop && CheckBounds(east, maxX, maxY) {
		FindAllConnectedCrops(Point{point.x + 1, point.y}, crop, cropMap, newCropGroup, maxX, maxY)
	}
	if cropMap[south] == crop && CheckBounds(south, maxX, maxY) {
		FindAllConnectedCrops(Point{point.x, point.y + 1}, crop, cropMap, newCropGroup, maxX, maxY)
	}
	if cropMap[west] == crop && CheckBounds(west, maxX, maxY) {
		FindAllConnectedCrops(Point{point.x - 1, point.y}, crop, cropMap, newCropGroup, maxX, maxY)
	}

	return newCropGroup

}

type Point struct {
	x int
	y int
}

func InCropGroupList(point Point, crop string, cropGroupList []map[Point]string) bool {
	for i := 0; i < len(cropGroupList); i++ {
		for p, c := range cropGroupList[i] {
			if p == point && c == crop {
				return true
			}
		}
	}
	return false
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

func PartTwo(filename string) string {
	input := readInput(filename)

	cropMap := make(map[Point]string)
	// Setup Map
	for y, line := range input {
		for x, char := range line {
			cropMap[Point{x, y}] = string(char)
		}
	}
	maxX := 0
	maxY := 0
	for p, _ := range cropMap {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	cropGroupList := make([]map[Point]string, 0)
	// Find each connected group of crops
	for point, crop := range cropMap {
		if !InCropGroupList(point, crop, cropGroupList) {
			// Lets find all crops connected to this
			newCropGroup := make(map[Point]string)
			FindAllConnectedCrops(point, crop, cropMap, newCropGroup, maxX, maxY)
			cropGroupList = append(cropGroupList, newCropGroup)
		}
	}

	// Work out area and perimeter
	total := 0
	for _, cropGroup := range cropGroupList {
		area := len(cropGroup)
		sides := FindSides(cropGroup, maxX, maxY)

		total += area * sides
	}

	return strconv.Itoa(total)
}

func FindSides(cropGroup map[Point]string, maxX, maxY int) int {
	sides := 0
	for point, crop := range cropGroup {

		// Check each direction
		north := Point{point.x, point.y - 1}
		east := Point{point.x + 1, point.y}
		south := Point{point.x, point.y + 1}
		west := Point{point.x - 1, point.y}

		north_east := Point{point.x + 1, point.y - 1}
		south_east := Point{point.x + 1, point.y + 1}
		south_west := Point{point.x - 1, point.y + 1}
		north_west := Point{point.x - 1, point.y - 1}

		// North-East corner
		if cropGroup[north] != crop && cropGroup[east] != crop {
			sides++
		}
		if cropGroup[north] == crop && cropGroup[east] == crop && cropGroup[north_east] != crop {
			sides++
		}

		// South-East corner
		if cropGroup[south] != crop && cropGroup[east] != crop {
			sides++
		}
		if cropGroup[south] == crop && cropGroup[east] == crop && cropGroup[south_east] != crop {
			sides++
		}

		// South-West corner
		if cropGroup[south] != crop && cropGroup[west] != crop {
			sides++
		}
		if cropGroup[south] == crop && cropGroup[west] == crop && cropGroup[south_west] != crop {
			sides++
		}

		// North-West corner
		if cropGroup[north] != crop && cropGroup[west] != crop {
			sides++
		}
		if cropGroup[north] == crop && cropGroup[west] == crop && cropGroup[north_west] != crop {
			sides++
		}

	}
	return sides
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
