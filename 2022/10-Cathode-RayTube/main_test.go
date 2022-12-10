package main

import (
	"testing"
)

func TestSilver(t *testing.T) {
	value := PartOne("large_example")
	expected := "13140"

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}

func TestGold(t *testing.T) {
	value := PartTwo("large_example")
	expected :=
	`##..##..##..##..##..##..##..##..##..##..
	###...###...###...###...###...###...###.
	####....####....####....####....####....
	#####.....#####.....#####.....#####.....
	######......######......######......####
	#######.......#######.......#######.....`

	if value != expected {
		t.Error("Got " + value + " expected " + expected)
	}
}
