package main

import (
	"bufio"
	"cmp"
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
		fmt.Println("Gold:" + PartTwo("input"))
	case "p1":
		fmt.Println("Silver:" + PartOne("input"))
	case "p2":
		fmt.Println("Gold:" + PartTwo("input"))
	}
}

func PartOne(filename string) string {
	input := readInput(filename)

	points := []Point{}
	distances := []DistanceBetweenPoints{}

	for _, lines := range input {
		splitPoint := strings.Split(lines, ",")
		xValue, _ := strconv.Atoi(splitPoint[0])
		yValue, _ := strconv.Atoi(splitPoint[1])
		point := Point{xValue, yValue}

		points = append(points, point)
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			distance := DistanceBetweenPoints{points[i], points[j], 0}
			xdistance := abs(points[i].x - points[j].x)
			ydistance := abs(points[i].y - points[j].y)
			distance.distance = xdistance + ydistance
			distances = append(distances, distance)
		}
	}

	slices.SortFunc(distances,
		func(a, b DistanceBetweenPoints) int {
			return cmp.Compare(a.distance, b.distance)
		})

	bestDistance := distances[len(distances)-1]
	bestRectangle := (abs(bestDistance.p1.x-bestDistance.p2.x) + 1) * (abs(bestDistance.p1.y-bestDistance.p2.y) + 1)

	return strconv.Itoa(bestRectangle)
}

func PartTwo(filename string) string {
	input := readInput(filename)

	grid := make(map[Point]int)
	points := []Point{}
	distances := []DistanceBetweenPoints{}

	// Read input and create points
	for _, lines := range input {
		splitPoint := strings.Split(lines, ",")
		xValue, _ := strconv.Atoi(splitPoint[0])
		yValue, _ := strconv.Atoi(splitPoint[1])
		point := Point{xValue, yValue}

		points = append(points, point)
	}

	// Create a grid with all valid points
	for i := 0; i < len(points)-1; i++ {
		grid[points[i]] = 1
		if points[i].x == points[i+1].x {
			// Vertical line
			for y := min(points[i].y, points[i+1].y); y <= max(points[i].y, points[i+1].y); y++ {
				grid[Point{points[i].x, y}] = 1
			}
		} else if points[i].y == points[i+1].y {
			// Horizontal line
			for x := min(points[i].x, points[i+1].x); x <= max(points[i].x, points[i+1].x); x++ {
				grid[Point{x, points[i].y}] = 1
			}
		}
	}

	// Hook up final and first point as well!
	if points[len(points)-1].x == points[0].x {
		// Vertical line
		for y := min(points[len(points)-1].y, points[0].y); y <= max(points[len(points)-1].y, points[0].y); y++ {
			grid[Point{points[len(points)-1].x, y}] = 1
		}
	} else if points[len(points)-1].y == points[0].y {
		// Horizontal line
		for x := min(points[len(points)-1].x, points[0].x); x <= max(points[len(points)-1].x, points[0].x); x++ {
			grid[Point{x, points[len(points)-1].y}] = 1
		}
	}

	//fmt.Println(PrintGrid(grid, 15, 15, false))

	// Flood fill from outside to find all invalid points
	grid = floodFill(grid, Point{0, 0})
	for x := 0; x < 15; x++ {
		for y := 0; y < 15; y++ {
			if grid[Point{x, y}] != -1  && grid[Point{x, y}] != 1 {
				grid[Point{x, y}] = 1
			}
		}
	}

	fmt.Println(PrintGrid(grid, 15, 15, true))

	// Make the best Distance cache
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			distance := DistanceBetweenPoints{points[i], points[j], 0}
			xdistance := abs(points[i].x - points[j].x)
			ydistance := abs(points[i].y - points[j].y)
			distance.distance = xdistance + ydistance
			distances = append(distances, distance)
		}
	}
	slices.SortFunc(distances,
		func(a, b DistanceBetweenPoints) int {
			return cmp.Compare(a.distance, b.distance)
		})

	// Find the best rectangle
	bestDistance := distances[len(distances)-1]
	for i := len(distances) - 1; i > 0; i-- {
		if isRectangleValid(grid, distances[i].p1, distances[i].p2) {
			bestDistance = distances[i]
			break
		}
	}
	bestRectangle := (abs(bestDistance.p1.x-bestDistance.p2.x) + 1) * (abs(bestDistance.p1.y-bestDistance.p2.y) + 1)

	return strconv.Itoa(bestRectangle)
}

func floodFill(grid map[Point]int, currentPoint Point) map[Point]int {
	if currentPoint.x < 0 || currentPoint.y < 0 || currentPoint.x >= 15 || currentPoint.y >= 15 {
		return grid // Out of bounds
	}

	// Start from the outside and fill in valid points
	grid[currentPoint] = -1

	up := Point{currentPoint.x, currentPoint.y - 1}
	down := Point{currentPoint.x, currentPoint.y + 1}
	left := Point{currentPoint.x - 1, currentPoint.y}
	right := Point{currentPoint.x + 1, currentPoint.y}

	if grid[up] != -1 && grid[up] != 1 {
		floodFill(grid, up)
	}

	if grid[down] != -1 && grid[down] != 1 {
		floodFill(grid, down)
	}

	if grid[left] != -1 && grid[left] != 1 {
		floodFill(grid, left)
	}

	if grid[right] != -1 && grid[right] != 1 {
		floodFill(grid, right)
	}

	return grid
}

func PrintGrid(grid map[Point]int, gridMaxX, gridMaxY int, emoji bool) string {
	// Print grid
	gridPrint := ""
	for y := 0; y <= gridMaxY; y++ {
		for x := 0; x <= gridMaxX; x++ {
			value, exists := grid[Point{x, y}]
			if exists && emoji {
				switch value {
				case 1:
					gridPrint += "⬜"
				case -1:
					gridPrint += "⬛"
				}
			}
			if exists && !emoji {
				switch value {
				case 1:
					gridPrint += "X"
				case -1:
					gridPrint += "."
				}
			}
		}
		gridPrint += "\n"
	}
	return gridPrint
}

func isRectangleValid(grid map[Point]int, p1, p2 Point) bool {
	rectanglePoints := []Point{}

	for x := min(p1.x, p2.x); x < max(p1.x, p2.x); x++ {
		for y := min(p1.y, p2.y); y < max(p1.y, p2.y); y++ {
			rectanglePoints = append(rectanglePoints, Point{x, y})
		}
	}

	valid := true
	for _, p := range rectanglePoints {
		if grid[p] != 1 {
			valid = false
			break
		}
	}

	return valid
}

type Point struct {
	x int
	y int
}

type DistanceBetweenPoints struct {
	p1       Point
	p2       Point
	distance int
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
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
