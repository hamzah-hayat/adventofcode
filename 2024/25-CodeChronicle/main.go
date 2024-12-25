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

	keys := []string{}
	locks := []string{}

	// Make our locks and keys, supper scuffed xd
	for i := 0; i < len(input); i = i + 8 {
		if input[i] == "#####" {
			// Lock
			firstLockNum := isHash(input[i+1][0]) + isHash(input[i+2][0]) + isHash(input[i+3][0]) + isHash(input[i+4][0]) + isHash(input[i+5][0])
			secondLockNum := isHash(input[i+1][1]) + isHash(input[i+2][1]) + isHash(input[i+3][1]) + isHash(input[i+4][1]) + isHash(input[i+5][1])
			thirdLockNum := isHash(input[i+1][2]) + isHash(input[i+2][2]) + isHash(input[i+3][2]) + isHash(input[i+4][2]) + isHash(input[i+5][2])
			fourthLockNum := isHash(input[i+1][3]) + isHash(input[i+2][3]) + isHash(input[i+3][3]) + isHash(input[i+4][3]) + isHash(input[i+5][3])
			fifthLockNum := isHash(input[i+1][4]) + isHash(input[i+2][4]) + isHash(input[i+3][4]) + isHash(input[i+4][4]) + isHash(input[i+5][4])
			lockStr := []string{strconv.Itoa(firstLockNum), strconv.Itoa(secondLockNum), strconv.Itoa(thirdLockNum), strconv.Itoa(fourthLockNum), strconv.Itoa(fifthLockNum)}
			locks = append(locks, strings.Join(lockStr, ","))
		} else {
			// Key
			firstkeyNum := 5 - (isDot(input[i+1][0]) + isDot(input[i+2][0]) + isDot(input[i+3][0]) + isDot(input[i+4][0]) + isDot(input[i+5][0]))
			secondkeyNum := 5 - (isDot(input[i+1][1]) + isDot(input[i+2][1]) + isDot(input[i+3][1]) + isDot(input[i+4][1]) + isDot(input[i+5][1]))
			thirdkeyNum := 5 - (isDot(input[i+1][2]) + isDot(input[i+2][2]) + isDot(input[i+3][2]) + isDot(input[i+4][2]) + isDot(input[i+5][2]))
			fourthkeyNum := 5 - (isDot(input[i+1][3]) + isDot(input[i+2][3]) + isDot(input[i+3][3]) + isDot(input[i+4][3]) + isDot(input[i+5][3]))
			fifthkeyNum := 5 - (isDot(input[i+1][4]) + isDot(input[i+2][4]) + isDot(input[i+3][4]) + isDot(input[i+4][4]) + isDot(input[i+5][4]))
			keyStr := []string{strconv.Itoa(firstkeyNum), strconv.Itoa(secondkeyNum), strconv.Itoa(thirdkeyNum), strconv.Itoa(fourthkeyNum), strconv.Itoa(fifthkeyNum)}
			keys = append(keys, strings.Join(keyStr, ","))
		}
	}

	// Now check combinations
	// Try every key with every lock
	itFits := 0
	for k := 0; k < len(keys); k++ {
		for l := 0; l < len(locks); l++ {
			if KeyFitsLock(keys[k], locks[l]) {
				itFits++
			}
		}
	}

	return strconv.Itoa(itFits)
}

func KeyFitsLock(key, lock string) bool {

	keyNums := strings.Split(key, ",")
	lockNums := strings.Split(lock, ",")
	for i := 0; i < 5; i++ {
		keyNum, _ := strconv.Atoi(keyNums[i])
		lockNum, _ := strconv.Atoi(lockNums[i])

		if keyNum+lockNum > 5 {
			return false
		}
	}
	return true
}

func isHash(character byte) int {
	if character == '#' {
		return 1
	} else {
		return 0
	}
}

func isDot(character byte) int {
	if character == '.' {
		return 1
	} else {
		return 0
	}
}

func PartTwo(filename string) string {
	return "There is no part two"
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
