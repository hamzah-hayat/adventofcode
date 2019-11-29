package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//PartOne()
	PartTwo()
}

func PartOne() {
	input := readInput()

	twoTimes := CountLettersWhichAppear(2, input)
	threeTimes := CountLettersWhichAppear(3, input)

	checksum := strconv.Itoa(twoTimes * threeTimes)

	fmt.Println("Checksum is " + checksum)
}

func PartTwo() {
	input := readInput()
	correctLetters := FindCorrectBoxes(input)
	fmt.Println("The correct letters are " + correctLetters)
}

// Read data from input.txt
// Load it into string array
func readInput() []string {

	var input []string

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() != "" {
			line := scanner.Text()
			input = append(input, line)
		}
	}
	return input
}

// Count how many times a line from the array has letters that apear numOfTimes times
func CountLettersWhichAppear(numOfTimes int, input []string) int {
	total := 0
	for _, boxName := range input {
		// for each box name, do a running total for each letter
		letterMap := make(map[rune]int)

		for _, char := range boxName {
			letterMap[char]++
		}

		// Check if any of the letters appear numOfTimes
		for _, num := range letterMap {
			if num == numOfTimes {
				total++
				break
			}
		}

	}
	return total
}

func FindCorrectBoxes(input []string) string {
	correctLetters := ""

	for _, firstBoxName := range input {
		for _, secondBoxName := range input {
			// Compare these two boxnames character by character
			// If they have one chracter difference, delete the different character then return the letters
			differences := 0
			for i, _ := range firstBoxName {
				if firstBoxName[i] != secondBoxName[i] {
					differences++
				}
			}
			if differences == 1 {
				// Found correct boxes, now build string
				for i, _ := range firstBoxName {
					if firstBoxName[i] == secondBoxName[i] {
						correctLetters += string(firstBoxName[i])
					}
				}
				return correctLetters
			}
		}
	}

	return correctLetters
}
