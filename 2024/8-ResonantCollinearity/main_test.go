package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "14"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestSilverSimple(t *testing.T) {
	value := PartOne("example_simple")
	expected := "4"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := "34"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold_Simple(t *testing.T) {
	value := PartTwo("example_gold_simple")
	expected := "9"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}