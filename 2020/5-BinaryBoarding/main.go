package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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

	highestSeatID := 0

	for _, v := range input {

		// Find row
		rows := make([]int, 128)
		for i := 0; i < 128; i++ {
			rows[i] = i
		}
		passRow := RecursiveFindRowOrColumn(v[0:7], rows)

		// Find column
		columns := make([]int, 8)
		for i := 0; i < 8; i++ {
			columns[i] = i
		}
		passColumn := RecursiveFindRowOrColumn(v[7:10], columns)

		seatID := passRow*8 + passColumn

		if seatID > highestSeatID {
			highestSeatID = seatID
		}

		fmt.Println("The Unique Seat ID for pass", v, "is", seatID, "its row is", passRow, "and its column is", passColumn)
	}
	fmt.Println("The highest Seat ID is", highestSeatID)
}

func PartTwo() {
	input := readInput()

	seatIDs := make([]int, 0)

	for _, v := range input {

		// Find row
		rows := make([]int, 128)
		for i := 0; i < 128; i++ {
			rows[i] = i
		}
		passRow := RecursiveFindRowOrColumn(v[0:7], rows)

		// Find column
		columns := make([]int, 8)
		for i := 0; i < 8; i++ {
			columns[i] = i
		}
		passColumn := RecursiveFindRowOrColumn(v[7:10], columns)

		seatIDs = append(seatIDs, passRow*8+passColumn)
	}

	sort.Ints(seatIDs)

	//fmt.Println(seatIDs)

	for i := 0; i < len(seatIDs); i++ {
		if !contains(seatIDs, i) {
			fmt.Println(i)
		}
	}

}

// shameleslly copied and pasted from https://stackoverflow.com/questions/10485743/contains-method-for-a-slice because i cba
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func RecursiveFindRowOrColumn(pass string, nums []int) int {

	if len(pass) == 0 {
		return nums[0]
	}

	// F and L are Lower half
	if pass[0] == 'F' || pass[0] == 'L' {
		return RecursiveFindRowOrColumn(pass[1:len(pass)], nums[:len(nums)/2])
	}

	// B and R are Upper half
	if pass[0] == 'B' || pass[0] == 'R' {
		return RecursiveFindRowOrColumn(pass[1:len(pass)], nums[len(nums)/2:])
	}

	return 0
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
