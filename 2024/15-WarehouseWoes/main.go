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

	warehouseGrid := make(map[Point]string)
	instructions := ""

	// Read in input
	readingGrid := true
	for y, line := range input {
		if line == "" {
			readingGrid = false
			continue
		}

		if readingGrid {
			for x := 0; x < len(line); x++ {
				warehouseGrid[Point{x, y}] = string(line[x])
			}
		} else {
			instructions += line
			instructions = strings.TrimRight(instructions, "\n")
		}
	}

	// Get MaxX/MaxY
	maxX := 0
	maxY := 0
	for p, _ := range warehouseGrid {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	//fmt.Println(PrintGrid(warehouseGrid, maxX, maxY,true))

	// Run robot movement
	for _, movement := range instructions {
		currentPoint := FindRobot(warehouseGrid)
		MoveObject(string(movement), currentPoint, warehouseGrid)
		//fmt.Println(PrintGrid(warehouseGrid, maxX, maxY,true))
	}

	total := 0
	// Score each box based on distance from top left
	// 100 times its distance from the top edge of the map plus its distance from the left edge of the map
	for p, obj := range warehouseGrid {
		if obj == "O" {
			total += p.x + p.y*100
		}
	}

	return strconv.Itoa(total)
}

func FindRobot(warehouseGrid map[Point]string) Point {
	for p, v := range warehouseGrid {
		if v == "@" {
			return p
		}
	}
	return Point{-1, -1}
}

// MoveObject moves our current point, while also trying to continue moving
func MoveObject(movement string, currentPoint Point, warehouseGrid map[Point]string) {

	nextPoint := Point{-1, -1}
	// Move
	switch movement {
	case "^":
		nextPoint = Point{currentPoint.x, currentPoint.y - 1}
	case ">":
		nextPoint = Point{currentPoint.x + 1, currentPoint.y}
	case "v":
		nextPoint = Point{currentPoint.x, currentPoint.y + 1}
	case "<":
		nextPoint = Point{currentPoint.x - 1, currentPoint.y}
	}

	// Move box in the way first if possible
	if warehouseGrid[nextPoint] == "O" {
		MoveObject(movement, nextPoint, warehouseGrid)
	}

	if warehouseGrid[nextPoint] == "#" {
		// We are done
		return
	} else if warehouseGrid[nextPoint] == "." {
		// Single movement
		warehouseGrid[nextPoint] = warehouseGrid[currentPoint]
		warehouseGrid[currentPoint] = "."
	}
}

func PrintGrid(grid map[Point]string, gridMaxX, gridMaxY int, emoji bool) string {
	// Print grid
	gridPrint := ""
	for y := 0; y <= gridMaxY; y++ {
		for x := 0; x <= gridMaxX; x++ {
			value, exists := grid[Point{x, y}]
			if exists && emoji {
				switch value {
				case "#":
					gridPrint += "🧱"
				case "O":
					gridPrint += "📦"
				case "[":
					gridPrint += "◀️"
				case "]":
					gridPrint += "▶️"
				case "@":
					gridPrint += "🤖"
				case ".":
					gridPrint += "⬛"
				}
			}
			if exists && !emoji {
				gridPrint += value
			}
		}
		gridPrint += "\n"
	}
	return gridPrint
}

type Point struct {
	x int
	y int
}

func PartTwo(filename string) string {
	input := readInput(filename)

	warehouseGrid := make(map[Point]string)
	instructions := ""

	// Read in input
	readingGrid := true
	for y, line := range input {
		if line == "" {
			readingGrid = false
			continue
		}

		if readingGrid {
			gridX := 0
			for x := 0; x < len(line); x++ {
				// If the tile is #, the new map contains ## instead.
				// If the tile is O, the new map contains [] instead.
				// If the tile is ., the new map contains .. instead.
				// If the tile is @, the new map contains @. instead.

				switch string(line[x]) {
				case "#":
					warehouseGrid[Point{gridX, y}] = "#"
					gridX++
					warehouseGrid[Point{gridX, y}] = "#"
				case "O":
					warehouseGrid[Point{gridX, y}] = "["
					gridX++
					warehouseGrid[Point{gridX, y}] = "]"
				case ".":
					warehouseGrid[Point{gridX, y}] = "."
					gridX++
					warehouseGrid[Point{gridX, y}] = "."
				case "@":
					warehouseGrid[Point{gridX, y}] = "@"
					gridX++
					warehouseGrid[Point{gridX, y}] = "."
				}
				gridX++
			}
		} else {
			instructions += line
			instructions = strings.TrimRight(instructions, "\n")
		}
	}

	// Get MaxX/MaxY
	maxX := 0
	maxY := 0
	for p, _ := range warehouseGrid {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	//fmt.Println(PrintGrid(warehouseGrid, maxX, maxY, false))

	// Run robot movement
	for _, movement := range instructions {
		// fmt.Print("Robot is going to move to the ")
		// switch string(movement) {
		// case "^":
		// 	fmt.Println("North")
		// case ">":
		// 	fmt.Println("East")
		// case "v":
		// 	fmt.Println("South")
		// case "<":
		// 	fmt.Println("West")
		// }
		currentPoint := FindRobot(warehouseGrid)
		MoveRobot(string(movement), currentPoint, warehouseGrid)
		//fmt.Println(PrintGrid(warehouseGrid, maxX, maxY, false))
	}

	total := 0
	// Score each box based on distance from top left
	// 100 times its distance from the top edge of the map plus its distance from the left edge of the map
	for p, obj := range warehouseGrid {
		if obj == "[" {
			total += p.x + p.y*100
		}
	}
	return strconv.Itoa(total)
}

func MoveRobot(movement string, robotPosition Point, warehouseGrid map[Point]string) {
	nextPoint := Point{-1, -1}

	// Move
	switch movement {
	case "^":
		nextPoint = Point{robotPosition.x, robotPosition.y - 1}
	case ">":
		nextPoint = Point{robotPosition.x + 1, robotPosition.y}
	case "v":
		nextPoint = Point{robotPosition.x, robotPosition.y + 1}
	case "<":
		nextPoint = Point{robotPosition.x - 1, robotPosition.y}
	}

	// Standard Robot movement
	if warehouseGrid[nextPoint] == "#" {
		// We are done
		return
	} else if (warehouseGrid[nextPoint] == "[" || warehouseGrid[nextPoint] == "]") && (movement == "<" || movement == ">") {
		// Horizontal Box movement
		MoveBigBoxHorizontal(movement, nextPoint, warehouseGrid)
	} else if warehouseGrid[nextPoint] == "[" && (movement == "^" || movement == "v") {
		leftBoxPos := nextPoint
		rightBoxPos := Point{nextPoint.x + 1, nextPoint.y}
		// Box movement
		MoveBigBoxVertical(movement, []Point{leftBoxPos, rightBoxPos}, warehouseGrid)
	} else if warehouseGrid[nextPoint] == "]" && (movement == "^" || movement == "v") {
		leftBoxPos := Point{nextPoint.x - 1, nextPoint.y}
		rightBoxPos := nextPoint
		// Box movement
		MoveBigBoxVertical(movement, []Point{leftBoxPos, rightBoxPos}, warehouseGrid)
	}

	if warehouseGrid[nextPoint] == "." {
		// Single movement
		warehouseGrid[nextPoint] = warehouseGrid[robotPosition]
		warehouseGrid[robotPosition] = "."
		return
	}
}

func MoveBigBoxHorizontal(movement string, currentPoint Point, warehouseGrid map[Point]string) {

	nextPoint := Point{-1, -1}

	// Move
	switch movement {
	case ">":
		nextPoint = Point{currentPoint.x + 1, currentPoint.y}
	case "<":
		nextPoint = Point{currentPoint.x - 1, currentPoint.y}
	}

	// Check if box in way firt
	if warehouseGrid[nextPoint] == "]" || warehouseGrid[nextPoint] == "[" {
		MoveBigBoxHorizontal(movement, nextPoint, warehouseGrid)
	}

	// Otherwise standard movement
	if warehouseGrid[nextPoint] == "#" {
		// We are done
		return
	} else if warehouseGrid[nextPoint] == "." {
		// Single movement
		warehouseGrid[nextPoint] = warehouseGrid[currentPoint]
		warehouseGrid[currentPoint] = "."
		return
	}
}

// We can only move if all points are valid
func MoveBigBoxVertical(movement string, currentPoints []Point, warehouseGrid map[Point]string) {

	maybeCanMove := true
	for _, currentPoint := range currentPoints {
		nextPoint := Point{-1, -1}
		switch movement {
		case "^":
			nextPoint = Point{currentPoint.x, currentPoint.y - 1}
		case "v":
			nextPoint = Point{currentPoint.x, currentPoint.y + 1}
		}

		// First check if we are obstructed
		if warehouseGrid[nextPoint] == "#" {
			// We are done
			maybeCanMove = false
			return
		}
	}

	// We've now checked all points, so no walls
	// Now check if any boxes, if so, group all points and run method again
	boxesToCheck := make([]Point, 0)
	if maybeCanMove {
		for _, currentPoint := range currentPoints {

			nextPoint := Point{-1, -1}
			switch movement {
			case "^":
				nextPoint = Point{currentPoint.x, currentPoint.y - 1}
			case "v":
				nextPoint = Point{currentPoint.x, currentPoint.y + 1}
			}

			// Check if box in way
			if warehouseGrid[nextPoint] == "[" {
				leftBoxPos := nextPoint
				rightBoxPos := Point{nextPoint.x + 1, nextPoint.y}

				containsLeft := false
				containsRight := false
				for _, b := range boxesToCheck {
					if b == leftBoxPos {
						containsLeft = true
					}
					if b == rightBoxPos {
						containsRight = true
					}
				}

				if !containsLeft {
					boxesToCheck = append(boxesToCheck, leftBoxPos)
				}
				if !containsRight {
					boxesToCheck = append(boxesToCheck, rightBoxPos)
				}
			} else if warehouseGrid[nextPoint] == "]" {
				leftBoxPos := Point{nextPoint.x - 1, nextPoint.y}
				rightBoxPos := nextPoint
				containsLeft := false
				containsRight := false
				for _, b := range boxesToCheck {
					if b == leftBoxPos {
						containsLeft = true
					}
					if b == rightBoxPos {
						containsRight = true
					}
				}

				if !containsLeft {
					boxesToCheck = append(boxesToCheck, leftBoxPos)
				}
				if !containsRight {
					boxesToCheck = append(boxesToCheck, rightBoxPos)
				}
			}
		}
	}

	// If we have to check some boxes, do that
	if len(boxesToCheck) > 0 {
		MoveBigBoxVertical(movement, boxesToCheck, warehouseGrid)
	}

	// Now check again if we can move all our points
	canMove := true
	for _, currentPoint := range currentPoints {
		nextPoint := Point{-1, -1}
		switch movement {
		case "^":
			nextPoint = Point{currentPoint.x, currentPoint.y - 1}
		case "v":
			nextPoint = Point{currentPoint.x, currentPoint.y + 1}
		}

		if warehouseGrid[nextPoint] != "." {
			canMove = false
		}
	}

	// Now, if we can move, we do so
	if canMove {
		for _, currentPoint := range currentPoints {
			nextPoint := Point{-1, -1}
			switch movement {
			case "^":
				nextPoint = Point{currentPoint.x, currentPoint.y - 1}
			case "v":
				nextPoint = Point{currentPoint.x, currentPoint.y + 1}
			}

			warehouseGrid[nextPoint] = warehouseGrid[currentPoint]
			warehouseGrid[currentPoint] = "."
		}
	}
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
