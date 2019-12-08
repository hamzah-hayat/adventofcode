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
	methodP := flag.String("method", "p2", "The method/part that should be run, valid are p1,p2 and test")
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
	input := stringInputToInt(readInput())

	width := 25
	height := 6

	pixels, maxLayers := buildPixelMap(input, height, width)

	// Now that we have the pixel grid
	// Layer all pixels over each other to find the "image layer"

	imageLayer := buildImageLayer(pixels, width, height, maxLayers)

	fmt.Print(printLayer(imageLayer, width, height))

}

func buildImageLayer(pixels map[pixel]int, width int, height int, maxLayers int) map[pixel]int {

	mergedLayer := make(map[pixel]int)

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {

			mergedLayers := getAllPixelsFromAllLayersForPoint(pixels, w, h, maxLayers)

			// Find first visible pixel
			for _, value := range mergedLayers {
				if value == 2 {
					continue
				} else if value == 1 {
					mergedLayer[pixel{x: w, y: h, layer: 0}] = 1
					break
				} else if value == 0 {
					mergedLayer[pixel{x: w, y: h, layer: 0}] = 0
					break
				}
			}
		}
	}

	fmt.Println(len(mergedLayer))

	return mergedLayer

}

func getAllPixelsFromAllLayersForPoint(pixels map[pixel]int, width int, height int, maxLayers int) []int {

	var allPixels []int

	for i := 0; i < maxLayers; i++ {
		allPixels = append(allPixels, pixels[pixel{x: width, y: height, layer: i}])
	}

	return allPixels
}

// Print out this image layer using unicode boxes
func printLayer(pixels map[pixel]int, width int, height int) string {

	image := ""

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			switch pixels[pixel{x: w, y: h, layer: 0}] {
			case 0:
				image += "\u25A0" // black
				break
			case 1:
				image += "\u25A1" // white
				break
			case 2:
				image += " " // transparent
				break
			}
		}
		image += "\n"
	}

	return image
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
