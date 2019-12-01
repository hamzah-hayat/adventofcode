package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p1", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()

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

	tf := 0

	for _, value := range input {

		mass, _ := strconv.Atoi(value)
		if mass == 0 {
			continue
		}
		tf += int(mass/3) - 2
	}

	fmt.Println("Fuel needed is", tf)
}

func PartTwo() {
	input := readInput()

	tf := 0

	for _, value := range input {

		mass, _ := strconv.Atoi(value)

		f := int(mass/3) - 2
		tf += f

		for {
			mass = f
			fuelNeeded := int(mass/3) - 2
			if fuelNeeded <= 0 {
				break
			} else {
				f = fuelNeeded
				tf += f
			}
		}
	}

	fmt.Println("Fuel needed is", tf)
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
