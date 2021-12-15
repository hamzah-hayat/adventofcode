package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	methodP *string
)

func parseFlags() {
	methodP = flag.String("method", "p2", "The method/part that should be run, valid are p1,p2 and test")
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
		break
	case "p2":
		fmt.Println("Gold:" + PartTwo("input"))
		break
	}
}

func PartOne(filename string) string {
	input := readInput(filename)

	riskMap := make(map[Point]int)

	maxX := 0
	maxY := 0

	for x, v := range input {
		for y, c := range v {
			num, _ := strconv.Atoi(string(c))
			riskMap[Point{x, y}] = num
			if y > maxY {
				maxY = y
			}
		}
		if x > maxX {
			maxX = x
		}
	}

	totalRisk := findShortestPath(Point{maxX, maxY}, riskMap)
	num := strconv.Itoa(totalRisk)

	return num
}

func PartTwo(filename string) string {
	input := readInput(filename)

	riskMap := make(map[Point]int)

	// We need to copy this into a 5x5 grid
	for x, v := range input {
		for y, c := range v {
			num, _ := strconv.Atoi(string(c))
			riskMap[Point{x, y}] = num
		}
	}

	maxX := 0
	maxY := 0
	// Initial max x and max y
	for p := range riskMap {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	expandedRiskMap := expandRiskMap(riskMap, 5, 5)

	// Larger max x and max y
	for p := range expandedRiskMap {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	totalRisk := findShortestPath(Point{maxX, maxY}, expandedRiskMap)
	num := strconv.Itoa(totalRisk)

	return num
}

type Point struct {
	X int
	Y int
}

type Node struct {
	point             Point
	distance          int
	distanceManhatten int
	prev              []Node
}

// 1  function Dijkstra(Graph, source):
//  2
//  3      create vertex set Q
//  4
//  5      for each vertex v in Graph:
//  6          dist[v] ← INFINITY
//  7          prev[v] ← UNDEFINED
//  8          add v to Q
//  9      dist[source] ← 0
// 10
// 11      while Q is not empty:
// 12          u ← vertex in Q with min dist[u]
// 13
// 14          remove u from Q
// 15
// 16          for each neighbor v of u still in Q:
// 17              alt ← dist[u] + length(u, v)
// 18              if alt < dist[v]:
// 19                  dist[v] ← alt
// 20                  prev[v] ← u
// 21
// 22      return dist[], prev[]

// Finds the shortest path
func findShortestPath(goal Point, riskMap map[Point]int) int {

	nodeList := []Node{}

	for point := range riskMap {
		if point.X == 0 && point.Y == 0 {
			nodeList = append(nodeList, Node{point, 0, 0, []Node{}})
		} else {
			nodeList = append(nodeList, Node{point, 100000000, 100000000 + ManHatten(point, goal), []Node{}})
		}
	}

	visitedNodes := []Node{}

	// while queue is not empty
	for len(nodeList) > 0 {

		currentNode := nodeList[0]
		nodeIndex := 0
		for i, v := range nodeList {
			if currentNode.distanceManhatten > v.distanceManhatten {
				currentNode = v
				nodeIndex = i
			}
		}
		nodeList = remove(nodeList, nodeIndex)

		// terminating search when reaching goal
		if currentNode.point.X == goal.X && currentNode.point.Y == goal.Y {
			break
		}

		neighbours := GetNeighbours(riskMap, currentNode.point)

		// iterating through all the vertices that we can reach from current vertex
		for p, v := range neighbours {

			inList := false
			index := 0
			for i, v := range nodeList {
				if p == v.point {
					index = i
					inList = true
				}
			}

			if inList {
				newDist := currentNode.distance + v
				if newDist < nodeList[index].distance {
					visitedNodes = append(visitedNodes, Node{p, newDist, newDist + ManHatten(p, goal), append([]Node{}, currentNode)})
					nodeList[index].distance = newDist
					nodeList[index].distanceManhatten = newDist + ManHatten(p, goal)
					nodeList[index].prev = append([]Node{}, currentNode)
				}
			}
		}
	}

	// Find goal node in VisitedNodes and return distance
	distance := 0
	for _, v := range visitedNodes {
		if v.point == goal {
			distance = v.distance
		}
	}

	return distance
}

// Get all possible neighbours off this point
func GetNeighbours(points map[Point]int, p Point) map[Point]int {

	neighbours := make(map[Point]int)

	// now we try and check a adjacent point
	// check if it exists
	// then add to neighbour list
	up := Point{p.X + 1, p.Y}
	if _, ok := points[up]; ok {
		neighbours[up] = points[up]
	}
	right := Point{p.X, p.Y + 1}
	if _, ok := points[right]; ok {
		neighbours[right] = points[right]
	}
	down := Point{p.X - 1, p.Y}
	if _, ok := points[down]; ok {
		neighbours[down] = points[down]
	}
	left := Point{p.X, p.Y - 1}
	if _, ok := points[left]; ok {
		neighbours[left] = points[left]
	}

	return neighbours
}

func expandRiskMap(riskMap map[Point]int, length, height int) map[Point]int {

	expandedRiskMap := make(map[Point]int)

	maxX := 0
	maxY := 0
	// Initial max x and max y
	for p := range riskMap {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	for p, v := range riskMap {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				newPoint := Point{p.X + (i * (maxX + 1)), p.Y + (j * (maxY + 1))}
				num := (v + i + j)
				for num > 9 {
					// lol what am i even doing?
					num -= 9
				}
				expandedRiskMap[newPoint] = num
			}
		}
	}

	return expandedRiskMap
}

func ManHatten(p1, p2 Point) int {
	return abs(p1.X - p2.X + p1.Y - p2.Y)
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

// remove element from array
func remove(s []Node, i int) []Node {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
