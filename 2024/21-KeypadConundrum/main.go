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

	// +---+---+---+
	// | 7 | 8 | 9 |
	// +---+---+---+
	// | 4 | 5 | 6 |
	// +---+---+---+
	// | 1 | 2 | 3 |
	// +---+---+---+
	//     | 0 | A |
	//     +---+---+
	bigKeypad := make(map[Point]string)
	bigKeypad[Point{0, 0}] = "7"
	bigKeypad[Point{1, 0}] = "8"
	bigKeypad[Point{2, 0}] = "9"
	bigKeypad[Point{0, 1}] = "4"
	bigKeypad[Point{1, 1}] = "5"
	bigKeypad[Point{2, 1}] = "6"
	bigKeypad[Point{0, 2}] = "1"
	bigKeypad[Point{1, 2}] = "2"
	bigKeypad[Point{2, 2}] = "3"
	bigKeypad[Point{1, 3}] = "0"
	bigKeypad[Point{2, 3}] = "A"

	// 	   +---+---+
	//     | ^ | A |
	// +---+---+---+
	// | < | v | > |
	// +---+---+---+
	smallKeypad := make(map[Point]string)
	smallKeypad[Point{1, 0}] = "^"
	smallKeypad[Point{2, 0}] = "A"
	smallKeypad[Point{0, 1}] = "<"
	smallKeypad[Point{1, 1}] = "v"
	smallKeypad[Point{2, 1}] = ">"

	total := 0
	for _, code := range input {
		totalPaths := RunRobotBigKeypad(bigKeypad, code)
		// Then run the mini Keypad two times
		robotPaths := make(map[string]int)
		for tp := range totalPaths {
			robotPathsNew := RunRobotSmallKeypad(smallKeypad, tp)
			for p := range robotPathsNew {
				robotPaths[p] += 1
			}
		}

		humanPaths := make(map[string]int)
		for rp := range robotPaths {
			humanPathsNew := RunRobotSmallKeypad(smallKeypad, rp)
			for p := range humanPathsNew {
				humanPaths[p] += 1
			}
		}

		lowest := math.MaxInt
		for hm := range humanPaths {
			if len(hm) < lowest {
				lowest = len(hm)
			}
		}

		keyPadNumber, _ := strconv.Atoi(code[:len(code)-1])

		total += keyPadNumber * lowest

	}

	return strconv.Itoa(total)
}

func RunRobotBigKeypad(bigKeypad map[Point]string, code string) map[string]int {
	start := FindNode("A", bigKeypad)
	totalPaths := make(map[string]int)
	totalPaths[""] = 0
	for _, codeSplit := range code {
		end := FindNode(string(codeSplit), bigKeypad)
		paths := PathFind(&start, &end, bigKeypad)
		start = FindNode(string(codeSplit), bigKeypad)

		newPaths := make(map[string]int)
		for tp := range totalPaths {
			for _, p := range paths {
				newPaths[tp+DirectionPath(p)+"A"] += 1
			}
			delete(totalPaths, tp)
		}
		totalPaths = newPaths
	}
	return totalPaths
}

func RunRobotSmallKeypad(smallKeypad map[Point]string, code string) map[string]int {
	start := FindNode("A", smallKeypad)
	totalPaths := make(map[string]int)
	totalPaths[""] = 0
	for _, codeSplit := range code {
		end := FindNode(string(codeSplit), smallKeypad)
		paths := PathFindSmallKeyPad(&start, &end, smallKeypad)
		start = FindNode(string(codeSplit), smallKeypad)

		newPaths := make(map[string]int)
		lowest := math.MaxInt
		for tp := range totalPaths {
			for _, p := range paths {
				newPath := tp + DirectionPath(p) + "A"
				if len(newPath) < lowest {
					lowest = len(newPath)
				}
				if len(newPath) == lowest {
					newPaths[newPath] += 1
				}
			}
			delete(totalPaths, tp)
		}
		totalPaths = newPaths
	}
	return totalPaths
}

func DirectionPath(node []*Node) string {
	directionStr := ""
	for i := 0; i < len(node); i++ {
		if node[i].direction != "" {
			directionStr += node[i].direction
		}
	}
	return directionStr
}

type Point struct {
	x int
	y int
}

type Node struct {
	point     Point
	cost      int
	direction string
	prev      *Node
}

func FindNode(nodeString string, keypad map[Point]string) Node {
	node := Node{}
	for i, v := range keypad {
		if v == nodeString {
			node.point = i
		}
	}
	return node
}

func PathFind(start, end *Node, keypad map[Point]string) [][]*Node {
	var closedSet []*Node
	var openSet = []*Node{start}
	start.cost = 0

	var paths [][]*Node

	lowestCost := math.MaxInt
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

			if len(path) < lowestCost {
				lowestCost = len(path)
			}
			if len(path) == lowestCost {
				paths = append(paths, path)
			}
		}
		closedSet = append(closedSet, current)

		if current.cost > lowestCost {
			continue
		}

		// Check each direction
		north := Point{current.point.x, current.point.y - 1}
		east := Point{current.point.x + 1, current.point.y}
		south := Point{current.point.x, current.point.y + 1}
		west := Point{current.point.x - 1, current.point.y}

		if keypad[north] != "" {
			newNode := Node{north, current.cost + 1, "^", current}
			// Check if in ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if keypad[east] != "" {
			newNode := Node{east, current.cost + 1, ">", current}
			// Check if in ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if keypad[south] != "" {
			newNode := Node{south, current.cost + 1, "v", current}
			// Check if in ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if keypad[west] != "" {
			newNode := Node{west, current.cost + 1, "<", current}
			// Check if in ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}

	}

	return paths
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

func PartTwo(filename string) string {
	input := readInput(filename)

	// +---+---+---+
	// | 7 | 8 | 9 |
	// +---+---+---+
	// | 4 | 5 | 6 |
	// +---+---+---+
	// | 1 | 2 | 3 |
	// +---+---+---+
	//     | 0 | A |
	//     +---+---+
	bigKeypad := make(map[Point]string)
	bigKeypad[Point{0, 0}] = "7"
	bigKeypad[Point{1, 0}] = "8"
	bigKeypad[Point{2, 0}] = "9"
	bigKeypad[Point{0, 1}] = "4"
	bigKeypad[Point{1, 1}] = "5"
	bigKeypad[Point{2, 1}] = "6"
	bigKeypad[Point{0, 2}] = "1"
	bigKeypad[Point{1, 2}] = "2"
	bigKeypad[Point{2, 2}] = "3"
	bigKeypad[Point{1, 3}] = "0"
	bigKeypad[Point{2, 3}] = "A"

	// 	   +---+---+
	//     | ^ | A |
	// +---+---+---+
	// | < | v | > |
	// +---+---+---+
	smallKeypad := make(map[Point]string)
	smallKeypad[Point{1, 0}] = "^"
	smallKeypad[Point{2, 0}] = "A"
	smallKeypad[Point{0, 1}] = "<"
	smallKeypad[Point{1, 1}] = "v"
	smallKeypad[Point{2, 1}] = ">"

	total := 0
	for _, code := range input {
		totalPaths := RunRobotBigKeypad(bigKeypad, code)

		lowestTupleScore := math.MaxInt
		for p := range totalPaths {
			// store as tuples
			tuples := make(map[string]int)
			startFromA := "A" + p
			for i := 0; i < len(startFromA)-1; i++ {
				tuples[string(startFromA[i])+string(startFromA[i+1])] += 1
			}

			// Run the small keypad 25 times, fug
			for i := 0; i < 2; i++ {
				// fmt.Println("Running Robot:", i)

				// For each tuple, divide into new tuples
				newTuples := make(map[string]int)
				for t, number := range tuples {
					subdivide := SubdivideTuple(string(t[0]), string(t[1]))
					for _, v := range subdivide {
						newTuples[v] += number
					}
				}
				tuples = newTuples
			}

			// Score our final tuple set
			total := 0
			for t, number := range tuples {
				total += SmallKeypadPathCost(string(t[0]), string(t[1])) * number
			}

			if total < lowestTupleScore {
				lowestTupleScore = total
			}
		}

		keyPadNumber, _ := strconv.Atoi(code[:len(code)-1])

		total += keyPadNumber * lowestTupleScore

	}

	return strconv.Itoa(total)
}

func SubdivideTuple(start, end string) []string {
	// 	   +---+---+
	//     | ^ | A |
	// +---+---+---+
	// | < | v | > |
	// +---+---+---+
	tuples := []string{}
	switch start {
	case "A":
		switch end {
		case "^":
			tuples = append(tuples, "<A")
		case ">":
			tuples = append(tuples, "vA")
		case "v":
			tuples = append(tuples, "<v")
			tuples = append(tuples, "vA")
		case "<":
			tuples = append(tuples, "v<")
			tuples = append(tuples, "<<")
			tuples = append(tuples, "<A")
		}
	case "^":
		switch end {
		case "A":
			tuples = append(tuples, ">A")
		case ">":
			tuples = append(tuples, ">v")
			tuples = append(tuples, "vA")
		case "v":
			tuples = append(tuples, "vA")
		case "<":
			tuples = append(tuples, "v<")
			tuples = append(tuples, "<A")
		}
	case ">":
		switch end {
		case "A":
			tuples = append(tuples, "^A")
		case "^":
			tuples = append(tuples, "<^")
			tuples = append(tuples, "^A")
		case "v":
			tuples = append(tuples, "<A")
		case "<":
			tuples = append(tuples, "<<")
			tuples = append(tuples, "<A")
		}
	case "v":
		switch end {
		case "A":
			tuples = append(tuples, ">^")
			tuples = append(tuples, "^A")
		case "^":
			tuples = append(tuples, "^A")
		case ">":
			tuples = append(tuples, ">A")
		case "<":
			tuples = append(tuples, "<A")
		}
	case "<":
		switch end {
		case "A":
			tuples = append(tuples, ">>")
			tuples = append(tuples, ">^")
			tuples = append(tuples, "^A")
		case "^":
			tuples = append(tuples, ">^")
			tuples = append(tuples, "^A")
		case "v":
			tuples = append(tuples, ">A")
		case ">":
			tuples = append(tuples, ">>")
			tuples = append(tuples, ">A")
		}
	}
	return tuples
}

func PathFindSmallKeyPad(start, end *Node, keypad map[Point]string) [][]*Node {
	var closedSet []*Node
	var openSet = []*Node{start}
	start.cost = 0

	var paths [][]*Node

	lowestCost := math.MaxInt
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

			pathCost := 0
			for i := 0; i < len(path)-1; i++ {
				pathCost += SmallKeypadPathCost(path[i].direction, path[i+1].direction)
			}

			if pathCost < lowestCost {
				lowestCost = pathCost
			}
			if pathCost == lowestCost {
				paths = append(paths, path)
			}
		}
		closedSet = append(closedSet, current)

		if current.cost > lowestCost {
			continue
		}

		// Check each direction
		north := Point{current.point.x, current.point.y - 1}
		east := Point{current.point.x + 1, current.point.y}
		south := Point{current.point.x, current.point.y + 1}
		west := Point{current.point.x - 1, current.point.y}

		if keypad[north] != "" {
			newNode := Node{north, current.cost + 1, "^", current}
			// Check if in ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if keypad[east] != "" {
			newNode := Node{east, current.cost + 1, ">", current}
			// Check if in ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if keypad[south] != "" {
			newNode := Node{south, current.cost + 1, "v", current}
			// Check if in ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}
		if keypad[west] != "" {
			newNode := Node{west, current.cost + 1, "<", current}
			// Check if in ClosedSet
			if !InSet(newNode, closedSet) {
				openSet = append(openSet, &newNode)
			}
		}

	}

	return paths
}

func TakeLowestPaths(robotPaths map[string]int) map[string]int {
	// We only want the shortest paths
	lowest := math.MaxInt
	for rp := range robotPaths {
		if len(rp) < lowest {
			lowest = len(rp)
		}
	}

	lowestPaths := make(map[string]int)
	counter := 0
	for p := range robotPaths {
		if counter >= 16 {
			break
		}
		if len(p) == lowest {
			lowestPaths[p] += 1
		}
		counter++
	}

	return lowestPaths
}

func SmallKeypadPathCost(start, end string) int {
	// 	   +---+---+
	//     | ^ | A |
	// +---+---+---+
	// | < | v | > |
	// +---+---+---+
	cost := 0
	switch start {
	case "A":
		switch end {
		case "^":
			cost = 1
		case ">":
			cost = 1
		case "v":
			cost = 2
		case "<":
			cost = 3
		}
	case "^":
		switch end {
		case "A":
			cost = 1
		case ">":
			cost = 2
		case "v":
			cost = 1
		case "<":
			cost = 2
		}
	case ">":
		switch end {
		case "A":
			cost = 1
		case "^":
			cost = 2
		case "v":
			cost = 1
		case "<":
			cost = 2
		}
	case "v":
		switch end {
		case "A":
			cost = 2
		case "^":
			cost = 1
		case ">":
			cost = 1
		case "<":
			cost = 1
		}
	case "<":
		switch end {
		case "A":
			cost = 3
		case "^":
			cost = 2
		case "v":
			cost = 1
		case ">":
			cost = 2
		}
	}
	return cost
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
