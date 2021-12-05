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

func init() {
	// Use Flags to run a part
	methodP = flag.String("method", "p2", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()
}

func main() {
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
	input := readInput()

	numBits := len(input[0])
	sum := make([]int, numBits)

	for _, lineInput := range input {
		for i, binaryVal := range lineInput {
			if string(binaryVal) == "1" {
				sum[i] += 1
			}
		}
	}

	gammaRate := ""
	epsilonRate := ""
	for _, v := range sum {
		if v > len(input)/2 {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	gammaRateInt, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilonRateInt, _ := strconv.ParseInt(epsilonRate, 2, 64)

	fmt.Println(gammaRateInt * epsilonRateInt)
}

func PartTwo() {
	input := readInput()

	// Find most common bit, then move right and discard from input
	var currentList []string
	for _, v := range input {
		currentList = append(currentList, v)
	}

	numBits := len(input[0])
	currentBit := 0

	for len(currentList) != 1 {
		sum := make([]int, numBits)

		for _, lineInput := range currentList {
			for i, binaryVal := range lineInput {
				if string(binaryVal) == "1" {
					sum[i] += 1
				}
			}
		}
		mostCommonIsOne := false
		if float64(sum[currentBit]) >= float64(len(currentList))/2 {
			mostCommonIsOne = true
		}

		// Now find most common bit (oxygen)
		// remove wrong lines
		for i := len(currentList) - 1; i >= 0; i-- {
			checkBit := currentList[i][currentBit]
			if mostCommonIsOne && string(checkBit) == "0" {
				currentList = remove(currentList, i)
				continue
			} else if !mostCommonIsOne && string(checkBit) == "1" {
				currentList = remove(currentList, i)
				continue
			}
		}

		currentBit++
	}

	oxygenGenRating, _ := strconv.ParseInt(currentList[0], 2, 64)

	currentBit = 0
	var currentList2 []string
	for _, v := range input {
		currentList2 = append(currentList2, v)
	}

	for len(currentList2) != 1 {
		sum := make([]int, numBits)

		for _, lineInput := range currentList2 {
			for i, binaryVal := range lineInput {
				if string(binaryVal) == "1" {
					sum[i] += 1
				}
			}
		}
		mostCommonIsOne := false
		if float64(sum[currentBit]) >= float64(len(currentList2))/2 {
			mostCommonIsOne = true
		}

		// Now find most common bit (oxygen)
		// remove wrong lines
		for i := len(currentList2) - 1; i >= 0; i-- {
			checkBit := currentList2[i][currentBit]
			if mostCommonIsOne && string(checkBit) == "1" {
				currentList2 = remove(currentList2, i)
				continue
			} else if !mostCommonIsOne && string(checkBit) == "0" {
				currentList2 = remove(currentList2, i)
				continue
			}
		}

		currentBit++
	}

	c02ScrubberRating, _ := strconv.ParseInt(currentList2[0], 2, 64)

	fmt.Println(oxygenGenRating * c02ScrubberRating)

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

// remove element from array
func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
