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

	var boxes [1000]box

	for i, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		boxData := strings.Split(line, "x")
		w, _ := strconv.Atoi(boxData[0])
		l, _ := strconv.Atoi(boxData[1])
		h, _ := strconv.Atoi(boxData[2])
		b := box{width: w, length: l, height: h}
		boxes[i] = b
	}

	// Got all the boxes
	wrappingPaper := 0
	for _, b := range boxes {
		wrappingPaper += b.paper()
	}

	fmt.Println("Wrapping paper needed is", wrappingPaper)

}

func PartTwo() {
	input := readInput()

	var boxes [1000]box

	for i, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		boxData := strings.Split(line, "x")
		w, _ := strconv.Atoi(boxData[0])
		l, _ := strconv.Atoi(boxData[1])
		h, _ := strconv.Atoi(boxData[2])
		b := box{width: w, length: l, height: h}
		boxes[i] = b
	}

	// Got all the boxes
	ribbon := 0
	for _, b := range boxes {
		ribbon += b.ribbon()
	}

	fmt.Println("Ribbon needed is", ribbon)
}

type box struct {
	width  int
	length int
	height int
}

func (b box) paper() int {
	areasize := 2*b.length*b.width + 2*b.width*b.height + 2*b.height*b.length
	slack := 0
	if b.length >= b.width && b.length >= b.height {
		slack += b.width * b.height
	} else if b.width >= b.length && b.width >= b.height {
		slack += b.length * b.height
	} else if b.height >= b.width && b.height >= b.length {
		slack += b.length * b.width
	}
	return areasize + slack
}

func (b box) ribbon() int {
	volume := b.length * b.width * b.height
	perimeter := 0
	if b.length >= b.width && b.length >= b.height {
		perimeter += b.width*2 + b.height*2
	} else if b.width >= b.length && b.width >= b.height {
		perimeter += b.length*2 + b.height*2
	} else if b.height >= b.width && b.height >= b.length {
		perimeter += b.length*2 + b.width*2
	}
	return volume + perimeter
}

// Read data from input.txt
// Return the string, so that we can deal with it however
func readInput() string {

	var input string

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}
	return input
}
