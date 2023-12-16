package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "1320"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestSilver_2(t *testing.T) {
	value := PartOne("example2")
	expected := "52"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := "145"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}
