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

	// Get random numbers
	randomBingoValues := strings.Split(input[0], ",")
	var randomBingoValuesNum []int
	for _, v := range randomBingoValues {
		num, _ := strconv.Atoi(v)
		randomBingoValuesNum = append(randomBingoValuesNum, num)
	}

	// Get boards
	var boards []Board
	for i := 2; i < len(input)-4; i = i + 6 {
		var b Board
		b.grid = make([][]int, 5)
		b.grid[0] = convertLineToIntArray(input[i])
		b.grid[1] = convertLineToIntArray(input[i+1])
		b.grid[2] = convertLineToIntArray(input[i+2])
		b.grid[3] = convertLineToIntArray(input[i+3])
		b.grid[4] = convertLineToIntArray(input[i+4])

		b.marked = make([][]bool, 5)
		b.marked[0] = []bool{false, false, false, false, false}
		b.marked[1] = []bool{false, false, false, false, false}
		b.marked[2] = []bool{false, false, false, false, false}
		b.marked[3] = []bool{false, false, false, false, false}
		b.marked[4] = []bool{false, false, false, false, false}

		boards = append(boards, b)
	}

	for _, randomNum := range randomBingoValuesNum {
		win := false
		for _, b := range boards {
			MarkBoard(b, randomNum)
			if CheckBoardWin(b) {
				fmt.Println(ScoreBoard(b, randomNum))
				win = true
				break
			}
		}
		if win {
			break
		}
	}

	//fmt.Println(boards[0])
}

func PartTwo() {
	input := readInput()

	// Get random numbers
	randomBingoValues := strings.Split(input[0], ",")
	var randomBingoValuesNum []int
	for _, v := range randomBingoValues {
		num, _ := strconv.Atoi(v)
		randomBingoValuesNum = append(randomBingoValuesNum, num)
	}

	// Get boards
	var boards []Board
	for i := 2; i < len(input)-4; i = i + 6 {
		var b Board
		b.grid = make([][]int, 5)
		b.grid[0] = convertLineToIntArray(input[i])
		b.grid[1] = convertLineToIntArray(input[i+1])
		b.grid[2] = convertLineToIntArray(input[i+2])
		b.grid[3] = convertLineToIntArray(input[i+3])
		b.grid[4] = convertLineToIntArray(input[i+4])

		b.marked = make([][]bool, 5)
		b.marked[0] = []bool{false, false, false, false, false}
		b.marked[1] = []bool{false, false, false, false, false}
		b.marked[2] = []bool{false, false, false, false, false}
		b.marked[3] = []bool{false, false, false, false, false}
		b.marked[4] = []bool{false, false, false, false, false}

		boards = append(boards, b)
	}

	winningBoards := 0
	for _, randomNum := range randomBingoValuesNum {
		win := false

		for i := len(boards) - 1; i > 0; i-- {
			MarkBoard(boards[i], randomNum)
			if CheckBoardWin(boards[i]) {
				fmt.Println(ScoreBoard(boards[i], randomNum))
				win = true
				boards = remove(boards, i)
			}
		}
		if win {
			winningBoards++
		}
	}

	//fmt.Println(winningBoards)
}

type Board struct {
	grid   [][]int
	marked [][]bool
}

// CheckBoardWin takes a bingo board and checks if its won
func CheckBoardWin(b Board) bool {
	// First check each row
	won := false
	for i := 0; i < len(b.marked); i++ {
		won = true
		for j := 0; j < len(b.marked[i]); j++ {
			if !b.marked[i][j] {
				won = false
			}
		}
		if won {
			return won
		}
	}

	// Check each column
	won = false
	for i := 0; i < len(b.marked); i++ {
		won = true
		for j := 0; j < len(b.marked[i]); j++ {
			if !b.marked[j][i] {
				won = false
			}
		}
		if won {
			return won
		}
	}

	return won
}

// MarkBoard takes a number and a board and marks the value, if it has it
func MarkBoard(b Board, num int) {
	for i, v := range b.grid {
		for i2, v2 := range v {
			if v2 == num {
				b.marked[i][i2] = true
			}
		}
	}
}

// ScoreBoard takes a board and scores it
func ScoreBoard(b Board, winNum int) int {
	unMarkedSum := 0

	for i := 0; i < len(b.marked); i++ {
		for j := 0; j < len(b.marked[i]); j++ {
			if !b.marked[i][j] {
				unMarkedSum += b.grid[i][j]
			}
		}
	}

	return winNum * unMarkedSum
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

func convertLineToIntArray(line string) []int {
	// Given a line of string numbers convert to int array
	var intArray []int

	// Remove extra spaces
	capture := regexp.MustCompile(`\s*([0-9]*)\s+([0-9]*)\s+([0-9]*)\s+([0-9]*)\s+([0-9]*)`)

	s := capture.FindStringSubmatch(line)

	for i := 1; i < len(s); i++ {
		num, _ := strconv.Atoi(s[i])
		intArray = append(intArray, num)
	}
	return intArray
}

// remove element from array
func remove(s []Board, i int) []Board {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
