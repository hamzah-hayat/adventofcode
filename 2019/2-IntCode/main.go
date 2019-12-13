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

	programStr := strings.Split(input[0], ",")
	var program []int

	for _, s := range programStr {
		i, _ := strconv.Atoi(s)
		program = append(program, i)
	}

	teriminate := false

	// setup input special
	program[1] = 12
	program[2] = 2

	// Loop over every 4 values starting from zero
	for opCodeI := 0; opCodeI < len(program); opCodeI = opCodeI + 4 {
		switch program[opCodeI] {
		case 1:
			fmt.Println("Added to", program[opCodeI+3], "the numbers", program[opCodeI+1], "and", program[opCodeI+2])
			program[program[opCodeI+3]] = program[program[opCodeI+1]] + program[program[opCodeI+2]]
			break
		case 2:
			fmt.Println("Multipled to", program[opCodeI+3])
			program[program[opCodeI+3]] = program[program[opCodeI+1]] * program[program[opCodeI+2]]
			break
		case 99:
			fmt.Println("teriminated")
			teriminate = true
			break
		default:
			fmt.Println("Error")
			break
		}
		if teriminate {
			break
		}
	}

	fmt.Println(program)

	fmt.Println("Positon zero is", program[0])

}

func partTwo() {
	input := readInput()

	programStr := strings.Split(input[0], ",")

	// look for the right noun and verb
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {

			// reset data
			var program []int
			for _, s := range programStr {
				i, _ := strconv.Atoi(s)
				program = append(program, i)
			}

			program[1] = i
			program[2] = j

			runProgram(program)

			if program[0] == 19690720 {
				fmt.Println("Noun is", i, " Verb is", j)
				break
			}
		}
	}
}

func runProgram(program []int) {

	teriminate := false
	// Loop over every 4 values starting from zero
	for opCodeI := 0; opCodeI < len(program); opCodeI = opCodeI + 4 {
		switch program[opCodeI] {
		case 1:
			//fmt.Println("Added to", program[opCodeI+3], "the numbers", program[opCodeI+1], "and", program[opCodeI+2])
			program[program[opCodeI+3]] = program[program[opCodeI+1]] + program[program[opCodeI+2]]
			break
		case 2:
			//fmt.Println("Multipled to", program[opCodeI+3])
			program[program[opCodeI+3]] = program[program[opCodeI+1]] * program[program[opCodeI+2]]
			break
		case 99:
			//fmt.Println("teriminated")
			teriminate = true
			break
		default:
			//fmt.Println("Error")
			break
		}
		if teriminate {
			break
		}
	}
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
