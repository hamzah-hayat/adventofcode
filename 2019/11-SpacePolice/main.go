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

	// First channel is for input
	inputChan := make(chan int)
	// Second channel is for output
	outputChan := make(chan int)
	// Terimnation channel
	t := make(chan bool)

	grid := make(map[space]int) // All spaces start as 0 implicity, aka black

	go intcode.RunIntCodeProgramWaitForTermination(program, inputChan, outputChan, t)

	// Run the robot
	painted := runRobotPainter(inputChan, outputChan, t, grid)

	paintedTiles := 0
	for _, val := range painted {
		if val {
			paintedTiles++
		}
	}

	fmt.Println("The number of painted tiles is", paintedTiles)
}

func runRobotPainter(inputChan chan int, outputChan chan int, t chan bool, grid map[space]int) map[space]bool {

	painted := make(map[space]bool) // The map of painted tiles, that were painted at least once by the robot

	x := 0
	y := 0
	d := 0

	teriminate := false
	for {

		select {
		case <-t:
			teriminate = true
			break
		case inputChan <- grid[space{x: x, y: y}]:
			paintOutput := <-outputChan
			moveOutput := <-outputChan

			// Paint the tile
			if paintOutput == 0 {
				grid[space{x: x, y: y}] = 0
				painted[space{x: x, y: y}] = true
			} else if paintOutput == 1 {
				grid[space{x: x, y: y}] = 1
				painted[space{x: x, y: y}] = true
			}

			// Then change direction and move
			if moveOutput == 0 {
				d = mod((d - 1), 4)
			} else if moveOutput == 1 {
				d = mod((d + 1), 4)
			}

			// Then move in that direction
			switch d {
			case 0:
				y++
				break
			case 1:
				x++
				break
			case 2:
				y--
				break
			case 3:
				x--
				break
			}

			if teriminate {
				break
			}
		}
		if teriminate {
			break
		}
	}
	return painted
}

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func partTwo() {
	input := readInput()
	program := convertToInts(input)

	// First channel is for input
	inputChan := make(chan int)
	// Second channel is for output
	outputChan := make(chan int)
	// Terimnation channel
	t := make(chan bool)

	grid := make(map[space]int) // All spaces start as 0 implicity, aka black

	go intcode.RunIntCodeProgramWaitForTermination(program, inputChan, outputChan, t)

	// Run the robot
	// Set inital tile to 1
	grid[space{x: 0, y: 0}] = 1
	paintedGrid := runRobotPainterReturnPaintedGrid(inputChan, outputChan, t, grid)

	// Find the xMin,xMax and ymin,yMax
	// Hardcode these for now
	xMin := 0
	xMax := 6
	yMin := 0
	yMax := 40

	// now print the grid

	picture := ""
	for i := xMin; i < xMax; i++ {
		for j := yMin; j < yMax; j++ {
			if (paintedGrid[space{x: i, y: j}] == 1) {
				picture += "#"
			} else {
				picture += "-"
			}
		}
		picture += "\n"
	}

	fmt.Println(picture)

}

func runRobotPainterReturnPaintedGrid(inputChan chan int, outputChan chan int, t chan bool, grid map[space]int) map[space]int {

	x := 0
	y := 0
	d := 0

	teriminate := false
	for {

		select {
		case <-t:
			teriminate = true
			break
		case inputChan <- grid[space{x: x, y: y}]:
			paintOutput := <-outputChan
			moveOutput := <-outputChan

			// Paint the tile
			if paintOutput == 0 {
				grid[space{x: x, y: y}] = 0
			} else if paintOutput == 1 {
				grid[space{x: x, y: y}] = 1
			}

			// Then change direction and move
			if moveOutput == 0 {
				d = mod((d - 1), 4)
			} else if moveOutput == 1 {
				d = mod((d + 1), 4)
			}

			// Then move in that direction
			switch d {
			case 0:
				x--
				break
			case 1:
				y++
				break
			case 2:
				x++
				break
			case 3:
				y--
				break
			}

			if teriminate {
				break
			}
		}
		if teriminate {
			break
		}
	}
	return grid
}

type space struct {
	x, y int
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
