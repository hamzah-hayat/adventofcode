package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "13"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := "1"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGoldLarge(t *testing.T) {
	value := PartTwo("large_example")
	expected := "36"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}
