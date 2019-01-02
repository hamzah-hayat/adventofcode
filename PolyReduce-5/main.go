package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

func main() {
	PartOne()
	//PartTwo()
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

	var wg sync.WaitGroup
	// Only run 5 threads at a time
	goroutineChannel := make(chan struct{}, 5)

	for _, letter := range allLetters {
		wg.Add(1)
		// First, take input and remove all letters of that type
		newInput := strings.Replace(input, strings.ToUpper(letter), "", -1)
		newInput = strings.Replace(newInput, strings.ToLower(letter), "", -1)
		goroutineChannel <- struct{}{}
		go func() {
			defer wg.Done()
			reduced := PolyReduction(newInput)
			thisLetter := letter
			answersList[thisLetter] = len(reduced)
			fmt.Println("Finished for letter " + thisLetter + " final value was " + strconv.Itoa(len(reduced)))
			<-goroutineChannel
		}()
	}
	fmt.Println(answersList)

	wg.Wait()

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

	foundReduction := true
	for foundReduction {
		fmt.Println("Current string length is " + strconv.Itoa(len(newPolymer)))
		foundReduction = false
		for index := 0; index < len(newPolymer); index++ {
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
