package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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

	dial := 50
	password := 0
	regexp := regexp.MustCompile(`([L|R])(\d+)`)

	for _, command := range input {

		regexSplit := regexp.FindStringSubmatch(command)

		direction := regexSplit[1]
		distance, _ := strconv.Atoi(regexSplit[2])

		if direction == "L" {
			dial -= distance
			dial = dial % 100
		} else if direction == "R" {
			dial += distance
			dial = dial % 100
		}

		if dial == 0 {
			password++
		}
	}

	return strconv.Itoa(password)
}

func PartTwo(filename string) string {
	input := readInput(filename)

	dial := 50
	password := 0
	regexp := regexp.MustCompile(`([L|R])(\d+)`)

	for _, command := range input {

		regexSplit := regexp.FindStringSubmatch(command)

		direction := regexSplit[1]
		distance, _ := strconv.Atoi(regexSplit[2])

		if direction == "L" {
			for i := 0; i < distance; i++ {
				dial--
				if dial == 0 {
					password++
				}
				if dial < 0 {
					dial += 100
				}
			}
		} else if direction == "R" {
			for i := 0; i < distance; i++ {
				dial++
				if dial == 100 {
					password++
					dial -= 100
				}
			}
		}
	}

	return strconv.Itoa(password)
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
