package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "4"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestSilver_Large(t *testing.T) {
	value := PartOne("example_large")
	expected := "2024"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

// No example for Gold
// func TestGold(t *testing.T) {
// 	value := PartTwo("example_gold_correct")
// 	expected := "z00,z01,z02,z05"

// 	if value != expected {
// 		t.Error("Got " + value + " expected " + expected)
// 	}
// }
