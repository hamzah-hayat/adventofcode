package intcode

import "fmt"

// RunIntCodeProgram runs an intcode program
func RunIntCodeProgram(program []int, input chan int, output chan int) {

	teriminate := false
	// Loop over opCodes, starting from zero
	for opCodeI := 0; opCodeI < len(program); {

		// Need to check what our opcode is, plus its param modes
		opCode, paramModes := readOpCode(program[opCodeI])

		//fmt.Println("Running", opCode, "with paramModes", paramModes)

		switch opCode {
		case 1:
			param1, param2 := getTwoParams(program, paramModes, opCodeI)
			//fmt.Println("Added to", program[opCodeI+3], "the numbers", param1, "and", param2)

			program[program[opCodeI+3]] = param1 + param2
			opCodeI = opCodeI + 4
			break
		case 2:
			param1, param2 := getTwoParams(program, paramModes, opCodeI)
			//fmt.Println("Multiplied to", program[opCodeI+3], "the numbers", param1, "and", param2)

			program[program[opCodeI+3]] = param1 * param2
			opCodeI = opCodeI + 4
			break
		case 3:
			//fmt.Println("Waiting for input")
			program[program[opCodeI+1]] = <-input
			opCodeI = opCodeI + 2
			break
		case 4:
			param1 := getOneParams(program, paramModes, opCodeI)
			//fmt.Println("The output is", param1)
			output <- param1
			opCodeI = opCodeI + 2
			break
		case 5:
			param1, param2 := getTwoParams(program, paramModes, opCodeI)
			//fmt.Println("Checking if", param1, " is non zero, will place answer in", param2)
			if param1 != 0 {
				opCodeI = param2
			} else {
				opCodeI = opCodeI + 3
			}
			break
		case 6:
			param1, param2 := getTwoParams(program, paramModes, opCodeI)
			//fmt.Println("Checking if", param1, " is zero, will place answer in", param2)
			if param1 == 0 {
				opCodeI = param2
			} else {
				opCodeI = opCodeI + 3
			}
			break
		case 7:
			param1, param2 := getTwoParams(program, paramModes, opCodeI)
			//fmt.Println("Checking if", param1, " is less then", param2, " will place answer in", program[program[opCodeI+3]])
			if param1 < param2 {
				program[program[opCodeI+3]] = 1
			} else {
				program[program[opCodeI+3]] = 0
			}
			opCodeI = opCodeI + 4
			break
		case 8:
			param1, param2 := getTwoParams(program, paramModes, opCodeI)
			//fmt.Println("Checking if", param1, " is equal to", param2, " will place answer in", program[program[opCodeI+3]])
			if param1 == param2 {
				program[program[opCodeI+3]] = 1
			} else {
				program[program[opCodeI+3]] = 0
			}
			opCodeI = opCodeI + 4
			break
		case 99:
			//fmt.Println("teriminated")
			teriminate = true
			break
		default:
			fmt.Println("Error")
			teriminate = true
			break
		}
		if teriminate {
			break
		}
	}
}

// ReadOpCode TODO: There has to be a better way to do this
func readOpCode(opCodeFull int) (int, [3]int) {

	opCode := opCodeFull % 100
	parameterModes := [3]int{0, 0, 0}

	params := (opCodeFull - opCode) / 100

	switch params {
	case 1:
		parameterModes[2] = 1
		break
	case 11:
		parameterModes[2] = 1
		parameterModes[1] = 1
		break
	case 10:
		parameterModes[1] = 1
		break
	}

	return opCode, parameterModes
}

func getOneParams(program []int, paramModes [3]int, opCodeI int) int {
	param1 := 0
	if paramModes[2] == 0 {
		param1 = program[program[opCodeI+1]]
	} else if paramModes[2] == 1 {
		param1 = program[opCodeI+1]
	}

	return param1
}

func getTwoParams(program []int, paramModes [3]int, opCodeI int) (int, int) {
	param1 := 0
	if paramModes[2] == 0 {
		param1 = program[program[opCodeI+1]]
	} else if paramModes[2] == 1 {
		param1 = program[opCodeI+1]
	}

	param2 := 0
	if paramModes[1] == 0 {
		param2 = program[program[opCodeI+2]]
	} else if paramModes[1] == 1 {
		param2 = program[opCodeI+2]
	}

	return param1, param2

}

func getThreeParams(program []int, paramModes [3]int, opCodeI int) (int, int, int) {

	param1 := 0
	if paramModes[2] == 0 {
		param1 = program[program[opCodeI+1]]
	} else if paramModes[2] == 1 {
		param1 = program[opCodeI+1]
	}

	param2 := 0
	if paramModes[1] == 0 {
		param2 = program[program[opCodeI+2]]
	} else if paramModes[1] == 1 {
		param2 = program[opCodeI+2]
	}

	param3 := 0
	if paramModes[0] == 0 {
		param3 = program[program[opCodeI+3]]
	} else if paramModes[0] == 1 {
		param3 = program[opCodeI+3]
	}
	return param1, param2, param3
}
