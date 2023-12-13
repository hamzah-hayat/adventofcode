package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
	input := readInput(filename)

	var result int

	// Grab each pattern
	var patterns [][]string
	var pattern []string
	for _, v := range input {
		if v != "" {
			pattern = append(pattern, v)
		} else {
			patterns = append(patterns, pattern)
			pattern = make([]string, 0)
		}
	}
	patterns = append(patterns, pattern)

	for _, p := range patterns {
		// Find the row or column reflection
		reflectRow := findRowReflection(p)
		if reflectRow != -1 {
			result += reflectRow
		}

		// Find the row or column reflection
		transpose := Transpose(p)
		reflectColumn := findRowReflection(transpose)
		if reflectColumn != -1 {
			result += reflectColumn * 100
		}
	}

	num := strconv.Itoa(result)

	return num
}

func findRowReflection(p []string) int {
	reflectRow := -1
	possibleReflectpoints := makeNumberList(1, len(p[0])-1)
	reflectMap := make(map[int]int)
	// Check each row
	for _, line := range p {
		// Can we "reflect" on each possibe Reflect points?
		for _, reflect := range possibleReflectpoints {
			leftString := line[:reflect]
			rightString := line[reflect:]

			if len(leftString) < len(rightString) {
				allMatch := true
				for i := 0; i < len(leftString); i++ {
					if leftString[len(leftString)-i-1] != rightString[i] {
						allMatch = false
					}
				}
				if allMatch {
					reflectMap[reflect] = reflectMap[reflect] + 1
				}
			}

			if len(rightString) < len(leftString) {
				allMatch := true
				for i := 0; i < len(rightString); i++ {
					if leftString[len(leftString)-i-1] != rightString[i] {
						allMatch = false
					}
				}
				if allMatch {
					reflectMap[reflect] = reflectMap[reflect] + 1
				}
			}
		}
	}

	// Find our answer
	// Check map for whichever key has number equal to len(p)
	for i, v := range reflectMap {
		if v == len(p) {
			reflectRow = i
		}
	}

	return reflectRow
}

func findRowReflectionIgnore(p []string, ignore int) int {
	reflectRow := -1
	possibleReflectpoints := makeNumberList(1, len(p[0])-1)
	reflectMap := make(map[int]int)
	// Check each row
	for _, line := range p {
		// Can we "reflect" on each possibe Reflect points?
		for _, reflect := range possibleReflectpoints {
			leftString := line[:reflect]
			rightString := line[reflect:]

			if len(leftString) < len(rightString) {
				allMatch := true
				for i := 0; i < len(leftString); i++ {
					if leftString[len(leftString)-i-1] != rightString[i] {
						allMatch = false
					}
				}
				if allMatch {
					reflectMap[reflect] = reflectMap[reflect] + 1
				}
			}

			if len(rightString) < len(leftString) {
				allMatch := true
				for i := 0; i < len(rightString); i++ {
					if leftString[len(leftString)-i-1] != rightString[i] {
						allMatch = false
					}
				}
				if allMatch {
					reflectMap[reflect] = reflectMap[reflect] + 1
				}
			}
		}
	}

	// Find our answer
	// Check map for whichever key has number equal to len(p)
	for i, v := range reflectMap {
		if v == len(p) && i != ignore {
			reflectRow = i
		}
	}

	return reflectRow
}

func makeNumberList(min, max int) []int {
	list := make([]int, max-min+1)
	for i := range list {
		list[i] = min + i
	}
	return list
}

func Transpose(s []string) []string {
	// Transpose
	var transpose []string
	for i := 0; i < len(s[0]); i++ {
		transposeLine := ""
		for _, v := range s {
			transposeLine += string(v[i])
		}
		transpose = append(transpose, transposeLine)
	}
	return transpose
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// To low 29983
func PartTwo(filename string) string {
	input := readInput(filename)

	var result int

	// Grab each pattern
	var patterns [][]string
	var pattern []string
	for _, v := range input {
		if v != "" {
			pattern = append(pattern, v)
		} else {
			patterns = append(patterns, pattern)
			pattern = make([]string, 0)
		}
	}
	patterns = append(patterns, pattern)

	for _, p := range patterns {
		//fmt.Println("Current Result is ", result, " while i is ", i)
		// Find the row or column reflection
		orgReflectRow := findRowReflection(p)

		// Find the row or column reflection
		transpose := Transpose(p)
		orgReflectColumn := findRowReflection(transpose)

		// Now that we have original reflection lines
		// we need to change a single part of p
		// brute force
		newReflectRow := -1
		newReflectColumn := -1

		for x := 0; x < len(p); x++ {
			for y := 0; y < len(p[0]); y++ {
				perm := makePossiblePatterns(p, x, y)
				// Find the row or column reflection
				reflectRow := findRowReflectionIgnore(perm, orgReflectRow)

				// Find the row or column reflection
				transpose2 := Transpose(perm)
				reflectColumn := findRowReflectionIgnore(transpose2, orgReflectColumn)

				if reflectRow != orgReflectRow && reflectRow != -1 {
					newReflectRow = reflectRow
				}

				if reflectColumn != orgReflectColumn && reflectColumn != -1 {
					newReflectColumn = reflectColumn
				}
			}
		}

		add := 0
		if newReflectRow != -1 {
			result += newReflectRow
			add += newReflectRow
		} else if newReflectColumn != -1 {
			result += newReflectColumn * 100
			add += newReflectColumn * 100
		}
	}

	num := strconv.Itoa(result)

	return num
}

func makePossiblePatterns(p []string, y, x int) []string {
	var copy []string

	for _, v := range p {
		copy = append(copy, v)
	}

	if copy[y][x] == '.' {
		copy[y] = copy[y][:x] + string('#') + copy[y][x+1:]
	} else if copy[y][x] == '#' {
		copy[y] = copy[y][:x] + string('.') + copy[y][x+1:]
	}

	return copy
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
