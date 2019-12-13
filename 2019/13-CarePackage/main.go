package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

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

	go intcode.RunIntCodeProgramWaitForTermination(program, inputChan, outputChan, t)

	screen := createGameScreen(outputChan, t)

	count := 0
	for _, val := range screen {
		if val == tIDBlock {
			count++
		}
	}

	fmt.Println("The number of block tiles is", count)
}

func createGameScreen(outputChan chan int, t chan bool) map[space]int {

	screen := make(map[space]int)

	teriminate := false
	for {
		select {
		case <-t:
			teriminate = true
			break
		case x := <-outputChan:
			y := <-outputChan
			tileID := <-outputChan

			screen[space{x: x, y: y}] = tileID
		}

		if teriminate {
			break
		}

	}

	return screen
}

type space struct {
	x, y int
}

const (
	tIDEmpty  int = 0 // 0 is an empty tile. No game object appears in this tile.
	tIDWall   int = 1 // 1 is a wall tile. Walls are indestructible barriers.
	tIDBlock  int = 2 // 2 is a block tile. Blocks can be broken by the ball.
	tIDPaddle int = 3 // 3 is a horizontal paddle tile. The paddle is indestructible.
	tIDBall   int = 4 // 4 is a ball tile. The ball moves diagonally and bounces off objects.
)

func partTwo() {
	input := readInput()
	program := convertToInts(input)

	// First channel is for input
	inputChan := make(chan int)
	// Second channel is for output
	outputChan := make(chan int)
	// Terimnation channel
	t := make(chan bool)

	// Put some money in!
	program[0] = 2

	go intcode.RunIntCodeProgramWaitForTermination(program, inputChan, outputChan, t)

	score := playGame(inputChan, outputChan, t)

	fmt.Println("The score is", score)

}

func playGame(inputChan chan int, outputChan chan int, t chan bool) int {

	screen := make(map[space]int)
	teriminate := false
	ballAndPaddleReady := false
	score := 0

	ballX := -1
	paddleX := -1

	for {
		select {
		case <-t:
			teriminate = true
			break
		case x := <-outputChan:
			y := <-outputChan
			tileID := <-outputChan

			if ballAndPaddleReady {
				cmd := exec.Command("cmd", "/c", "cls")
				cmd.Stdout = os.Stdout
				cmd.Run()
				fmt.Println(printScreen(screen))
				time.Sleep(time.Millisecond * 50)
				// Now move towards the ball
				if ballX < paddleX {
					inputChan <- -1
				} else if ballX > paddleX {
					inputChan <- 1
				} else {
					inputChan <- 0
				}
			}

			if x == -1 && y == 0 {
				score = tileID
				continue
			}

			screen[space{x: x, y: y}] = tileID

			if tileID == tIDBall {
				ballX = x
			}
			if tileID == tIDPaddle {
				paddleX = x
			}

			if ballX != -1 && paddleX != -1 && !ballAndPaddleReady {
				ballAndPaddleReady = true
			}
		}

		if teriminate {
			break
		}

	}

	return score
}

func findObject(screen map[space]int, tID int) space {

	for s, val := range screen {
		if val == tID {
			return s
		}
	}

	// Cant find this object
	return space{x: -1, y: -1}
}

func printScreen(screen map[space]int) string {
	// Find the xMin,xMax and ymin,yMax
	// Hardcode these for now
	xMin := 0
	xMax := 23
	yMin := 0
	yMax := 50

	// now print the grid
	picture := ""
	for i := xMin; i < xMax; i++ {
		for j := yMin; j < yMax; j++ {
			switch screen[space{x: j, y: i}] {
			case 0:
				picture += " "
			case 1:
				picture += "|"
			case 2:
				picture += "█"
			case 3:
				picture += "_"
			case 4:
				picture += "O"
			}
		}
		picture += "\n"
	}

	return picture
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
