package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
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

var pathNames map[string]bool

func PartOne() {
	pathNames = make(map[string]bool)
	input := readInput()

	caves := CreateCaveMap(input)

	CreatePathsToExit(caves)

	fmt.Println(len(pathNames))

}

func PartTwo() {
	pathNames = make(map[string]bool)
	input := readInput()

	caves := CreateCaveMap(input)

	CreatePathsToExitSpecialCave(caves)

	fmt.Println(len(pathNames))
}

type Cave struct {
	name        string
	large       bool
	linkedCaves []string
}

// Each path is formated like
// start,A,b,A,c,A,end
// start,A,b,A,end
// etc

func CreatePathsToExit(caves []Cave) {
	// Starting from first cave
	startCave := getCave(caves, "start")
	currentPath := []Cave{startCave}
	currentPathStr := "start"
	FindPathRecursive(caves, currentPath, currentPathStr)
}

func FindPathRecursive(caves []Cave, currentPath []Cave, currentPathStr string) {

	// If we are at end, return
	if currentPath[len(currentPath)-1].name == "end" {
		pathNames[currentPathStr] = true
		return
	}

	// Otherwise look at our connections
	for _, v := range currentPath[len(currentPath)-1].linkedCaves {
		// Check cave is already if currentPath if not large, otherwise use it
		cave := getCave(caves, v)
		if cave.large {
			FindPathRecursive(caves, append(currentPath, cave), currentPathStr+","+cave.name)
		} else if !cave.large && !hasCave(currentPath, cave.name) {
			FindPathRecursive(caves, append(currentPath, cave), currentPathStr+","+cave.name)
		}
	}
}

func CreatePathsToExitSpecialCave(caves []Cave) {
	// Starting from first cave
	startCave := getCave(caves, "start")
	currentPath := []Cave{startCave}
	currentPathStr := "start"

	for _, c := range caves {
		if c.name == "start" || c.name == "end" || c.large == true {
			continue
		}
		FindPathRecursiveSpecialCave(caves, currentPath, currentPathStr, c)
	}
}

func FindPathRecursiveSpecialCave(caves []Cave, currentPath []Cave, currentPathStr string, superSpecialSmallCave Cave) {

	// If we are at end, return
	if currentPath[len(currentPath)-1].name == "end" {
		pathNames[currentPathStr] = true
		return
	}

	// Otherwise look at our connections
	for _, v := range currentPath[len(currentPath)-1].linkedCaves {
		// Check cave is already if currentPath if not large, otherwise use it
		cave := getCave(caves, v)
		if cave.large {
			FindPathRecursiveSpecialCave(caves, append(currentPath, cave), currentPathStr+","+cave.name, superSpecialSmallCave)
		} else if !cave.large && superSpecialSmallCave.name == cave.name && !(hasCaveSum(currentPath, cave.name) > 1) {
			FindPathRecursiveSpecialCave(caves, append(currentPath, cave), currentPathStr+","+cave.name, superSpecialSmallCave)
		} else if !cave.large && !hasCave(currentPath, cave.name) {
			FindPathRecursiveSpecialCave(caves, append(currentPath, cave), currentPathStr+","+cave.name, superSpecialSmallCave)
		}
	}
}

func CreateCaveMap(input []string) []Cave {
	var caveMap []Cave

	// For each line
	// First find all caves
	for _, v := range input {
		splitStr := strings.Split(v, "-")
		// Check both sides
		if !hasCave(caveMap, splitStr[0]) {
			caveMap = append(caveMap, Cave{splitStr[0], IsUpper(splitStr[0]), nil})
		}
		if !hasCave(caveMap, splitStr[1]) {
			caveMap = append(caveMap, Cave{splitStr[1], IsUpper(splitStr[1]), nil})
		}
	}

	// Now that we have all Caves, check connections
	for _, v := range input {
		splitStr := strings.Split(v, "-")
		addCaveConnection(caveMap, splitStr[0], splitStr[1])
		addCaveConnection(caveMap, splitStr[1], splitStr[0])
	}

	return caveMap
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

func hasCave(caves []Cave, caveName string) bool {
	for _, c := range caves {
		if c.name == caveName {
			return true
		}
	}
	return false
}

func hasCaveSum(caves []Cave, caveName string) int {
	num := 0
	for _, c := range caves {
		if c.name == caveName {
			num++
		}
	}
	return num
}

func addCaveConnection(caves []Cave, startCaveName, endCaveName string) {
	for i, c := range caves {
		newCave := c
		if newCave.name == startCaveName {
			newCave.linkedCaves = append(c.linkedCaves, endCaveName)
		}
		caves[i] = newCave
	}
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func getCave(caves []Cave, name string) Cave {
	for _, c := range caves {
		if c.name == name {
			return c
		}
	}
	return Cave{"notfound", false, nil}
}
