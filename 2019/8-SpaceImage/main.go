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
	input := stringInputToInt(readInput())

	width := 25
	height := 6

	pixels, maxLayers := buildPixelMap(input, height, width)

	// Now that we have the pixel grid
	// Find the layer with lowest amount of zeros

	lowestLayer := findLayerWithLowestNumber(pixels, 0, maxLayers)

	//fmt.Println("The lowest layer is", lowestLayer)
	fmt.Println(len(getLayer(pixels, lowestLayer)))

	num1 := findNumberCountInLayer(pixels, 1, lowestLayer)
	num2 := findNumberCountInLayer(pixels, 2, lowestLayer)

	fmt.Println("The number of 1s multiped by 2s is", num1*num2)

}

func buildPixelMap(input []int, height int, width int) (map[pixel]int, int) {

	pixels := make(map[pixel]int)

	currentWidth := 0
	currentHeight := 0
	currentLayer := 0

	for _, value := range input {
		// Three cases
		// 1. Go to next layer
		// 2. Go down a row
		// 3. Place pixel down
		if currentWidth == width-1 && currentHeight == height-1 {
			pixels[pixel{x: currentWidth, y: currentHeight, layer: currentLayer}] = value
			currentLayer++
			currentWidth = 0
			currentHeight = 0
		} else if currentWidth == width-1 {
			pixels[pixel{x: currentWidth, y: currentHeight, layer: currentLayer}] = value
			currentHeight++
			currentWidth = 0
		} else {
			pixels[pixel{x: currentWidth, y: currentHeight, layer: currentLayer}] = value
			currentWidth++
		}
	}

	// Currentlayer at this point is the max number of layers
	return pixels, currentLayer
}

// Find the layer with the lowest number of number
func findLayerWithLowestNumber(pixels map[pixel]int, number int, maxLayers int) int {
	lowestLayer := 0
	lowestCount := 10000
	for layer := 0; layer < maxLayers; layer++ {
		total := 0
		for i, value := range pixels {
			if i.layer == layer && value == number {
				total++
			}
		}
		if total < lowestCount {
			lowestCount = total
			lowestLayer = layer
		}
	}
	return lowestLayer
}

// Find the count of a number in a layer
func findNumberCountInLayer(pixels map[pixel]int, number int, layer int) int {
	num := 0
	for i, pixelColour := range pixels {
		if i.layer == layer {
			if pixelColour == number {
				num++
			}
		}
	}
	return num
}

// Get a layer and return it
func getLayer(pixels map[pixel]int, layer int) map[pixel]int {
	newLayer := make(map[pixel]int)
	for i, pixelColour := range pixels {
		if i.layer == layer {
			newLayer[i] = pixelColour
		}
	}
	return newLayer
}

func partTwo() {
	//input := readInput()
}

type pixel struct {
	x, y, layer int
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

// Turn a string input to a int input
func stringInputToInt(stringInput []string) []int {
	var intInput []int
	for _, val := range strings.Split(stringInput[0], "") {
		i, _ := strconv.Atoi(val)
		intInput = append(intInput, i)
	}
	return intInput
}
