package main

import "testing"

func TestManHattenDistance(t *testing.T) {

	inputs1 := []Point{Point{1, 2}}
	inputs2 := []Point{Point{2, 2}}
	outputs := []int{1}

	for i, input := range inputs {
		t.Run(input, func(t *testing.T) {
			testOutput := PolyReduction(input)
			if testOutput != outputs[i] {
				t.Error("Failed on input " + input)
				t.Error("Expected result " + outputs[i])
				t.Error("Instead got " + testOutput)
				t.Fail()
			}
		})
	}
}
