package main

import "testing"

func TestReducePolymer(t *testing.T) {

	inputs := []string{"a", "aA", "aa", "acCCcA", "aaaaaacCcCAAAAAA", "VbGiIbBKkzZeEukK", "AafEeF", "UuUuWmMPpwfEeFoODdblLoOFqQf"}
	outputs := []string{"a", "", "aa", "", "", "VbGu", "", "b"}

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
