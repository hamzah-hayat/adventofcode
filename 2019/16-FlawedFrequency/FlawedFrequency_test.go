package main

import (
	"fmt"
	"testing"
)

func TestFlawedFrequency_TestPhaseSmall(t *testing.T) {

	input := []string{"12345678"}
	expected := "01029498"

	phase := convertToInts(input)

	outputNumbers := runPhases(phase, 4)

	if expected != string(outputNumbers[3][0:8]) {
		t.Error(fmt.Sprint("Exected output ", expected, " but got ", string(outputNumbers[3][0:8])))
	}

}

func TestFlawedFrequency_TestPhaseSmall1(t *testing.T) {

	input := []string{"80871224585914546619083218645595"}
	expected := "24176176"

	phase := convertToInts(input)

	outputNumbers := runPhases(phase, 100)

	if expected != string(outputNumbers[99][0:8]) {
		t.Error(fmt.Sprint("Exected output ", expected, " but got ", string(outputNumbers[99][0:8])))
	}

}

func TestFlawedFrequency_TestPhaseSmall2(t *testing.T) {

	input := []string{"19617804207202209144916044189917"}
	expected := "73745418"

	phase := convertToInts(input)

	outputNumbers := runPhases(phase, 100)

	if expected != string(outputNumbers[99][0:8]) {
		t.Error(fmt.Sprint("Exected output ", expected, " but got ", string(outputNumbers[99][0:8])))
	}

}

func TestFlawedFrequency_TestPhaseSmall3(t *testing.T) {

	input := []string{"69317163492948606335995924319873"}
	expected := "52432133"

	phase := convertToInts(input)

	outputNumbers := runPhases(phase, 100)

	if expected != string(outputNumbers[99][0:8]) {
		t.Error(fmt.Sprint("Exected output ", expected, " but got ", string(outputNumbers[99][0:8])))
	}

}

func TestFlawedFrequency_TestPhaseLarge0(t *testing.T) {

	input := []string{"03036732577212944063491565474664"}
	expected := "84462026"

	phase := convertToInts(input)

	phase = multiplyPhases(phase, 10000)
	offset := getPhaseOffset(phase)

	outputNumbers := runPhases(phase, 100)

	if expected != string(outputNumbers[99][0+offset:8+offset]) {
		t.Error(fmt.Sprint("Exected output ", expected, " but got ", string(outputNumbers[99][0+offset:8+offset])))
	}

}
