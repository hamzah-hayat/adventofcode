package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	direction := "N"
	x := 0
	y := 0
	capture := regexp.MustCompile(`([RL])([0-9]*)`)
	for _, v := range strings.Split(input[0], ", ") {
		r := capture.FindStringSubmatch(v)
		// First change direction
		if r[1] == "L" {
			switch direction {
			case "N":
				direction = "W"
			case "E":
				direction = "N"
			case "S":
				direction = "E"
			case "W":
				direction = "S"
			}
		} else {
			switch direction {
			case "N":
				direction = "E"
			case "E":
				direction = "S"
			case "S":
				direction = "W"
			case "W":
				direction = "N"
			}
		}
		// Then move
		num, _ := strconv.Atoi(r[2])
		switch direction {
		case "N":
			y += num
		case "E":
			x += num
		case "S":
			y -= num
		case "W":
			x -= num
		}
	}

	fmt.Println(x + y)
}

func PartTwo() {
	input := readInput()

	direction := "N"
	x := 0
	y := 0
	visted := make(map[string]int)
	capture := regexp.MustCompile(`([RL])([0-9]*)`)
	for _, v := range strings.Split(input[0], ", ") {
		r := capture.FindStringSubmatch(v)
		// First change direction
		if r[1] == "L" {
			switch direction {
			case "N":
				direction = "W"
			case "E":
				direction = "N"
			case "S":
				direction = "E"
			case "W":
				direction = "S"
			}
		} else {
			switch direction {
			case "N":
				direction = "E"
			case "E":
				direction = "S"
			case "S":
				direction = "W"
			case "W":
				direction = "N"
			}
		}
		// Then move
		num, _ := strconv.Atoi(r[2])
		xMulti := 0
		yMulti := 0
		switch direction {
		case "N":
			yMulti = 1
		case "E":
			xMulti = 1
		case "S":
			yMulti = -1
		case "W":
			xMulti = -1
		}

		found := false
		for i := 0; i < num; i++ {
			pointStr := strconv.Itoa(x+(i*xMulti)) + "," + strconv.Itoa(y+(i*yMulti))
			visted[pointStr] += 1
			if visted[pointStr] == 2 {
				found = true
				break
			}
		}
		if found {
			break
		}

		switch direction {
		case "N":
			y += num
		case "E":
			x += num
		case "S":
			y -= num
		case "W":
			x -= num
		}

	}

	for i, v := range visted {
		if v == 2 {
			fmt.Println(i)
		}
	}

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
