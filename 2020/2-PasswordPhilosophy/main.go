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
	methodP := flag.String("method", "p1", "The method/part that should be run, valid are p1,p2 and test")
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
	input := readInput()

	validPasswords := 0

	for _, v := range input {
		// for each string we need
		// the lower value
		// the upper value
		// the needed letter
		// the password

		// format is lower-upper letter: passwordtext

		// use regex to make it easier
		// (\d*)-(\d*) (\w): (\w+)

		re := regexp.MustCompile(`(\d*)-(\d*) (\w): (\w+)`)
		match := re.FindAllStringSubmatch(v, -1)

		// Check our match
		lower, _ := strconv.Atoi(match[0][1])
		upper, _ := strconv.Atoi(match[0][2])
		letter := match[0][3]
		password := match[0][4]

		numletters := strings.Count(password, letter)

		if lower <= numletters && numletters <= upper {
			validPasswords++
		}

	}

	fmt.Println("The total number of valid passwords is ", validPasswords)

}

func PartTwo() {
	input := readInput()

	validPasswords := 0

	for _, v := range input {
		// for each string we need
		// the lower value
		// the upper value
		// the needed letter
		// the password

		// format is lower-upper letter: passwordtext

		// use regex to make it easier
		// (\d*)-(\d*) (\w): (\w+)

		re := regexp.MustCompile(`(\d*)-(\d*) (\w): (\w+)`)
		match := re.FindAllStringSubmatch(v, -1)

		// Check our match
		lower, _ := strconv.Atoi(match[0][1])
		upper, _ := strconv.Atoi(match[0][2])
		letter := match[0][3]
		password := match[0][4]

		if (password[lower-1] == letter[0]) != (password[upper-1] == letter[0]) {
			validPasswords++
		}

	}

	fmt.Println("The total number of valid passwords is ", validPasswords)
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
