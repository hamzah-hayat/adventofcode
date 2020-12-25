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
	input := readInputInt()

	found := false

	for _, v := range input {
		for _, v2 := range input {

			if v2+v == 2020 {
				fmt.Println("The values are ", v, " and ", v2, ", which when multiplied together produce ", v*v2)
				found = true
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}
}

func PartTwo() {
	input := readInputInt()

	found := false

	for _, v := range input {
		for _, v2 := range input {
			for _, v3 := range input {

				if v3+v2+v == 2020 {
					fmt.Println("The values are ", v, ", ", v2, " and ", v3, " which when multiplied together produce ", v*v2*v3)
					found = true
				}
				if found {
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}
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
