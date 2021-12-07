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

	var crabs []int
	for _, v := range strings.Split(input[0], ",") {
		num, _ := strconv.Atoi(v)
		crabs = append(crabs, num)
	}

	min := 10000
	max := 0
	for _, v := range crabs {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}

	minFuelNeeded := 1000000000
	for i := min; i < max; i++ {
		fuelNeeded := 0
		for _, v := range crabs {
			fuelNeeded += abs(i - v)
		}
		if minFuelNeeded > fuelNeeded {
			minFuelNeeded = fuelNeeded
		}
	}

	fmt.Println(minFuelNeeded)
}

func PartTwo() {
	input := readInput()

	var crabs []int
	for _, v := range strings.Split(input[0], ",") {
		num, _ := strconv.Atoi(v)
		crabs = append(crabs, num)
	}

	min := 10000
	max := 0
	for _, v := range crabs {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}

	minFuelNeeded := 1000000000
	for i := min; i < max; i++ {
		fuelNeeded := 0
		for _, v := range crabs {
			for j := 1; j < abs(i-v)+1; j++ {
				fuelNeeded += j
			}
		}
		if minFuelNeeded > fuelNeeded {
			minFuelNeeded = fuelNeeded
		}
	}

	fmt.Println(minFuelNeeded)
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

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
