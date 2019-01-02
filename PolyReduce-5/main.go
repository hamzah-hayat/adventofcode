package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	PartOne()
	//PartTwo()
}

func PartOne() {
	//input := readInput()
	fmt.Println("New Poly is:")
	fmt.Println(PolyReduction("aAbBCd"))
}

func PartTwo() {
	//input := readInput()
}

// Reduce the polymerstring
// Any characters with a opposite case
func PolyReduction(polymerString string) string {
	newPolymer := polymerString

	foundReduction := true
	for foundReduction {
		foundReduction = false
		for index, character := range newPolymer {
			// First check if we are out of range
			if index+1 > len(newPolymer) {
				break
			}

			if unicode.IsLower(character) {
				if unicode.IsUpper([]rune(newPolymer)[index+1]) && unicode.ToUpper(character) == unicode.ToUpper([]rune(newPolymer)[index+1]) {
					// Reduce these two
					newPolymer = string([]rune(newPolymer)[:index-1]) + string([]rune(newPolymer)[index+1:])
					foundReduction = true
					break
				}
			}

			if unicode.IsUpper(character) {
				if unicode.IsLower([]rune(newPolymer)[index+1]) && unicode.ToUpper(character) == unicode.ToUpper([]rune(newPolymer)[index+1]) {
					// Reduce these two
					newPolymer = string([]rune(newPolymer)[:index-1]) + string([]rune(newPolymer)[index+1:])
					foundReduction = true
					break
				}
			}

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
