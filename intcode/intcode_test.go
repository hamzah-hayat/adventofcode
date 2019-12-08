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
