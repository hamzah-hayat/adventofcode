package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "3"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := "6"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold2(t *testing.T) {
	value := PartTwo("example2")
	expected := "10"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}
