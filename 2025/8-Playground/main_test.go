package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example", 10)
	expected := "40"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := "25272"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}
