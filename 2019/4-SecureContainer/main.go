package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p2", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()

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
	//input := readInput()
	// Get Min and Max ranges for password
	input := "307237-769058"
	ranges := strings.Split(input, "-")
	rMin, _ := strconv.Atoi(ranges[0])
	rMax, _ := strconv.Atoi(ranges[1])

	possiblePasswords := 0

	for i := rMin; i < rMax; i++ {
		// Check each password for rules

		// Rules:
		// 1. six digits (all will be six digits already)
		// 2. Within range of rMin and rMax
		// 3. two adjacant digits are the same
		// 4. going from left to right, the digits never decrease, only increase

		if checkAdjacant(i) && checkAscending(i) {
			possiblePasswords++
		}
	}

	fmt.Println("The number of possible passwords is", possiblePasswords)
}

func checkAdjacant(i int) bool {
	passwordStr := strconv.Itoa(i)

	matched, _ := regexp.MatchString(`11|22|33|44|55|66|77|88|99|00`, passwordStr)
	if !matched {
		return false
	}

	return true
}

func checkAscending(i int) bool {
	passwordStr := strconv.Itoa(i)

	lowest := 0
	for _, value := range passwordStr {
		num, _ := strconv.Atoi(string(value))
		if num < lowest {
			return false
		} else {
			lowest = num
		}
	}

	return true
}

func PartTwo() {
	//input := readInput()
	// Get Min and Max ranges for password
	input := "307237-769058"
	ranges := strings.Split(input, "-")
	rMin, _ := strconv.Atoi(ranges[0])
	rMax, _ := strconv.Atoi(ranges[1])

	possiblePasswords := 0

	for i := rMin; i < rMax; i++ {
		// Check each password for rules

		// Rules:
		// 1. six digits (all will be six digits already)
		// 2. Within range of rMin and rMax
		// 3. two adjacant digits are the same (however, not part of larger group of digits)
		// 4. going from left to right, the digits never decrease, only increase

		if checkAdjacantWithoutLargerGroup(i) && checkAscending(i) {
			possiblePasswords++
		}
	}

	fmt.Println("The number of possible passwords is", possiblePasswords)
}

func checkAdjacantWithoutLargerGroup(i int) bool {
	passwordStr := strconv.Itoa(i)
	// Have to add the X for regex usage
	passwordStr = "X" + passwordStr + "X"

	regexTesters1 := [10]string{"[^1]11[^1]", "[^2]22[^2]", "[^3]33[^3]", "[^4]44[^4]", "[^5]55[^5]", "[^6]66[^6]", "[^7]77[^7]", "[^8]88[^8]", "[^9]99[^9]", "[^0]00[^0]"}

	for _, test := range regexTesters1 {
		match1, _ := regexp.MatchString(test, passwordStr)

		if match1 {
			return true
		}
	}
	return false
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
