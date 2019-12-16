package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p1", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()

	switch *methodP {
	case "p1":
		partOne()
		break
	case "p2":
		partTwo()
		break
	case "test":
		break
	}
}

func partOne() {
	input := readInput()
	phase := convertToInts(input)

	outputNumbers := runPhases(phase, 100)

	fmt.Println("The first eight digits of the 100th phase are:", string(outputNumbers[99][0:8]))
}

func runPhases(phase []int, runAmount int) []string {
	output := make([]string, runAmount)
	for run := 0; run < runAmount; run++ {
		newPhase := make([]int, len(phase))
		for i := 0; i < len(phase); i++ {
			// Build Base Pattern
			basePattern := buildBasePattern(i + 1)
			// Ignore the first element of the pattern
			patternNum := 1
			outputNum := 0
			for _, val := range phase {
				outputNum += val * basePattern[patternNum%len(basePattern)] // multiply with pattern number
				patternNum++
			}
			outputNum = abs(outputNum) % 10 //make positive and single digit
			newPhase[i] = outputNum
		}

		o := ""
		for _, val := range newPhase {
			o += strconv.Itoa(val)
		}
		output[run] = o
		phase = newPhase
	}

	return output
}

func buildBasePattern(runAmount int) []int {
	newBasePattern := make([]int, 4*runAmount)
	for i := 0; i < runAmount; i++ {
		newBasePattern[i] = 0
	}
	for i := runAmount; i < runAmount*2; i++ {
		newBasePattern[i] = 1
	}
	for i := runAmount * 2; i < runAmount*3; i++ {
		newBasePattern[i] = 0
	}
	for i := runAmount * 3; i < runAmount*4; i++ {
		newBasePattern[i] = -1
	}

	return newBasePattern
}

func partTwo() {
	input := readInput()
	phase := convertToInts(input)

	phase = multiplyPhases(phase, 10000)
	offset := getPhaseOffset(phase)

	outputNumbers := runPhases(phase, 100)

	fmt.Println("The first eight digits of the 100th phase (with offset) are:", string(outputNumbers[99][0+offset:8+offset]))
}

func multiplyPhases(phases []int, multiplyNum int) []int {
	newPhases := make([]int, len(phases)*multiplyNum)
	for i := 0; i < multiplyNum; i++ {
		for j := 0; j < len(phases); j++ {
			newPhases[j*(i+1)] = phases[j]
		}
	}
	return newPhases
}

func getPhaseOffset(phases []int) int {
	str := ""
	for i := 0; i < 7; i++ {
		str += strconv.Itoa(phases[i])
	}
	num, _ := strconv.Atoi(str)
	return num
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

func convertToInts(input []string) []int {
	programStr := strings.Split(input[0], "")
	var program []int

	for _, s := range programStr {
		i, _ := strconv.Atoi(s)
		program = append(program, i)
	}
	return program
}

// Absoulute value of Int
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
