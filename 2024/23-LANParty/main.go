package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
		fmt.Println("Gold:" + PartTwo("input"))
	case "p1":
		fmt.Println("Silver:" + PartOne("input"))
	case "p2":
		fmt.Println("Gold:" + PartTwo("input"))
	}
}

func PartOne(filename string) string {
	input := readInput(filename)

	connections := make(map[string]bool)
	tripleConnections := make(map[string]bool)
	computerList := make(map[string]bool)
	for _, line := range input {
		computers := strings.Split(line, "-")
		slices.Sort(computers)
		stringJoin := strings.Join(computers, ",")
		connections[stringJoin] = true
		computerList[computers[0]] = true
		computerList[computers[1]] = true
	}

	// Now with all connections, find any triples
	for conn := range connections {
		computers := strings.Split(conn, ",")
		// Is there a triple with another computer?
		for comp := range computerList {
			if comp != computers[0] && comp != computers[1] {
				// Find A and B if they are in our connections
				checkA := []string{comp, computers[0]}
				slices.Sort(checkA)
				checkAStr := strings.Join(checkA, ",")
				checkB := []string{comp, computers[1]}
				slices.Sort(checkB)
				checkBStr := strings.Join(checkB, ",")

				if connections[checkAStr] && connections[checkBStr] {
					sortedComputers := []string{comp, computers[0], computers[1]}
					slices.Sort(sortedComputers)
					sortedComputersString := strings.Join(sortedComputers, ",")
					tripleConnections[sortedComputersString] = true
				}
			}
		}
	}

	// Now find any tripple connections with a computer that starts with t
	total := 0
	for tc := range tripleConnections {
		split := strings.Split(tc, ",")

		if split[0][0] == 't' || split[1][0] == 't' || split[2][0] == 't' {
			total++
		}
	}

	return strconv.Itoa(total)
}

func PartTwo(filename string) string {
	input := readInput(filename)

	connections := make(map[string]bool)
	computerList := make(map[string]bool)
	for _, line := range input {
		computers := strings.Split(line, "-")
		slices.Sort(computers)
		stringJoin := strings.Join(computers, ",")
		connections[stringJoin] = true
		computerList[computers[0]] = true
		computerList[computers[1]] = true
	}

	largestConnectedGroup := ""
	// For each computer, check what other computers its connected to
	// Return largest connection
	for comp := range computerList {
		connectionGroup := comp
		for compCheck := range computerList {
			// First make sure we aren't checking something already in the group
			compsInGroup := strings.Split(connectionGroup, ",")
			alreadyIn := false
			for _, checkNotInComp := range compsInGroup {
				if compCheck == checkNotInComp {
					alreadyIn = true
				}
			}
			if alreadyIn {
				continue
			}

			// Check if we are connected to every other comp in this group
			connectedToAllComps := true
			for _, checkConnectedTo := range compsInGroup {
				checkConn := []string{compCheck, checkConnectedTo}
				slices.Sort(checkConn)
				checkConnStr := strings.Join(checkConn, ",")
				if !connections[checkConnStr] {
					connectedToAllComps = false
				}
			}

			if connectedToAllComps {
				connectionGroup += "," + compCheck
			}
		}

		if len(connectionGroup) > len(largestConnectedGroup) {
			largestConnectedGroup = connectionGroup
		}
	}

	largestCompGroupSplit := strings.Split(largestConnectedGroup, ",")
	slices.Sort(largestCompGroupSplit)
	largestConnectedGroup = strings.Join(largestCompGroupSplit, ",")

	return largestConnectedGroup
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
