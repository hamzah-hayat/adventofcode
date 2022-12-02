package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	methodP *string
)

func parseFlags() {
	methodP = flag.String("method", "all", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()
}

func main() {

	parseFlags()

	switch *methodP {
	case "all":
		fmt.Println("Silver:" + PartOne("input"))
		fmt.Println("Gold:" + PartTwo("input"))
	case "p1":
		fmt.Println("Silver:" + PartOne("input"))
	case "p2":
		fmt.Println("Gold:" + PartTwo("input"))
	}
}

func PartOne(filename string) string {
	input := readInput(filename)

	totalScore := 0
	for _, v := range input {
		game := strings.Split(v, " ")
		totalScore += calculateScore(game[0], game[1])
	}

	return strconv.Itoa(totalScore)
}

func PartTwo(filename string) string {
	input := readInput(filename)

	totalScore := 0
	for _, v := range input {
		game := strings.Split(v, " ")
		totalScore += calculateScoreWithMatch(game[0], game[1])
	}

	return strconv.Itoa(totalScore)
}

func calculateScore(opponentChoice, playerChoice string) int {
	totalScore := 0

	// score based on what we picked
	switch playerChoice {
	case "X":
		totalScore += 1
		break
	case "Y":
		totalScore += 2
		break
	case "Z":
		totalScore += 3
		break
	}

	// score based on round win/loss
	if (opponentChoice == "A" && playerChoice == "Y") || (opponentChoice == "B" && playerChoice == "Z") || (opponentChoice == "C" && playerChoice == "X") {
		totalScore += 6
	}
	if (opponentChoice == "A" && playerChoice == "X") || (opponentChoice == "B" && playerChoice == "Y") || (opponentChoice == "C" && playerChoice == "Z") {
		totalScore += 3
	}

	return totalScore
}

func calculateScoreWithMatch(opponentChoice, playerMatchResult string) int {
	totalScore := 0

	// Figure out what we need to pick
	// X means lose
	// Y means we draw
	// Z means we win

	// A and X mean rock
	// B and Y mean paper
	// C and Z mean scissors
	// here we convert and use original function

	// Do lose,draw, win
	if opponentChoice == "A" && playerMatchResult == "X" {
		totalScore = calculateScore("A", "Z")
	} else if opponentChoice == "A" && playerMatchResult == "Y" {
		totalScore = calculateScore("A", "X")
	} else if opponentChoice == "A" && playerMatchResult == "Z" {
		totalScore = calculateScore("A", "Y")
	}

	if opponentChoice == "B" && playerMatchResult == "X" {
		totalScore = calculateScore("B", "X")
	} else if opponentChoice == "B" && playerMatchResult == "Y" {
		totalScore = calculateScore("B", "Y")
	} else if opponentChoice == "B" && playerMatchResult == "Z" {
		totalScore = calculateScore("B", "Z")
	}

	if opponentChoice == "C" && playerMatchResult == "X" {
		totalScore = calculateScore("C", "Y")
	} else if opponentChoice == "C" && playerMatchResult == "Y" {
		totalScore = calculateScore("C", "Z")
	} else if opponentChoice == "C" && playerMatchResult == "Z" {
		totalScore = calculateScore("C", "X")
	}

	return totalScore
}

// Read data from input.txt
// Return the string, so that we can deal with it however
func readInput(filename string) []string {

	var input []string

	f, _ := os.Open(filename + ".txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

// Read data from input.txt
// Return the string as int
func readInputInt(filename string) []int {

	var input []int

	f, _ := os.Open(filename + ".txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		input = append(input, num)
	}
	return input
}
