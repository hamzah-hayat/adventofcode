package intcode

import "fmt"

// RunIntCodeProgram runs an intcode program
func RunIntCodeProgram(programInt []int, input chan int, output chan int, message chan Message) {

	// Use maps instead of int arrays
	program := make(map[int]int)
	for i, val := range programInt {
		program[i] = val
	}

	teriminate := false
	relativeBase := 0
	// Loop over opCodes, starting from zero
	for opCodeI := 0; opCodeI < len(program); {

		// Need to check what our opcode is, plus its param modes
		opCode, paramModes := readOpCode(program[opCodeI])

		//fmt.Println("Running", opCode, "with paramModes", paramModes)

		switch opCode {
		case 1:
			param1, param2 := getTwoParams(program, paramModes, opCodeI, relativeBase)
			//fmt.Println("Added to", program[opCodeI+3], "the numbers", param1, "and", param2)
			placeParam(program, paramModes, 2, opCodeI, relativeBase, param1+param2)
			opCodeI = opCodeI + 4
			break
		case 2:
			param1, param2 := getTwoParams(program, paramModes, opCodeI, relativeBase)
			//fmt.Println("Multiplied to", program[opCodeI+3], "the numbers", param1, "and", param2)
			placeParam(program, paramModes, 2, opCodeI, relativeBase, param1*param2)
			opCodeI = opCodeI + 4
			break
		case 3:
			//fmt.Println("Waiting for input")
			if message != nil {
				message <- Message{MessageType: 0}
			}
			placeParam(program, paramModes, 0, opCodeI, relativeBase, <-input)
			opCodeI = opCodeI + 2
			break
		case 4:
			if message != nil {
				message <- Message{MessageType: 1}
			}
			param1 := getOneParams(program, paramModes, opCodeI, relativeBase)
			//fmt.Println("The output is", param1)
			output <- param1
			opCodeI = opCodeI + 2
			break
		case 5:
			param1, param2 := getTwoParams(program, paramModes, opCodeI, relativeBase)
			//fmt.Println("Checking if", param1, " is non zero, will place answer in", param2)
			if param1 != 0 {
				opCodeI = param2
			} else {
				opCodeI = opCodeI + 3
			}
			break
		case 6:
			param1, param2 := getTwoParams(program, paramModes, opCodeI, relativeBase)
			//fmt.Println("Checking if", param1, " is zero, will place answer in", param2)
			if param1 == 0 {
				opCodeI = param2
			} else {
				opCodeI = opCodeI + 3
			}
			break
		case 7:
			param1, param2 := getTwoParams(program, paramModes, opCodeI, relativeBase)
			//fmt.Println("Checking if", param1, " is less then", param2, " will place answer in", program[program[opCodeI+3]])
			if param1 < param2 {
				placeParam(program, paramModes, 2, opCodeI, relativeBase, 1)
			} else {
				placeParam(program, paramModes, 2, opCodeI, relativeBase, 0)
			}
			opCodeI = opCodeI + 4
			break
		case 8:
			param1, param2 := getTwoParams(program, paramModes, opCodeI, relativeBase)
			//fmt.Println("Checking if", param1, " is equal to", param2, " will place answer in", program[program[opCodeI+3]])
			if param1 == param2 {
				placeParam(program, paramModes, 2, opCodeI, relativeBase, 1)
			} else {
				placeParam(program, paramModes, 2, opCodeI, relativeBase, 0)
			}
			opCodeI = opCodeI + 4
			break
		case 9:
			param1 := getOneParams(program, paramModes, opCodeI, relativeBase)
			//fmt.Println("The output is", param1)
			relativeBase += param1
			opCodeI = opCodeI + 2
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

// RunIntCodeProgramWaitForTermination runs an intcode program then sends a true signal to the t channel
func RunIntCodeProgramWaitForTermination(program []int, input chan int, output chan int, t chan bool, message chan Message) {
	RunIntCodeProgram(program, input, output, message)
	t <- true
}

// ReadOpCode TODO: There has to be a better way to do this
func readOpCode(opCodeFull int) (int, [3]int) {

	opCode := opCodeFull % 100
	parameterModes := [3]int{0, 0, 0}

	params := (opCodeFull - opCode) / 100

	parameterModes[0] = params % 10
	parameterModes[1] = (params / 10) % 10
	parameterModes[2] = (params / 100) % 10

	return opCode, parameterModes
}

func getOneParams(program map[int]int, paramModes [3]int, opCodeI int, relativeBase int) int {
	param1 := 0
	if paramModes[0] == 0 {
		param1 = program[program[opCodeI+1]]
	} else if paramModes[0] == 1 {
		param1 = program[opCodeI+1]
	} else if paramModes[0] == 2 {
		param1 = program[program[opCodeI+1]+relativeBase]
	}

	return param1
}

func getTwoParams(program map[int]int, paramModes [3]int, opCodeI int, relativeBase int) (int, int) {
	param1 := 0
	if paramModes[0] == 0 {
		param1 = program[program[opCodeI+1]]
	} else if paramModes[0] == 1 {
		param1 = program[opCodeI+1]
	} else if paramModes[0] == 2 {
		param1 = program[program[opCodeI+1]+relativeBase]
	}

	param2 := 0
	if paramModes[1] == 0 {
		param2 = program[program[opCodeI+2]]
	} else if paramModes[1] == 1 {
		param2 = program[opCodeI+2]
	} else if paramModes[1] == 2 {
		param2 = program[program[opCodeI+2]+relativeBase]
	}

	return param1, param2

}

// Place a parameter into the program based on relative or position mode
func placeParam(program map[int]int, paramModes [3]int, outputParam int, opCodeI int, relativeBase int, value int) {
	if paramModes[outputParam] == 0 {
		program[program[opCodeI+outputParam+1]] = value
	} else if paramModes[outputParam] == 2 {
		program[program[opCodeI+outputParam+1]+relativeBase] = value
	}
}

// Message is used to signal whether the intcode computer is waiting for input or waiting for output
// We can use this to decide whether to send input or take output from the program
// MessageTypes are 0 for input, 1 for output
type Message struct {
	MessageType int
}
