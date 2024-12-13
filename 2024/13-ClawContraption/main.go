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

	totalTokens := 0
	buttonARegex := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	buttonBRegex := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	for i := 0; i < len(input); i = i + 4 {
		buttonA := input[0+i]
		buttonB := input[1+i]
		prize := input[2+i]

		buttonAMatch := buttonARegex.FindAllStringSubmatch(buttonA, -1)
		buttonBMatch := buttonBRegex.FindAllStringSubmatch(buttonB, -1)
		prizeMatch := prizeRegex.FindAllStringSubmatch(prize, -1)

		buttonA_X, _ := strconv.Atoi(buttonAMatch[0][1])
		buttonA_Y, _ := strconv.Atoi(buttonAMatch[0][2])

		buttonB_X, _ := strconv.Atoi(buttonBMatch[0][1])
		buttonB_Y, _ := strconv.Atoi(buttonBMatch[0][2])

		prize_X, _ := strconv.Atoi(prizeMatch[0][1])
		prize_Y, _ := strconv.Atoi(prizeMatch[0][2])

		// Try and find the correct button presses
		solutions := make([][]int, 0)
		for buttonAPresses := 0; buttonAPresses < 100; buttonAPresses++ {
			for buttonBPresses := 0; buttonBPresses < 100; buttonBPresses++ {
				totalX := buttonAPresses*buttonA_X + buttonBPresses*buttonB_X
				totalY := buttonAPresses*buttonA_Y + buttonBPresses*buttonB_Y
				if totalX == prize_X && totalY == prize_Y {
					solutions = append(solutions, []int{buttonAPresses, buttonBPresses})
				}
			}
		}

		if len(solutions) > 0 {
			smallest := 800
			for _, sol := range solutions {
				if sol[0]*3+sol[1] < smallest {
					smallest = sol[0]*3 + sol[1]
				}
			}

			totalTokens += smallest
		}
	}

	return strconv.Itoa(totalTokens)
}

func PartTwo(filename string) string {
	input := readInput(filename)

	totalTokens := 0
	buttonARegex := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	buttonBRegex := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	for i := 0; i < len(input); i = i + 4 {
		buttonA := input[0+i]
		buttonB := input[1+i]
		prize := input[2+i]

		buttonAMatch := buttonARegex.FindAllStringSubmatch(buttonA, -1)
		buttonBMatch := buttonBRegex.FindAllStringSubmatch(buttonB, -1)
		prizeMatch := prizeRegex.FindAllStringSubmatch(prize, -1)

		buttonA_X, _ := strconv.Atoi(buttonAMatch[0][1])
		buttonA_Y, _ := strconv.Atoi(buttonAMatch[0][2])

		buttonB_X, _ := strconv.Atoi(buttonBMatch[0][1])
		buttonB_Y, _ := strconv.Atoi(buttonBMatch[0][2])

		prize_X, _ := strconv.Atoi(prizeMatch[0][1])
		prize_Y, _ := strconv.Atoi(prizeMatch[0][2])

		// Oh lawd
		prize_X += 10000000000000
		prize_Y += 10000000000000

		// Try and find the correct button presses
		// Simultaneous equations
		// prize_X = buttonAPresses*buttonA_X + buttonBPresses*buttonB_X
		// prize_Y = buttonAPresses*buttonA_Y + buttonBPresses*buttonB_Y

		// It just works
		ANum := buttonB_Y*prize_X - prize_Y*buttonB_X
		ADen := buttonA_X*buttonB_Y - buttonA_Y*buttonB_X
		buttonAPresses, checkRemainder := ANum/ADen, ANum%ADen

		BNum := buttonA_X*prize_Y - prize_X*buttonA_Y
		BDen := buttonA_X*buttonB_Y - buttonA_Y*buttonB_X
		buttonBPresses, checkRemainder2 := BNum/BDen, BNum%BDen

		if checkRemainder != 0 || checkRemainder2 != 0 {
			// No match here
		} else {
			totalTokens += int(buttonAPresses*3) + int(buttonBPresses)
		}
	}

	return strconv.Itoa(int(totalTokens))
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
