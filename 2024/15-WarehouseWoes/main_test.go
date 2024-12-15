package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "10092"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestSilver_Tiny(t *testing.T) {
	value := PartOne("example_tiny")
	expected := "2028"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("example")
	expected := "9021"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold_Tiny(t *testing.T) {
	value := PartTwo("example_gold_tiny")
	expected := "618"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold_Tiny_2(t *testing.T) {
	value := PartTwo("example_gold_tiny_2")
	expected := "2038"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold_Tiny_3(t *testing.T) {
	value := PartTwo("example_gold_tiny_3")
	expected := "2035"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}