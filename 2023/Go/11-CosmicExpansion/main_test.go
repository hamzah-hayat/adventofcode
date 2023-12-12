package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "374"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example", 10)
	expected := "1030"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold_Big(t *testing.T) {
	value := PartTwo("example", 100)
	expected := "8410"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}
