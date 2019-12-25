package main

import (
	"fmt"
	"testing"
)

func TestSlamShuffle_MakeDeck(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	deck := makeDeck(10)

	for i := 0; i < len(expected); i++ {
		if expected[i] != deck[i] {
			t.Error(fmt.Sprint("Exected output ", expected[i], " but got ", deck[i]))
		}
	}

}

func TestSlamShuffle_DealStack(t *testing.T) {

	input := []string{
		"deal into new stack",
	}
	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	deck := makeDeck(10)

	deck = runHand(input, deck)

	for i := 0; i < len(expected)-1; i++ {
		if expected[i] != deck[i] {
			t.Error(fmt.Sprint("Exected output ", expected[i], " but got ", deck[i]))
		}
	}

}

func TestSlamShuffle_DealStackIncrement(t *testing.T) {

	input := []string{
		"deal with increment 3",
	}
	expected := []int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}

	deck := makeDeck(10)

	deck = runHand(input, deck)

	for i := 0; i < len(expected)-1; i++ {
		if expected[i] != deck[i] {
			t.Error(fmt.Sprint("Exected output ", expected[i], " but got ", deck[i]))
		}
	}

}

func TestSlamShuffle_CutPositive(t *testing.T) {

	input := []string{
		"cut 3",
	}
	expected := []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}

	deck := makeDeck(10)

	deck = runHand(input, deck)

	for i := 0; i < len(expected)-1; i++ {
		if expected[i] != deck[i] {
			t.Error(fmt.Sprint("Exected output ", expected[i], " but got ", deck[i]))
		}
	}

}

func TestSlamShuffle_CutNegative(t *testing.T) {

	input := []string{
		"cut -4",
	}
	expected := []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}

	deck := makeDeck(10)

	deck = runHand(input, deck)

	for i := 0; i < len(expected)-1; i++ {
		if expected[i] != deck[i] {
			t.Error(fmt.Sprint("Exected output ", expected[i], " but got ", deck[i]))
		}
	}

}

func TestSlamShuffle_MultipleShuffles1(t *testing.T) {

	input := []string{
		"deal with increment 7",
		"deal into new stack",
		"deal into new stack",
	}
	expected := []int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}

	deck := makeDeck(10)

	deck = runHand(input, deck)

	for i := 0; i < len(expected)-1; i++ {
		if expected[i] != deck[i] {
			t.Error(fmt.Sprint("Exected output ", expected[i], " but got ", deck[i]))
		}
	}

}

func TestSlamShuffle_MultipleShuffles2(t *testing.T) {

	input := []string{
		"cut 6",
		"deal with increment 7",
		"deal into new stack",
	}
	expected := []int{3, 0, 7, 4, 1, 8, 5, 2, 9, 6}

	deck := makeDeck(10)

	deck = runHand(input, deck)

	for i := 0; i < len(expected)-1; i++ {
		if expected[i] != deck[i] {
			t.Error(fmt.Sprint("Exected output ", expected[i], " but got ", deck[i]))
		}
	}

}

func TestSlamShuffle_MultipleShuffles3(t *testing.T) {

	input := []string{
		"deal with increment 7",
		"deal with increment 9",
		"cut -2",
	}
	expected := []int{6, 3, 0, 7, 4, 1, 8, 5, 2, 9}

	deck := makeDeck(10)

	deck = runHand(input, deck)

	for i := 0; i < len(expected)-1; i++ {
		if expected[i] != deck[i] {
			t.Error(fmt.Sprint("Exected output ", expected[i], " but got ", deck[i]))
		}
	}

}

func TestSlamShuffle_MultipleShuffles4(t *testing.T) {

	input := []string{
		"deal into new stack",
		"cut -2",
		"deal with increment 7",
		"cut 8",
		"cut -4",
		"deal with increment 7",
		"cut 3",
		"deal with increment 9",
		"deal with increment 3",
		"cut -1",
	}
	expected := []int{9, 2, 5, 8, 1, 4, 7, 0, 3, 6}

	deck := makeDeck(10)

	deck = runHand(input, deck)

	for i := 0; i < len(expected)-1; i++ {
		if expected[i] != deck[i] {
			t.Error(fmt.Sprint("Exected output ", expected[i], " but got ", deck[i]))
		}
	}

}

func TestSlamShuffle_TestIncrementMutlipleRuns(t *testing.T) {

	input := []string{
		"deal with increment 7",
		"deal with increment 3",
	}
	expected := []int{}

	deck := makeDeck(15)

	deck = runHand(input, deck)

	for i := 0; i < len(expected)-1; i++ {
		if expected[i] != deck[i] {
			t.Error(fmt.Sprint("Exected output ", expected[i], " but got ", deck[i]))
		}
	}

}
