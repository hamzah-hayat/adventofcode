package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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
		fmt.Println("Gold:\n" + PartTwo("input"))
	case "p1":
		fmt.Println("Silver:" + PartOne("input"))
	case "p2":
		fmt.Println("Gold:\n" + PartTwo("input"))
	}
}

func PartOne(filename string) string {
	input := readInput(filename)

	monkeyList := CreateMonkeyList(input)

	// Run 20 rounds
	for i := 0; i < 20; i++ {
		// run each monkey in turn
		for _, monkey := range monkeyList {
			// For each Item eg
			// Monkey inspects an item with a worry level of 79.
			// Worry level is multiplied by 19 to 1501.
			// Monkey gets bored with item. Worry level is divided by 3 to 500.
			// Current worry level is not divisible by 23.
			// Item with worry level 500 is thrown to monkey 3.

			for _, item := range monkey.items {
				// Inspect
				item = monkey.operation(item)
				// Bored
				item = item / 3
				// Throw
				nextMonkey := monkey.test(item)
				monkeyList[nextMonkey].items = append(monkeyList[nextMonkey].items, item)
				// Remove from this monkeys items
				monkey.items = monkey.items[1:]
				// Add interaction
				monkey.interactions++
			}

		}

	}

	// find top two monkeys by interactions
	monkeyInteractions := []int{}

	for _, m := range monkeyList {
		monkeyInteractions = append(monkeyInteractions, m.interactions)
	}
	sort.Ints(monkeyInteractions)

	monkeyBusiness := monkeyInteractions[len(monkeyInteractions)-1] * monkeyInteractions[len(monkeyInteractions)-2]
	return strconv.Itoa(monkeyBusiness)
}

func PartTwo(filename string) string {
	input := readInput(filename)

	monkeyList := CreateMonkeyList(input)

	// Run 20 rounds
	for i := 0; i < 10000; i++ {
		// run each monkey in turn
		for _, monkey := range monkeyList {
			// For each Item eg
			// Monkey inspects an item with a worry level of 79.
			// Worry level is multiplied by 19 to 1501.
			// Monkey gets bored with item. Worry level is divided by 3 to 500.
			// Current worry level is not divisible by 23.
			// Item with worry level 500 is thrown to monkey 3.

			for _, item := range monkey.items {
				// Inspect
				item = monkey.operation(item)
				// Throw
				nextMonkey := monkey.test(item)
				monkeyList[nextMonkey].items = append(monkeyList[nextMonkey].items, item)
				// Remove from this monkeys items
				monkey.items = monkey.items[1:]
				// Add interaction
				monkey.interactions++
			}

		}

	}

	// find top two monkeys by interactions
	monkeyInteractions := []int{}

	for _, m := range monkeyList {
		monkeyInteractions = append(monkeyInteractions, m.interactions)
	}
	sort.Ints(monkeyInteractions)

	monkeyBusiness := monkeyInteractions[len(monkeyInteractions)-1] * monkeyInteractions[len(monkeyInteractions)-2]
	return strconv.Itoa(monkeyBusiness)
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

type Monkey struct {
	number       int
	items        []int
	operation    func(int) int
	test         func(int) int
	interactions int
}

func CreateMonkeyList(input []string) []*Monkey {
	monkeyList := make([]*Monkey, 0)

	// Create monkeys
	for i := 0; i < len(input); i = i + 7 {

		// Each 5 lines is a monkey
		regexMonkeyNumber := regexp.MustCompile(`Monkey ([0-9]):`)
		regexStartingItems := regexp.MustCompile(`  Starting items: (.+)`)
		regexOperation := regexp.MustCompile(`  Operation: new = old ([*|+]) ([0-9]+|old)`)
		regexTest := regexp.MustCompile(`  Test: divisible by ([0-9]+)`)
		regexTestTrue := regexp.MustCompile(`    If true: throw to monkey ([0-9])`)
		regexTestFalse := regexp.MustCompile(`    If false: throw to monkey ([0-9])`)

		monkeyNumString := regexMonkeyNumber.FindStringSubmatch(input[i])
		startingItemsString := regexStartingItems.FindStringSubmatch(input[i+1])
		operationString := regexOperation.FindStringSubmatch(input[i+2])
		testString := regexTest.FindStringSubmatch(input[i+3])
		testTrueString := regexTestTrue.FindStringSubmatch(input[i+4])
		testFalseString := regexTestFalse.FindStringSubmatch(input[i+5])

		monkeyNum, _ := strconv.Atoi(monkeyNumString[1])
		var startingItems []int
		startingItemsSplit := strings.Split(startingItemsString[1], ",")
		for _, item := range startingItemsSplit {
			itemNum, _ := strconv.Atoi(strings.TrimSpace(item))
			startingItems = append(startingItems, itemNum)
		}

		var operation func(i int) int
		if operationString[2] == "old" {
			operation = func(i int) int {
				if operationString[1] == "+" {
					return i + i
				} else {
					return i * i
				}

			}
		} else {
			operationNum, _ := strconv.Atoi(operationString[2])
			operation = func(i int) int {
				if operationString[1] == "+" {
					return i + operationNum
				} else {
					return i * operationNum
				}

			}
		}

		testCaseNum, _ := strconv.Atoi(testString[1])
		testCaseTrueNum, _ := strconv.Atoi(testTrueString[1])
		testCaseFalseNum, _ := strconv.Atoi(testFalseString[1])
		test := func(i int) int {
			if i%testCaseNum == 0 {
				return testCaseTrueNum
			} else {
				return testCaseFalseNum
			}
		}

		newMonkey := Monkey{
			number:       monkeyNum,
			items:        startingItems,
			operation:    operation,
			test:         test,
			interactions: 0,
		}

		monkeyList = append(monkeyList, &newMonkey)

	}

	return monkeyList
}
