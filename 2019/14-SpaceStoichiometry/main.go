package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p2", "The method/part that should be run, valid are p1,p2 and test")
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
	input := readInput()

	// Create reactions list
	reactionList := createReactions(input)

	oreNeededForFuel := convertOreToFuel(reactionList)

	fmt.Println("The amount of ore needed for 1 fuel is", oreNeededForFuel)
}

func createReactions(input []string) []reaction {
	var reactionList []reaction

	for _, val := range input {

		// Turn
		// 10 ORE => 10 A
		// 1 ORE => 1 B
		// 7 A, 1 B => 1 C
		// 7 A, 1 C => 1 D
		// 7 A, 1 D => 1 E
		// 7 A, 1 E => 1 FUEL
		// into reactions

		splitReaction := strings.Split(val, "=>")

		// Get the ingredients
		var ingredients []chemical
		for _, valI := range strings.Split(splitReaction[0], ",") {
			trimed := strings.Trim(valI, " ")
			finalSplit := strings.Split(trimed, " ")
			i, _ := strconv.Atoi(finalSplit[0])
			chemName := finalSplit[1]
			ingredients = append(ingredients, chemical{number: i, name: chemName})
		}

		// Get the results
		var results []chemical
		for _, valI := range strings.Split(splitReaction[1], ",") {
			trimed := strings.Trim(valI, " ")
			finalSplit := strings.Split(trimed, " ")
			i, _ := strconv.Atoi(finalSplit[0])
			chemName := finalSplit[1]
			results = append(results, chemical{number: i, name: chemName})
		}

		reactionList = append(reactionList, reaction{ingredients: ingredients, results: results})

	}

	return reactionList
}

func convertOreToFuel(reactionList []reaction) int {
	oreNeeded := 0

	return oreNeeded
}

func partTwo() {
	//input := readInput()
}

type reaction struct {
	ingredients []chemical
	results     []chemical
}

type chemical struct {
	number int
	name   string
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
