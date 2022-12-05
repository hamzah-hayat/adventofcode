package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
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

	stacks := make(map[int]*Stack)

	numStacks := (len(input[0]) + 1) / 4
	for i := 0; i < numStacks; i++ {
		stacks[i] = &Stack{}
	}

	endLine := 0
	// First process our input
	for _, line := range input {
		// Starting from index 2 and plus 4 each time, take each box and push onto its
		for i := 0; i < len(line); i++ {
			if unicode.IsLetter(rune(line[i])) {
				// get right index then push onto right stack
				// sequence is 1,5,9,13 etc
				// we want 0,1,2,3
				// so (i-1) /4
				stackToPushTo := (i - 1) / 4
				stacks[stackToPushTo].Push(string(line[i]))
			}
		}

		endLine++
		// Check if we are on number line
		if line[1] == '1' {
			break
		}
	}

	// Reverse our stacks
	for _, s := range stacks {
		s.Reverse()
	}

	// Now read instructions
	for i := endLine + 1; i < len(input); i++ {
		// move 1 from 2 to 1
		// regex to get nums then push and pop
		regex := regexp.MustCompile(`move ([0-9]+) from ([0-9]+) to ([0-9]+)`)

		result := regex.FindStringSubmatch(input[i])

		numToMove, _ := strconv.Atoi(result[1])
		startStack, _ := strconv.Atoi(result[2])
		endStack, _ := strconv.Atoi(result[3])

		for i := 0; i < numToMove; i++ {
			item := stacks[startStack-1].Pop()
			stacks[endStack-1].Push(item)
		}
	}

	outputString := ""
	// get our output
	for i := 0; i < numStacks; i++ {
		outputString += stacks[i].Peak()
	}

	return outputString
}

func PartTwo(filename string) string {
	input := readInput(filename)

	stacks := make(map[int]*Stack)

	numStacks := (len(input[0]) + 1) / 4
	for i := 0; i < numStacks; i++ {
		stacks[i] = &Stack{}
	}

	endLine := 0
	// First process our input
	for _, line := range input {
		// Starting from index 2 and plus 4 each time, take each box and push onto its
		for i := 0; i < len(line); i++ {
			if unicode.IsLetter(rune(line[i])) {
				// get right index then push onto right stack
				// sequence is 1,5,9,13 etc
				// we want 0,1,2,3
				// so (i-1) /4
				stackToPushTo := (i - 1) / 4
				stacks[stackToPushTo].Push(string(line[i]))
			}
		}

		endLine++
		// Check if we are on number line
		if line[1] == '1' {
			break
		}
	}

	// Reverse our stacks
	for _, s := range stacks {
		s.Reverse()
	}

	// Now read instructions
	for i := endLine + 1; i < len(input); i++ {
		// move 1 from 2 to 1
		// regex to get nums then push and pop
		regex := regexp.MustCompile(`move ([0-9]+) from ([0-9]+) to ([0-9]+)`)

		result := regex.FindStringSubmatch(input[i])

		numToMove, _ := strconv.Atoi(result[1])
		startStack, _ := strconv.Atoi(result[2])
		endStack, _ := strconv.Atoi(result[3])

		var itemsToMove []string
		for i := 0; i < numToMove; i++ {
			itemsToMove = append(itemsToMove, stacks[startStack-1].Pop())
		}

		for i := numToMove; i > 0; i-- {
			stacks[endStack-1].Push(itemsToMove[i-1])
		}
	}

	outputString := ""
	// get our output
	for i := 0; i < numStacks; i++ {
		outputString += stacks[i].Peak()
	}

	return outputString
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

// Stack implementation
type Stack []string

// Pop an item off top of stack and return it
func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	} else {
		// Get top item and pop it off, then return
		topItem := len(*s) - 1
		item := (*s)[topItem]
		*s = (*s)[:topItem]
		return item
	}
}

// Push an item on top of the stack
func (s *Stack) Push(item string) {
	*s = append(*s, item)
}

// Peak at the top item of the stack
func (s *Stack) Peak() string {
	return (*s)[len(*s)-1]
}

// Reverse the stack
func (s *Stack) Reverse() {
	var items []string

	for {
		item := s.Pop()
		if item == "" {
			break
		}
		items = append(items, item)
	}

	for _, i := range items {
		s.Push(i)
	}
}
