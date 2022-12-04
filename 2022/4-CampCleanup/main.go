package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"reflect"
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

	totalPairs := 0

	for _, v := range input {
		// get both assignments
		assignments := strings.Split(v, ",")
		firstElf := assignments[0]
		secondElf := assignments[1]

		firstElfAssignments := strings.Split(firstElf, "-")
		firstElfStart, _ := strconv.Atoi(firstElfAssignments[0])
		firstElfEnd, _ := strconv.Atoi(firstElfAssignments[1])
		firstElfMap := make(map[int]bool)
		for i := firstElfStart; i < firstElfEnd+1; i++ {
			firstElfMap[i] = true
		}

		secondElfAssignments := strings.Split(secondElf, "-")
		secondElfStart, _ := strconv.Atoi(secondElfAssignments[0])
		secondElfEnd, _ := strconv.Atoi(secondElfAssignments[1])
		secondElfMap := make(map[int]bool)
		for i := secondElfStart; i < secondElfEnd+1; i++ {
			secondElfMap[i] = true
		}

		// Now get intersection
		intersectionMap := make(map[int]bool)
		for section := range firstElfMap {
			if secondElfMap[section] {
				intersectionMap[section] = true
			}
		}

		// check if it is equal to either of the elfs assignments
		if reflect.DeepEqual(intersectionMap, firstElfMap) || reflect.DeepEqual(intersectionMap, secondElfMap) {
			totalPairs++
		}
	}

	return strconv.Itoa(totalPairs)
}

func PartTwo(filename string) string {
	input := readInput(filename)

	totalPairs := 0

	for _, v := range input {
		// get both assignments
		assignments := strings.Split(v, ",")
		firstElf := assignments[0]
		secondElf := assignments[1]

		firstElfAssignments := strings.Split(firstElf, "-")
		firstElfStart, _ := strconv.Atoi(firstElfAssignments[0])
		firstElfEnd, _ := strconv.Atoi(firstElfAssignments[1])
		firstElfMap := make(map[int]bool)
		for i := firstElfStart; i < firstElfEnd+1; i++ {
			firstElfMap[i] = true
		}

		secondElfAssignments := strings.Split(secondElf, "-")
		secondElfStart, _ := strconv.Atoi(secondElfAssignments[0])
		secondElfEnd, _ := strconv.Atoi(secondElfAssignments[1])
		secondElfMap := make(map[int]bool)
		for i := secondElfStart; i < secondElfEnd+1; i++ {
			secondElfMap[i] = true
		}

		// Now get intersection
		intersectionMap := make(map[int]bool)
		for section := range firstElfMap {
			if secondElfMap[section] {
				intersectionMap[section] = true
			}
		}

		// check if we have any overlap
		if len(intersectionMap) > 0 {
			totalPairs++
		}
	}

	return strconv.Itoa(totalPairs)
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
