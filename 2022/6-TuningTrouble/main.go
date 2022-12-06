package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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
	start := ""

	for i := 0; i < len(input[0])-3; i++ {
		// check four chars at a time
		firstChar := string(input[0][i])
		secondChar := string(input[0][i+1])
		thirdChar := string(input[0][i+2])
		fourthChar := string(input[0][i+3])

		if firstChar != secondChar && firstChar != thirdChar && firstChar != fourthChar && secondChar != thirdChar && secondChar != fourthChar && thirdChar != fourthChar {
			start = strconv.Itoa(i + 4)
			break
		}

	}

	return start
}

func PartTwo(filename string) string {
	input := readInput(filename)
	start := ""

	for i := 0; i < len(input[0])-13; i++ {

		// Check 14 characters at a time
		// grab a slice and make into set, if set and slice are same size each character is unique
		mapChars := make(map[string]bool)
		for charNum := i; charNum < i+13; charNum++ {
			mapChars[string(input[0][charNum])] = true
		}

		if len(mapChars) == 13 {
			start = strconv.Itoa(i + 14)
			break
		}

	}

	return start
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
