package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
	"strconv"
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

	lines := convertToLines(input)

	sum := make(map[string]int)
	for _, v := range lines {
		for key, point := range v.points {
			sum[key] += point
		}
	}

	overlap := 0
	for _, v := range sum {
		if v >= 2 {
			overlap++
		}
	}

	width := 1000
	height := 1000

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	green := color.RGBA{0, 255, 0, 0xff}
	blue := color.RGBA{0, 0, 255, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pointStr := strconv.Itoa(x) + "," + strconv.Itoa(y)
			if sum[pointStr] == 0 {
				img.Set(x, y, color.Black)
			} else if sum[pointStr] == 1 {
				img.Set(x, y, color.White)
			} else if sum[pointStr] == 2 {
				img.Set(x, y, green)
			} else if sum[pointStr] > 2 {
				img.Set(x, y, blue)
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)

	fmt.Println(overlap)
}

func PartTwo() {
	input := readInput()

	lines := convertToLinesComplex(input)

	sum := make(map[string]int)
	for _, v := range lines {
		for key, point := range v.points {
			sum[key] += point
		}
	}

	overlap := 0
	for _, v := range sum {
		if v >= 2 {
			overlap++
		}
	}

	width := 1000
	height := 1000

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	green := color.RGBA{0, 255, 0, 0xff}
	blue := color.RGBA{0, 0, 255, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pointStr := strconv.Itoa(x) + "," + strconv.Itoa(y)
			if sum[pointStr] == 0 {
				img.Set(x, y, color.Black)
			} else if sum[pointStr] == 1 {
				img.Set(x, y, color.White)
			} else if sum[pointStr] == 2 {
				img.Set(x, y, green)
			} else if sum[pointStr] > 2 {
				img.Set(x, y, blue)
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("imageComplex.png")
	png.Encode(f, img)

	fmt.Println(overlap)
}

type Line struct {
	startX int
	startY int
	endX   int
	endY   int
	points map[string]int
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

func convertToLines(input []string) []Line {

	// Given a line of string numbers convert to Lines
	var lineArray []Line

	// Remove extra spaces
	capture := regexp.MustCompile(`([0-9]*),([0-9]*) -> ([0-9]*),([0-9]*)`)

	for _, v := range input {
		s := capture.FindStringSubmatch(v)

		startX, _ := strconv.Atoi(s[1])
		startY, _ := strconv.Atoi(s[2])
		endX, _ := strconv.Atoi(s[3])
		endY, _ := strconv.Atoi(s[4])

		p := make(map[string]int)

		if startX == endX {
			//Straight Horizontal Line
			for i := min(startY, endY); i < max(startY, endY)+1; i++ {
				pointStr := strconv.Itoa(startX) + "," + strconv.Itoa(i)
				p[pointStr] = 1
			}
		}
		if startY == endY {
			//Straight Horizontal Line
			for i := min(startX, endX); i < max(startX, endX)+1; i++ {
				pointStr := strconv.Itoa(i) + "," + strconv.Itoa(startY)
				p[pointStr] = 1
			}
		}

		l := Line{startX, startY, endX, endY, p}

		lineArray = append(lineArray, l)
	}

	return lineArray
}

func convertToLinesComplex(input []string) []Line {

	// Given a line of string numbers convert to Lines
	var lineArray []Line

	// Remove extra spaces
	capture := regexp.MustCompile(`([0-9]*),([0-9]*) -> ([0-9]*),([0-9]*)`)

	for _, v := range input {
		s := capture.FindStringSubmatch(v)

		startX, _ := strconv.Atoi(s[1])
		startY, _ := strconv.Atoi(s[2])
		endX, _ := strconv.Atoi(s[3])
		endY, _ := strconv.Atoi(s[4])

		p := make(map[string]int)

		if startX == endX {
			//Straight Horizontal Line
			for i := min(startY, endY); i < max(startY, endY)+1; i++ {
				pointStr := strconv.Itoa(startX) + "," + strconv.Itoa(i)
				p[pointStr] = 1
			}
		}
		if startY == endY {
			//Straight Horizontal Line
			for i := min(startX, endX); i < max(startX, endX)+1; i++ {
				pointStr := strconv.Itoa(i) + "," + strconv.Itoa(startY)
				p[pointStr] = 1
			}
		}

		// Find diagonal
		if abs(startX-endX) == abs(startY-endY) {
			// 45 Degree line
			// Start from top left point (min point)
			xMulti := 0
			yMulti := 0
			if startX < endX {
				xMulti = 1
			} else {
				xMulti = -1
			}
			if startY < endY {
				yMulti = 1
			} else {
				yMulti = -1
			}

			for i := 0; i < abs(startX-endX)+1; i++ {
				pointStr := strconv.Itoa(startX+(i*xMulti)) + "," + strconv.Itoa(startY+(i*yMulti))
				p[pointStr] = 1
			}

		}

		l := Line{startX, startY, endX, endY, p}

		lineArray = append(lineArray, l)
	}

	return lineArray
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
