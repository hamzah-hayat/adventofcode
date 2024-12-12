package main

import (
	"testing"
)

func TestSilver_Small(t *testing.T) {
	value := PartOne("example_small")
	expected := "140"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestSilver_Small_2(t *testing.T) {
	value := PartOne("example_small_2")
	expected := "772"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "1930"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold_small(t *testing.T) {
	value := PartTwo("example_small")
	expected := "80"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold_small_2(t *testing.T) {
	value := PartTwo("example_small_2")
	expected := "436"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold_small_3(t *testing.T) {
	value := PartTwo("example_small_3")
	expected := "236"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold_small_4(t *testing.T) {
	value := PartTwo("example_small_4")
	expected := "368"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}