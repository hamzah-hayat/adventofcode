package main

import (
	"strings"
	"testing"
)

func TestOneStep(t *testing.T) {
	input := readInput("example")

	startTemplate := input[0]
	var formula []Formula

	for i := 2; i < len(input); i++ {
		split := strings.Split(input[i], "->")
		formula = append(formula, Formula{strings.Trim(split[0], " "), strings.Trim(split[1], " ")})
	}

	startTemplate = RunFormula(startTemplate, formula)

	if startTemplate != "NCNBCHB" {
		t.Errorf("Expected NCNBCHB, received %v", startTemplate)
	}
}

func TestTwoStep(t *testing.T) {
	input := readInput("example")

	startTemplate := input[0]
	var formula []Formula

	for i := 2; i < len(input); i++ {
		split := strings.Split(input[i], "->")
		formula = append(formula, Formula{strings.Trim(split[0], " "), strings.Trim(split[1], " ")})
	}

	startTemplate = RunFormula(startTemplate, formula)
	startTemplate = RunFormula(startTemplate, formula)

	if startTemplate != "NBCCNBBBCBHCB" {
		t.Errorf("Expected NBCCNBBBCBHCB, received %v", startTemplate)
	}
}

func TestThreeStep(t *testing.T) {
	input := readInput("example")

	startTemplate := input[0]
	var formula []Formula

	for i := 2; i < len(input); i++ {
		split := strings.Split(input[i], "->")
		formula = append(formula, Formula{strings.Trim(split[0], " "), strings.Trim(split[1], " ")})
	}

	startTemplate = RunFormula(startTemplate, formula)
	startTemplate = RunFormula(startTemplate, formula)
	startTemplate = RunFormula(startTemplate, formula)

	if startTemplate != "NBBBCNCCNBBNBNBBCHBHHBCHB" {
		t.Errorf("Expected NBBBCNCCNBBNBNBBCHBHHBCHB, received %v", startTemplate)
	}
}

func TestFourStep(t *testing.T) {
	input := readInput("example")

	startTemplate := input[0]
	var formula []Formula

	for i := 2; i < len(input); i++ {
		split := strings.Split(input[i], "->")
		formula = append(formula, Formula{strings.Trim(split[0], " "), strings.Trim(split[1], " ")})
	}

	startTemplate = RunFormula(startTemplate, formula)
	startTemplate = RunFormula(startTemplate, formula)
	startTemplate = RunFormula(startTemplate, formula)
	startTemplate = RunFormula(startTemplate, formula)

	if startTemplate != "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB" {
		t.Errorf("Expected NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB, received %v", startTemplate)
	}
}
