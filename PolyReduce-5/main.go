package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//PartOne()
	PartTwo()
}

func PartOne() {
	input := readInput()
	reduced := PolyReduction(input)
	fmt.Printf("The Length of the new Polymer is %v\n", len(reduced))

}

func PartTwo() {
	input := readInput()
	allLetters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "y", "x", "z"}

	// Need to remove a letter then polyreduce the file, then store the length of the file with the letter
	answersList := make(map[string]int)

	for _, letter := range allLetters {
		// First, take input and remove all letters of that type
		newInput := strings.Replace(input, strings.ToUpper(letter), "", -1)
		newInput = strings.Replace(newInput, strings.ToLower(letter), "", -1)
		reduced := PolyReduction(newInput)
		answersList[letter] = len(reduced)

	}
	fmt.Println(answersList)

	lowestValue := 50000
	lowestLetter := ""
	for letter, value := range answersList {
		if value < lowestValue {
			lowestValue = value
			lowestLetter = letter
		}
	}

	fmt.Println("Removing the letter " + lowestLetter + " reduces the polymer the most")

}

// Reduce the polymerstring
// Any characters with a opposite case
func PolyReduction(polymerString string) string {
	newPolymer := polymerString
	allLetters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "y", "x", "z"}

	newLength := len(newPolymer) + 1
	for newLength > len(newPolymer) {
		newLength = len(newPolymer)
		for _, letter := range allLetters {
			// Replace any lowercase + uppercase letters in the polymer, use length check to break out
			newPolymer = strings.Replace(newPolymer, strings.ToUpper(letter)+strings.ToLower(letter), "", -1)
			newPolymer = strings.Replace(newPolymer, strings.ToLower(letter)+strings.ToUpper(letter), "", -1)
		}
	}
	return newPolymer
}

// Read data from input.txt
// Load it into claim array
func readInput() string {

	var input string

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() != "" {
			input = scanner.Text()
		}
	}
	return input
}
