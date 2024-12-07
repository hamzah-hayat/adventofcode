package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mowshon/iterium"
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
	total := 0

	for _, v := range input {
		total += SolveEquation(v)
	}

	return strconv.Itoa(total)
}

// Find all possible solves of this equation
func SolveEquation(v string) int {
	total := 0
	// Input
	target := strings.Split(v, ":")[0]
	numbers := strings.Split(strings.TrimSpace(strings.SplitAfter(v, ":")[1]), " ")
	targetNum, _ := strconv.Atoi(target)
	numbersNums := make([]int, 0)
	for _, n := range numbers {
		num, _ := strconv.Atoi(n)
		numbersNums = append(numbersNums, num)
	}

	// Work out all possible combinations
	permutations := iterium.Product([]string{"+", "*","||"}, len(numbersNums)-1)
	permSlice, _ := permutations.Slice()

	for _, perm := range permSlice {
		currentTotal := numbersNums[0]
		for i, p := range perm {
			if p == "+" {
				currentTotal += numbersNums[i+1]
			} else if p == "*" {
				currentTotal *= numbersNums[i+1]
			} else if p == "||" {
				// Concat numbers together
				currentTotalStr := strconv.Itoa(currentTotal)
				numStr := strconv.Itoa(numbersNums[i+1])
				newStr := currentTotalStr + numStr

				// Back into int
				currentTotal, _ = strconv.Atoi(newStr)
			}
		}
		if currentTotal == targetNum {
			total = currentTotal
		}
	}

	return total
}

func PartTwo(filename string) string {
	input := readInput(filename)
	total := 0

	for _, v := range input {
		total += SolveEquation(v)
	}

	return strconv.Itoa(total)
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
