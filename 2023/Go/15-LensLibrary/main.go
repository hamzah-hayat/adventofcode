package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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

	split := strings.Split(input[0], ",")

	// Work out value of each sequence
	totalResult := 0
	for _, sequence := range split {
		totalResult += HolidayHash(sequence)
	}

	num := strconv.Itoa(totalResult)

	return num
}

type Lens struct {
	label string
	focal int
}

func PartTwo(filename string) string {
	input := readInput(filename)

	split := strings.Split(input[0], ",")

	var boxes [][]Lens

	for i := 0; i < 256; i++ {
		boxes = append(boxes, make([]Lens, 0))
	}

	for _, sequence := range split {
		re, _ := regexp.Compile(`(\w+)(=|-)(\d*)`)
		regexMatch := re.FindStringSubmatch(sequence)

		label := regexMatch[1]
		expression := regexMatch[2]

		if expression == "=" {
			lensFocal, _ := strconv.Atoi(regexMatch[3])
			boxes = AddToBox(boxes, label, lensFocal)
		} else if expression == "-" {
			boxes = RemoveFromBox(boxes, label)
		}

	}

	// Now score each box
	totalPower := 0
	for i := 0; i < len(boxes); i++ {
		for j := 0; j < len(boxes[i]); j++ {
			totalPower += (1 + i) * (j + 1) * boxes[i][j].focal
		}
	}

	num := strconv.Itoa(totalPower)

	return num
}

func RemoveFromBox(boxes [][]Lens, label string) [][]Lens {
	boxNum := HolidayHash(label)
	for i := 0; i < len(boxes[boxNum]); i++ {
		if boxes[boxNum][i].label == label {
			// Remove from Box
			boxes[boxNum] = append(boxes[boxNum][:i], boxes[boxNum][i+1:]...)
		}
	}
	return boxes
}

func AddToBox(boxes [][]Lens, label string, lensFocal int) [][]Lens {
	boxNum := HolidayHash(label)
	found := false
	for i := 0; i < len(boxes[boxNum]); i++ {
		if boxes[boxNum][i].label == label {
			// Replace lens
			boxes[boxNum][i] = Lens{label, lensFocal}
			found = true;
		}
	}

	if !found {
		boxes[boxNum] = append(boxes[boxNum], Lens{label, lensFocal})
	}

	return boxes
}

func HolidayHash(input string) int {
	result := 0
	for _, char := range input {
		result += int(char)
		result = result * 17
		result = result % 256
	}
	return result
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
func readInputInt() []int {

	var input []int

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		input = append(input, num)
	}
	return input
}
