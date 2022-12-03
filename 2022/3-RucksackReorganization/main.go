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

func parseFlags() {
	methodP = flag.String("method", "all", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()
}

func main() {

	parseFlags()

	switch *methodP {
	case "all":
		fmt.Println("Silver:" + PartOne("input"))
		fmt.Println("Gold:" + PartTwo("input"))
	case "p1":
		fmt.Println("Silver:" + PartOne("input"))
	case "p2":
		fmt.Println("Gold:" + PartTwo("input"))
	}
}

func PartOne(filename string) string {
	input := readInput(filename)

	totalSum := 0

	for _, v := range input {
		// split line in half
		firstCompartment := v[:len(v)/2]
		secondCompartment := v[len(v)/2:]

		// fmt.Println("first: " + firstCompartment)
		// fmt.Println("first: " + secondCompartment)
		// fmt.Println("----------------------------")

		// Make our sets for each compartment
		firstCompartmentSet := make(map[string]bool)
		secondCompartmentSet := make(map[string]bool)

		for _, item := range firstCompartment {
			firstCompartmentSet[string(item)] = true
		}

		for _, item := range secondCompartment {
			secondCompartmentSet[string(item)] = true
		}

		commonItem := ""
		// now find the common item
		for item := range firstCompartmentSet {
			if secondCompartmentSet[item] {
				commonItem = item
			}
		}

		totalSum += getPriority(commonItem)

	}

	return strconv.Itoa(totalSum)
}

func PartTwo(filename string) string {
	input := readInput(filename)

	totalSum := 0

	for line := 0; line < len(input); line += 3 {
		// check three bags for common item
		// Make our sets for each bag
		firstBagLine := input[line]
		secondBagLine := input[line+1]
		thirdBagLine := input[line+2]

		firstBagSet := make(map[string]bool)
		secondBagSet := make(map[string]bool)
		thirdBagSet := make(map[string]bool)

		for _, item := range firstBagLine {
			firstBagSet[string(item)] = true
		}

		for _, item := range secondBagLine {
			secondBagSet[string(item)] = true
		}

		for _, item := range thirdBagLine {
			thirdBagSet[string(item)] = true
		}

		badge := ""
		// Find the badge, common item in all three
		for item := range firstBagSet {
			if secondBagSet[item] && thirdBagSet[item] {
				badge = item
			}
		}

		totalSum += getPriority(badge)

	}

	return strconv.Itoa(totalSum)
}

// getPriority gets number from a string
// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
func getPriority(item string) int {
	valueArray := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(valueArray, item) + 1
}

// Read data from input.txt
// Return the string, so that we can deal with it however
func readInput(filename string) []string {

	var input []string

	f, _ := os.Open(filename + ".txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
