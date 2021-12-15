package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//PartOne()
	PartTwo()
}

func PartOne() {
	input := readInput()
	fmt.Println(FindOverlap(input))
}

func PartTwo() {
	input := readInput()
	fmt.Println(FindNonOverlapClaim(input))
}

type claim struct {
	id     int
	left   int
	top    int
	width  int
	height int
}

// Read data from input.txt
// Load it into claim array
func readInput() []claim {

	var input []claim

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() != "" {
			line := scanner.Text()

			fields := strings.Fields(line)

			var newClaim claim
			newClaim.id, _ = strconv.Atoi(strings.TrimLeft(fields[0], "#"))
			newClaim.left, _ = strconv.Atoi(strings.Split(strings.TrimRight(fields[2], ":"), ",")[0])
			newClaim.top, _ = strconv.Atoi(strings.Split(strings.TrimRight(fields[2], ":"), ",")[1])
			newClaim.width, _ = strconv.Atoi(strings.Split(fields[3], "x")[0])
			newClaim.height, _ = strconv.Atoi(strings.Split(fields[3], "x")[1])

			input = append(input, newClaim)
		}
	}
	return input
}

// Find the amount of overlap between all the claims
func FindOverlap(input []claim) int {
	overlap := 0
	var claimedCloth [1000][1000]int

	for _, claim := range input {
		// Claim some cloth
		for i := 0; i < claim.width; i++ {
			for j := 0; j < claim.height; j++ {
				claimedCloth[claim.left+i][claim.top+j]++
			}
		}
	}

	// Now iterate over finished array and find anything which has been claimed more then once
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if claimedCloth[i][j] > 1 {
				overlap++
			}
		}
	}

	return overlap
}

// Find the Claim which doesnt overlap with anything else
func FindNonOverlapClaim(input []claim) string {
	claimID := ""
	var claimedCloth [1000][1000]int

	for _, claim := range input {
		// Claim some cloth
		for i := 0; i < claim.width; i++ {
			for j := 0; j < claim.height; j++ {
				claimedCloth[claim.left+i][claim.top+j]++
			}
		}
	}

	// Find a claim that doesnt overlap
	for _, claim := range input {
		noOverlap := true

		// Check the claim
		for i := 0; i < claim.width; i++ {
			for j := 0; j < claim.height; j++ {
				if claimedCloth[claim.left+i][claim.top+j] > 1 {
					noOverlap = false
				}
			}
		}

		if noOverlap {
			claimID = strconv.Itoa(claim.id)
		}
	}
	return claimID
}
