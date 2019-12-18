package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

	maze := createMaze(input)

	lowestSteps := solveMaze(maze)

	fmt.Println("The lowest number of steps to solve the maze is", lowestSteps)

}

func createMaze(input []string) map[space]mazeObject {
	maze := make(map[space]mazeObject)

	x :=0
	y := 0
	for _, mazeLine := range input {
		for _, c := range mazeLine {
			
		}
		
	}

	return maze
}

func solveMaze(maze map[space]mazeObject) int {
	lowestSteps := 0

	return lowestSteps
}

func partTwo() {
	//input := readInput()
}

type space struct {
	x, y int
}

type mazeObject struct {
	name       string // used when this is a key or door
	objectType int    // 0 is empty space, 1 is wall, 2 is entrance, 3 is key, 4 is door
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
