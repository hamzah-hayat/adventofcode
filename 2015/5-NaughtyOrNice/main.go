package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p1", "The method/part that should be run, valid are p1,p2 and test")
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
	input := readInput()

	nice := 0

	for _, text := range strings.Split(input, "\n") {
		// Check the string
		if CheckNiceString(text) {
			nice++
		}
	}

	fmt.Println("The total number of strings is", len(strings.Split(input, "\n")))
	fmt.Println("The number of nice strings is", nice)
}

func CheckNiceString(text string) bool {
	// Check the string
	//fmt.Println("Checking string", text)
	// check for ab, cd, pq, or xy
	matched, _ := regexp.MatchString(`ab|cd|pq|xy`, text)
	if matched {
		//fmt.Println("Failed test 1")
		return false
	}

	// one letter that appears twice in a row
	// This is a nice requirement so we negate the match
	matched, _ = regexp.MatchString(`aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz`, text)
	if !matched {
		//fmt.Println("Failed test 2")
		return false
	}

	// At least three vowels
	// This is a nice requirement so we negate the match
	matched, _ = regexp.MatchString(`[aeiou].*[aeiou].*[aeiou]`, text)
	if !matched {
		//fmt.Println("Failed test 3")
		return false
	}

	return true
}

func PartTwo() {
	input := readInput()

	nice := 0

	for _, text := range strings.Split(input, "\n") {
		// Check the string
		if CheckNiceStringV2(text) {
			nice++
		}
	}

	fmt.Println("The total number of strings is", len(strings.Split(input, "\n")))
	fmt.Println("The number of nice strings is", nice)
}

func CheckNiceStringV2(text string) bool {
	// Check the string

	// A nice string needs a pair of characters that are not overlapping
	alphabet := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	// We cant use backsearch in go so build the regex expression manually, quite costly but it works
	matchedAtLeastOnce := false
	for _, a := range alphabet {
		for _, b := range alphabet {
			matched, _ := regexp.MatchString(a+b+".*"+a+b, text)
			if matched {
				matchedAtLeastOnce = true
				//fmt.Println("match text was", text, " using ", a, b)
				break
			}
		}
	}

	if !matchedAtLeastOnce {
		//fmt.Println("Failed test 1")
		return false
	}

	// contains a repeating character with at least a single character between them
	// This is a nice requirement so we negate the match
	matched, _ := regexp.MatchString(`a.a|b.b|c.c|d.d|e.e|f.f|g.g|h.h|i.i|j.j|k.k|l.l|m.m|n.n|o.o|p.p|q.q|r.r|s.s|t.t|u.u|v.v|w.w|x.x|y.y|z.z`, text)
	if !matched {
		//fmt.Println("Failed test 2")
		return false
	}

	return true
}

// Read data from input.txt
// Return the string, so that we can deal with it however
func readInput() string {

	var input string

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}
	return input
}
