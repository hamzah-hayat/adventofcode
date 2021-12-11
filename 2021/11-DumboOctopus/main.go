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

func init() {
	// Use Flags to run a part
	methodP = flag.String("method", "p1", "The method/part that should be run, valid are p1,p2 and test")
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

	var jellyfish map[Point]JellyFish
	jellyfish = make(map[Point]JellyFish)
	for x, v := range input {
		for y, c := range v {
			num, _ := strconv.Atoi(string(c))
			jellyfish[Point{x, y}] = JellyFish{num, false}
		}
	}

	// Now do steps
	sumFlashes := 0
	for i := 0; i < 100; i++ {
		sumFlashes += StepJellyfish(jellyfish)
	}

	fmt.Println(sumFlashes)
}

func PartTwo() {
	input := readInput()

	var jellyfish map[Point]JellyFish
	jellyfish = make(map[Point]JellyFish)
	for x, v := range input {
		for y, c := range v {
			num, _ := strconv.Atoi(string(c))
			jellyfish[Point{x, y}] = JellyFish{num, false}
		}
	}

	// Now do steps
	for i := 0; i < 1000; i++ {
		sumFlashes := StepJellyfish(jellyfish)
		if sumFlashes == len(jellyfish) {
			fmt.Println(i + 1)
			break
		}
	}
}

func StepJellyfish(jellyfish map[Point]JellyFish) int {
	flashs := 0

	// First add one to everything
	for p := range jellyfish {
		j := jellyfish[p]
		j.energy++
		jellyfish[p] = j
	}

	// Now check for flashes
	hasBeenFlash := true
	for hasBeenFlash {
		hasBeenFlash = false
		for p, v := range jellyfish {
			if v.energy > 9 && !v.hasFlashed {
				flashs++
				AddToSurroundingJellyFish(p, jellyfish)
				hasBeenFlash = true
				j := jellyfish[p]
				j.hasFlashed = true
				jellyfish[p] = j
			}
		}
	}

	// reset any jelly fish that flashed
	for p := range jellyfish {
		if jellyfish[p].hasFlashed {
			j := jellyfish[p]
			j.hasFlashed = false
			j.energy = 0
			jellyfish[p] = j
		}
	}

	return flashs
}

func AddToSurroundingJellyFish(p Point, jellyfish map[Point]JellyFish) bool {
	hasBeenFlash := false

	up := Point{p.X + 1, p.Y}
	if _, ok := jellyfish[up]; ok {
		hasBeenFlash = true
		j := jellyfish[up]
		j.energy++
		jellyfish[up] = j
	}
	upLeft := Point{p.X + 1, p.Y - 1}
	if _, ok := jellyfish[upLeft]; ok {
		hasBeenFlash = true
		j := jellyfish[upLeft]
		j.energy++
		jellyfish[upLeft] = j
	}
	upRight := Point{p.X + 1, p.Y + 1}
	if _, ok := jellyfish[upRight]; ok {
		hasBeenFlash = true
		j := jellyfish[upRight]
		j.energy++
		jellyfish[upRight] = j
	}

	right := Point{p.X, p.Y + 1}
	if _, ok := jellyfish[right]; ok {
		hasBeenFlash = true
		j := jellyfish[right]
		j.energy++
		jellyfish[right] = j
	}
	down := Point{p.X - 1, p.Y}
	if _, ok := jellyfish[down]; ok {
		hasBeenFlash = true
		j := jellyfish[down]
		j.energy++
		jellyfish[down] = j
	}
	downLeft := Point{p.X - 1, p.Y - 1}
	if _, ok := jellyfish[downLeft]; ok {
		hasBeenFlash = true
		j := jellyfish[downLeft]
		j.energy++
		jellyfish[downLeft] = j
	}
	downRight := Point{p.X - 1, p.Y + 1}
	if _, ok := jellyfish[downRight]; ok {
		hasBeenFlash = true
		j := jellyfish[downRight]
		j.energy++
		jellyfish[downRight] = j
	}
	left := Point{p.X, p.Y - 1}
	if _, ok := jellyfish[left]; ok {
		hasBeenFlash = true
		j := jellyfish[left]
		j.energy++
		jellyfish[left] = j
	}

	return hasBeenFlash
}

type Point struct {
	X int
	Y int
}

type JellyFish struct {
	energy     int
	hasFlashed bool
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
