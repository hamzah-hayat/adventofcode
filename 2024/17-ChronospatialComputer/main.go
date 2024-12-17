package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	methodP   *string
	RegisterA int
	RegisterB int
	RegisterC int
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

	// Grab input
	regexRegA := regexp.MustCompile(`Register A: (\d+)`)
	regexRegB := regexp.MustCompile(`Register B: (\d+)`)
	regexRegC := regexp.MustCompile(`Register C: (\d+)`)
	regexProgram := regexp.MustCompile(`Program: (.+)`)

	RegisterA, _ = strconv.Atoi(regexRegA.FindAllStringSubmatch(input[0], -1)[0][1])
	RegisterB, _ = strconv.Atoi(regexRegB.FindAllStringSubmatch(input[1], -1)[0][1])
	RegisterC, _ = strconv.Atoi(regexRegC.FindAllStringSubmatch(input[2], -1)[0][1])
	program := regexProgram.FindAllStringSubmatch(input[4], -1)[0][1]

	programSplit := strings.Split(program, ",")
	output := []string{}
	currentInstructionCounter := 0
	for currentInstructionCounter < len(programSplit) {
		opcode, _ := strconv.Atoi(programSplit[currentInstructionCounter])
		operand, _ := strconv.Atoi(programSplit[currentInstructionCounter+1])

		currentInstructionCounter, output = RunOperand(output, currentInstructionCounter, opcode, operand)
	}

	return strings.Join(output, ",")
}

func RunOperand(output []string, currentInstructionCounter, opcode, operand int) (int, []string) {

	switch opcode {
	case 0:
		// adv
		dem := ComboOperator(operand)
		RegisterA = RegisterA / int(math.Pow(2, float64(dem)))
		return currentInstructionCounter + 2, output
	case 1:
		// bxl
		RegisterB = RegisterB ^ operand
		return currentInstructionCounter + 2, output
	case 2:
		// bst
		comboOperand := ComboOperator(operand)
		RegisterB = comboOperand % 8
	case 3:
		// jnz
		if RegisterA != 0 {
			return operand, output
		} else {
			return currentInstructionCounter + 2, output
		}
	case 4:
		// bxc
		RegisterB = RegisterB ^ RegisterC
		return currentInstructionCounter + 2, output
	case 5:
		// out
		val := strconv.Itoa(ComboOperator(operand) % 8)
		output = append(output, val)
		return currentInstructionCounter + 2, output
	case 6:
		// bdv
		dem := ComboOperator(operand)
		RegisterB = RegisterA / int(math.Pow(2, float64(dem)))
		return currentInstructionCounter + 2, output
	case 7:
		// cdv
		dem := ComboOperator(operand)
		RegisterC = RegisterA / int(math.Pow(2, float64(dem)))
		return currentInstructionCounter + 2, output
	}
	return currentInstructionCounter + 2, output
}

// Combo operands 0 through 3 represent literal values 0 through 3.
// Combo operand 4 represents the value of register A.
// Combo operand 5 represents the value of register B.
// Combo operand 6 represents the value of register C.
// Combo operand 7 is reserved and will not appear in valid programs.
func ComboOperator(operand int) int {
	switch operand {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return RegisterA
	case 5:
		return RegisterB
	case 6:
		return RegisterC
	}
	return -1
}

func PartTwo(filename string) string {
	input := readInput(filename)

	// Grab input
	regexRegA := regexp.MustCompile(`Register A: (\d+)`)
	regexRegB := regexp.MustCompile(`Register B: (\d+)`)
	regexRegC := regexp.MustCompile(`Register C: (\d+)`)
	regexProgram := regexp.MustCompile(`Program: (.+)`)

	RegisterA, _ = strconv.Atoi(regexRegA.FindAllStringSubmatch(input[0], -1)[0][1])
	RegisterB, _ = strconv.Atoi(regexRegB.FindAllStringSubmatch(input[1], -1)[0][1])
	RegisterC, _ = strconv.Atoi(regexRegC.FindAllStringSubmatch(input[2], -1)[0][1])
	program := regexProgram.FindAllStringSubmatch(input[4], -1)[0][1]

	programSplit := strings.Split(program, ",")
	startA := int(math.Pow(8, float64(len(programSplit)-1))) - 2
	for {
		// Setup
		RegisterA = startA
		currentInstructionCounter := 0
		output := []string{}

		// Run program
		for currentInstructionCounter < len(programSplit) {
			opcode, _ := strconv.Atoi(programSplit[currentInstructionCounter])
			operand, _ := strconv.Atoi(programSplit[currentInstructionCounter+1])

			currentInstructionCounter, output = RunOperand(output, currentInstructionCounter, opcode, operand)
		}

		// Check output
		if strings.Join(output, ",") == program {
			break
		}

		if len(output) == len(programSplit) {
			for i := len(output) - 1; i >= 0; i-- {
				numOut, _ := strconv.Atoi(output[i])
				numPro, _ := strconv.Atoi(string(programSplit[i]))
				if numOut != numPro {
					if i == 0 {
						startA++
					} else {
						startA += int(math.Pow(8, (float64(i - 1))))
					}
					fmt.Println("Checking Start A Register:", startA)
					break
				}
			}
		} else {
			startA += int(math.Pow(8, float64(len(programSplit)-1)))
		}
	}

	// We've now found an input that returns the program, but is it the lowest one?
	startAUpperBound := 202322936867370
	answer := startA
	for {
		// Setup
		RegisterA = startAUpperBound
		currentInstructionCounter := 0
		output := []string{}

		// Run program
		for currentInstructionCounter < len(programSplit) {
			opcode, _ := strconv.Atoi(programSplit[currentInstructionCounter])
			operand, _ := strconv.Atoi(programSplit[currentInstructionCounter+1])

			currentInstructionCounter, output = RunOperand(output, currentInstructionCounter, opcode, operand)
		}

		// Check output
		if strings.Join(output, ",") == program {
			answer = startAUpperBound
		}

		// Should we stop looking?
		if startAUpperBound < int(math.Pow(8, float64(len(programSplit)-1))) {
			break
		}

		startAUpperBound -= 4096
	}

	return strconv.Itoa(answer)
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
