package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
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
	point    Point
	distance int
	prev     []Node
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

	dist := make(map[Point]int)
	visited := make(map[Point]bool)
	prev := make(map[Point]Point)

	q := NodeQueue{}
	pq := q.NewQ()

	pq.Enqueue(Node{Point{0, 0}, 0, []Node{}})

	for point := range riskMap {
		dist[point] = math.MaxInt64
	}
	// We never "enter the starting point"
	dist[Point{0, 0}] = 0

	// while queue is not empty
	for !pq.IsEmpty() {

		n := pq.Dequeue()

		if visited[n.point] {
			continue
		}
		visited[n.point] = true

		// terminating search when reaching goal
		if n.point.X == goal.X && n.point.Y == goal.Y {
			dist[goal] = n.distance
			break
		}

		neighbours := GetNeighbours(riskMap, n.point)

		// iterating through all the vertices that we can reach from current vertex
		for p, v := range neighbours {
			if !visited[p] {
				newDist := n.distance + v
				if newDist < dist[p] {
					prev[p] = n.point
					pq.Enqueue(Node{p, newDist, append([]Node{}, *n)})
					dist[n.point] = newDist
				}
			}
		}
	}

	// Find goal node in VisitedNodes and return distance
	distance := dist[goal]

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

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type NodeQueue struct {
	Items []Node
	Lock  sync.RWMutex
}

// Enqueue adds an Node to the end of the queue
func (s *NodeQueue) Enqueue(t Node) {
	s.Lock.Lock()
	if len(s.Items) == 0 {
		s.Items = append(s.Items, t)
		s.Lock.Unlock()
		return
	}
	var insertFlag bool
	for k, v := range s.Items {
		if t.distance < v.distance {
			if k > 0 {
				s.Items = append(s.Items[:k+1], s.Items[k:]...)
				s.Items[k] = t
				insertFlag = true
			} else {
				s.Items = append([]Node{t}, s.Items...)
				insertFlag = true
			}
		}
		if insertFlag {
			break
		}
	}
	if !insertFlag {
		s.Items = append(s.Items, t)
	}
	//s.items = append(s.items, t)
	s.Lock.Unlock()
}

// Dequeue removes an Node from the start of the queue
func (s *NodeQueue) Dequeue() *Node {
	s.Lock.Lock()
	item := s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	s.Lock.Unlock()
	return &item
}

//NewQ Creates New Queue
func (s *NodeQueue) NewQ() *NodeQueue {
	s.Lock.Lock()
	s.Items = []Node{}
	s.Lock.Unlock()
	return s
}

// IsEmpty returns true if the queue is empty
func (s *NodeQueue) IsEmpty() bool {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items) == 0
}

// Size returns the number of Nodes in the queue
func (s *NodeQueue) Size() int {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items)
}
