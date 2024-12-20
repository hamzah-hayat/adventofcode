package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example", 10)
	expected := "10"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example", 70)
	expected := "41"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}
