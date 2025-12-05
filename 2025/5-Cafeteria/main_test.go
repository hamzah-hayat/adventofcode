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
	expected := "14"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold2(t *testing.T) {
	value := PartTwo("example2")
	expected := "14"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}
