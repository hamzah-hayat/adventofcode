package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
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
		fmt.Println("Gold:" + PartTwo("input"))
	case "p1":
		fmt.Println("Silver:" + PartOne("input"))
	case "p2":
		fmt.Println("Gold:" + PartTwo("input"))
	}
}

func PartOne(filename string) string {
	input := readInput(filename)

	maze := make(map[Point]string)

	// Read in input
	for y, line := range input {
		for x := 0; x < len(line); x++ {
			maze[Point{x, y}] = string(line[x])
		}
	}

	reindeerScore := 0
	startNode := FindNode("S", maze)
	endNode := FindNode("E", maze)
	path := PathFind(&startNode, &endNode, maze)
	if len(path) != 0 {
		reindeerScore = PathCost(path)
	}

	return strconv.Itoa(reindeerScore)
}

func FindNode(nodeString string, maze map[Point]string) Node {
	node := Node{}
	for i, v := range maze {
		if v == nodeString {
			node.point = i
		}
	}
	return node
}

func PathFind(start, end *Node, maze map[Point]string) []*Node {
	var closedSet []*Node
	var openSet = []*Node{start}
	start.cost = 0
	start.direction = East

	var paths [][]*Node

	for len(openSet) > 0 {
		current, leastCostNodeIndex := LeastCostNode(openSet)
		openSet = append(openSet[:leastCostNodeIndex], openSet[leastCostNodeIndex+1:]...)
		if current.point == end.point {
			var path []*Node
			pathNode := current
			for pathNode != nil {
				path = append([]*Node{pathNode}, path...)
				pathNode = pathNode.prev
			}
			paths = append(paths, path)
		}
		closedSet = append(closedSet, current)

		// Check each direction
		north := Point{current.point.x, current.point.y - 1}
		east := Point{current.point.x + 1, current.point.y}
		south := Point{current.point.x, current.point.y + 1}
		west := Point{current.point.x - 1, current.point.y}

		if maze[north] != "#" {
			newNode := Node{north, current.cost + 1, North, current}
			newNode.cost += ChangeDirectionCost(current.direction, North)

			// Check if in OpenSet/ClosedSet
			if InSet(newNode, openSet) {
				ReplaceIfBetterCost(newNode, openSet)
			} else if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if maze[east] != "#" {
			newNode := Node{east, current.cost + 1, East, current}
			newNode.cost += ChangeDirectionCost(current.direction, East)
			// Check if in OpenSet/ClosedSet
			if InSet(newNode, openSet) {
				ReplaceIfBetterCost(newNode, openSet)
			} else if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if maze[south] != "#" {
			newNode := Node{south, current.cost + 1, South, current}
			newNode.cost += ChangeDirectionCost(current.direction, South)
			// Check if in OpenSet/ClosedSet
			if InSet(newNode, openSet) {
				ReplaceIfBetterCost(newNode, openSet)
			} else if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if maze[west] != "#" {
			newNode := Node{west, current.cost + 1, West, current}
			newNode.cost += ChangeDirectionCost(current.direction, West)
			// Check if in OpenSet/ClosedSet
			if InSet(newNode, openSet) {
				ReplaceIfBetterCost(newNode, openSet)
			} else if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}

	}

	lowestCost := 0
	lowestCostPath := paths[0]
	for _, path := range paths {
		if PathCost(path) < lowestCost {
			lowestCost = PathCost(path)
			lowestCostPath = path
		}
	}

	return lowestCostPath
}

// Pop least cost node
func LeastCostNode(set []*Node) (*Node, int) {
	leastCost := math.MaxInt
	leastCostNode := set[0]
	leastCostNodeIndex := 0
	for i, n := range set {
		if n.cost < leastCost {
			leastCost = n.cost
			leastCostNode = n
			leastCostNodeIndex = i
		}
	}
	return leastCostNode, leastCostNodeIndex
}

// Check if the newNode is already in this set
func InSet(newNode Node, set []*Node) bool {
	for _, n := range set {
		if newNode.point == n.point && newNode.direction == n.direction {
			return true
		}
	}
	return false
}

func ReplaceIfBetterCost(newNode Node, set []*Node) {
	for _, n := range set {
		if newNode.point == n.point && newNode.direction == n.direction {
			if newNode.cost < n.cost {
				n.cost = newNode.cost
				n.prev = newNode.prev
			}
		}
	}
}

func ChangeDirectionCost(currentDirection, newDirection int) int {
	// Same dir
	if currentDirection == newDirection {
		return 0
	}

	// Different
	// Pure 💩 code
	switch currentDirection {
	case North:
		if newDirection == South {
			return 2000
		} else {
			return 1000
		}
	case East:
		if newDirection == West {
			return 2000
		} else {
			return 1000
		}
	case South:
		if newDirection == North {
			return 2000
		} else {
			return 1000
		}
	case West:
		if newDirection == East {
			return 2000
		} else {
			return 1000
		}
	}

	return 0
}

func PathCost(path []*Node) int {
	return path[len(path)-1].cost
}

type Point struct {
	x int
	y int
}

type Node struct {
	point     Point
	cost      int
	direction int
	prev      *Node
}

const (
	North = iota
	East
	South
	West
)

func PartTwo(filename string) string {
	input := readInput(filename)

	maze := make(map[Point]string)

	// Read in input
	for y, line := range input {
		for x := 0; x < len(line); x++ {
			maze[Point{x, y}] = string(line[x])
		}
	}

	startNode := FindNode("S", maze)
	endNode := FindNode("E", maze)
	paths := PathFindAllPaths(&startNode, &endNode, maze)

	sittingSpots := make(map[Point]int)
	for _, path := range paths {
		for _, node := range path {
			sittingSpots[node.point] += 1
		}
	}

	return strconv.Itoa(len(sittingSpots))
}

func PathFindAllPaths(start, end *Node, maze map[Point]string) [][]*Node {
	var closedSet []*Node
	var openSet = []*Node{start}
	start.cost = 0
	start.direction = East

	var paths [][]*Node

	for len(openSet) > 0 {
		current, leastCostNodeIndex := LeastCostNode(openSet)
		openSet = append(openSet[:leastCostNodeIndex], openSet[leastCostNodeIndex+1:]...)

		if current.point == end.point {
			var path []*Node
			pathNode := current
			for pathNode != nil {
				path = append([]*Node{pathNode}, path...)
				pathNode = pathNode.prev
			}
			paths = append(paths, path)
		}
		closedSet = append(closedSet, current)

		// Check each direction
		north := Point{current.point.x, current.point.y - 1}
		east := Point{current.point.x + 1, current.point.y}
		south := Point{current.point.x, current.point.y + 1}
		west := Point{current.point.x - 1, current.point.y}

		if maze[north] != "#" {
			newNode := Node{north, current.cost + 1, North, current}
			newNode.cost += ChangeDirectionCost(current.direction, North)

			// Check if in OpenSet/ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if maze[east] != "#" {
			newNode := Node{east, current.cost + 1, East, current}
			newNode.cost += ChangeDirectionCost(current.direction, East)
			// Check if in OpenSet/ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if maze[south] != "#" {
			newNode := Node{south, current.cost + 1, South, current}
			newNode.cost += ChangeDirectionCost(current.direction, South)
			// Check if in OpenSet/ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if maze[west] != "#" {
			newNode := Node{west, current.cost + 1, West, current}
			newNode.cost += ChangeDirectionCost(current.direction, West)
			// Check if in OpenSet/ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}

	}

	lowestCost := math.MaxInt
	for _, path := range paths {
		if PathCost(path) < lowestCost {
			lowestCost = PathCost(path)
		}
	}

	var bestPaths [][]*Node
	for _, path := range paths {
		if PathCost(path) == lowestCost {
			bestPaths = append(bestPaths, path)
		}
	}

	return bestPaths
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

// Read data from input.txt
// Return the string as int
func readInputInt() []int {

	var input []int

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		input = append(input, num)
	}
	return input
}
