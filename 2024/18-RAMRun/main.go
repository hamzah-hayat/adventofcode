package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
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
		fmt.Println("Silver:" + PartOne("input", 71, 71, 1024))
		fmt.Println("Gold:" + PartTwo("input", 71, 71))
	case "p1":
		fmt.Println("Silver:" + PartOne("input", 71, 71, 1024))
	case "p2":
		fmt.Println("Gold:" + PartTwo("input", 71, 71))
	}
}

func PartOne(filename string, gridMaxX, gridMaxY, stopAfterBytes int) string {
	input := readInput(filename)

	maze := make(map[Point]string)

	// Set Grid initial
	for y := 0; y < gridMaxY; y++ {
		for x := 0; x < gridMaxX; x++ {
			maze[Point{x, y}] = "."
		}
	}

	// Set blockers
	for i := 0; i < stopAfterBytes; i++ {
		splitLine := strings.Split(input[i], ",")
		x, _ := strconv.Atoi(splitLine[0])
		y, _ := strconv.Atoi(splitLine[1])
		maze[Point{x, y}] = "#"
	}

	// print Grid
	// fmt.Println(PrintGrid(maze, gridMaxX, gridMaxY))

	runScore := 0
	startNode := FindNodeWithPoint(Point{0, 0}, maze)
	endNode := FindNodeWithPoint(Point{gridMaxX - 1, gridMaxY - 1}, maze)
	path := PathFind(&startNode, &endNode, maze)
	if len(path) != 0 {
		runScore = PathCost(path)
	}

	return strconv.Itoa(runScore)
}

func PrintGrid(grid map[Point]string, gridMaxX, gridMaxY int) string {
	// Print grid
	gridPrint := ""
	for y := 0; y < gridMaxY; y++ {
		for x := 0; x < gridMaxX; x++ {
			if grid[Point{x, y}] == "." {
				gridPrint += "⬜"
			} else {
				gridPrint += "⬛"
			}
		}
		gridPrint += "\n"
	}
	return gridPrint
}

func FindNodeWithPoint(nodePoint Point, maze map[Point]string) Node {
	node := Node{}
	for i := range maze {
		if i == nodePoint {
			node.point = i
		}
	}
	return node
}

func PathFind(start, end *Node, maze map[Point]string) []*Node {
	var closedSet []*Node
	var openSet = []*Node{start}
	start.cost = 0

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

		if maze[north] == "." {
			newNode := Node{north, current.cost + 1, current}
			// Check if in OpenSet/ClosedSet
			if InSet(newNode, openSet) {
				ReplaceIfBetterCost(newNode, openSet)
			} else if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if maze[east] == "." {
			newNode := Node{east, current.cost + 1, current}
			// Check if in OpenSet/ClosedSet
			if InSet(newNode, openSet) {
				ReplaceIfBetterCost(newNode, openSet)
			} else if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if maze[south] == "." {
			newNode := Node{south, current.cost + 1, current}
			// Check if in OpenSet/ClosedSet
			if InSet(newNode, openSet) {
				ReplaceIfBetterCost(newNode, openSet)
			} else if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if maze[west] == "." {
			newNode := Node{west, current.cost + 1, current}
			// Check if in OpenSet/ClosedSet
			if InSet(newNode, openSet) {
				ReplaceIfBetterCost(newNode, openSet)
			} else if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}

	}

	if len(paths) != 0 {
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

	// Otherwise no path, return nothing
	return []*Node{}
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
		if newNode.point == n.point {
			return true
		}
	}
	return false
}

func ReplaceIfBetterCost(newNode Node, set []*Node) {
	for _, n := range set {
		if newNode.point == n.point {
			if newNode.cost < n.cost {
				n.cost = newNode.cost
				n.prev = newNode.prev
			}
		}
	}
}

func PathCost(path []*Node) int {
	return path[len(path)-1].cost
}

type Point struct {
	x int
	y int
}

type Node struct {
	point Point
	cost  int
	prev  *Node
}

const (
	North = iota
	East
	South
	West
)

func PartTwo(filename string, gridMaxX, gridMaxY int) string {
	input := readInput(filename)

	// Set blockers
	counter := 0
	lastPoint := ""
	for {
		maze := make(map[Point]string)

		// Set Grid inital
		for y := 0; y < gridMaxY; y++ {
			for x := 0; x < gridMaxX; x++ {
				maze[Point{x, y}] = "."
			}
		}
		for i := 0; i < counter; i++ {
			lastPoint = input[i]
			splitLine := strings.Split(input[i], ",")
			x, _ := strconv.Atoi(splitLine[0])
			y, _ := strconv.Atoi(splitLine[1])
			maze[Point{x, y}] = "#"
		}

		// print Grid
		// fmt.Println(PrintGrid(maze, gridMaxX, gridMaxY))

		startNode := FindNodeWithPoint(Point{0, 0}, maze)
		endNode := FindNodeWithPoint(Point{gridMaxX - 1, gridMaxY - 1}, maze)
		path := PathFind(&startNode, &endNode, maze)
		if len(path) == 0 {
			break
		} else {
			counter++
		}
	}

	return lastPoint
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
