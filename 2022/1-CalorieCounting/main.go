package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
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
	input := readInputInt(filename)

	var elfCal []int

	sum := 0
	for _, v := range input {
		if v == 0 {
			elfCal = append(elfCal, sum)
			sum = 0
		} else {
			sum += v
		}
	}
	// Remember to grab last elf!
	elfCal = append(elfCal, sum)

	largest := 0
	for _, v := range elfCal {
		if v > largest {
			largest = v
		}
	}

	return strconv.Itoa(largest)
}

func PartTwo(filename string) string {
	input := readInputInt(filename)

	var elfCal []int

	sum := 0
	for _, v := range input {
		if v == 0 {
			elfCal = append(elfCal, sum)
			sum = 0
		} else {
			sum += v
		}
	}
	// Remember to grab last elf!
	elfCal = append(elfCal, sum)

	// sort list and get top three
	sort.Ints(elfCal)
	sumTopThree := elfCal[len(elfCal)-1] + elfCal[len(elfCal)-2] + elfCal[len(elfCal)-3]

	return strconv.Itoa(sumTopThree)
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
