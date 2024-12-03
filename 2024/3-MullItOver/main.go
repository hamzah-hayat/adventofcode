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

	r, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	r2, _ := regexp.Compile(`([0-9]+),([0-9]+)`)

	total := 0
	for _, v := range input {
		matches := r.FindAllString(v, -1)

		for _, mul := range matches {
			mulMatch := r2.FindAllStringSubmatch(mul, -1)

			num1, _ := strconv.Atoi(mulMatch[0][1])
			num2, _ := strconv.Atoi(mulMatch[0][2])

			total += num1 * num2
		}
	}

	return strconv.Itoa(total)
}

func PartTwo(filename string) string {
	input := readInput(filename)

	r, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
	r2, _ := regexp.Compile(`([0-9]+),([0-9]+)`)

	total := 0
	mulEnabled := true
	for _, v := range input {
		matches := r.FindAllString(v, -1)

		for _, mul := range matches {

			if mul=="do()" {
				mulEnabled = true
				continue
			}

			if mul=="don't()" {
				mulEnabled = false
				continue
			}

			if mulEnabled {
				mulMatch := r2.FindAllStringSubmatch(mul, -1)

				num1, _ := strconv.Atoi(mulMatch[0][1])
				num2, _ := strconv.Atoi(mulMatch[0][2])
	
				total += num1 * num2
			}
		}
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
