package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p1", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()

	switch *methodP {
	case "p1":
		partOne()
		break
	case "p2":
		partTwo()
		break
	case "test":
		break
	}
}

func partOne() {
	input := readInput()

	maze := createMaze(input)

	//fmt.Println(printMaze(maze))

	lowestSteps := solveMaze(maze)

	fmt.Println("The lowest number of steps to solve the maze is", lowestSteps)

}

func createMaze(input []string) map[space]mazeObject {
	maze := make(map[space]mazeObject)

	x := 0
	y := 0
	for _, mazeLine := range input {
		for _, c := range mazeLine {
			obj := makeMazeObject(c)
			maze[space{x: x, y: y}] = obj
			x++
		}
		x = 0
		y++
	}

	return maze
}

func makeMazeObject(character rune) mazeObject {
	switch character {
	case '.':
		return mazeObject{char: ".", objectType: 0}
	case '#':
		return mazeObject{char: "#", objectType: 1}
	case '@':
		return mazeObject{char: "@", objectType: 2}
	}

	// This is a door/key
	if unicode.IsLower(character) {
		// is key
		return mazeObject{char: string(character), objectType: 3}
	}

	// is door
	return mazeObject{char: string(character), objectType: 4}
}

var solveCalls int

func solveMaze(maze map[space]mazeObject) int {

	// Find the entrance
	var entranceSpace space
	for i, obj := range maze {
		if obj.objectType == 2 {
			entranceSpace = i
		}
	}
	solveCalls = 0

	return findShortestPathToGetAllKeys(maze, entranceSpace, 0)
}

// Find all keys
func findShortestPathToGetAllKeys(maze map[space]mazeObject, currentSpace space, pathlength int) int {
	solveCalls++

	// Copy the Maze
	newMaze := make(map[space]mazeObject)
	for k, v := range maze {
		newMaze[k] = v
	}

	keys := keysInMaze(newMaze)
	// If we are now on a key, get rid of it and plus our pathLength
	for i, k := range keys {
		if currentSpace == i {
			doorSpace := findDoorForKey(newMaze, k)
			delete(newMaze, doorSpace) // delete the door
			delete(newMaze, i)         // delete the key
			// Add an empty space at each location
			newMaze[doorSpace] = mazeObject{char: ".", objectType: 0}
			newMaze[i] = mazeObject{char: ".", objectType: 0}
		}
	}

	// If we get all keys return
	keys = keysInMaze(newMaze)
	if len(keys) == 0 {
		return pathlength
	}

	// See what keys we can reach from here
	reachableKeys := make(map[space]mazeObject)
	for i, k := range keys {
		length := findShortestPath(newMaze, currentSpace, i)
		if length > 0 {
			reachableKeys[i] = k
		}
	}

	// Now that we know what keys we can get to, try going to each one
	lowestPath := 100000000000000
	for i := range reachableKeys {
		length := findShortestPathToGetAllKeys(newMaze, i, pathlength+findShortestPath(newMaze, currentSpace, i))
		if length < lowestPath {
			lowestPath = length
		}
	}

	return lowestPath
}

// Find the shortest path between space a and b
func findShortestPath(maze map[space]mazeObject, a space, b space) int {

	correctPath := make(map[space]bool)
	wasHere := make(map[space]bool)

	found, correctPath := recursiveSolve(maze, correctPath, wasHere, a, b, -100, 100, -100, 100)

	if found {
		return len(correctPath)
	}

	return -1
}

func recursiveSolve(mapTiles map[space]mazeObject, correctPath map[space]bool, wasHere map[space]bool, currentTile space, endTile space, xMin, xMax, yMin, yMax int) (bool, map[space]bool) {

	if currentTile.x == endTile.x && currentTile.y == endTile.y {
		return true, correctPath
	}
	if mapTiles[currentTile].objectType == 1 || mapTiles[currentTile].objectType == 4 || wasHere[currentTile] {
		return false, nil
	}

	wasHere[currentTile] = true

	if currentTile.x != xMin {
		newTile := currentTile
		newTile.x--
		found, correctPath := recursiveSolve(mapTiles, correctPath, wasHere, newTile, endTile, xMin, xMax, yMin, yMax)
		if found {
			correctPath[currentTile] = true
			return true, correctPath
		}
	}

	if currentTile.x != xMax-1 {
		newTile := currentTile
		newTile.x++
		found, correctPath := recursiveSolve(mapTiles, correctPath, wasHere, newTile, endTile, xMin, xMax, yMin, yMax)
		if found {
			correctPath[currentTile] = true
			return true, correctPath
		}
	}

	if currentTile.x != yMin {
		newTile := currentTile
		newTile.y++
		found, correctPath := recursiveSolve(mapTiles, correctPath, wasHere, newTile, endTile, xMin, xMax, yMin, yMax)
		if found {
			correctPath[currentTile] = true
			return true, correctPath
		}
	}

	if currentTile.x != yMax-1 {
		newTile := currentTile
		newTile.y--
		found, correctPath := recursiveSolve(mapTiles, correctPath, wasHere, newTile, endTile, xMin, xMax, yMin, yMax)
		if found {
			correctPath[currentTile] = true
			return true, correctPath
		}
	}
	return false, nil
}

func keysInMaze(maze map[space]mazeObject) map[space]mazeObject {
	keys := make(map[space]mazeObject)
	for i, obj := range maze {
		if obj.objectType == 3 {
			keys[i] = obj
		}

	}
	return keys
}

func findDoorForKey(maze map[space]mazeObject, key mazeObject) space {
	r := []rune(key.char)
	doorStr := unicode.ToUpper(r[0])

	var doorSpace space
	for i, d := range maze {
		if d.char == string(doorStr) {
			doorSpace = i
		}
	}
	return doorSpace
}

func partTwo() {
	//input := readInput()
}

func printMaze(maze map[space]mazeObject) string {
	mazeStr := ""

	// Find max x and y
	maxX := 0
	maxY := 0
	for s := range maze {
		if s.x > maxX {
			maxX = s.x + 1
		}
		if s.y > maxY {
			maxY = s.y + 1
		}
	}

	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			mazeStr += maze[space{x: i, y: j}].char
		}
		mazeStr += "\n"
	}
	return mazeStr
}

type space struct {
	x, y int
}

type mazeObject struct {
	char       string // the printable character (or key/door name)
	objectType int    // 0 is empty space, 1 is wall, 2 is entrance, 3 is key, 4 is door
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
