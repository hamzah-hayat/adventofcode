package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/hamzah-hayat/adventofcode/intcode"
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
	program := convertToInts(input)

	// First channel is for input
	inputChan := make(chan int)
	// Second channel is for output
	outputChan := make(chan int)
	// Terimnation channel
	t := make(chan bool)
	// Message channel
	message := make(chan intcode.Message)

	go intcode.RunIntCodeProgramWaitForTermination(program, inputChan, outputChan, t, message)

	mapTiles, oxygenSystemTile := buildMap(inputChan, outputChan, t, message)

	distanceFromStartToOxygenSystem := findShortestPath(mapTiles, oxygenSystemTile, tile{x: 0, y: 0})

	fmt.Println("The number of moves to the oxygen system is", distanceFromStartToOxygenSystem)
}

func partTwo() {
	input := readInput()
	program := convertToInts(input)

	// First channel is for input
	inputChan := make(chan int)
	// Second channel is for output
	outputChan := make(chan int)
	// Terimnation channel
	t := make(chan bool)
	// Message channel
	message := make(chan intcode.Message)

	go intcode.RunIntCodeProgramWaitForTermination(program, inputChan, outputChan, t, message)

	mapTiles, oxygenSystemTile := buildMapRandom(inputChan, outputChan, t, message)

	// Get distance from oyxgen Tile to every other open tile
	// take highest amount
	highest := 0
	for tile, val := range mapTiles {
		if val != 1 {
			distanceFromOxygenSystem := findShortestPath(mapTiles, oxygenSystemTile, tile)
			if highest < distanceFromOxygenSystem {
				highest = distanceFromOxygenSystem
			}
		}
	}

	fmt.Println("The time it will take to fill all tiles is", highest)

}

func buildMap(inputChan chan int, outputChan chan int, t chan bool, message chan intcode.Message) (map[tile]int, tile) {
	found := false

	mapTiles := make(map[tile]int)

	currentX := 0
	currentY := 0
	mapTiles[tile{x: currentX, y: currentY}] = 2
	var oxygenSystemTile tile

	for !found {

		var possibleDirections []int

		// check directions
		// north (1), south (2), west (3), and east (4)
		if (mapTiles[tile{x: currentX, y: currentY - 1}] != 1) {
			possibleDirections = append(possibleDirections, 1)
		}

		if (mapTiles[tile{x: currentX, y: currentY + 1}] != 1) {
			possibleDirections = append(possibleDirections, 2)
		}

		if (mapTiles[tile{x: currentX - 1, y: currentY}] != 1) {
			possibleDirections = append(possibleDirections, 3)
		}

		if (mapTiles[tile{x: currentX + 1, y: currentY}] != 1) {
			possibleDirections = append(possibleDirections, 4)
		}

		// Now move randomly

		max := len(possibleDirections)
		randomDirection := rand.Intn(max)

		<-message

		inputChan <- possibleDirections[randomDirection]

		<-message

		// Read the output
		output := <-outputChan

		if output == 0 {
			setTile(mapTiles, output, possibleDirections[randomDirection], &currentX, &currentY)
		} else if output == 1 {
			setTile(mapTiles, output, possibleDirections[randomDirection], &currentX, &currentY)
		} else if output == 2 {
			setTile(mapTiles, output, possibleDirections[randomDirection], &currentX, &currentY)
			//fmt.Println("Found!")
			oxygenSystemTile = tile{x: currentX, y: currentY}
			found = true
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Println(printTiles(mapTiles, currentX, currentY))
			time.Sleep(time.Millisecond)
		} else {
			fmt.Println("Error")
			found = true
		}

	}

	return mapTiles, oxygenSystemTile
}

func buildMapRandom(inputChan chan int, outputChan chan int, t chan bool, message chan intcode.Message) (map[tile]int, tile) {

	mapTiles := make(map[tile]int)

	currentX := 0
	currentY := 0
	mapTiles[tile{x: currentX, y: currentY}] = 2
	var oxygenSystemTile tile

	for i := 0; i < 1000000; i++ {

		var possibleDirections []int

		// check directions
		// north (1), south (2), west (3), and east (4)
		if (mapTiles[tile{x: currentX, y: currentY - 1}] != 1) {
			possibleDirections = append(possibleDirections, 1)
		}

		if (mapTiles[tile{x: currentX, y: currentY + 1}] != 1) {
			possibleDirections = append(possibleDirections, 2)
		}

		if (mapTiles[tile{x: currentX - 1, y: currentY}] != 1) {
			possibleDirections = append(possibleDirections, 3)
		}

		if (mapTiles[tile{x: currentX + 1, y: currentY}] != 1) {
			possibleDirections = append(possibleDirections, 4)
		}

		// Now move randomly

		max := len(possibleDirections)
		randomDirection := rand.Intn(max)

		<-message

		inputChan <- possibleDirections[randomDirection]

		<-message

		// Read the output
		output := <-outputChan

		if output == 0 {
			setTile(mapTiles, output, possibleDirections[randomDirection], &currentX, &currentY)
		} else if output == 1 {
			setTile(mapTiles, output, possibleDirections[randomDirection], &currentX, &currentY)
		} else if output == 2 {
			setTile(mapTiles, output, possibleDirections[randomDirection], &currentX, &currentY)
			//fmt.Println("Found!")
			oxygenSystemTile = tile{x: currentX, y: currentY}
		} else {
			fmt.Println("Error")
		}

	}

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println(printTiles(mapTiles, currentX, currentY))
	time.Sleep(time.Millisecond)

	return mapTiles, oxygenSystemTile
}

// Find the shortest path between a and b using mapTiles
func findShortestPath(mapTiles map[tile]int, a tile, b tile) int {

	correctPath := make(map[tile]bool)
	wasHere := make(map[tile]bool)

	found, correctPath := recursiveSolve(mapTiles, correctPath, wasHere, a, b, -25, 25, -25, 25)

	if found {
		return len(correctPath)
	}

	return -1
}

func recursiveSolve(mapTiles map[tile]int, correctPath map[tile]bool, wasHere map[tile]bool, currentTile tile, endTile tile, xMin, xMax, yMin, yMax int) (bool, map[tile]bool) {

	if currentTile.x == endTile.x && currentTile.y == endTile.y {
		return true, correctPath
	}
	if mapTiles[currentTile] == 1 || mapTiles[currentTile] == 0 || wasHere[currentTile] {
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

	// if (x == endX && y == endY) return true; // If you reached the end
	// if (maze[x][y] == 2 || wasHere[x][y]) return false;
	// // If you are on a wall or already were here
	// wasHere[x][y] = true;
	// if (x != 0) // Checks if not on left edge
	//     if (recursiveSolve(x-1, y)) { // Recalls method one to the left
	//         correctPath[x][y] = true; // Sets that path value to true;
	//         return true;
	//     }
	// if (x != width - 1) // Checks if not on right edge
	//     if (recursiveSolve(x+1, y)) { // Recalls method one to the right
	//         correctPath[x][y] = true;
	//         return true;
	//     }
	// if (y != 0)  // Checks if not on top edge
	//     if (recursiveSolve(x, y-1)) { // Recalls method one up
	//         correctPath[x][y] = true;
	//         return true;
	//     }
	// if (y != height - 1) // Checks if not on bottom edge
	//     if (recursiveSolve(x, y+1)) { // Recalls method one down
	//         correctPath[x][y] = true;
	//         return true;
	//     }
	// return false;
	return false, nil
}

func setTile(mapTiles map[tile]int, output int, direction int, currentX *int, currentY *int) {
	switch direction {
	case 1:
		mapTiles[tile{x: *currentX, y: *currentY - 1}] = output + 1
		break
	case 2:
		mapTiles[tile{x: *currentX, y: *currentY + 1}] = output + 1
		break
	case 3:
		mapTiles[tile{x: *currentX - 1, y: *currentY}] = output + 1
		break
	case 4:
		mapTiles[tile{x: *currentX + 1, y: *currentY}] = output + 1
		break
	}

	if output != 0 {
		switch direction {
		case 1:
			*currentY--
			break
		case 2:
			*currentY++
			break
		case 3:
			*currentX--
			break
		case 4:
			*currentX++
			break
		}
	}
}

// A tile can be either
//	0: The repair droid hit a wall. Its position has not changed.
//	1: The repair droid has moved one step in the requested direction.
//	2: The repair droid has moved one step in the requested direction; its new position is the location of the oxygen system.
// 1 is wall, 2 is empty space, 3 is oxygen system

func printTiles(mapTiles map[tile]int, currentX int, currentY int) string {
	// Find the xMin,xMax and ymin,yMax
	// Hardcode these for now
	xMin := -25
	xMax := 25
	yMin := -25
	yMax := 25

	// now print the grid

	picture := ""
	for i := xMin; i < xMax; i++ {
		for j := yMin; j < yMax; j++ {
			if i == currentX && j == currentY {
				picture += "B"
				continue
			}
			if i == 0 && j == 0 {
				picture += "S"
				continue
			}
			if (mapTiles[tile{x: i, y: j}] == 1) {
				picture += "O"
			} else if (mapTiles[tile{x: i, y: j}] == 2) {
				picture += " "
			} else if (mapTiles[tile{x: i, y: j}] == 3) {
				picture += "X"
			} else if (mapTiles[tile{x: i, y: j}] == 0) {
				picture += "-"
			}
		}
		picture += "\n"
	}
	return picture
}

type tile struct {
	x, y int
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

func convertToInts(input []string) []int {
	programStr := strings.Split(input[0], ",")
	var program []int

	for _, s := range programStr {
		i, _ := strconv.Atoi(s)
		program = append(program, i)
	}
	return program
}
