package intcode

import (
	"fmt"
	"testing"
)

// Tests from Day 5 (https://adventofcode.com/2019/day/5)
func TestIntCode_EqualTo8_True_Position(t *testing.T) {

	program := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	expected := 1

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 8
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_EqualTo8_False_Position(t *testing.T) {

	program := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	expected := 0

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 10
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_LessThan8_True_Position(t *testing.T) {

	program := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	expected := 1

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 3
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_LessThan8_False_Position(t *testing.T) {

	program := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	expected := 0

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 10
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_EqualTo8_True_Immediate(t *testing.T) {

	program := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	expected := 1

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 8
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_EqualTo8_False_Immediate(t *testing.T) {

	program := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	expected := 0

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 3
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_LessThan8_True_Immediate(t *testing.T) {

	program := []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	expected := 1

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 3
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_LessThan8_False_Immediate(t *testing.T) {

	program := []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	expected := 0

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 12
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_CheckZero_False_Position(t *testing.T) {

	program := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	expected := 0

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 0
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_CheckZero_True_Position(t *testing.T) {

	program := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	expected := 1

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 31
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_CheckZero_False_Immediate(t *testing.T) {

	program := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	expected := 0

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 0
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_CheckZero_True_Immediate(t *testing.T) {

	program := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	expected := 1

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 31
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_CheckNumberAgainst8_Lower(t *testing.T) {

	program := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	expected := 999

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 3
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_CheckNumberAgainst8_Equal(t *testing.T) {

	program := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	expected := 1000

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 8
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_CheckNumberAgainst8_Greater(t *testing.T) {

	program := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	expected := 1001

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 12
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

// Test Relative Mode

func TestIntCode_RelativeMode_Simple(t *testing.T) {

	program := []int{9, 3, 203, 10, 204, 10, 99}
	expected := 99

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	input <- 99
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

// Extra features

func TestIntCode_PrintLargeNumber(t *testing.T) {

	program := []int{104, 1125899906842624, 99}
	expected := 1125899906842624

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_LargeMemorySpace(t *testing.T) {

	program := []int{1101, 1, 1, 10000000, 4, 10000000, 99}
	expected := 2

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_Print16DigitNumber(t *testing.T) {

	program := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	expected := 1219070632396864

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)
	result := <-output

	if expected != result {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestIntCode_ProduceCopyOfSelf(t *testing.T) {

	program := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	expected := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go RunIntCodeProgram(program, input, output)

	var result [16]int
	for i := 0; i < 15; i++ {
		result[i] = <-output

		if expected[i] != result[i] {
			t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
		}
	}

}

// OPCode reading tests

func TestOPCodeRead_simple(t *testing.T) {

	testOpCodes := []int{101, 1001, 204}
	expectedOpCodes := []int{1, 1, 4}
	expectedParamModes := [][3]int{{1, 0, 0}, {0, 1, 0}, {2, 0, 0}}

	for i, test := range testOpCodes {

		opCode, paramModes := readOpCode(test)

		if expectedOpCodes[i] != opCode || expectedParamModes[i] != paramModes {
			t.Error(fmt.Sprint("Expected ", expectedOpCodes[i], " and", expectedParamModes[i], " but got ", opCode, " and", paramModes))
		}

	}

}
