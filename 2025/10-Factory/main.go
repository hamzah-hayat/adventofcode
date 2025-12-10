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
	case "p2":
		fmt.Println("Gold:" + PartTwo("input"))
	}
}

func PartOne(filename string) string {
	input := readInput(filename)
	machines := []Machine{}

	// Read input and create machines
	machineRegex := regexp.MustCompile(`(\d+),(\d+)`)
	for _, lines := range input {
		

		machines = append(machines, machine)
	}

	return strconv.Itoa(len(input))
}

func PartTwo(filename string) string {
	input := readInput(filename)

	return strconv.Itoa(len(input))
}

type Machine struct {
	lights        []bool
	correctLights []bool
	buttons       []Button
}

type Button struct {
	lights []int
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
