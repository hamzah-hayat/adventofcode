package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p1", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()

	switch *methodP {
	case "p1":
		partOne()
		break
	case "p2":
		partTwo()
		break
	case "test":
		break
	}
}

func partOne() {
	input := readInput()

	deck := makeDeck(10006)

	deck = runHand(input, deck)

	for i, val := range deck {
		if val == -1 {
			fmt.Println(i, ":", val)
		}
	}

}

func makeDeck(deckSize int) []int {
	deck := []int{}
	for i := 0; i < deckSize; i++ {
		deck = append(deck, i)
	}
	return deck
}

func partTwo() {
	//input := readInput()
}

func runHand(input []string, startDeck []int) []int {
	newDeck := []int{}
	newDeck = startDeck

	for _, val := range input {

		// Figure out what method we are using, then do it
		if strings.Contains(val, "deal into new stack") {
			newDeck = dealNewStack(newDeck)
			continue
		}

		if strings.Contains(val, "deal with increment") {
			incrementStr := strings.SplitAfter(val, "increment ")[1]
			increment, _ := strconv.Atoi(incrementStr)
			newDeck = dealWithIncrement(newDeck, increment)
			continue
		}

		if strings.Contains(val, "cut") {
			cutStr := strings.SplitAfter(val, "cut ")[1]
			cut, _ := strconv.Atoi(cutStr)
			if cut < 0 {
				newDeck = cutDeckNegative(newDeck, cut*-1)
			} else {
				newDeck = cutDeckPositive(newDeck, cut)
			}
			continue
		}

	}

	return newDeck
}

func dealNewStack(deck []int) []int {
	newDeck := []int{}
	for i := len(deck) - 1; i >= 0; i-- {
		newDeck = append(newDeck, deck[i])
	}
	return newDeck
}

func dealWithIncrement(deck []int, increment int) []int {
	newDeck := []int{}
	// setup array
	for i := 0; i < len(deck); i++ {
		newDeck = append(newDeck, -1)
	}

	startDeckCounter := 0
	newDeckCounter := 0
	for startDeckCounter < len(deck) {
		newDeck[newDeckCounter] = deck[startDeckCounter]
		startDeckCounter++
		newDeckCounter = mod(newDeckCounter+increment, len(deck))
	}
	for i, val := range newDeck {
		if val == -1 {
			fmt.Println(i, ":", val)
		}
	}
	return newDeck
}

func cutDeckPositive(deck []int, cut int) []int {
	newDeck := []int{}
	// First take bottom half of cards and place them on top
	// then take "cut" cards and place them on bottom
	for i := cut; i < len(deck); i++ {
		newDeck = append(newDeck, deck[i])
	}

	for i := 0; i < cut; i++ {
		newDeck = append(newDeck, deck[i])
	}
	return newDeck
}

func cutDeckNegative(deck []int, cut int) []int {
	newDeck := []int{}
	// Opposite of cut deck positive
	for i := len(deck) - cut; i < len(deck); i++ {
		newDeck = append(newDeck, deck[i])
	}

	for i := 0; i < len(deck)-cut; i++ {
		newDeck = append(newDeck, deck[i])
	}
	return newDeck
}

// Read data from input.txt
// Return the string, so that we can deal with it however
func readInput() []string {

	var input []string

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
