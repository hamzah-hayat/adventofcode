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
		break
	case "p2":
		fmt.Println("Gold:" + PartTwo("input"))
		break
	}
}

func PartOne(filename string) string {
	input := readInput(filename)

	snailPairs := []NumberPair{}
	for _, v := range input {
		snailPairs = append(snailPairs, *CreateNumberPair(v))
	}

	totalPair := snailPairs[0]
	for i := 1; i < len(snailPairs); i++ {
		totalPair = AddPairs(totalPair, snailPairs[i])
		ProcessPairs(&totalPair)
	}
	score := ScorePair(&totalPair)
	num := strconv.Itoa(score)

	return num
}

func PartTwo(filename string) string {
	input := readInput(filename)

	num := strconv.Itoa(len(input))

	return num
}

func AddPairs(pairOne NumberPair, pairTwo NumberPair) NumberPair {
	return NumberPair{X: &pairOne, Y: &pairTwo}
}

func ProcessPairs(numPair *NumberPair) *NumberPair {
	actionTaken := true

	for actionTaken {
		actionTaken = false
		numPair, actionTaken = ExplodePairs(numPair)
		if actionTaken {
			continue
		}

		numPair, actionTaken = SplitPair(numPair)
		if actionTaken {
			continue
		}

	}

	return numPair
}

func ExplodePairs(numPair *NumberPair) (*NumberPair, bool) {
	explode := false

	return numPair, explode
}

func SplitPair(numPair *NumberPair) (*NumberPair, bool) {
	split := false

	if numPair.value == 0 && numPair.X != nil {
		// go down left
		_, split = SplitPair(numPair.X)
	}
	if numPair.value == 0 && numPair.Y != nil {
		// Go down right
		_, split = SplitPair(numPair.Y)
	}

	if numPair.value >= 10 {
		if numPair.value%2 == 0 {
			// Even number
			numPair.X = &NumberPair{value: numPair.value / 2}
			numPair.Y = &NumberPair{value: numPair.value / 2}
		} else {
			// Odd number
			numPair.X = &NumberPair{value: numPair.value / 2}
			numPair.Y = &NumberPair{value: (numPair.value / 2) + 1}
		}
		numPair.value = 0
		return numPair, true
	}
	return numPair, split
}

// Return the score of this pair
func ScorePair(numPair *NumberPair) int {

	currentValue := numPair.value

	if numPair.X != nil {
		// go down left
		currentValue += ScorePair(numPair.X) * 3
	}
	if numPair.Y != nil {
		// Go down right
		currentValue += ScorePair(numPair.Y) * 2
	}

	return currentValue
}

func CreateNumberPair(input string) *NumberPair {
	// Read the string and form number pairs
	numPair := NumberPair{}
	input = strings.TrimPrefix(strings.TrimSuffix(input, "]"), "[")

	if len(input) == 1 {
		//Just a regular value
		num, _ := strconv.Atoi(string(input[0]))
		numPair.value = num

	} else if len(input) == 3 {
		// two regular numbers
		num1, _ := strconv.Atoi(string(input[0]))
		num2, _ := strconv.Atoi(string(input[2]))
		return &NumberPair{X: &NumberPair{value: num1}, Y: &NumberPair{value: num2}}
	} else {
		// We need to find the inner comma to split on
		// count number of brackets
		leftBrackets := 0
		rightBrackets := 0

		splitCommaIndex := 0
		for i, c := range input {
			if string(c) == "]" {
				leftBrackets++
			} else if string(c) == "[" {
				rightBrackets++
			}
			if leftBrackets == rightBrackets {
				splitCommaIndex = i + 1
				break
			}
		}

		XString := input[:splitCommaIndex]
		YString := input[splitCommaIndex+1:]

		numPair.X = CreateNumberPair(XString)
		numPair.Y = CreateNumberPair(YString)
	}
	return &numPair
}

type NumberPair struct {
	X     *NumberPair
	Y     *NumberPair
	value int
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
