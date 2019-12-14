package main

import (
	"fmt"
	"testing"
)

func TestSpaceStoichiometry_TestOreNeededSmall(t *testing.T) {

	input := []string{
		"10 ORE => 10 A",
		"1 ORE => 1 B",
		"7 A, 1 B => 1 C",
		"7 A, 1 C => 1 D",
		"7 A, 1 D => 1 E",
		"7 A, 1 E => 1 FUEL",
	}
	expected := 31

	reactionList := createReactions(input)

	oreNeededForFuel := convertOreToFuel(reactionList)

	if expected != oreNeededForFuel {
		t.Error(fmt.Sprint("Exected output ", expected, " but got ", oreNeededForFuel))
	}

}