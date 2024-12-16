package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "7036"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestSilver_2(t *testing.T) {
	value := PartOne("example_2")
	expected := "11048"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := "45"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold_2(t *testing.T) {
	value := PartTwo("example_2")
	expected := "64"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}
