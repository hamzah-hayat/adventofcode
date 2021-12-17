package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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
		break
	case "p2":
		fmt.Println("Gold:" + PartTwo("input"))
		break
	}
}

func PartOne(filename string) string {
	input := readInput(filename)

	targetZone, lowestPointInGrid := createTargetZone(input[0])
	currentMaxHeight := 0

	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			maxHeight, didWeHitTarget := fireProbe(x, y, targetZone, lowestPointInGrid)

			if didWeHitTarget && maxHeight > currentMaxHeight {
				// Find point with highest hit
				currentMaxHeight = maxHeight
			}
		}
	}

	num := strconv.Itoa(currentMaxHeight)
	return num
}

func PartTwo(filename string) string {
	input := readInput(filename)

	targetZone, lowestPointInGrid := createTargetZone(input[0])
	hits := 0

	for x := 0; x < 300; x++ {
		for y := -300; y < 300; y++ {
			_, didWeHitTarget := fireProbe(x, y, targetZone, lowestPointInGrid)

			if didWeHitTarget {
				// Find point with highest hit
				hits++
			}
		}
	}

	num := strconv.Itoa(hits)
	return num
}

type Point struct {
	X int
	Y int
}

func fireProbe(xVelo, yVelo int, targetGrid map[Point]bool, lowestPointInTarget int) (int, bool) {
	hitTarget := false
	maxHeight := 0

	// Check if we hit target based on xVelo and yVelo
	// The probe's x position increases by its x velocity.
	// The probe's y position increases by its y velocity.
	// Due to drag, the probe's x velocity changes by 1 toward the value 0; that is, it decreases by 1 if it is greater than 0, increases by 1 if it is less than 0, or does not change if it is already 0.
	// Due to gravity, the probe's y velocity decreases by 1.
	currentX := 0
	currentY := 0

	for !hitTarget {
		// run a step
		currentX += xVelo
		currentY += yVelo
		if xVelo < 0 {
			xVelo++
		} else if xVelo > 0 {
			xVelo--
		}
		yVelo--

		if currentY > maxHeight {
			maxHeight = currentY
		}
		// Check if we hit target
		if targetGrid[Point{currentX, currentY}] {
			hitTarget = true
			break
		}

		// If we are below target, we will never hit, so break out
		if currentY < lowestPointInTarget {
			break
		}
	}

	return maxHeight, hitTarget
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

func createTargetZone(input string) (map[Point]bool, int) {
	pointsInGrid := make(map[Point]bool)

	// target area: x=185..221, y=-122..-74
	gridRegex := regexp.MustCompile("target area: x=([0-9]+)..([0-9]+), y=(-[0-9]+)..(-[0-9]+)")

	gridString := gridRegex.FindAllStringSubmatch(input, 1)

	xMin, _ := strconv.Atoi(gridString[0][1])
	xMax, _ := strconv.Atoi(gridString[0][2])
	yMin, _ := strconv.Atoi(gridString[0][3])
	yMax, _ := strconv.Atoi(gridString[0][4])

	for x := xMin; x < xMax+1; x++ {
		for y := yMin; y < yMax+1; y++ {
			pointsInGrid[Point{x, y}] = true
		}
	}

	return pointsInGrid, yMin
}
