package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "3"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

// There is no part two
// func TestGold(t *testing.T) {
// 	value := PartTwo("example")
// 	expected := ""

// 	if value != expected {
// 		t.Error("Got " + value + " expected " + expected)
// 	}
// }
