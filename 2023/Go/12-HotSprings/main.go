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

	possiblePatterns := 0
	for _, v := range input {
		split := strings.Split(v, " ")
		gears := split[0]
		pattern := split[1]

		patterns, _ := solveGearPatterns(gears, pattern)

		possiblePatterns += patterns
	}

	num := strconv.Itoa(possiblePatterns)

	return num
}

func solveGearPatterns(gears, pattern string) (int, []string) {
	var matchs []string
	patterns := 0

	possibleGearStrings := expandGears(gears)

	regexString := ""
	split := strings.Split(pattern, ",")
	regexString += `^\.*`
	for i, v := range split {
		regexString += `#{` + string(v) + `}`

		if i < len(split)-1 {
			regexString += `\.+`
		}
	}
	regexString += `\.*$`

	r, _ := regexp.Compile(regexString)

	for _, v := range possibleGearStrings {
		if r.Match([]byte(v)) {
			matchs = append(matchs, v+" "+pattern)
			patterns++
		}
	}

	return patterns, matchs
}

func expandGears(gears string) []string {
	var possibleGearStrings []string

	possibleGearStrings = expandGearsRecurse(gears)

	return possibleGearStrings
}

func expandGearsRecurse(gear string) []string {
	var possibleGearStrings []string
	if strings.Contains(gear, "?") {
		notGear := strings.Replace(gear, "?", ".", 1)
		gear := strings.Replace(gear, "?", "#", 1)
		possibleGearStrings = append(possibleGearStrings, expandGearsRecurse(notGear)...)
		possibleGearStrings = append(possibleGearStrings, expandGearsRecurse(gear)...)
	} else {
		possibleGearStrings = append(possibleGearStrings, gear)
	}
	return possibleGearStrings
}

// To low 7405945528484
// To low 23741531685985
// Wrong 31837419160301
func PartTwo(filename string) string {
	input := readInput(filename)

	possiblePatterns := 0
	for i, v := range input {
		split := strings.Split(v, " ")
		gears := split[0]
		pattern := split[1]

		r := solveGearPatternsPart2R(gears, pattern)
		l := solveGearPatternsPart2L(gears, pattern)

		if r > l {
			possiblePatterns += r
		} else {
			possiblePatterns += l
		}

		println(i)
	}

	num := strconv.Itoa(possiblePatterns)

	return num
}

// Correct, but never finishes :(
// func PartTwo(filename string) string {
// 	input := readInput(filename)

// 	possiblePatterns := 0
// 	for _, v := range input {
// 		split := strings.Split(v, " ")
// 		gears := split[0]
// 		pattern := split[1]

// 		p, _ := solveGearPatterns(gears, pattern)
// 		p2, _ := solveGearPatterns(gears+"#", pattern)
// 		p3, _ := solveGearPatterns(gears+".", pattern)

// 		multiplier := p2 + p3
// 		possiblePatterns += p * multiplier * multiplier * multiplier * multiplier
// 	}

// 	num := strconv.Itoa(possiblePatterns)

// 	return num
// }

func unfoldGear(gear string, size int) string {

	split := strings.Split(gear, " ")

	gears := split[0]
	pattern := split[1]

	// Unfold by copying gears by 5 (with seperator ?)
	// Copy pattern by 5 (with seperator ,)
	newGears := ""
	newPattern := ""

	for i := 0; i < size; i++ {
		newGears += gears + "?"
		newPattern += pattern + ","
	}

	return newGears[:len(newGears)-1] + " " + newPattern[:len(newPattern)-1]
}

func unfoldGears(input []string, size int) []string {
	var unfoldedGears []string

	for _, v := range input {
		unfoldedGears = append(unfoldedGears, unfoldGear(v, size))
	}

	return unfoldedGears
}

func solveGearPatternsPart2R(gears, pattern string) int {
	patterns := 0

	_, firstMatchs := solveGearPatterns(gears, pattern)

	for _, m := range firstMatchs {
		secondUnfold := unfoldGearWithOrgR(gears, m, 2)
		split := strings.Split(secondUnfold, " ")
		gears := split[0]
		pattern := split[1]
		_, secondMatchs := solveGearPatterns(gears, pattern)
		for _, m := range secondMatchs {
			thirdUnfold := unfoldGearWithOrgR(gears, m, 2)
			split := strings.Split(thirdUnfold, " ")
			gears := split[0]
			pattern := split[1]
			_, thirdMatchs := solveGearPatterns(gears, pattern)
			for _, m := range thirdMatchs {
				fourthUnfold := unfoldGearWithOrgR(gears, m, 2)
				split := strings.Split(fourthUnfold, " ")
				gears := split[0]
				pattern := split[1]
				_, fourtMatchs := solveGearPatterns(gears, pattern)
				for _, m := range fourtMatchs {
					fifthUnfold := unfoldGearWithOrgR(gears, m, 2)
					split := strings.Split(fifthUnfold, " ")
					gears := split[0]
					pattern := split[1]
					_, fifthMatchs := solveGearPatterns(gears, pattern)
					patterns += len(firstMatchs) * len(secondMatchs) * len(thirdMatchs) * len(fourtMatchs) * len(fifthMatchs)
					break
				}
				break
			}
			break
		}
		break
	}

	return patterns
}

func solveGearPatternsPart2L(gears, pattern string) int {
	patterns := 0

	_, firstMatchs := solveGearPatterns(gears, pattern)

	for _, m := range firstMatchs {
		secondUnfold := unfoldGearWithOrgL(gears, m, 2)
		split := strings.Split(secondUnfold, " ")
		gears := split[0]
		pattern := split[1]
		_, secondMatchs := solveGearPatterns(gears, pattern)
		for _, m := range secondMatchs {
			thirdUnfold := unfoldGearWithOrgL(gears, m, 2)
			split := strings.Split(thirdUnfold, " ")
			gears := split[0]
			pattern := split[1]
			_, thirdMatchs := solveGearPatterns(gears, pattern)
			for _, m := range thirdMatchs {
				fourthUnfold := unfoldGearWithOrgL(gears, m, 2)
				split := strings.Split(fourthUnfold, " ")
				gears := split[0]
				pattern := split[1]
				_, fourtMatchs := solveGearPatterns(gears, pattern)
				for _, m := range fourtMatchs {
					fifthUnfold := unfoldGearWithOrgL(gears, m, 2)
					split := strings.Split(fifthUnfold, " ")
					gears := split[0]
					pattern := split[1]
					_, fifthMatchs := solveGearPatterns(gears, pattern)
					patterns += len(firstMatchs) * len(secondMatchs) * len(thirdMatchs) * len(fourtMatchs) * len(fifthMatchs)
					break
				}
				break
			}
			break
		}
		break
	}

	return patterns
}

func unfoldGearWithOrgR(org string, gear string, size int) string {

	split := strings.Split(gear, " ")

	g := split[0]
	pattern := split[1]

	// Unfold by copying gears by size (with seperator ?)
	// Copy pattern by size (with seperator ,)
	newGears := g + "?"
	newPattern := pattern + ","

	for i := 1; i < size; i++ {
		newGears += org + "?"
		newPattern += pattern + ","
	}

	return newGears[:len(newGears)-1] + " " + newPattern[:len(newPattern)-1]
}

func unfoldGearWithOrgL(org string, gear string, size int) string {

	split := strings.Split(gear, " ")

	g := split[0]
	pattern := split[1]

	// Unfold by copying gears by size (with seperator ?)
	// Copy pattern by size (with seperator ,)
	newGears := "?" + g
	newPattern := "," + pattern

	for i := 1; i < size; i++ {
		newGears = "?" + org + newGears
		newPattern = "," + pattern + newPattern
	}

	return newGears[1:] + " " + newPattern[1:]
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
