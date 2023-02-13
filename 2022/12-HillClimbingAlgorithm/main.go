package main

import (
	"bufio"
	"flag"
	"fmt"
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

	start := findStart(input)
	nodeList := []Point{start}
	came_from := make(map[Point]Point)

	for len(nodeList) > 0 {
		currentNode := nodeList[0]
		nodeList = nodeList[1:]
		for _, node := range getNeighbours(currentNode, input) {
			_, hasNode := came_from[node]
			if !hasNode {
				nodeList = append(nodeList, node)
				came_from[node] = currentNode
			}
		}
	}

	currentNode := findGoal(input)
	path := []Point{}
	for currentNode != start {
		path = append(path, currentNode)
		currentNode = came_from[currentNode]
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	pathSize := strconv.Itoa(len(path))
	return pathSize
}

func findStart(input []string) Point {
	// Find S point in input
	y := 0
	for _, line := range input {
		x := strings.Index(line, "S")
		if x != -1 {
			return Point{x, y}
		}
		y++
	}
	return Point{-1, -1}
}

func findGoal(input []string) Point {
	// Find E point in input
	y := 0
	for _, line := range input {
		x := strings.Index(line, "E")
		if x != -1 {
			return Point{x, y}
		}
		y++
	}
	return Point{-1, -1}
}

// Get all neighbours of currentNode
func getNeighbours(currentNode Point, input []string) []Point {
	points := []Point{}
	// Up
	if currentNode.y != 0 {
		upNode := Point{currentNode.x, currentNode.y - 1}
		if canClimbTo(currentNode, upNode, input) {
			points = append(points, upNode)
		}
	}
	// Down
	if currentNode.y != len(input)-1 {
		downNode := Point{currentNode.x, currentNode.y + 1}
		if canClimbTo(currentNode, downNode, input) {
			points = append(points, downNode)
		}
	}
	// Left
	if currentNode.x != 0 {
		leftNode := Point{currentNode.x - 1, currentNode.y}
		if canClimbTo(currentNode, leftNode, input) {
			points = append(points, leftNode)
		}
	}
	// Right
	if currentNode.x != len(input[0])-1 {
		rightNode := Point{currentNode.x + 1, currentNode.y}
		if canClimbTo(currentNode, rightNode, input) {
			points = append(points, rightNode)
		}
	}

	return points
}

func PartTwo(filename string) string {
	input := readInput(filename)

	starts := findScenicStarts(input)
	smallestPath := 10000000000

	for _, start := range starts {
		nodeList := []Point{start}
		came_from := make(map[Point]Point)

		for len(nodeList) > 0 {
			currentNode := nodeList[0]
			nodeList = nodeList[1:]
			for _, node := range getNeighbours(currentNode, input) {
				_, hasNode := came_from[node]
				if !hasNode {
					nodeList = append(nodeList, node)
					came_from[node] = currentNode
				}
			}
		}

		currentNode := findGoal(input)
		path := []Point{}
		for currentNode != start {
			path = append(path, currentNode)
			currentNode = came_from[currentNode]
		}

		if len(path) < smallestPath {
			smallestPath = len(path)
		}
	}

	return strconv.Itoa(smallestPath)
}

func findScenicStarts(input []string) []Point {
	// Find S point in input
	points := []Point{}
	y := 0
	for _, line := range input {
		x := strings.Index(line, "a")
		if x != -1 {
			points = append(points, Point{x, y})
		}
		y++
	}
	return points
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

type Point struct {
	x, y int
}

// getNum returns num from string
func getNum(space string) int {
	// special case for S and E, S has elevation a, E has elevation z
	if space == "S" {
		space = "a"
	}
	if space == "E" {
		space = "z"
	}

	valueArray := "abcdefghijklmnopqrstuvwxyz"
	return strings.Index(valueArray, strings.ToLower(space)) + 1
}

func canClimbTo(start, end Point, input []string) bool {

	startClimbStr := string(input[start.y][start.x])
	endClimbStr := string(input[end.y][end.x])

	climbDif := getNum(startClimbStr) - getNum(endClimbStr)

	if climbDif == -1 || climbDif >= 0 {
		return true
	}

	return false
}
