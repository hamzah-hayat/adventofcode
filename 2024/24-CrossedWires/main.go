package main

import (
	"bufio"
	"cmp"
	"flag"
	"fmt"
	"os"
	"regexp"
	"slices"
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
		fmt.Println("Gold:" + PartTwo("input_fixed"))
	case "p1":
		fmt.Println("Silver:" + PartOne("input"))
	case "p2":
		fmt.Println("Gold:" + PartTwo("input_fixed"))
	}
}

func PartOne(filename string) string {
	input := readInput(filename)

	// Start by reading wires/gates
	wires := make([]Wire, 0)
	gates := make([]Gate, 0)

	readingWires := true
	for _, line := range input {
		if line == "" {
			readingWires = false
			continue
		}

		if readingWires {
			wireSplit := strings.Split(line, ":")
			wireName := wireSplit[0]
			wireNum, _ := strconv.Atoi(strings.TrimSpace(wireSplit[1]))

			if wireNum == 0 {
				wires = append(wires, Wire{wireName: wireName, wireValue: false})
			} else {
				wires = append(wires, Wire{wireName: wireName, wireValue: true})
			}
		} else {
			gateRegex := regexp.MustCompile("(.{3}) (XOR|OR|AND) (.{3}) -> (.{3})")
			match := gateRegex.FindAllStringSubmatch(line, 1)
			gates = append(gates, Gate{match[0][1], match[0][3], match[0][4], match[0][2]})
		}
	}

	// Then process the gates
	completedGates := make([]Gate, 0)
	for len(completedGates) != len(gates) {
		for _, gate := range gates {
			if slices.Contains(completedGates, gate) {
				continue
			}
			completed := false
			wires, completed = RunGate(gate, wires)
			if completed {
				completedGates = append(completedGates, gate)
			}
		}
	}

	// Find all wires starting with z and form binary number
	slices.SortFunc(wires,
		func(a, b Wire) int {
			return cmp.Compare(a.wireName, b.wireName)
		})

	numStr := ""
	for _, wireBool := range wires {
		if wireBool.wireName[0] == 'z' {
			if wireBool.wireValue {
				numStr += "1"
			} else {
				numStr += "0"
			}
		}
	}

	numBinary, _ := strconv.ParseInt(Reverse(numStr), 2, 64)

	return strconv.Itoa(int(numBinary))
}

func RunGate(gate Gate, wires []Wire) ([]Wire, bool) {

	// Find wires first
	firstWire := Wire{}
	secondWire := Wire{}
	finalWire := Wire{}
	for _, w := range wires {
		if w.wireName == gate.first {
			firstWire = w
		}
		if w.wireName == gate.second {
			secondWire = w
		}
		if w.wireName == gate.final {
			finalWire = w
		}
	}

	if firstWire.wireName == "" || secondWire.wireName == "" {
		return wires, false
	}

	switch gate.instruction {
	case "AND":
		if finalWire.wireName != "" {
			finalWire.wireValue = firstWire.wireValue && secondWire.wireValue
		} else {
			wires = append(wires, Wire{gate.final, firstWire.wireValue && secondWire.wireValue})
		}
	case "OR":
		if finalWire.wireName != "" {
			finalWire.wireValue = firstWire.wireValue || secondWire.wireValue
		} else {
			wires = append(wires, Wire{gate.final, firstWire.wireValue || secondWire.wireValue})
		}
	case "XOR":
		if finalWire.wireName != "" {
			finalWire.wireValue = firstWire.wireValue != secondWire.wireValue
		} else {
			wires = append(wires, Wire{gate.final, firstWire.wireValue != secondWire.wireValue})
		}
	}

	return wires, true
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

type Wire struct {
	wireName  string
	wireValue bool
}

type Gate struct {
	first       string
	second      string
	final       string
	instruction string
}

func PartTwo(filename string) string {
	//input := readInput(filename)
	//Wrong Gates
	// I solved it manually
	// // Any gate that produces a Z output must use an XOR
	// y22 AND x22 -> z22//
	// grd OR rpq -> z29//
	// vqp AND frr -> z08//

	// // Using XOR but not with x,y,z
	// cmn XOR cdf -> hwq//
	// frr XOR vqp -> thm//
	// bfq XOR dcf -> gbs//

	// // Start of section where binary num doesn't match
	// x14 AND y14 -> wss
	// x14 XOR y14 -> wrm

	// // Wires
	// thm,gbs,hwq,z08,z29,z22,wrm,wss

	// // Im going to swap
	// vqp AND frr -> z08 AND frr XOR vqp -> thm
	// grd OR rpq -> z29 AND bfq XOR dcf -> gbs
	// y22 AND x22 -> z22 AND cmn XOR cdf -> hwq
	// x14 AND y14 -> wss AND x14 XOR y14 -> wrm
	finalWires := []string{"thm", "gbs", "hwq", "z08", "z29", "z22", "wrm", "wss"}
	slices.Sort(finalWires)

	return strings.Join(finalWires, ",")

	// for {
	// 	//Start by reading wires/gates
	// 	gates, wires := GetInitialGatesAndWires(input)

	// 	_, finalValue := RunGatesAndCheck(gates, wires)
	// }
}

// Runs a full set of gates, returning whether x values + y values = z values
func RunGatesAndCheck(gates []Gate, wires []Wire) (bool, int) {

	// Then process the gates
	completedGates := make([]Gate, 0)
	for len(completedGates) != len(gates) {
		for _, gate := range gates {
			if slices.Contains(completedGates, gate) {
				continue
			}
			completed := false
			wires, completed = RunGate(gate, wires)
			if completed {
				completedGates = append(completedGates, gate)
			}
		}
	}

	// Find all wires starting with x,y and z and form binary numbers
	slices.SortFunc(wires,
		func(a, b Wire) int {
			return cmp.Compare(a.wireName, b.wireName)
		})

	numXStr := ""
	numYStr := ""
	numZStr := ""
	for _, wireBool := range wires {
		if wireBool.wireName[0] == 'x' {
			if wireBool.wireValue {
				numXStr += "1"
			} else {
				numXStr += "0"
			}
		}
		if wireBool.wireName[0] == 'y' {
			if wireBool.wireValue {
				numYStr += "1"
			} else {
				numYStr += "0"
			}
		}
		if wireBool.wireName[0] == 'z' {
			if wireBool.wireValue {
				numZStr += "1"
			} else {
				numZStr += "0"
			}
		}
	}

	numXBinary, _ := strconv.ParseInt(Reverse(numXStr), 2, 64)
	numYBinary, _ := strconv.ParseInt(Reverse(numYStr), 2, 64)
	numZBinary, _ := strconv.ParseInt(Reverse(numZStr), 2, 64)

	// Check each Z value here
	carry := 0
	for i := 0; i < len(numXStr); i++ {
		numXb, _ := strconv.Atoi(string(numXStr[i]))
		numYb, _ := strconv.Atoi(string(numYStr[i]))
		numZb, _ := strconv.Atoi(string(numZStr[i]))

		if (numXb+numYb+carry)%2 != numZb {
			fmt.Println("The value of z", i, "is sus")
		}
		if numXb+numYb+carry >= 2 {
			carry = 1
		} else {
			carry = 0
		}
	}

	return numXBinary+numYBinary == numZBinary, int(numZBinary)
}

func GetInitialGatesAndWires(input []string) ([]Gate, []Wire) {
	// Start by reading wires/gates
	wires := make([]Wire, 0)
	gates := make([]Gate, 0)

	readingWires := true
	for _, line := range input {
		if line == "" {
			readingWires = false
			continue
		}

		if readingWires {
			wireSplit := strings.Split(line, ":")
			wireName := wireSplit[0]
			wireNum, _ := strconv.Atoi(strings.TrimSpace(wireSplit[1]))

			if wireNum == 0 {
				wires = append(wires, Wire{wireName: wireName, wireValue: false})
			} else {
				wires = append(wires, Wire{wireName: wireName, wireValue: true})
			}
		} else {
			gateRegex := regexp.MustCompile("(.{3}) (XOR|OR|AND) (.{3}) -> (.{3})")
			match := gateRegex.FindAllStringSubmatch(line, 1)
			gates = append(gates, Gate{match[0][1], match[0][3], match[0][4], match[0][2]})
		}
	}

	return gates, wires
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
