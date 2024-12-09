package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "1928"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := "2858"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}