package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	methodP *string
)

func main() {
	methodP = flag.String("method", "p2", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()
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
	input := readInput("input")

	startTemplate := input[0]
	var formula []Formula

	for i := 2; i < len(input); i++ {
		split := strings.Split(input[i], "->")
		formula = append(formula, Formula{strings.Trim(split[0], " "), strings.Trim(split[1], " ")})
	}

	for i := 0; i < 10; i++ {
		startTemplate = RunFormula(startTemplate, formula)
	}

	// count each element
	counts := make(map[string]int)
	for _, c := range startTemplate {
		counts[string(c)]++
	}

	// find least and most
	//least := "A"
	leastSum := 10000
	//most := "A"
	mostSum := 0

	for _, s := range counts {
		if leastSum > s {
			leastSum = s
		}
		if mostSum < s {
			mostSum = s
		}
	}

	fmt.Println(mostSum - leastSum)

}

func PartTwo() {
	input := readInput("input")

	startTemplate := input[0]
	var formula []Formula

	for i := 2; i < len(input); i++ {
		split := strings.Split(input[i], "->")
		formula = append(formula, Formula{strings.Trim(split[0], " "), strings.Trim(split[1], " ")})
	}

	// Instead lets do it like how we did Lanternfish
	// Put pairs in map instead
	pairMap := ConvertStringToMap(startTemplate)

	for i := 0; i < 40; i++ {
		pairMap = RunFormulaMap(pairMap, formula)
	}

	// count each element
	counts := make(map[string]int)
	for s, i := range pairMap {
		counts[string(s[0])] += i
	}

	// Count final element once
	// lol
	counts[string(startTemplate[len(startTemplate)-1])]++

	// find least and most
	//least := "A"
	leastSum := 1000000000000000000
	//most := "A"
	mostSum := 0

	for _, s := range counts {
		if leastSum > s {
			leastSum = s
		}
		if mostSum < s {
			mostSum = s
		}
	}

	fmt.Println(mostSum - leastSum)
}

// Run the formula and output
func RunFormula(currentString string, formula []Formula) string {
	// Regex test
	// look for matchs with LookElement, if true insert InsertElement
	newString := currentString
	for _, v := range formula {
		checkMatch := regexp.MustCompile(v.LookElement)
		replace := string(v.LookElement[0]) + strings.ToLower(v.InsertElement) + string(v.LookElement[1])
		replaceBool := true
		for replaceBool {
			replaceBool = false
			checkNewString := newString
			newString = checkMatch.ReplaceAllLiteralString(newString, replace)
			if checkNewString != newString {
				replaceBool = true
			}
		}
	}
	return strings.ToUpper(newString)
}

func RunFormulaMap(pairMap map[string]int, formula []Formula) map[string]int {
	// Just directly check pairMap
	newMap := make(map[string]int)
	for _, v := range formula {
		for s, i := range pairMap {
			if v.LookElement == s {
				pair1 := string(s[0]) + v.InsertElement
				pair2 := v.InsertElement + string(s[1])
				newMap[pair1] += i
				newMap[pair2] += i
				delete(pairMap, s)
			}
		}
	}

	// Then add maps together
	for i, v := range newMap {
		pairMap[i] = v
	}

	return pairMap
}

// Convert each pair of characters to map
func ConvertStringToMap(str string) map[string]int {
	pairMap := make(map[string]int)
	for i := 0; i < len(str)-1; i++ {
		pair := string(str[i]) + string(str[i+1])
		pairMap[pair]++
	}
	return pairMap
}

type Formula struct {
	LookElement   string
	InsertElement string
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
