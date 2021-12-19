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

	num := strconv.Itoa(len(input))

	snailPairs := []NumberPair{}
	for _, v := range input {
		snailPairs = append(snailPairs, *CreateNumberPair(v))
	}

	return num
}

func PartTwo(filename string) string {
	input := readInput(filename)

	num := strconv.Itoa(len(input))

	return num
}

func ProcessPairs(numPair NumberPair) NumberPair {

	return numPair
}

func ExplodePairs(numPair *NumberPair, nested int) *NumberPair {

	// Look for any pairs that are nested 4 times deep
	// If there are, explode it
	if nested != 4 {
		numPair.X = ExplodePairs(numPair.X, nested+1)
		numPair.Y = ExplodePairs(numPair.Y, nested+1)
	} else {
		// Once we are 4 in, explode
		// First we add the X value to the pair to the left of us
		// Then add Y value to the pair to the right of us
		// Then replace this numberPair with 0 value

	}

	return numPair
}

func SplitPair(numPair *NumberPair) *NumberPair {
	if numPair.value == 0 && numPair.X != nil {
		// go down left
		SplitPair(numPair.X)
	}
	if numPair.value == 0 && numPair.Y != nil {
		// Go down right
		SplitPair(numPair.Y)
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
	}
	return numPair
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

		// leftBracketCountregex := regexp.MustCompile(`\[+`)
		// leftBrackets := len(leftBracketCountregex.FindString(input))
		// rightBracketCountregex := regexp.MustCompile(`\]+`)
		// rightBrackets := len(rightBracketCountregex.FindString(Reverse(input)))
		leftBrackets := 0
		rightBrackets := 0

		// split left
		splitCommaIndex := 0
		for i, c := range input {
			if string(c) == "]" {
				leftBrackets++
			} else if string(c)=="["{
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

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
