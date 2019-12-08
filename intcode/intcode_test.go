package intcode

import (
	"fmt"
	"testing"
)

// Tests from Day 5 (https://adventofcode.com/2019/day/5)
// 3,9,8,9,10,9,4,9,99,-1,8 - Using position mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
func TestIntCode_EqualTo8_True(t *testing.T) {

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

func TestIntCode_EqualTo8_False(t *testing.T) {

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
