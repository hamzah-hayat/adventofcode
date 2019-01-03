package main

import (
	"strconv"
	"testing"
)

func TestFindInfinitePoints(t *testing.T) {
	inputs := [][]Point{[]Point{Point{1, 1}, Point{1, 6}, Point{8, 3}, Point{3, 4}, Point{5, 5}, Point{8, 9}}}
	outputs := [][]Point{[]Point{Point{3, 4}, Point{5, 5}}}

	for i := range inputs {
		t.Run("Test using Inputs "+strconv.Itoa(i), func(t *testing.T) {
			testOutput := FindInfinitePoints(inputs[i])
			if CheckSlicesOutput(testOutput, outputs[i]) {
				t.Errorf("Failed on input %v", inputs[i])
				t.Errorf("Expected result %v", outputs[i])
				t.Errorf("Instead got %v", testOutput)
				t.Fail()
			} else {
				t.Logf("Passed on input %v", inputs[i])
				t.Logf("Expected result %v", outputs[i])
				t.Logf("Value got was %v", testOutput)
			}
		})
	}
}

func TestNearestNeighbours(t *testing.T) {
	inputs1 := []Point{Point{18, 20}, Point{0, 0}}
	inputs2 := [][]Point{[]Point{Point{12, 18}, Point{1, 2}}, []Point{Point{1, 1}}}
	outputs := []Point{Point{12, 18}, Point{1, 1}}

	for i := range inputs1 {
		t.Run("Test using Inputs "+strconv.Itoa(i), func(t *testing.T) {
			testOutput := FindClosestNeighbour(inputs1[i], inputs2[i])
			if testOutput != outputs[i] {
				t.Errorf("Failed on input %v and %v", inputs1[i], inputs2[i])
				t.Errorf("Expected result %v", outputs[i])
				t.Errorf("Instead got %v", testOutput)
				t.Fail()
			} else {
				t.Logf("Passed on input %v and %v", inputs1[i], inputs2[i])
				t.Logf("Expected result %v", outputs[i])
				t.Logf("Value got was %v", testOutput)
			}
		})
	}
}

func TestManHattenDistance(t *testing.T) {

	inputs1 := []Point{Point{1, 2}, Point{0, 35}, Point{200, 0}}
	inputs2 := []Point{Point{2, 2}, Point{0, 0}, Point{0, -1}}
	outputs := []int{1, 35, 201}

	for i := range inputs1 {
		t.Run("Test using Inputs "+strconv.Itoa(i), func(t *testing.T) {
			testOutput := ManhattenDistance(inputs1[i], inputs2[i])
			if testOutput != outputs[i] {
				t.Errorf("Failed on input %v and %v", inputs1[i], inputs2[i])
				t.Errorf("Expected result %v", outputs[i])
				t.Errorf("Instead got %v", testOutput)
				t.Fail()
			} else {
				t.Logf("Passed on input %v and %v", inputs1[i], inputs2[i])
				t.Logf("Expected result %v", outputs[i])
				t.Logf("Value got was %v", testOutput)
			}
		})
	}
}

func CheckSlicesOutput(testOutput []Point, outputChecked []Point) bool {
	// Check length first
	if len(testOutput) != len(outputChecked) {
		return true
	}

	for _, point := range testOutput {
		found := false
		for _, point2 := range outputChecked {
			if point == point2 {
				found = true
				continue
			}
		}
		if !found {
			return true
		}
	}
	return false
}
