package main

import (
	"bufio"
	"cmp"
	"flag"
	"fmt"
	"os"
	"slices"
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

	stripes := strings.Split(input[0], ",")
	for i := 0; i < len(stripes); i++ {
		stripes[i] = strings.TrimSpace(stripes[i])
	}

	slices.SortFunc(stripes,
		func(a, b string) int {
			return cmp.Compare(len(b), len(a))
		})

	totalPossibleTowels := 0
	for towel := 2; towel < len(input); towel++ {
		if TowelIsPossible(input[towel], stripes) {
			totalPossibleTowels++
		}
	}

	return strconv.Itoa(totalPossibleTowels)
}

func TowelIsPossible(towel string, stripes []string) bool {

	// Try and build the towel using the stripes
	possibleTowels := make(map[string]bool)
	possibleTowels[towel] = true
	towelsLength := 1
	for {
		for pt := range possibleTowels {
			for _, s := range stripes {
				if len(s) <= len(pt) {
					if s == pt[:len(s)] {
						possibleTowels[pt[len(s):]] = true
					}
				}
			}
		}

		// Did we get more possible towels?
		if len(possibleTowels) == towelsLength {
			return false
		} else {
			towelsLength = len(possibleTowels)
		}

		// Did we finish?
		for t := range possibleTowels {
			if len(t) == 0 {
				return true
			}
		}
	}
}

func PartTwo(filename string) string {
	input := readInput(filename)

	stripes := strings.Split(input[0], ",")
	for i := 0; i < len(stripes); i++ {
		stripes[i] = strings.TrimSpace(stripes[i])
	}

	slices.SortFunc(stripes,
		func(a, b string) int {
			return cmp.Compare(len(b), len(a))
		})

	totalPossibleTowels := 0
	for towel := 2; towel < len(input); towel++ {
		totalPossibleTowels += AllTowelIsPossible(input[towel], stripes)
	}

	return strconv.Itoa(totalPossibleTowels)
}

func AllTowelIsPossible(towel string, stripes []string) int {

	// Try and build the towel using the stripes
	possibleTowels := make(map[string]int)
	possibleTowels[""] = 1
	for {
		for pt := range possibleTowels {
			for _, s := range stripes {
				if len(pt)+len(s) <= len(towel) {
					if pt+s == towel[:len(pt+s)] {
						newTowel := pt + s
						possibleTowels[newTowel] += possibleTowels[pt]
					}
				}
			}
			if pt != towel {
				delete(possibleTowels, pt)
			}
		}

		// Are we done (no Solution)
		if len(possibleTowels) == 0 {
			return 0
		}

		// Did we finish?
		if len(possibleTowels) == 1 && possibleTowels[towel] > 0 {
			return possibleTowels[towel]
		}
	}
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

// Read data from input.txt
// Return the string as int
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
