package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hamzah-hayat/adventofcode/intcode"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p2", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()

	switch *methodP {
	case "p1":
		partOne()
		break
	case "p2":
		partTwo()
		break
	case "test":
		break
	}
}

func partOne() {
	input := readInput()
	program := convertToInts(input)

	inputChan, outputChan, terminateChan, _ := makeChannels()

	go intcode.RunIntCodeProgramWaitForTermination(program, inputChan, outputChan, terminateChan, nil)

	cameraView := buildCameraView(outputChan, terminateChan)

	intersections := findIntersections(cameraView)

	total := 0

	for _, s := range intersections {
		total += (s.x) * (s.y)
	}

	fmt.Println("The total sum of the alignment parameters is", total)
}

func buildCameraView(outputChan chan int, terminateChan chan bool) map[space]string {
	cameraView := make(map[space]string)

	t := false
	x := 0
	y := 0
	for {
		select {
		case <-terminateChan:
			t = true
			break
		case input := <-outputChan:
			if input == 10 {
				x++
				y = 0
				break
			} else {
				cameraView[space{x: x, y: y}] = string(input)
				y++
				break
			}
		}
		if t {
			break
		}
	}

	return cameraView
}

func findIntersections(cameraView map[space]string) []space {

	intersections := make([]space, 0)
	for i, val := range cameraView {

		if val == "#" {
			x := i.x
			y := i.y

			intersection := cameraView[space{x: x + 1, y: y}] == "#" && cameraView[space{x: x - 1, y: y}] == "#" && cameraView[space{x: x, y: y + 1}] == "#" && cameraView[space{x: x, y: y - 1}] == "#"

			if intersection {
				intersections = append(intersections, i)
			}

		}

	}
	return intersections
}

func printCamera(cameraView map[space]string) string {
	cameraViewStr := ""

	// Find max x and y
	maxX := 0
	maxY := 0
	for s := range cameraView {
		if s.x+1 > maxX {
			maxX = s.x + 1
		}
		if s.y+1 > maxY {
			maxY = s.y + 1
		}
	}

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			cameraViewStr += cameraView[space{x: x, y: y}]
		}
		cameraViewStr += "\n"
	}
	return cameraViewStr
}

func partTwo() {
	input := readInput()
	program := convertToInts(input)

	inputChan, outputChan, terminateChan, _ := makeChannels()

	go intcode.RunIntCodeProgramWaitForTermination(program, inputChan, outputChan, terminateChan, nil)

	cameraView := buildCameraView(outputChan, terminateChan)

	// try to pathfind on scaffold to end
	fmt.Println(printCamera(cameraView))

	moves := solveMaze(cameraView)

	fmt.Println(moves)
}

func solveMaze(cameraView map[space]string) string {
	moves := ""
	moveLength := 0

	currentSpace := space{x: 0, y: 0}
	direction := 0 // directions are 0,1,2,3 for north,east,south,west

	// Find starting space
	for i, val := range cameraView {
		if val == "^" {
			currentSpace = i
		}
	}

	for {
		forwardSpace := getTileInDirection(cameraView, currentSpace, direction)
		leftSpace := getTileInDirection(cameraView, currentSpace, mod(direction-1, 4))
		rightSpace := getTileInDirection(cameraView, currentSpace, mod(direction+1, 4))

		// starting from the direciton we are facing, are we at the end?
		if cameraView[forwardSpace] == "." && cameraView[leftSpace] == "." && cameraView[rightSpace] == "." {
			moves += strconv.Itoa(moveLength)
			break
		}

		// Try to move forward
		if cameraView[forwardSpace] == "#" {
			moveLength++
			currentSpace = forwardSpace
			continue
		}

		// If not, try to turn and go forward
		if cameraView[leftSpace] == "#" {
			moves += strconv.Itoa(moveLength) + "," + "L" + ","
			moveLength = 1
			// Turn left
			direction = mod(direction-1, 4)
			currentSpace = leftSpace
		} else if cameraView[rightSpace] == "#" {
			moves += strconv.Itoa(moveLength) + "," + "R" + ","
			moveLength = 1
			// Turn Right
			direction = mod(direction+1, 4)
			currentSpace = rightSpace
		}
	}

	return moves
}

func getTileInDirection(cameraView map[space]string, currentSpace space, direction int) space {

	tile := space{}

	switch direction {
	case 0:
		tile = space{x: currentSpace.x, y: currentSpace.y - 1}
		break
	case 1:
		tile = space{x: currentSpace.x + 1, y: currentSpace.y}
		break
	case 2:
		tile = space{x: currentSpace.x, y: currentSpace.y + 1}
		break
	case 3:
		tile = space{x: currentSpace.x - 1, y: currentSpace.y}
		break
	}

	return tile

}

type space struct {
	x, y int
}

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
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

func convertToInts(input []string) []int {
	programStr := strings.Split(input[0], ",")
	var program []int

	for _, s := range programStr {
		i, _ := strconv.Atoi(s)
		program = append(program, i)
	}
	return program
}
func makeChannels() (chan int, chan int, chan bool, chan intcode.Message) {
	// First channel is for input
	inputChan := make(chan int)
	// Second channel is for output
	outputChan := make(chan int)
	// Terimnation channel
	terminateChan := make(chan bool)
	// Message channel
	messageChan := make(chan intcode.Message)

	return inputChan, outputChan, terminateChan, messageChan
}
