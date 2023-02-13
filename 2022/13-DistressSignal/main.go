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
	indexSum := 0

	currentIndex := 1
	for i := 0; i < len(input)-1; i = i + 3 {
		leftList := input[i]
		rightList := input[i+1]

		if InRightOrder(leftList, rightList) {
			indexSum += currentIndex
		}
		currentIndex++
	}

	return strconv.Itoa(indexSum)
}

// Check if lists are in order
func InRightOrder(leftList, rightList string) bool {
	// First we need to split on commas, while keeping any lists intact
	leftListSplit := SplitLineIntoLists(leftList)
	rightListSplit := SplitLineIntoLists(rightList)

	// // if one is a list and the other isnt, convert the odd one out to a list by adding brackets
	// if len(leftList) == 1 && len(rightList) > 1 {
	// 	leftList = "[" + leftList + "]"
	// } else if len(leftList) > 1 && len(rightList) == 1 {
	// 	rightList = "[" + rightList + "]"
	// }

	// If we have two lists like such [1,2,3,4] [1,2,3,4]
	// aka only one set of brackets, then we compare numbers
	// If both values are integers, the lower integer should come first.
	// If the left integer is lower than the right integer, the inputs are in the right order.
	// If the left integer is higher than the right integer, the inputs are not in the right order.
	// Otherwise, the inputs are the same integer; continue checking the next part of the input
	if HasNoListsInside(leftListSplit) && HasNoListsInside(rightListSplit) {
		// Check if one list is bigger then the other
		if len(leftListSplit) > len(rightListSplit) {
			for i := 0; i < len(rightListSplit); i++ {
				leftNum, _ := strconv.Atoi(leftListSplit[i])
				rightNum, _ := strconv.Atoi(rightListSplit[i])
				if leftNum < rightNum {
					return true
				}
				if leftNum > rightNum {
					return false
				}
			}
			return false
		} else {
			for i := 0; i < len(leftListSplit); i++ {
				leftNum, _ := strconv.Atoi(leftListSplit[i])
				rightNum, _ := strconv.Atoi(rightListSplit[i])
				if leftNum < rightNum {
					return true
				}
				if leftNum > rightNum {
					return false
				}
			}
			return true
		}
	}

	// We need to go deeper
	order := true
	for j := 0; j < len(leftListSplit); j++ {
		if !InRightOrder(leftListSplit[j], rightListSplit[j]) {
			order = false
		}
	}

	return order
}

func HasNoListsInside(listSplit []string) bool {
	// make sure there are no lists inside this
	for _, v := range listSplit {
		if strings.ContainsAny(v, "[]") {
			return false
		}
	}
	return true
}

func SplitLineIntoLists(stringList string) []string {
	var list []string

	if stringList == "" {
		return nil
	}
	if BracketNumber(stringList) == 0 {
		stringList = "[" + stringList + "]"
	}

	// First trim brackets
	trimmed := stringList[1 : len(stringList)-1]

	// Then we find all the commas to split on
	addStr := ""
	bracketNum := 0

	for _, v := range trimmed {
		if string(v) == "," && bracketNum == 0 {
			// Add this string to list and continue
			list = append(list, addStr)
			addStr = ""
			continue
		}
		if string(v) == "[" {
			bracketNum++
		} else if string(v) == "]" {
			bracketNum--
		}
		addStr += string(v)
	}
	// add last list
	list = append(list, addStr)

	return list
}

func StripBrackets(str string) string {
	newStr := strings.TrimLeft(str, "[")
	newStr = strings.TrimRight(newStr, "]")
	return newStr
}

func BracketNumber(str string) int {
	brackets := 0
	for _, v := range str {
		if (string(v) == "[") || (string(v) == "]") {
			brackets += 1
		}
	}
	return brackets
}

func PartTwo(filename string) string {
	input := readInput(filename)

	return input[0]
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
