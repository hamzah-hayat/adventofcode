package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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
	//input := readInput()
}

// Reduce the polymerstring
// Any characters with a opposite case
func PolyReduction(polymerString string) string {
	newPolymer := polymerString

	foundReduction := true
	for foundReduction {
		fmt.Println("Current string length is " + strconv.Itoa(len(newPolymer)))
		foundReduction = false
		for index := range newPolymer {
			// First check if we are out of range
			if index+1 >= len(newPolymer) {
				break
			}
			currentChar := []rune(newPolymer)[index]
			nextChar := []rune(newPolymer)[index+1]
			//fmt.Println("Current Char is " + string(currentChar) + " and next char is " + string(nextChar))

			if unicode.IsLower(currentChar) {
				if unicode.IsUpper(nextChar) && unicode.ToUpper(currentChar) == unicode.ToUpper(nextChar) {
					// Reduce these two
					if index == 0 {
						newPolymer = string([]rune(newPolymer)[index+2:])
					} else {
						newPolymer = string([]rune(newPolymer)[:index]) + string([]rune(newPolymer)[index+2:])
					}
					foundReduction = true
					continue
				}
			}

			if unicode.IsUpper(currentChar) {
				if unicode.IsLower(nextChar) && unicode.ToUpper(currentChar) == unicode.ToUpper(nextChar) {
					// Reduce these two
					if index == 0 {
						newPolymer = string([]rune(newPolymer)[index+2:])
					} else {
						newPolymer = string([]rune(newPolymer)[:index]) + string([]rune(newPolymer)[index+2:])
					}
					foundReduction = true
					continue
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
