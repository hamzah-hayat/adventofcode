package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "36"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestSilver_Small(t *testing.T) {
	value := PartOne("example_small")
	expected := "1"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := "81"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}
