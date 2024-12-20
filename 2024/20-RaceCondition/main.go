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
		fmt.Println("Silver:" + PartOne("input", 100))
		fmt.Println("Gold:" + PartTwo("input", 100))
	case "p1":
		fmt.Println("Silver:" + PartOne("input", 100))
	case "p2":
		fmt.Println("Gold:" + PartTwo("input", 100))
	}
}

func PartOne(filename string, goal int) string {
	input := readInput(filename)

	raceGrid := make(map[Point]string)

	// Read in input
	for y, line := range input {
		for x := 0; x < len(line); x++ {
			raceGrid[Point{x, y}] = string(line[x])
		}
	}

	// Find gridMaxX,gridMaxY
	gridMaxX := 0
	gridMaxY := 0
	for p := range raceGrid {
		if p.x > gridMaxX {
			gridMaxX = p.x
		}
		if p.y > gridMaxY {
			gridMaxY = p.y
		}
	}

	// print Grid
	// fmt.Println(PrintGrid(raceGrid, gridMaxX, gridMaxY))

	// Find our normal path cost
	startNode := FindNode("S", raceGrid)
	endNode := FindNode("E", raceGrid)
	normalPath := PathFindNormal(&startNode, &endNode, raceGrid)

	// Now find all cheatPaths that have a cost that is a goal less then normal path cost
	cheatPathsLess := PathFindCheat(raceGrid, goal, normalPath)

	return strconv.Itoa(cheatPathsLess)
}

type Point struct {
	x int
	y int
}

type Node struct {
	point      Point
	cost       int
	cheatPoint Point
	hasCheated bool
	prev       *Node
}

func PrintGrid(grid map[Point]string, gridMaxX, gridMaxY int) string {
	// Print grid
	gridPrint := ""
	for y := 0; y < gridMaxY; y++ {
		for x := 0; x < gridMaxX; x++ {
			if grid[Point{x, y}] == "." {
				gridPrint += "⬜"
			} else if grid[Point{x, y}] == "S" {
				gridPrint += "🚥"
			} else if grid[Point{x, y}] == "E" {
				gridPrint += "🏁"
			} else {
				gridPrint += "⬛"
			}
		}
		gridPrint += "\n"
	}
	return gridPrint
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

func PathFindCheat(maze map[Point]string, goal int, standardPath []*Node) int {
	var openSet = standardPath
	cheatPaths := 0

	for len(openSet) > 0 {
		current := openSet[0]
		openSet = openSet[1:]

		// Check each direction
		north := Point{current.point.x, current.point.y - 1}
		east := Point{current.point.x + 1, current.point.y}
		south := Point{current.point.x, current.point.y + 1}
		west := Point{current.point.x - 1, current.point.y}

		north2 := Point{current.point.x, current.point.y - 2}
		east2 := Point{current.point.x + 2, current.point.y}
		south2 := Point{current.point.x, current.point.y + 2}
		west2 := Point{current.point.x - 2, current.point.y}

		if maze[north] == "#" && (maze[north2] == "." || maze[north2] == "E") {
			newNode := Node{north2, current.cost + 2, north, true, current}
			// Check if in the normalPath
			if !CheatDoesntSaveEnoughTime(&newNode, goal, standardPath) {
				cheatPaths++
			}
		}
		if maze[east] == "#" && (maze[east2] == "." || maze[east2] == "E") {
			newNode := Node{east2, current.cost + 2, east, true, current}
			// Check if in the normalPath
			if !CheatDoesntSaveEnoughTime(&newNode, goal, standardPath) {
				cheatPaths++
			}
		}
		if maze[south] == "#" && (maze[south2] == "." || maze[south2] == "E") {
			newNode := Node{south2, current.cost + 2, south, true, current}
			// Check if in the normalPath
			if !CheatDoesntSaveEnoughTime(&newNode, goal, standardPath) {
				cheatPaths++
			}
		}
		if maze[west] == "#" && (maze[west2] == "." || maze[west2] == "E") {
			newNode := Node{west2, current.cost + 2, west, true, current}
			// Check if in the normalPath
			if !CheatDoesntSaveEnoughTime(&newNode, goal, standardPath) {
				cheatPaths++
			}
		}
	}
	return cheatPaths
}

// If this node (compared to the standard path)
// Doesn't save enough time, we want to prune
func CheatDoesntSaveEnoughTime(current *Node, goal int, standardPath []*Node) bool {
	for _, n := range standardPath {
		if current.point == n.point {
			// Did we save enough time?
			if current.cost+goal <= n.cost {
				return false
			}
		}
	}
	return true
}

func PathFindNormal(start, end *Node, maze map[Point]string) []*Node {
	var closedSet []*Node
	var openSet = []*Node{start}
	start.cost = 0

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
			return path
		}
		closedSet = append(closedSet, current)

		// Check each direction
		north := Point{current.point.x, current.point.y - 1}
		east := Point{current.point.x + 1, current.point.y}
		south := Point{current.point.x, current.point.y + 1}
		west := Point{current.point.x - 1, current.point.y}

		if maze[north] == "." || maze[north] == "E" {
			newNode := Node{north, current.cost + 1, current.cheatPoint, current.hasCheated, current}
			// Check if in OpenSet/ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if maze[east] == "." || maze[east] == "E" {
			newNode := Node{east, current.cost + 1, current.cheatPoint, current.hasCheated, current}
			// Check if in OpenSet/ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if maze[south] == "." || maze[south] == "E" {
			newNode := Node{south, current.cost + 1, current.cheatPoint, current.hasCheated, current}
			// Check if in OpenSet/ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if maze[west] == "." || maze[west] == "E" {
			newNode := Node{west, current.cost + 1, current.cheatPoint, current.hasCheated, current}
			// Check if in OpenSet/ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}

	}

	return nil
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
		if newNode.point == n.point && newNode.cheatPoint == n.cheatPoint && newNode.hasCheated == n.hasCheated {
			return true
		}
	}
	return false
}

func ReplaceIfBetterCost(newNode Node, set []*Node) {
	for _, n := range set {
		if newNode.point == n.point && newNode.hasCheated == n.hasCheated {
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

func PartTwo(filename string, goal int) string {
	input := readInput(filename)

	raceGrid := make(map[Point]string)

	// Read in input
	for y, line := range input {
		for x := 0; x < len(line); x++ {
			raceGrid[Point{x, y}] = string(line[x])
		}
	}

	// Find gridMaxX,gridMaxY
	gridMaxX := 0
	gridMaxY := 0
	for p := range raceGrid {
		if p.x > gridMaxX {
			gridMaxX = p.x
		}
		if p.y > gridMaxY {
			gridMaxY = p.y
		}
	}

	// print Grid
	// fmt.Println(PrintGrid(raceGrid, gridMaxX, gridMaxY))

	// Find our normal path cost
	startNode := FindNode("S", raceGrid)
	endNode := FindNode("E", raceGrid)
	normalPath := PathFindNormal(&startNode, &endNode, raceGrid)

	// Now find all cheatPaths that have a cost that is a goal less then normal path cost
	cheatPathsLess := PathFindSuperCheat(raceGrid, goal, normalPath)

	return strconv.Itoa(cheatPathsLess)
}

func PathFindSuperCheat(maze map[Point]string, goal int, standardPath []*Node) int {
	var openSet = standardPath
	cheatPaths := 0

	for len(openSet) > 0 {
		current := openSet[0]
		openSet = openSet[1:]

		// Check every node that is in Manhatten distance of 20 or less
		for _, node := range standardPath {

			distance := ManhattenDistance(current.point, node.point)

			if distance <= 20 && (maze[node.point] == "." || maze[node.point] == "E") {
				newNode := Node{node.point, current.cost + distance, node.point, true, current}
				// Check this node
				if !CheatDoesntSaveEnoughTime(&newNode, goal, standardPath) {
					cheatPaths++
				}
			}
		}
	}
	return cheatPaths
}

// Figure out the Manhatten Distance between two points
func ManhattenDistance(firstPoint Point, secondPoint Point) int {
	x := abs(firstPoint.x - secondPoint.x)
	y := abs(firstPoint.y - secondPoint.y)
	return x + y
}

// Absoulute value of Int
func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
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
