package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example", 11, 7)
	expected := "12"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

// Cannot be tested lmao
func TestGold(t *testing.T) {
	value := PartTwo("example", 11, 7)
	expected := "0"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}
