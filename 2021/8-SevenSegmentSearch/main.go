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

func init() {
	// Use Flags to run a part
	methodP = flag.String("method", "p2", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()
}

func main() {
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

	// Count number of times 1,4,7,8 appear in output values
	// 1 = len(2)
	// 4 = len(4)
	// 7 = len(3)
	// 8 = len(7)
	sum := 0
	for _, v := range input {
		output := strings.SplitAfter(v, "|")
		outputs := strings.Split(strings.Trim(output[1], " "), " ")
		for _, num := range outputs {
			//output nums
			if len(num) == 2 || len(num) == 3 || len(num) == 4 || len(num) == 7 {
				sum++
			}
		}
	}

	fmt.Println(sum)

}

func PartTwo() {
	input := readInput()

	total := 0

	// Count number of times 1,4,7,8 appear in output values
	// 1 = len(2)
	// 4 = len(4)
	// 7 = len(3)
	// 8 = len(7)
	for _, v := range input {
		var number1 string
		var number4 string
		var number7 string
		var number8 string
		var rest []string

		output := strings.SplitAfter(v, "|")
		outputs := strings.Split(strings.Trim(output[0], " "), " ")

		// First the easy solves
		for _, num := range outputs {
			//output nums
			switch len(num) {
			case 2:
				number1 = num
			case 3:
				number7 = num
			case 4:
				number4 = num
			case 7:
				number8 = num
			default:
				rest = append(rest, num)
			}
		}

		a := "a"
		b := "b"
		c := "c"
		d := "d"
		e := "e"
		f := "f"
		g := "g"

		// Now solve for rest
		// try and work out what each character is
		// a is the character missing between one and seven
		a = strings.Replace(number7, string(number1[0]), "", -1)
		a = strings.Replace(a, string(number1[1]), "", -1)

		// c appears 8 times overall and also in 1 (appears 4 times in rest)
		// f appears 9 times overall and also in 1 (appears 5 times in rest)
		// count number1[0] and number1[1] and put them in above
		sumcf := 0
		if strings.Contains(rest[0], string(number1[0])) {
			sumcf++
		}
		if strings.Contains(rest[1], string(number1[0])) {
			sumcf++
		}
		if strings.Contains(rest[2], string(number1[0])) {
			sumcf++
		}
		if strings.Contains(rest[3], string(number1[0])) {
			sumcf++
		}
		if strings.Contains(rest[4], string(number1[0])) {
			sumcf++
		}
		if strings.Contains(rest[5], string(number1[0])) {
			sumcf++
		}

		if sumcf == 4 {
			c = string(number1[0])
			f = string(number1[1])
		} else {
			c = string(number1[1])
			f = string(number1[0])
		}

		// Now do same for b and d
		// b and d are in 4 (-c and f)
		bd := strings.Replace(number4, c, "", -1)
		bd = strings.Replace(bd, f, "", -1)

		// b appears 4 times in rest and also in bd
		// d appears 5 times in rest and also in bd
		// count bd and put them in above
		sumbd := 0
		if strings.Contains(rest[0], string(bd[0])) {
			sumbd++
		}
		if strings.Contains(rest[1], string(bd[0])) {
			sumbd++
		}
		if strings.Contains(rest[2], string(bd[0])) {
			sumbd++
		}
		if strings.Contains(rest[3], string(bd[0])) {
			sumbd++
		}
		if strings.Contains(rest[4], string(bd[0])) {
			sumbd++
		}
		if strings.Contains(rest[5], string(bd[0])) {
			sumbd++
		}

		if sumbd == 4 {
			b = string(bd[0])
			d = string(bd[1])
		} else {
			b = string(bd[1])
			d = string(bd[0])
		}

		// Do eg now
		// e and g are in 8 (-abcdf)
		eg := strings.ReplaceAll(number8, a, "")
		eg = strings.ReplaceAll(eg, b, "")
		eg = strings.ReplaceAll(eg, c, "")
		eg = strings.ReplaceAll(eg, d, "")
		eg = strings.ReplaceAll(eg, f, "")

		// e appears 3 times in rest and also in eg
		// g appears 6 times in rest and also in ge
		// count bd and put them in above
		sumeg := 0
		if strings.Contains(rest[0], string(eg[0])) {
			sumeg++
		}
		if strings.Contains(rest[1], string(eg[0])) {
			sumeg++
		}
		if strings.Contains(rest[2], string(eg[0])) {
			sumeg++
		}
		if strings.Contains(rest[3], string(eg[0])) {
			sumeg++
		}
		if strings.Contains(rest[4], string(eg[0])) {
			sumeg++
		}
		if strings.Contains(rest[5], string(eg[0])) {
			sumeg++
		}

		if sumeg == 3 {
			e = string(eg[0])
			g = string(eg[1])
		} else {
			e = string(eg[1])
			g = string(eg[0])
		}

		//fmt.Println(a + b + c + d + e + f + g)

		// now that we have all characters, we should be able to work out the outputs
		checks := strings.Split(strings.Trim(output[1], " "), " ")

		numString := ""

		for _, checkNumber := range checks {
			numString += strconv.Itoa(checkNum(a, b, c, d, e, f, g, checkNumber))
			//fmt.Println(checkNumber, "=", checkNum(a, b, c, d, e, f, g, checkNumber))
		}

		backtoInt, _ := strconv.Atoi(numString)
		total += backtoInt
		//fmt.Println(backtoInt)

	}

	fmt.Println(total)

}

// Given a set of input values and a checkNum, work out what this number is
func checkNum(a, b, c, d, e, f, g string, checkNum string) int {
	hasa := strings.ContainsAny(checkNum, a)
	hasb := strings.ContainsAny(checkNum, b)
	hasc := strings.ContainsAny(checkNum, c)
	hasd := strings.ContainsAny(checkNum, d)
	hase := strings.ContainsAny(checkNum, e)
	hasf := strings.ContainsAny(checkNum, f)
	hasg := strings.ContainsAny(checkNum, g)

	// Check each number
	if hasa && hasb && hasc && hase && hasf && hasg && len(checkNum) == 6 {
		return 0
	}

	if hasc && hasf && len(checkNum) == 2 {
		return 1
	}

	if hasa && hasc && hasd && hase && hasg && len(checkNum) == 5 {
		return 2
	}

	if hasa && hasc && hasd && hasf && hasg && len(checkNum) == 5 {
		return 3
	}

	if hasb && hasc && hasd && hasf && len(checkNum) == 4 {
		return 4
	}

	if hasa && hasb && hasd && hasf && hasg && len(checkNum) == 5 {
		return 5
	}

	if hasa && hasb && hasd && hase && hasf && hasg && len(checkNum) == 6 {
		return 6
	}

	if hasa && hasc && hasf && len(checkNum) == 3 {
		return 7
	}

	if hasa && hasb && hasc && hasd && hase && hasf && hasg && len(checkNum) == 7 {
		return 8
	}

	return 9
}

// Read data from input.txt
// Return the string, so that we can deal with it however
func readInput() []string {

	var input []string

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
