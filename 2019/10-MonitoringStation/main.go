package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p1", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()

	switch *methodP {
	case "p1":
		partOne()
		break
	case "p2":
		partTwo()
		break
	case "test":
		break
	}
}

func partOne() {
	input := readInput()

	asteroidMap := createAsteroidMap(input)

	highest := 0
	var highestSpace space

	for i, value := range asteroidMap {
		if value {
			num := numberOfAsteroidsSeenFromSpace(asteroidMap, i)
			if num > highest {
				highest = num
				highestSpace = i
			}
		}
	}

	fmt.Println("The best space is", highestSpace, ", A total of", highest, "asteroids can be seen from here")

}

// Check how many asteriods can be seen from this space
func numberOfAsteroidsSeenFromSpace(asteroidMap map[space]bool, searchSpot space) int {

	// Grab just our asteroids
	asteroids := make(map[space]bool)
	for i, asteroid := range asteroidMap {
		if asteroid {
			asteroids[i] = true
		}
	}

	asteroidsSeeable := 0
	// For each space, check against another one
	for i := range asteroids {
		colinear := false
		asteroidsSeeable = 0
		for j := range asteroids {
			if i == j || i == searchSpot || j == searchSpot {
				continue
			}
			// Also make sure that both i and j are on the same "side"
			if i.x > searchSpot.x && j.x < searchSpot.x {
				continue
			}
			if i.y > searchSpot.y && j.y < searchSpot.y {
				continue
			}
			if i.x < searchSpot.x && j.x > searchSpot.x {
				continue
			}
			if i.y < searchSpot.y && j.y > searchSpot.y {
				continue
			}

			// Check if these two points are colinear (with searchspot as well)
			// 1st point is searchSpot
			// 2nd point is i
			// 3rd point is j
			area := searchSpot.x*(i.y-j.y) + i.x*(j.y-searchSpot.y) + j.x*(searchSpot.y-i.y)
			if area == 0 {
				colinear = true
				break
			}
		}
		if !colinear {
			asteroidsSeeable++
		}
	}

	return asteroidsSeeable
}

func partTwo() {
	//input := readInput()
}

func createAsteroidMap(input []string) map[space]bool {
	asteroidMap := make(map[space]bool)

	for i := 0; i < len(input); i++ {
		for i2, s := range strings.Split(input[i], "") {
			if s == "." {
				asteroidMap[space{x: i2, y: i}] = false
			} else {
				asteroidMap[space{x: i2, y: i}] = true
			}
		}
	}
	return asteroidMap
}

type space struct {
	x, y int
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

func convertToInts(input []string) []int {
	programStr := strings.Split(input[0], ",")
	var program []int

	for _, s := range programStr {
		i, _ := strconv.Atoi(s)
		program = append(program, i)
	}
	return program
}
