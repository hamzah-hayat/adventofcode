package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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

	pointsVisitedByTail := make(map[Point]bool)
	currentTail := Point{x: 0, y: 0}
	currentHead := Point{x: 0, y: 0}

	for _, line := range input {
		// directions are U,R,D,L Up,Right,Down,Left
		// steps are 1-99
		regex := regexp.MustCompile(`([U|R|D|L]) ([0-9]+)`)
		regexResult := regex.FindStringSubmatch(line)

		direction := regexResult[1]
		steps, _ := strconv.Atoi(regexResult[2])

		// Now calculate steps
		for i := 0; i < steps; i++ {
			// Move the head
			switch direction {
			case "U":
				currentHead.y = currentHead.y + 1
			case "R":
				currentHead.x = currentHead.x + 1
			case "D":
				currentHead.y = currentHead.y - 1
			case "L":
				currentHead.x = currentHead.x - 1
			}

			// Do we need to move the tail?
			distx := 0
			disty := 0
			if currentTail.x >= currentHead.x {
				distx = currentTail.x - currentHead.x
			} else {
				distx = currentHead.x - currentTail.x
			}
			if currentTail.y >= currentHead.y {
				disty = currentTail.y - currentHead.y
			} else {
				disty = currentHead.y - currentTail.y
			}

			distance := distx + disty
			if distance == 2 {
				// Cardinal direction
				/*
					1. Up
					2. Right
					3. Down
					4. Left
				*/
				if currentTail.y+2 == currentHead.y {
					currentTail.y = currentTail.y + 1
				}
				if currentTail.x+2 == currentHead.x {
					currentTail.x = currentTail.x + 1
				}
				if currentTail.y-2 == currentHead.y {
					currentTail.y = currentTail.y - 1
				}
				if currentTail.x-2 == currentHead.x {
					currentTail.x = currentTail.x - 1
				}
			}
			if distance == 3 {
				//diagonal direction
				/*
					1. Up-Right
					2. Down-Right
					3. Down-Left
					4. Up-Left
				*/
				if currentTail.y < currentHead.y && currentTail.x < currentHead.x {
					currentTail.y = currentTail.y + 1
					currentTail.x = currentTail.x + 1
				}
				if currentTail.y > currentHead.y && currentTail.x < currentHead.x {
					currentTail.y = currentTail.y - 1
					currentTail.x = currentTail.x + 1
				}
				if currentTail.y > currentHead.y && currentTail.x > currentHead.x {
					currentTail.y = currentTail.y - 1
					currentTail.x = currentTail.x - 1
				}
				if currentTail.y < currentHead.y && currentTail.x > currentHead.x {
					currentTail.y = currentTail.y + 1
					currentTail.x = currentTail.x - 1
				}
			}
			// set point in memory
			newPoint := Point{currentTail.x, currentTail.y}
			pointsVisitedByTail[newPoint] = true
		}
	}

	numVisitedStr := strconv.Itoa(len(pointsVisitedByTail))
	return numVisitedStr
}

func PartTwo(filename string) string {
	input := readInput(filename)

	pointsVisitedByTail := make(map[Point]bool)
	knots := make(map[int]Point) // we have nine tails now
	for i := 0; i < 10; i++ {
		knots[i] = Point{0, 0}
	}

	for _, line := range input {
		// directions are U,R,D,L Up,Right,Down,Left
		regex := regexp.MustCompile(`([U|R|D|L]) ([0-9]+)`)
		regexResult := regex.FindStringSubmatch(line)

		direction := regexResult[1]
		steps, _ := strconv.Atoi(regexResult[2])

		for i := 0; i < steps; i++ {
			// move head
			switch direction {
			case "U":
				newPoint := Point{knots[0].x, knots[0].y + 1}
				knots[0] = newPoint
			case "R":
				newPoint := Point{knots[0].x + 1, knots[0].y}
				knots[0] = newPoint
			case "D":
				newPoint := Point{knots[0].x, knots[0].y - 1}
				knots[0] = newPoint
			case "L":
				newPoint := Point{knots[0].x - 1, knots[0].y}
				knots[0] = newPoint
			}

			for tailNum := 1; tailNum < 10; tailNum++ {
				// Do we need to move this tail?
				distx := 0
				disty := 0
				if knots[tailNum].x >= knots[tailNum-1].x {
					distx = knots[tailNum].x - knots[tailNum-1].x
				} else {
					distx = knots[tailNum-1].x - knots[tailNum].x
				}
				if knots[tailNum].y >= knots[tailNum-1].y {
					disty = knots[tailNum].y - knots[tailNum-1].y
				} else {
					disty = knots[tailNum-1].y - knots[tailNum].y
				}

				distance := distx + disty
				if distance == 2 {
					// Cardinal direction
					/*
						1. Up
						2. Right
						3. Down
						4. Left
					*/
					if knots[tailNum].y+2 == knots[tailNum-1].y {
						newPoint := Point{knots[tailNum].x, knots[tailNum].y + 1}
						knots[tailNum] = newPoint
					}
					if knots[tailNum].x+2 == knots[tailNum-1].x {
						newPoint := Point{knots[tailNum].x + 1, knots[tailNum].y}
						knots[tailNum] = newPoint
					}
					if knots[tailNum].y-2 == knots[tailNum-1].y {
						newPoint := Point{knots[tailNum].x, knots[tailNum].y - 1}
						knots[tailNum] = newPoint
					}
					if knots[tailNum].x-2 == knots[tailNum-1].x {
						newPoint := Point{knots[tailNum].x - 1, knots[tailNum].y}
						knots[tailNum] = newPoint
					}
				}
				if distance > 2 {
					//diagonal direction
					/*
						1. Up-Right
						2. Down-Right
						3. Down-Left
						4. Up-Left
					*/
					if knots[tailNum].y < knots[tailNum-1].y && knots[tailNum].x < knots[tailNum-1].x {
						newPoint := Point{knots[tailNum].x + 1, knots[tailNum].y + 1}
						knots[tailNum] = newPoint
					}
					if knots[tailNum].y > knots[tailNum-1].y && knots[tailNum].x < knots[tailNum-1].x {
						newPoint := Point{knots[tailNum].x + 1, knots[tailNum].y - 1}
						knots[tailNum] = newPoint
					}
					if knots[tailNum].y > knots[tailNum-1].y && knots[tailNum].x > knots[tailNum-1].x {
						newPoint := Point{knots[tailNum].x - 1, knots[tailNum].y - 1}
						knots[tailNum] = newPoint
					}
					if knots[tailNum].y < knots[tailNum-1].y && knots[tailNum].x > knots[tailNum-1].x {
						newPoint := Point{knots[tailNum].x - 1, knots[tailNum].y + 1}
						knots[tailNum] = newPoint
					}
				}
				// set point in memory
				if tailNum == 9 {
					newPoint := Point{knots[tailNum].x, knots[tailNum].y}
					pointsVisitedByTail[newPoint] = true
				}
			}
		}
	}

	numVisitedStr := strconv.Itoa(len(pointsVisitedByTail))
	return numVisitedStr
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

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
