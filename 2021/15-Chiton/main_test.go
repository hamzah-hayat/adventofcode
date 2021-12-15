package main

import (
	"strconv"
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "40"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := "315"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestExpandedMap(t *testing.T) {

	input := readInput("example")
	expanded := readInput("exampleExpanded")

	riskMap := make(map[Point]int)
	riskMapExpandedReal := make(map[Point]int)

	// small input
	for x, v := range input {
		for y, c := range v {
			num, _ := strconv.Atoi(string(c))
			riskMap[Point{x, y}] = num
		}
	}

	// big input
	for x, v := range expanded {
		for y, c := range v {
			num, _ := strconv.Atoi(string(c))
			riskMapExpandedReal[Point{x, y}] = num
		}
	}

	// Now lets expand small input and see if it is same to big input
	expandedRiskMap := expandRiskMap(riskMap, 5, 5)

	// expandedRiskMap and riskMapExpandedReal should be the same
	for p, v := range riskMapExpandedReal {
		if expandedRiskMap[p] != v {
			t.Errorf("Expected %v but got %v at point %v", v, expandedRiskMap[p], p)
		}
	}

}
