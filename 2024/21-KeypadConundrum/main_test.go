package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "126384"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := "154115708116294"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}
