package main

import (
	"fmt"
	"testing"
)

func TestSpaceStoichiometry_CreateReactions(t *testing.T) {

	input := []string{
		"10 ORE => 10 A",
		"1 ORE => 1 B",
		"7 A, 1 B => 1 C",
		"7 A, 1 C => 1 D",
		"7 A, 1 D => 1 E",
		"7 A, 1 E => 1 FUEL",
	}

	reactionsList := createReactions(input)

	firstchem := []chemical{chemical{number: 10, name: "ORE"}}
	secondchem := []chemical{chemical{number: 10, name: "A"}}
	reaction := reaction{ingredients: firstchem, results: secondchem}

	if reactionsList[0].ingredients[0].number != reaction.ingredients[0].number || reactionsList[0].ingredients[0].name != reaction.ingredients[0].name {
		t.Error(fmt.Sprint("Exected output ", reactionsList[0].ingredients[0], " but got ", reaction.ingredients[0]))
	}

}

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
