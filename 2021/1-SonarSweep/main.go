package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	methodP *string
)

func init() {
	// Use Flags to run a part
	methodP = flag.String("method", "p1", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()
}

func main() {
	switch *methodP {
	case "p1":
		PartOne()
		break
	case "p2":
		PartTwo()
		break
	case "test":
		break
	}
}

func PartOne() {
	input := readInputInt()

	increases := 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increases++
		}
	}

	fmt.Println(increases)
}

func PartTwo() {
	input := readInputInt()

	increases := 0

	for i := 1; i < len(input)-2; i++ {
		val1 := input[i-1] + input[i] + input[i+1]
		val2 := input[i] + input[i+1] + input[i+2]
		if val2 > val1 {
			increases++
		}
	}

	fmt.Println(increases)
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

// Read data from input.txt
// Return the string, so that we can deal with it however
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
