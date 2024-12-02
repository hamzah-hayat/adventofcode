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
	case "p2":
		fmt.Println("Gold:" + PartTwo("input"))
	}
}

func PartOne(filename string) string {
	input := readInput(filename)

	safeReports := 0

	for _, v := range input {
		// A report is safe if each number differs by at least 1 and by at most 3
		split := strings.Split(v, " ")

		// Check if increasing or decreasing first
		startNum, _ := strconv.Atoi(split[0])
		firstNum, _ := strconv.Atoi(split[1])

		decreasing := false
		if startNum >= firstNum {
			decreasing = true
		}

		safe := false
		if decreasing {
			for i := 1; i < len(split); i++ {
				num, _ := strconv.Atoi(split[i])
				if num-startNum <= -1 && num-startNum >= -3 {
					startNum = num
				} else {
					break
				}
				if i == len(split)-1 {
					safe = true
					break
				}
			}
		} else {
			for i := 1; i < len(split); i++ {
				num, _ := strconv.Atoi(split[i])
				if num-startNum >= 1 && num-startNum <= 3 {
					startNum = num
				} else {
					break
				}
				if i == len(split)-1 {
					safe = true
					break
				}
			}
		}

		if safe {
			safeReports++
		}

	}

	return strconv.Itoa(safeReports)
}

func PartTwo(filename string) string {
	input := readInput(filename)

	safeReports := 0

	for _, v := range input {
		// A report is safe if each number differs by at least 1 and by at most 3
		split := strings.Split(v, " ")

		// Brute force and check every possible combination of the report
		// Including one missing digit in each position or no missing digits
		for i := 0; i < len(split)+1; i++ {
			newLine := make([]int, 0)
			for j, v := range split {
				if i == j {
					// Do nothing
				} else {
					num, _ := strconv.Atoi(v)
					newLine = append(newLine, num)
				}
			}
			if checkReport(newLine) {
				safeReports++
				break
			}
		}
	}

	return strconv.Itoa(safeReports)
}

func checkReport(input []int) bool {
	// Check if increasing or decreasing first
	startNum := input[0]
	firstNum := input[1]

	decreasing := false
	if startNum >= firstNum {
		decreasing = true
	}

	safe := false
	if decreasing {
		for i := 1; i < len(input); i++ {
			num := input[i]
			if num-startNum <= -1 && num-startNum >= -3 {
				startNum = num
			} else {
				break
			}
			if i == len(input)-1 {
				safe = true
				break
			}
		}
	} else {
		for i := 1; i < len(input); i++ {
			num := input[i]
			if num-startNum >= 1 && num-startNum <= 3 {
				startNum = num
			} else {
				break
			}
			if i == len(input)-1 {
				safe = true
				break
			}
		}
	}

	if safe {
		return true
	}
	return false
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
