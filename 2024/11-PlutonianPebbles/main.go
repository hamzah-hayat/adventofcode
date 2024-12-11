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

	splitLine := strings.Split(input[0], " ")

	numMap := make(map[int]int)

	for _, v := range splitLine {
		num, _ := strconv.Atoi(v)
		numMap[num] += 1
	}

	// Run simulation
	numMap = Blink(25, numMap)

	total := 0
	for _, num := range numMap {
		total += num
	}

	return strconv.Itoa(total)
}

// Blink simulates our stone ruleset
func Blink(blinks int, numMap map[int]int) map[int]int {

	for i := 0; i < blinks; i++ {
		newNumMap := make(map[int]int)
		for num, totalNum := range numMap {

			numStr := strconv.Itoa(num)

			delete(numMap, num)

			if num == 0 {
				newNumMap[1] += totalNum
			} else if len(numStr)%2 == 0 {

				left, _ := strconv.Atoi(numStr[:len(numStr)/2])
				right, _ := strconv.Atoi(numStr[len(numStr)/2:])

				newNumMap[left] += totalNum
				newNumMap[right] += totalNum

			} else {
				newNumMap[num*2024] += totalNum
			}
		}
		numMap = newNumMap
		// total := 0
		// for _, num := range numMap {
		// 	total += num
		// }
		// // Total stones
		// fmt.Println("After blinking", i+1, "times, we have", total, "stones")

	}
	return numMap
}

func PartTwo(filename string) string {
	input := readInput(filename)

	splitLine := strings.Split(input[0], " ")

	numMap := make(map[int]int)

	for _, v := range splitLine {
		num, _ := strconv.Atoi(v)
		numMap[num] += 1
	}

	// Run simulation
	numMap = Blink(75, numMap)

	total := 0
	for _, num := range numMap {
		total += num
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
