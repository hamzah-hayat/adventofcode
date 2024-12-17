package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("example")
	expected := "4,6,3,5,6,3,5,2,1,0"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestSilver_Tiny_1(t *testing.T) {
	value := PartOne("example_tiny_1")
	expected := 1

	if RegisterB != expected {
		t.Error("Got ", value, " expected ", expected)
	}
}

func TestSilver_Tiny_2(t *testing.T) {
	value := PartOne("example_tiny_2")
	expected := "0,1,2"

	if value != expected {
		t.Error("Got ", value, " expected ", expected)
	}
}

func TestSilver_Tiny_3(t *testing.T) {
	value := PartOne("example_tiny_3")
	expected := "4,2,5,6,7,7,7,7,3,1,0"

	if value != expected {
		t.Error("Got ", value, " expected ", expected)
	}
	if RegisterA != 0 {
		t.Error("Got ", value, " expected ", 0)
	}
}

func TestSilver_Tiny_4(t *testing.T) {
	value := PartOne("example_tiny_4")
	expected := 26

	if RegisterB != expected {
		t.Error("Got ", value, " expected ", expected)
	}
}

func TestSilver_Tiny_5(t *testing.T) {
	value := PartOne("example_tiny_5")
	expected := 44354

	if RegisterB != expected {
		t.Error("Got ", value, " expected ", expected)
	}
}

// https://www.youtube.com/watch?v=vu2NK5REvWM
// func TestGold(t *testing.T) {
// 	value := PartTwo("example_gold")
// 	expected := "117440"
// 	if value != expected {
// 		t.Error("Got ", value, " expected ", expected)
// 	}
// }

// func TestGold_Input(t *testing.T) {
// 	value := PartTwo("example_gold")
// 	expected := "117440"

// 	if value != expected {
// 		t.Error("Got ", value, " expected ", expected)
// 	}
// }