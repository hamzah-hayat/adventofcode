package main

import (
	"strconv"
	"testing"
)

func TestNearestNeighbours(t *testing.T) {
	inputs1 := []Point{Point{1, 2}}
	inputs2 := [][]Point{[]Point{Point{1, 1}, Point{1, 2}}}
	outputs := []Point{Point{1, 2}}

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
