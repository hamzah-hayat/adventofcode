package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/hamzah-hayat/adventofcode/intcode"
	"os"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p1", "The method/part that should be run, valid are p1,p2 and test")
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
	//input := readInput()

	program := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	//highest := RunRocketAmplifiers(program)

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	// Run the five go progams
	// check the output
	go intcode.RunIntCodeProgram(program, input, output)
	input <- 4
	input <- 0

	out := <-output

	go intcode.RunIntCodeProgram(program, input, output)
	input <- 4
	input <- out

	out = <-output

	go intcode.RunIntCodeProgram(program, input, output)
	input <- 4
	input <- out

	out = <-output

	go intcode.RunIntCodeProgram(program, input, output)
	input <- 4
	input <- out

	out = <-output

	go intcode.RunIntCodeProgram(program, input, output)
	input <- 4
	input <- out

	out = <-output

	fmt.Println("Highest output is", out)

}

// func partOne() {
// 	//input := readInput()

// 	program := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
// 	highest := RunRocketAmplifiers(program)

// 	fmt.Println("Highest output is", highest)

// }

func partTwo() {
	//input := readInput()
}

//RunRocketAmplifiers takes a program, runs it through 5 different amplifiers, then outputs the highest result
func RunRocketAmplifiers(program []int) int {
	highestResult := 0

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	// Run a nested for loop so that we can run each phase setting
	// Then run the program and check the result
	// find the highest result and return it

	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			for c := 0; c < 5; c++ {
				for d := 0; d < 5; d++ {
					for e := 0; e < 5; e++ {

						// Run the five go progams
						// check the output
						go intcode.RunIntCodeProgram(program, input, output)
						input <- a
						input <- 0

						out := <-output

						go intcode.RunIntCodeProgram(program, input, output)
						input <- b
						input <- out

						out = <-output

						go intcode.RunIntCodeProgram(program, input, output)
						input <- c
						input <- out

						out = <-output

						go intcode.RunIntCodeProgram(program, input, output)
						input <- d
						input <- out

						out = <-output

						go intcode.RunIntCodeProgram(program, input, output)
						input <- e
						input <- out

						out = <-output

						if out > highestResult {
							highestResult = out
						}
					}
				}
			}
		}
	}

	return highestResult
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
