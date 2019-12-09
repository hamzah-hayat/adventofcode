package main

import (
	"bufio"
	"fmt"
	"flag"
	"github.com/hamzah-hayat/adventofcode/intcode"
	"os"
	"strconv"
	"strings"
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

	//intcode.RunIntCodeProgram()

	a :=3

	if(a==0){
		fmt.Println("zero")
	} else if a==1 {
		fmt.Println("two")
	}

}

func partTwo() {
	//input := readInput()
	//program := convertToInts(input)

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
