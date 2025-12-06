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
	mathTotal := 0

	numberRegex := regexp.MustCompile(`(\d+)`)
	signRegex := regexp.MustCompile(`(\+|\*)`)

	// Find out how many problems we have
	firstLine := input[0]
	numbers := numberRegex.FindAllString(firstLine, -1)

	problems := []MathProblem{}
	for i := 0; i < len(numbers); i++ {
		problems = append(problems, MathProblem{})
	}

	for _, lines := range input {
		if strings.Contains(lines, "+") {
			// Final line
			signs := signRegex.FindAllString(lines, -1)
			for i := 0; i < len(signs); i++ {
				switch signs[i] {
				case "+":
					problems[i].operation = "+"
				case "*":
					problems[i].operation = "*"
				}
			}
			break
		}

		numbers := numberRegex.FindAllString(lines, -1)
		for i := 0; i < len(numbers); i++ {
			numberInt, _ := strconv.Atoi(numbers[i])
			problems[i].numbers = append(problems[i].numbers, numberInt)
		}

	}

	// Solve each problem
	for _, p := range problems {
		mathTotal += solveProblem(p)
	}

	return strconv.Itoa(mathTotal)
}

func PartTwo(filename string) string {
	input := readInput(filename)
	mathTotal := 0

	numberRegex := regexp.MustCompile(`(\d+)`)
	signRegex := regexp.MustCompile(`(\+|\*)`)

	// key is index, string is the number
	numberStringMap := make(map[int]string)

	// Find out how many problems we have
	firstLine := input[0]
	numbers := numberRegex.FindAllString(firstLine, -1)

	problems := []MathProblem{}
	for i := 0; i < len(numbers); i++ {
		problems = append(problems, MathProblem{})
	}

	for _, lines := range input {
		if strings.Contains(lines, "+") {
			// Final line
			signs := signRegex.FindAllString(lines, -1)
			for i := 0; i < len(signs); i++ {
				switch signs[i] {
				case "+":
					problems[i].operation = "+"
				case "*":
					problems[i].operation = "*"
				}
			}
			break
		}

		currentProblem := 0
		currentProblemNums := 0
		for i := 0; i < len(lines); i++ {
			if columnEmpty(i, input) {
				problems[currentProblem].totalNums = currentProblemNums
				currentProblem++
				currentProblemNums = 0
			} else if lines[i] != ' ' {
				currentProblemNums++
				numberStringMap[i] += string(lines[i])
			} else {
				currentProblemNums++
			}
		}
	}

	maxNumber := 0
	for i, _ := range numberStringMap {
		if i > maxNumber {
			maxNumber = i
		}
	}

	currentProblem := 0
	currentProblemNums := problems[0].totalNums
	for i := 0; i <= maxNumber; i++ {
		num, _ := strconv.Atoi(numberStringMap[i])
		if currentProblemNums == 0 {
			currentProblem++
			currentProblemNums = problems[currentProblem].totalNums
		} else {
			currentProblemNums--
			problems[currentProblem].numbers = append(problems[currentProblem].numbers, num)
		}
	}

	// Solve each problem
	for _, p := range problems {
		mathTotal += solveProblem(p)
	}

	return strconv.Itoa(mathTotal)
}

func columnEmpty(i int, input []string) bool {
	for _, lines := range input {
		if lines[i] != ' ' {
			return false
		}
	}
	return true
}

type MathProblem struct {
	numbers   []int
	operation string
	totalNums int
}

func solveProblem(p MathProblem) int {
	total := 0
	if p.operation == "+" {
		for _, n := range p.numbers {
			total += n
		}
	} else if p.operation == "*" {
		total = 1
		for _, n := range p.numbers {
			total *= n
		}
	}
	return total
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
