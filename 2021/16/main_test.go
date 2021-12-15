package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "0"

	if value != "0" {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := "0"

	if value != "0" {
		t.Error("Got " + value + " expected " + expected)
	}
}
