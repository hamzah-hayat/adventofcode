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

	tractorView := buildTractorBeamView(program, 0, 0, 50, 50)

	tractorBeamTiles := countTractorBeamTiles(tractorView)

	fmt.Println("The number of tractor beam tiles is", tractorBeamTiles)

}

func buildTractorBeamView(program []int, startX, startY, xMax, yMax int) map[space]int {

	tractorBeamView := make(map[space]int)
	outputs := make(chan [3]int)

	leftX := startX
	for y := startY; y < yMax; y++ {
		fmt.Println(y)
		empty := false
		// find first beam location from left side
		for x := leftX; x < xMax; x++ {
			go tractorBeamBot(program, x, y, outputs)
			output := <-outputs
			if output[2] == 1 {
				// Found the leftX
				tractorBeamView[space{x: x, y: y}] = output[2]
				leftX = x
				break
			}
			if x == xMax-1 {
				empty = true
				break
			}
		}
		if empty {
			continue
		}

		for x := leftX + 1; x < xMax; x++ {
			go tractorBeamBot(program, x, y, outputs)
			output := <-outputs
			//tractorBeamView[space{x: x, y: y}] = output[2]
			if output[2] == 0 {
				// Found rightSide
				tractorBeamView[space{x: x - 1, y: y}] = 1
				break
			}
		}

	}

	return tractorBeamView
}

func tractorBeamBot(program []int, x, y int, out chan [3]int) {

	inputChan, outputChan, terminateChan, _ := makeChannels()

	go intcode.RunIntCodeProgramWaitForTermination(program, inputChan, outputChan, terminateChan, nil)

	inputChan <- x
	inputChan <- y

	output := <-outputChan

	outputInts := [3]int{x, y, output}

	out <- outputInts

}

func countTractorBeamTiles(tractorBeamView map[space]int) int {
	tractorCount := 0
	for _, t := range tractorBeamView {
		if t == 1 {
			tractorCount++
		}
	}
	return tractorCount
}

func printTractorBeam(tractorView map[space]int) string {
	tractorViewStr := ""

	// Find max x and y
	maxX := 0
	maxY := 0
	for s := range tractorView {
		if s.x > maxX {
			maxX = s.x + 1
		}
		if s.y > maxY {
			maxY = s.y + 1
		}
	}

	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			if (tractorView[space{x: i, y: j}] == 0) {
				tractorViewStr += "."
			} else {
				tractorViewStr += "#"
			}
		}
		tractorViewStr += "\n"
	}
	return tractorViewStr
}

func partTwo() {
	input := readInput()
	program := convertToInts(input)

	// tractorView := buildTractorBeamView(program, 0, 0, 10000, 10000)

	// startSpace := findSquareSize(tractorView, 100, 100)
	var startSpace space
	for i := 0; i < 10000; i = i + 1000 {
		tractorView := buildTractorBeamView(program, i, i, i+1000, i+1000)
		//fmt.Println(printTractorBeam(tractorView))
		startSpace = findSquareSize(tractorView, i, i, i+1000, i+1000, 100, 100)
		if startSpace.x != -1 && startSpace.y != -1 {
			break
		}
	}

	//fmt.Println(tractorView)

	//fmt.Println(printTractorBeam(tractorView))

	fmt.Println("The starting space for the square is", startSpace)
}

func findSquareSize(tractorView map[space]int, startx, startY, endX, endY, sizeX, sizeY int) space {

	for y := startY; y < endY; y++ {
		for x := startx; x < endX; x++ {
			// check topRight and bottomLeft corners
			topRight := tractorView[space{x: x + sizeX, y: y}]
			bottomLeft := tractorView[space{x: x, y: y + sizeY}]

			if topRight == 1 && bottomLeft == 1 {
				return space{x: x, y: y}
			}
		}
	}
	return space{x: -1, y: -1}
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
