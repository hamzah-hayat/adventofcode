package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "18"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestSilver_Empty(t *testing.T) {
	value := PartOne("example_empty")
	expected := "18"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestSilver_Small(t *testing.T) {
	value := PartOne("example_small")
	expected := "4"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example_mas")
	expected := "9"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}
