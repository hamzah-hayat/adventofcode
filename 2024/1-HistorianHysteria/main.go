package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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

type Point struct {
	x int
	y int
}

func PartOne(filename string) string {
	input := readInput(filename)

	firstList := make([]int, 0)
	secondList := make([]int, 0)

	for _, v := range input {
		split := strings.Split(v, "   ")
		num1, _ := strconv.Atoi(split[0])
		num2, _ := strconv.Atoi(split[1])
		firstList = append(firstList, num1)
		secondList = append(secondList, num2)
	}

	sort.Ints(firstList)
	sort.Ints(secondList)

	total := 0
	for i := 0; i < len(firstList); i++ {
		total = total + abs(secondList[i]-firstList[i])
	}

	return strconv.Itoa(total)
}

// Absoulute value of Int
func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func PartTwo(filename string) string {
	input := readInput(filename)

	firstList := make([]int, 0)
	secondList := make([]int, 0)

	for _, v := range input {
		split := strings.Split(v, "   ")
		num1, _ := strconv.Atoi(split[0])
		num2, _ := strconv.Atoi(split[1])
		firstList = append(firstList, num1)
		secondList = append(secondList, num2)
	}

	sort.Ints(firstList)
	sort.Ints(secondList)

	total := 0
	for i := 0; i < len(firstList); i++ {
		multiplier := 0
		for j := 0; j < len(secondList); j++ {
			if firstList[i] == secondList[j] {
				multiplier++
			}
		}
		total = total + (multiplier * firstList[i])
	}

	return strconv.Itoa(total)
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
