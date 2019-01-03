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
	checkPoints := FindInfinitePoints(points)
	// Find number of closest neighbours for point
	sizeMap := make(map[Point]int)

	for _, point := range checkPoints {
		areaSize := FindNumClosestNeighbours(point, points)
		sizeMap[point] = areaSize
		if areaSize > largestArea {
			largestArea = areaSize
		}
	}
	fmt.Println(sizeMap)

	return largestArea
}

// Remove any points that have an "infinite area", check if they are closest to one of the edge points of the grid
func FindInfinitePoints(points []Point) []Point {
	checkPoints := make([]Point, len(points))
	copy(checkPoints, points)
	// First find out the area we are dealing with, dont want to consider any points that have infinite area
	// Figure out max X and Y, then do closest neighbours on all edge points, and remove the points that are associated with them
	xmax := FindMax(1, points)
	ymax := FindMax(2, points)

	edgePoints := make([]Point, 0)
	for i := 0; i <= xmax; i++ {
		for j := 0; j <= ymax; j++ {
			if i == 0 || i == xmax {
				//Add this point to check points list
				edgePoints = append(edgePoints, Point{i, j})
			} else if j == 0 || j == ymax {
				//Add this point to check points list
				edgePoints = append(edgePoints, Point{i, j})
			}
		}
	}

	// Now that we have all the edge points
	// Find the list of nearest neighbours for each of those points, then remove them
	for _, point := range edgePoints {
		closest := FindClosestNeighbour(point, points)
		// Remove this from our checkPoints list
		for i, checkPoint := range checkPoints {
			if checkPoint == closest {
				checkPoints = append(checkPoints[:i], checkPoints[i+1:]...)
			}
		}
	}
	return checkPoints
}

// FindMax for x and y
func FindMax(field int, points []Point) int {

	highest := 0

	switch field {
	case 1:
		for _, point := range points {
			if point.x > highest {
				highest = point.x
			}
		}
	case 2:
		for _, point := range points {
			if point.y > highest {
				highest = point.y
			}
		}
	}

	return highest

}

// Find number of Closest neighbours for a point
func FindNumClosestNeighbours(neighbour Point, points []Point) int {
	xmax := FindMax(1, points)
	ymax := FindMax(2, points)
	areaSize := 0

	for i := 0; i < xmax; i++ {
		for j := 0; j < ymax; j++ {
			if (FindClosestNeighbour(Point{i, j}, points) == neighbour) {
				areaSize++
			}
		}

	}

	return areaSize
}

// Find out which is the closest neighbour to the gridpoint
func FindClosestNeighbour(gridPoint Point, neighbours []Point) Point {
	var closestPoint Point
	closestPoint = neighbours[0]

	for _, point := range neighbours[1:] {
		delta := ManhattenDistance(point, gridPoint) - ManhattenDistance(closestPoint, gridPoint)
		if delta < 0 {
			closestPoint = point
		} else if delta == 0 {
			// This Point is considered to be not close to any point as two points are
			// Equal distance away
			return Point{-1, -1}
		}
	}

	return closestPoint
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
