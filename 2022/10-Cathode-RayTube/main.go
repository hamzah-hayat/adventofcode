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
		fmt.Println("Gold:\n" + PartTwo("input"))
	case "p1":
		fmt.Println("Silver:" + PartOne("input"))
	case "p2":
		fmt.Println("Gold:\n" + PartTwo("input"))
	}
}

func PartOne(filename string) string {
	input := readInput(filename)

	commands := make(map[int]int)
	xValue := 1
	signalStrengthSum := 0
	inputLine := 0

	for cycle := 0; cycle < 240; cycle++ {
		addCycle := false

		// We can run a command
		regexAdd := regexp.MustCompile(`addx (-?[0-9]+)`)
		addMatch := regexAdd.FindStringSubmatch(input[inputLine])

		if addMatch != nil {
			// add to our map so we can do it later
			num, _ := strconv.Atoi(addMatch[1])
			commands[cycle+2] += num
			addCycle = true
		}

		// Get our signal strength
		switch cycle {
		case 20:
			signalStrengthSum += xValue * cycle
		case 60:
			signalStrengthSum += xValue * cycle
		case 100:
			signalStrengthSum += xValue * cycle
		case 140:
			signalStrengthSum += xValue * cycle
		case 180:
			signalStrengthSum += xValue * cycle
		case 220:
			signalStrengthSum += xValue * cycle
		}
		if addCycle {
			// now check the map and add to x if necessary
			// in case we do this during skip
			if commands[cycle] != 0 {
				xValue += commands[cycle]
				delete(commands, cycle)
			}
			cycle++
			switch cycle {
			case 20:
				signalStrengthSum += xValue * cycle
			case 60:
				signalStrengthSum += xValue * cycle
			case 100:
				signalStrengthSum += xValue * cycle
			case 140:
				signalStrengthSum += xValue * cycle
			case 180:
				signalStrengthSum += xValue * cycle
			case 220:
				signalStrengthSum += xValue * cycle
			}
		}
		// now check the map and add to x if necessary
		// no danger of dupes because we delete from map after
		if commands[cycle] != 0 {
			xValue += commands[cycle]
			delete(commands, cycle)
		}
		inputLine++
	}

	return strconv.Itoa(signalStrengthSum)
}

func PartTwo(filename string) string {
	input := readInput(filename)

	commands := make(map[int]int)
	xValue := 1
	inputLine := 0
	CRTOutput := ""

	for cycle := 0; cycle < 240; cycle++ {
		addCycle := false

		// We can run a command
		regexAdd := regexp.MustCompile(`addx (-?[0-9]+)`)
		addMatch := regexAdd.FindStringSubmatch(input[inputLine])

		if addMatch != nil {
			// add to our map so we can do it later
			num, _ := strconv.Atoi(addMatch[1])
			commands[cycle+2] += num
			addCycle = true
		}
		// now check the map and add to x if necessary
		if commands[cycle] != 0 {
			xValue += commands[cycle]
			delete(commands, cycle)
		}
		// Check CRT pixel
		xValueMod40 := xValue % 40
		cycleMod40 := cycle % 40
		if xValueMod40+1 == cycleMod40 || xValueMod40 == cycleMod40 || xValueMod40-1 == cycleMod40 {
			CRTOutput += "#"
		} else {
			CRTOutput += "."
		}

		// Get our signal strength
		switch cycle {
		case 40:
			CRTOutput += "\n"
		case 80:
			CRTOutput += "\n"
		case 120:
			CRTOutput += "\n"
		case 160:
			CRTOutput += "\n"
		case 200:
			CRTOutput += "\n"
		case 240:
			CRTOutput += "\n"
		}
		if addCycle {
			// now check the map and add to x if necessary
			if commands[cycle] != 0 {
				xValue += commands[cycle]
				delete(commands, cycle)
			}
			cycle++
			// Check CRT pixel
			xValueMod40 := xValue % 40
			cycleMod40 := cycle % 40
			if xValueMod40+1 == cycleMod40 || xValueMod40 == cycleMod40 || xValueMod40-1 == cycleMod40 {
				CRTOutput += "#"
			} else {
				CRTOutput += "."
			}
			switch cycle {
			case 40:
				CRTOutput += "\n"
			case 80:
				CRTOutput += "\n"
			case 120:
				CRTOutput += "\n"
			case 160:
				CRTOutput += "\n"
			case 200:
				CRTOutput += "\n"
			case 240:
				CRTOutput += "\n"
			}
		}
		inputLine++
	}
	CRTOutput += "\n"
	return CRTOutput
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
