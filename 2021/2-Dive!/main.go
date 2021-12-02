package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	input := readInput()

	horizontal := 0
	depth := 0

	for _, v := range input {
		splitString := strings.Split(v, " ")
		num, _ := strconv.Atoi(splitString[1])
		switch splitString[0] {
		case "forward":
			horizontal += num
			break
		case "down":
			depth += num
			break
		case "up":
			depth -= num
			break
		}
	}

	fmt.Println(horizontal * depth)
}

func PartTwo() {
	input := readInput()

	horizontal := 0
	depth := 0
	aim := 0

	for _, v := range input {
		splitString := strings.Split(v, " ")
		num, _ := strconv.Atoi(splitString[1])
		switch splitString[0] {
		case "forward":
			horizontal += num
			depth += aim * num
			break
		case "down":
			aim += num
			break
		case "up":
			aim -= num
			break
		}
	}

	fmt.Println(horizontal * depth)
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
