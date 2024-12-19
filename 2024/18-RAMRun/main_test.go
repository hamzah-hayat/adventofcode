package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example", 7, 7, 12)
	expected := "22"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example", 7, 7)
	expected := "6,1"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}
