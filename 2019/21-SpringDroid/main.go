package main

import (
	"bufio"
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/hamzah-hayat/adventofcode/intcode"
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
	//program := convertToInts(input)
}

func partTwo() {
	//input := readInput()
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
