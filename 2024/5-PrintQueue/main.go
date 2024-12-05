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

	total := 0
	pageRules := make([]PageRule, 0)

	// Read Page rules
	ReadingUpdates := false
	for _, v := range input {
		if !ReadingUpdates {
			if v == "" {
				ReadingUpdates = true
				continue
			}
			splitPageRule := strings.Split(v, "|")
			pageRules = append(pageRules, PageRule{splitPageRule[0], splitPageRule[1]})

		} else {
			if checkUpdate(v, pageRules) {
				// Get Middle page and add to total
				updateSplit := strings.Split(v, ",")
				middlePage := updateSplit[(len(updateSplit)-1)/2]
				middlePageNum, _ := strconv.Atoi(middlePage)
				total += middlePageNum
			}
		}
	}

	return strconv.Itoa(total)
}

func checkUpdate(update string, pageRules []PageRule) bool {
	splitUpdate := strings.Split(update, ",")

	for currentPage, page := range splitUpdate {
		for _, rule := range pageRules {
			if page == rule.before {
				for j := currentPage; j >= 0; j-- {
					if rule.after == splitUpdate[j] {
						return false
					}
				}
			}
		}
	}

	return true
}

// PageRule is a struct that shows which page must be printed before another
type PageRule struct {
	before string
	after  string
}

func PartTwo(filename string) string {
	input := readInput(filename)

	total := 0
	pageRules := make([]PageRule, 0)

	// Read Page rules
	ReadingUpdates := false
	for _, v := range input {
		if !ReadingUpdates {
			if v == "" {
				ReadingUpdates = true
				continue
			}
			splitPageRule := strings.Split(v, "|")
			pageRules = append(pageRules, PageRule{splitPageRule[0], splitPageRule[1]})

		} else {
			if checkUpdate(v, pageRules) {
				// // Get Middle page and add to total
				// updateSplit := strings.Split(v, ",")
				// middlePage := updateSplit[(len(updateSplit)-1)/2]
				// middlePageNum, _ := strconv.Atoi(middlePage)
				// total += middlePageNum
			} else {
				// Fix incorrect updates then get total
				updateSplit := strings.Split(fixUpdate(v, pageRules), ",")
				middlePage := updateSplit[(len(updateSplit)-1)/2]
				middlePageNum, _ := strconv.Atoi(middlePage)
				total += middlePageNum
			}
		}
	}

	return strconv.Itoa(total)
}

func fixUpdate(update string, pageRules []PageRule) string {

	updateSplit := strings.Split(update, ",")

	// https://hectorcorrea.com/blog/2022-08-30/i-can-t-believe-it-can-sort-visualized
	// Sort the list using the rules
	for i := 0; i < len(updateSplit); i++ {
		for j := 0; j < len(updateSplit); j++ {
			for _, r := range pageRules {
				if updateSplit[i] == r.before {
					if updateSplit[j] == r.after {
						// Swap
						temp := updateSplit[j]
						updateSplit[j] = updateSplit[i]
						updateSplit[i] = temp
					}
				}
			}
		}
	}

	sortedSplit := strings.Join(updateSplit, ",")

	return sortedSplit
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
