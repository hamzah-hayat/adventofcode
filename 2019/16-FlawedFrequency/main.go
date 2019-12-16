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
	methodP := flag.String("method", "p2", "The method/part that should be run, valid are p1,p2 and test")
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
			// Ignore the first element of the pattern
			patternNum := 1
			outputNum := 0
			for _, val := range phase {
				outputNum += val * basePattern(i+1, patternNum)
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

func basePattern(i, patterNum int) int {

	patterNum = patterNum % (i * 4)

	if patterNum >= 3*i {
		return -1
	}

	if patterNum >= 2*i {
		return 0
	}

	if patterNum >= i {
		return 1
	}

	if patterNum < i {
		return 0
	}

	return 0
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

	finalPhase := runPhasesFastWithMoreThanHalfOffset(phase[offset:], 100)

	// Convert final Phase into string
	finalPhaseStr := ""
	for i := len(finalPhase) - 1; i > 0; i-- {
		finalPhaseStr += strconv.Itoa(finalPhase[i])
	}

	fmt.Println("The first eight digits of the 100th phase (with offset) are:", finalPhaseStr[0:8])

}

func runPhasesFastWithMoreThanHalfOffset(phases []int, runAmount int) []int {

	// Since we started from more then halfway, we can reverse the list and add each number sequentially
	// This will give us each correct digit from the last digit to the offset digit
	// This counts as a full phase
	// we can then keep going this
	phases = reverseArray(phases)
	finalPhase := phases

	for run := 0; run < runAmount; run++ {
		newPhase := make([]int, len(phases))
		sum := 0
		for i := 0; i < len(phases); i++ {
			// Since we start from the back, we can just sum
			sum += phases[i]
			// This is a correct digit
			newPhase[i] = sum % 10
			sum = sum % 10
		}
		phases = newPhase
		finalPhase = phases
	}

	return finalPhase
}

func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func reverseArrayStr(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func multiplyPhases(phases []int, multiplyNum int) []int {
	newPhases := make([]int, len(phases)*multiplyNum)
	for i := 0; i < len(phases)*multiplyNum; i++ {
		newPhases[i] = phases[i%len(phases)]
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
