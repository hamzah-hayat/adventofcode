package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
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
		if asteroid && i != searchSpot {
			asteroids[i] = true
		}
	}

	asteroidsSeeable := 0
	// For each space, check against another one
	for i := range asteroids {
		colinear := false
		for j := range asteroids {
			if i == j {
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

			// And that I is is closer then j
			if math.Abs(manhatten(searchSpot, i)) < math.Abs(manhatten(searchSpot, j)) {
				continue
			}

			// Check if these two points are colinear (with searchspot as well)
			// 1st point is searchSpot
			// 2nd point is i
			// 3rd point is j
			area := searchSpot.x*(i.y-j.y) + i.x*(j.y-searchSpot.y) + j.x*(searchSpot.y-i.y)
			if area == 0 {
				colinear = true
			}
		}
		if !colinear {
			asteroidsSeeable++
		}
	}

	return asteroidsSeeable
}

func partTwo() {
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

	// Best space is highestSpace
	the200thAsteroid := vapeAsteroids(asteroidMap, highestSpace, 200)

	fmt.Println("The 200th space to be destroyed is", the200thAsteroid)
}

// Starting from the lazer space, face north and destroy each asteroid in order, keep doing this until the 200th asteroid, return that one
func vapeAsteroids(asteroidMap map[space]bool, lazer space, destroyedNum int) spaceWithCenter {

	// Sort the asteroids, then start destroying them
	// Use sort
	seeableAsteroids := getAsteroidsSeenFromSpace(asteroidMap, lazer)
	var sortedAsteroids spaces
	for i, s := range seeableAsteroids {
		if s == true {
			sortedAsteroids = append(sortedAsteroids, spaceWithCenter{x: i.x, y: i.y, center: lazer})
		}
	}

	sort.Sort(sortedAsteroids)

	return sortedAsteroids[destroyedNum-1]
}

// try and sort with center
type spaces []spaceWithCenter

func (a spaces) Len() int      { return len(a) }
func (a spaces) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Check if this space is more counter clockwise then another
func (spaceArray spaces) Less(i int, j int) bool {

	a := spaceArray[i]
	b := spaceArray[j]

	if a.x-a.center.x >= 0 && b.x-a.center.x < 0 {
		return true
	}
	if a.x-a.center.x < 0 && b.x-a.center.x >= 0 {
		return false
	}
	if a.x-a.center.x == 0 && b.x-a.center.x == 0 {
		if a.y-a.center.y >= 0 || b.y-a.center.y >= 0 {
			return a.y > b.y
		}
		return b.y < a.y
	}

	// compute the cross product of vectors (center -> a) x (center -> b)
	// det := ((a.x - a.center.x) * (b.y - a.center.y)) - ((b.x - a.center.x) * (a.y - a.center.y))
	deta := (a.x - a.center.x) * (b.y - a.center.y)
	detb := (b.x - a.center.x) * (a.y - a.center.y)
	det := deta - detb
	if det > 0 {
		return true
	}
	if det < 0 {
		return false
	}

	// points a and b are on the same line from the center
	// check which point is closer to the center
	d1 := (a.x-a.center.x)*(a.x-a.center.x) + (a.y-a.center.y)*(a.y-a.center.y)
	d2 := (b.x-a.center.x)*(b.x-a.center.x) + (b.y-a.center.y)*(b.y-a.center.y)
	return d1 < d2
}

// Get the asteroids we can see from the search space
func getAsteroidsSeenFromSpace(asteroidMap map[space]bool, searchSpot space) map[space]bool {

	// Grab just our asteroids
	asteroids := make(map[space]bool)
	for i, asteroid := range asteroidMap {
		if asteroid && i != searchSpot {
			asteroids[i] = true
		}
	}

	seeable := make(map[space]bool)
	// For each space, check against another one
	for i := range asteroids {
		colinear := false
		for j := range asteroids {
			if i == j {
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

			// And that I is is closer then j
			d1 := (i.x-searchSpot.x)*(i.x-searchSpot.x) + (i.y-searchSpot.y)*(i.y-searchSpot.y)
			d2 := (j.x-searchSpot.x)*(j.x-searchSpot.x) + (j.y-searchSpot.y)*(j.y-searchSpot.y)
			if d1 < d2 {
				continue
			}

			// Check if these two points are colinear (with searchspot as well)
			// 1st point is searchSpot
			// 2nd point is i
			// 3rd point is j
			area := searchSpot.x*(i.y-j.y) + i.x*(j.y-searchSpot.y) + j.x*(searchSpot.y-i.y)
			if area == 0 {
				colinear = true
			}
		}
		if !colinear {
			seeable[i] = true
		}
	}

	return seeable
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

type spaceWithCenter struct {
	x, y   int
	center space
}

func manhatten(first space, second space) float64 {
	return float64(first.x - second.x + first.y - second.y)
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
