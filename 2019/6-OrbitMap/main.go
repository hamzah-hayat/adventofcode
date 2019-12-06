package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p2", "The method/part that should be run, valid are p1,p2 and test")
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

	var orbitTarget []string
	var orbitObject []string

	for _, orbit := range input {
		// Split string again
		orbitItems := strings.Split(orbit, ")")

		// First find root item
		orbitTarget = append(orbitTarget, orbitItems[0])
		orbitObject = append(orbitObject, orbitItems[1])
	}

	// Root node is COM

	nodesToCheck := []*Tree{}
	root := Tree{name: "COM", orbits: 0}
	nodesToCheck = append(nodesToCheck, &root)

	for len(nodesToCheck) > 0 {

		for i, orbitT := range orbitTarget {
			if orbitT == nodesToCheck[0].name {
				newTree := Tree{name: orbitObject[i], orbits: nodesToCheck[0].orbits + 1}
				nodesToCheck = append(nodesToCheck, &newTree)
				nodesToCheck[0].children = append(nodesToCheck[0].children, &newTree)
			}
		}

		// now remove first node from nodesToCheck

		// Remove the element at index i from a.
		copy(nodesToCheck[0:], nodesToCheck[1:])          // Shift a[i+1:] left one index.
		nodesToCheck[len(nodesToCheck)-1] = &Tree{}       // Erase last element (write zero value).
		nodesToCheck = nodesToCheck[:len(nodesToCheck)-1] // Truncate slice.

	}

	fmt.Println("The number of orbits is ", Walk(&root))

}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// Walk traverses a tree depth-first,
func Walk(t *Tree) int {
	if t.children == nil {
		return t.orbits
	} else {
		total := t.orbits
		for _, child := range t.children {
			total += Walk(child)
		}
		return total
	}
}

func PartTwo() {
	input := readInput()

	root := BuildTree(input)

	// Find shortest
	lowestSanta := 10000
	lowestYou := 10000

	currentNode := root

	for {
		currentNodeName := currentNode.name

		for _, child := range currentNode.children {
			distanceToSanta, found := FindDistance(child, "SAN")
			distanceToYou, found2 := FindDistance(child, "YOU")

			if found && found2 && distanceToSanta < lowestSanta && distanceToYou < lowestYou {
				lowestSanta = distanceToSanta
				lowestYou = distanceToYou
				currentNode = *child
				break
			}
		}
		if currentNodeName == currentNode.name {
			break
		}
	}
	fmt.Println(currentNode)
	s, _ := FindDistance(&currentNode, "SAN")
	y, _ := FindDistance(&currentNode, "YOU")
	fmt.Println("Distance to Santa is", s+y)
}

// FindDistance looks for a node, and returns the distance to reach it
func FindDistance(t *Tree, nodeName string) (int, bool) {

	if t.name == nodeName {
		return -1, true
	} else if t.children == nil {
		return 0, false
	} else {
		total := 1
		f := false
		for _, child := range t.children {
			t, found := FindDistance(child, nodeName)
			if found {
				total += t
				f = true
			}
		}
		return total, f
	}
}

func BuildTree(input []string) Tree {
	var orbitTarget []string
	var orbitObject []string

	for _, orbit := range input {
		// Split string again
		orbitItems := strings.Split(orbit, ")")

		// First find root item
		orbitTarget = append(orbitTarget, orbitItems[0])
		orbitObject = append(orbitObject, orbitItems[1])
	}

	// Root node is COM

	nodesToCheck := []*Tree{}
	root := Tree{name: "COM", orbits: 0}
	nodesToCheck = append(nodesToCheck, &root)

	for len(nodesToCheck) > 0 {

		for i, orbitT := range orbitTarget {
			if orbitT == nodesToCheck[0].name {
				newTree := Tree{name: orbitObject[i], orbits: nodesToCheck[0].orbits + 1}
				nodesToCheck = append(nodesToCheck, &newTree)
				nodesToCheck[0].children = append(nodesToCheck[0].children, &newTree)
			}
		}

		// now remove first node from nodesToCheck

		// Remove the element at index i from a.
		copy(nodesToCheck[0:], nodesToCheck[1:])          // Shift a[i+1:] left one index.
		nodesToCheck[len(nodesToCheck)-1] = &Tree{}       // Erase last element (write zero value).
		nodesToCheck = nodesToCheck[:len(nodesToCheck)-1] // Truncate slice.

	}
	return root
}

type Tree struct {
	name     string
	orbits   int
	children []*Tree
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
