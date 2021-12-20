package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

var (
	methodP *string
)

func parseFlags() {
	methodP = flag.String("method", "p1", "The method/part that should be run, valid are p1,p2 and test")
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

	enchancmentAlgorithm := input[0]

	imageMap := make(map[Point]bool)

	for y, line := range input[2:] {
		for x, char := range line {
			point := Point{x, y}
			if char == '#' {
				imageMap[point] = true
			} else {
				imageMap[point] = false
			}
		}
	}

	outside := false
	DrawPNG(imageMap, "input")

	for i := 0; i < 2; i++ {
		imageMap = EnhanceImage(imageMap, enchancmentAlgorithm, outside)
		DrawPNG(imageMap, fmt.Sprint("enhance", i+1))

		// Do we flip bits?
		if !outside && enchancmentAlgorithm[0] == '#' {
			outside = !outside
			continue
		}
		if outside && enchancmentAlgorithm[len(enchancmentAlgorithm)-1] == '.' {
			outside = !outside
			continue
		}
	}

	// count number of lit pixels
	litPixels := 0
	for _, v := range imageMap {
		if v {
			litPixels++
		}
	}
	num := strconv.Itoa(litPixels)

	return num
}

func PartTwo(filename string) string {
	input := readInput(filename)

	enchancmentAlgorithm := input[0]

	imageMap := make(map[Point]bool)

	for y, line := range input[2:] {
		for x, char := range line {
			point := Point{x, y}
			if char == '#' {
				imageMap[point] = true
			} else {
				imageMap[point] = false
			}
		}
	}

	outside := false
	DrawPNG(imageMap, "input")

	for i := 0; i < 50; i++ {
		imageMap = EnhanceImage(imageMap, enchancmentAlgorithm, outside)
		//DrawPNG(imageMap, fmt.Sprint("enhance", i+1))

		// Do we flip bits?
		if !outside && enchancmentAlgorithm[0] == '#' {
			outside = !outside
			continue
		}
		if outside && enchancmentAlgorithm[len(enchancmentAlgorithm)-1] == '.' {
			outside = !outside
			continue
		}
	}

	// count number of lit pixels
	litPixels := 0
	for _, v := range imageMap {
		if v {
			litPixels++
		}
	}
	num := strconv.Itoa(litPixels)

	return num
}

func EnhanceImage(imageMap map[Point]bool, enhancementAlgorithm string, outside bool) map[Point]bool {

	imageMap = AddPadding(imageMap, outside)
	newMap := make(map[Point]bool)

	for p := range imageMap {
		// for each pixel, we look at the nine squares around it
		binaryNum := ""
		for y := -1; y < 2; y++ {
			for x := -1; x < 2; x++ {
				if val, ok := imageMap[Point{p.X + x, p.Y + y}]; ok {
					if val {
						binaryNum += "1"
					} else {
						binaryNum += "0"
					}
				} else {
					if outside {
						binaryNum += "1"
					} else {
						binaryNum += "0"
					}
				}
			}
		}
		// convert binaryNum to index
		indexNum, _ := strconv.ParseInt(binaryNum, 2, 64)

		// Now look at how we enhance
		if enhancementAlgorithm[indexNum] == '#' {
			newMap[p] = true
		} else {
			newMap[p] = false
		}
	}

	return newMap
}

// Add padding around imagemap
func AddPadding(imageMap map[Point]bool, outside bool) map[Point]bool {
	newMap := make(map[Point]bool)
	width, height := GetWidthAndHeightOfMap(imageMap)

	for x := 0; x < width+2; x++ {
		for y := 0; y < height+2; y++ {
			newMap[Point{x, y}] = outside
		}
	}

	for p, v := range imageMap {
		newMap[Point{p.X + 1, p.Y + 1}] = v
	}

	return newMap
}

func DrawPNG(imageMap map[Point]bool, filename string) {

	width, height := GetWidthAndHeightOfMap(imageMap)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	// green := color.RGBA{0, 255, 0, 0xff}
	// blue := color.RGBA{0, 0, 255, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if imageMap[Point{x, y}] == true {
				img.Set(x, y, color.White)
			} else {
				img.Set(x, y, color.Black)
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create(filename + ".png")
	png.Encode(f, img)
}

func GetWidthAndHeightOfMap(imageMap map[Point]bool) (int, int) {
	width := 0
	height := 0
	for p := range imageMap {
		if width <= p.X {
			width = p.X + 1
		}
		if height <= p.Y {
			height = p.Y + 1
		}
	}

	return width, height
}

type Point struct {
	X, Y int
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
