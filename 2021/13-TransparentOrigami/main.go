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

	dotMap := make(map[Point]bool)
	var instructions []Fold

	// Get input
	for _, v := range input {
		dotRegex := regexp.MustCompile("([0-9]+),([0-9]+)")
		instructionRegex := regexp.MustCompile("fold along ([xy])=([0-9]+)")

		dot := dotRegex.FindStringSubmatch(v)
		if dot != nil {
			x, _ := strconv.Atoi(dot[1])
			y, _ := strconv.Atoi(dot[2])
			dotMap[Point{x, y}] = true
		}

		fold := instructionRegex.FindStringSubmatch(v)
		if fold != nil {
			num, _ := strconv.Atoi(fold[2])
			if fold[1] == "y" {
				instructions = append(instructions, Fold{true, num})
			} else {
				instructions = append(instructions, Fold{false, num})
			}
		}
	}

	printGrid(dotMap, 15, 15, "start")
	// Do the fold
	for _, foldMethod := range instructions {
		FoldPaper(dotMap, foldMethod)
		// For part one, we just do this once
		break
	}
	printGrid(dotMap, 15, 15, "end")

	dotsSum := 0
	for _, v := range dotMap {
		if v {
			dotsSum++
		}
	}
	fmt.Println(dotsSum)
}

func FoldPaper(dots map[Point]bool, foldMethod Fold) {
	// Fold the paper
	if foldMethod.isUpFold {
		// Fold up along the y axis
		for d := range dots {
			if d.Y > foldMethod.foldNum {
				// Fold this dot
				// its new num is = foldMethod.foldNum - (d.Y - foldMethod.foldNum)
				dots[d] = false
				newDot := Point{d.X, foldMethod.foldNum - (d.Y - foldMethod.foldNum)}
				dots[newDot] = true
			}
		}
	} else {
		// fold left along x axis
		// Fold up along the y axis
		for d := range dots {
			if d.X > foldMethod.foldNum {
				// Fold this dot
				// its new num is = d.X - foldMethod.foldNum
				delete(dots, d)
				newDot := Point{foldMethod.foldNum - (d.X - foldMethod.foldNum), d.Y}
				dots[newDot] = true
			}
		}
	}
}

func printGrid(dots map[Point]bool, maxX, maxY int, name string) {
	width := maxX
	height := maxY

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if dots[Point{x, y}] {
				img.Set(x, y, color.Black)
			} else {
				img.Set(x, y, color.White)
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create(name + ".png")
	png.Encode(f, img)

}

type Point struct {
	X int
	Y int
}

type Fold struct {
	isUpFold bool
	foldNum  int
}

func PartTwo() {
	input := readInput()

	dotMap := make(map[Point]bool)
	var instructions []Fold

	// Get input
	for _, v := range input {
		dotRegex := regexp.MustCompile("([0-9]+),([0-9]+)")
		instructionRegex := regexp.MustCompile("fold along ([xy])=([0-9]+)")

		dot := dotRegex.FindStringSubmatch(v)
		if dot != nil {
			x, _ := strconv.Atoi(dot[1])
			y, _ := strconv.Atoi(dot[2])
			dotMap[Point{x, y}] = true
		}

		fold := instructionRegex.FindStringSubmatch(v)
		if fold != nil {
			num, _ := strconv.Atoi(fold[2])
			if fold[1] == "y" {
				instructions = append(instructions, Fold{true, num})
			} else {
				instructions = append(instructions, Fold{false, num})
			}
		}
	}

	printGrid(dotMap, 1000, 1000, "start")
	// Do the fold
	for _, foldMethod := range instructions {
		FoldPaper(dotMap, foldMethod)
	}
	printGrid(dotMap, 40, 6, "end")

	fmt.Println("Go look at end.png!")
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
