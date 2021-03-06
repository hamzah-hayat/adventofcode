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

	oreNeededForFuel := convertOreToFuel(reactionList, 1)

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
		ingredients := make(map[string]int)
		for _, valI := range strings.Split(splitReaction[0], ",") {
			trimed := strings.Trim(valI, " ")
			finalSplit := strings.Split(trimed, " ")
			i, _ := strconv.Atoi(finalSplit[0])
			chemName := finalSplit[1]
			ingredients[chemName] = i
		}

		// Get the results
		results := make(map[string]int)
		for _, valI := range strings.Split(splitReaction[1], ",") {
			trimed := strings.Trim(valI, " ")
			finalSplit := strings.Split(trimed, " ")
			i, _ := strconv.Atoi(finalSplit[0])
			chemName := finalSplit[1]
			results[chemName] = i
		}

		reactionList = append(reactionList, reaction{ingredients: ingredients, results: results})

	}

	return reactionList
}

func convertOreToFuel(reactionList []reaction, fuelNeeded int) int {

	// start with what I want, aka a set amount of fuel
	wantedChemicals, _ := chemicalsNeededForResult(reactionList, fuelNeeded, "FUEL")

	// now, work backwards and find the wanted chems for each method, untill we only have ore
	for {
		changed := false
		for chemName, chemNeeded := range wantedChemicals {
			if chemName == "ORE" || chemNeeded < 0 {
				continue
			}
			newWantedChemicals, chemsOutputted := chemicalsNeededForResult(reactionList, chemNeeded, chemName)
			newWantedChemicals[chemName] = -chemsOutputted

			// Merge this new list with our old list
			wantedChemicals = mergeMaps(wantedChemicals, newWantedChemicals)
			if wantedChemicals[chemName] == 0 {
				delete(wantedChemicals, chemName)
			}
			changed = true
		}
		if !changed {
			break
		}
	}

	return wantedChemicals["ORE"]
}

func chemicalsNeededForResult(reactionList []reaction, numNeeded int, chemName string) (map[string]int, int) {
	wantedChemicals := make(map[string]int)
	chemsProduced := 0
	for _, val := range reactionList {
		for i, val2 := range val.results {
			if i == chemName {
				// Check how many times we need this
				multiplier := numNeeded / val2
				if numNeeded%val2 > 0 {
					multiplier++
				}
				chemsProduced = multiplier * val2
				for i2, val3 := range val.ingredients {
					wantedChemicals[i2] = val3 * multiplier
				}
			}
		}

	}
	return wantedChemicals, chemsProduced
}

func mergeMaps(a map[string]int, b map[string]int) map[string]int {
	newMap := make(map[string]int)
	for k, v := range a {
		newMap[k] += v
	}
	for k, v := range b {
		newMap[k] += v
	}
	return newMap
}

func partTwo() {
	input := readInput()

	// Create reactions list
	reactionList := createReactions(input)

	oreNeededForFuel := 0
	fuelTest := 2876992
	for {
		oreNeededForFuel = convertOreToFuel(reactionList, fuelTest)
		// Found an upper limit
		if oreNeededForFuel > 1000000000000 {
			fuelTest--
			break
		}
		fuelTest++
		fmt.Println("Ore needed is", oreNeededForFuel)
		fmt.Println("Fuel Produced is", fuelTest)
	}

	fmt.Println("The amount of fuel that can be created with a trillion ore is", fuelTest)
}

type reaction struct {
	ingredients map[string]int
	results     map[string]int
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
