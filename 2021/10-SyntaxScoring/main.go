package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
)

var (
	methodP *string
)

func init() {
	// Use Flags to run a part
	methodP = flag.String("method", "p2", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()
}

func main() {
	switch *methodP {
	case "p1":
		PartOne()
		break
	case "p2":
		PartTwo()
		break
	case "test":
		break
	}
}

func PartOne() {
	input := readInput()

	sum := 0
	for _, v := range input {
		sum += ScoreLine(v)
	}
	fmt.Println(sum)
}

func PartTwo() {
	input := readInput()

	var incompleteLines []string
	for _, l := range input {
		if ScoreLine(l) == 0 {
			incompleteLines = append(incompleteLines, l)
		}
	}

	var sums []int
	for _, v := range incompleteLines {
		sums = append(sums, ScoreCompletionLine(v))
	}

	sort.Ints(sums)

	fmt.Println(sums[(len(sums)/2)])

}

func ScoreLine(line string) int {

	roundRegex, _ := regexp.Compile(`\(\)`)
	squareRegex, _ := regexp.Compile(`\[\]`)
	curlyRegex, _ := regexp.Compile("{}")
	triangleRegex, _ := regexp.Compile("<>")
	currentLine := line
	match := true

	for match {
		match = false
		if len(currentLine) == 0 {
			break
		}

		if roundRegex.MatchString(currentLine) {
			currentLine = roundRegex.ReplaceAllString(currentLine, "")
			match = true
		}
		if squareRegex.MatchString(currentLine) {
			currentLine = squareRegex.ReplaceAllString(currentLine, "")
			match = true
		}
		if curlyRegex.MatchString(currentLine) {
			currentLine = curlyRegex.ReplaceAllString(currentLine, "")
			match = true
		}
		if triangleRegex.MatchString(currentLine) {
			currentLine = triangleRegex.ReplaceAllString(currentLine, "")
			match = true
		}

	}

	for _, c := range currentLine {
		switch c {
		case ')':
			return 3
		case ']':
			return 57
		case '}':
			return 1197
		case '>':
			return 25137
		}
	}

	return 0
}

func ScoreCompletionLine(line string) int {

	roundRegex, _ := regexp.Compile(`\(\)`)
	squareRegex, _ := regexp.Compile(`\[\]`)
	curlyRegex, _ := regexp.Compile("{}")
	triangleRegex, _ := regexp.Compile("<>")
	currentLine := line
	match := true

	for match {
		match = false
		if len(currentLine) == 0 {
			break
		}

		if roundRegex.MatchString(currentLine) {
			currentLine = roundRegex.ReplaceAllString(currentLine, "")
			match = true
		}
		if squareRegex.MatchString(currentLine) {
			currentLine = squareRegex.ReplaceAllString(currentLine, "")
			match = true
		}
		if curlyRegex.MatchString(currentLine) {
			currentLine = curlyRegex.ReplaceAllString(currentLine, "")
			match = true
		}
		if triangleRegex.MatchString(currentLine) {
			currentLine = triangleRegex.ReplaceAllString(currentLine, "")
			match = true
		}

	}

	// Now we have our incomplete line
	// To complete it, reverse and flip the characters
	reversed := Reverse(currentLine)

	score := 0
	for _, c := range reversed {
		score = score * 5
		if c == '(' {
			score = score + 1
		}
		if c == '[' {
			score = score + 2
		}
		if c == '{' {
			score = score + 3
		}
		if c == '<' {
			score = score + 4
		}
	}

	return score
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

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
