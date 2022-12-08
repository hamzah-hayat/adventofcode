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

	maxX := len(input[0])
	maxY := len(input)

	treeGrid := make(map[int]map[int]TreeSpace)
	for x := 0; x < maxX; x++ {
		treeGrid[x] = make(map[int]TreeSpace)
	}

	// Form a 2D grid with bool for visible
	y := 0
	for _, line := range input {
		for x := 0; x < len(line); x++ {
			heightInt, _ := strconv.Atoi(string(line[x]))
			treeGrid[x][y] = TreeSpace{Visible: false, Height: heightInt}
		}
		y++
	}

	// now mark outside edges as visible
	// 1. top row
	// 2. bottom row
	// 3. left side
	// 4. right side
	for x := 0; x < maxX; x++ {
		gridSpace := TreeSpace{Visible: true, Height: treeGrid[x][0].Height}
		treeGrid[x][0] = gridSpace
	}
	for x := 0; x < maxX; x++ {
		gridSpace := TreeSpace{Visible: true, Height: treeGrid[x][maxY-1].Height}
		treeGrid[x][maxY-1] = gridSpace
	}
	for y := 0; y < maxY; y++ {
		gridSpace := TreeSpace{Visible: true, Height: treeGrid[0][y].Height}
		treeGrid[0][y] = gridSpace
	}
	for y := 0; y < maxY; y++ {
		gridSpace := TreeSpace{Visible: true, Height: treeGrid[maxX-1][y].Height}
		treeGrid[maxX-1][y] = gridSpace
	}

	// fmt.Println(printGrid(treeGrid, maxX, maxY))

	// now check each row and column from each side
	// 1. top to bottom
	// 2. bottom to top
	// 3. left to right
	// 4. right to left
	CheckGridVisibleFromEdge(treeGrid, 0, maxX, maxY)
	CheckGridVisibleFromEdge(treeGrid, 1, maxX, maxY)
	CheckGridVisibleFromEdge(treeGrid, 2, maxX, maxY)
	CheckGridVisibleFromEdge(treeGrid, 3, maxX, maxY)

	// return number of trees visible
	treesVisible := countVisible(treeGrid, maxX, maxY)
	return strconv.Itoa(treesVisible)
}

func PartTwo(filename string) string {
	input := readInput(filename)

	maxX := len(input[0])
	maxY := len(input)

	treeGrid := make(map[int]map[int]TreeSpace)
	for x := 0; x < maxX; x++ {
		treeGrid[x] = make(map[int]TreeSpace)
	}

	// Form a 2D grid with bool for visible
	y := 0
	for _, line := range input {
		for x := 0; x < len(line); x++ {
			heightInt, _ := strconv.Atoi(string(line[x]))
			treeGrid[x][y] = TreeSpace{Visible: false, Height: heightInt}
		}
		y++
	}

	// fmt.Println(printGrid(treeGrid, maxX, maxY))

	// now check each space, looking for best "scenary"
	bestScenary := 0
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			downScore := CheckVisibleTreesFromTree(treeGrid, x, y, 0, maxX, maxY)
			upScore := CheckVisibleTreesFromTree(treeGrid, x, y, 1, maxX, maxY)
			rightScore := CheckVisibleTreesFromTree(treeGrid, x, y, 2, maxX, maxY)
			leftScore := CheckVisibleTreesFromTree(treeGrid, x, y, 3, maxX, maxY)
			score := upScore * downScore * rightScore * leftScore
			if score > bestScenary {
				bestScenary = score
			}
		}
	}

	return strconv.Itoa(bestScenary)
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

type TreeSpace struct {
	Visible bool
	Height  int
}

type Vector struct {
	x, y int
}

// now check each row and column from each side
// 1. top to bottom
// 2. bottom to top
// 3. left to right
// 4. right to left
func CheckGridVisibleFromEdge(grid map[int]map[int]TreeSpace, direction, maxX, maxY int) {

	switch direction {
	case 0:
		for x := 0; x < maxX; x++ {
			currentMaxHeight := 0
			for y := 0; y < maxY; y++ {
				if grid[x][y].Height > currentMaxHeight {
					currentMaxHeight = grid[x][y].Height
					gridSpace := TreeSpace{Visible: true, Height: grid[x][y].Height}
					grid[x][y] = gridSpace
				}
			}
		}
		break
	case 1:
		for x := 0; x < maxX; x++ {
			currentMaxHeight := 0
			for y := maxY; y > 0; y-- {
				if grid[x][y].Height > currentMaxHeight {
					currentMaxHeight = grid[x][y].Height
					gridSpace := TreeSpace{Visible: true, Height: grid[x][y].Height}
					grid[x][y] = gridSpace
				}
			}
		}
		break
	case 2:
		for y := 0; y < maxY; y++ {
			currentMaxHeight := 0
			for x := 0; x < maxX; x++ {
				if grid[x][y].Height > currentMaxHeight {
					currentMaxHeight = grid[x][y].Height
					gridSpace := TreeSpace{Visible: true, Height: grid[x][y].Height}
					grid[x][y] = gridSpace
				}
			}
		}
		break
	case 3:
		for y := 0; y < maxY; y++ {
			currentMaxHeight := 0
			for x := maxX; x > 0; x-- {
				if grid[x][y].Height > currentMaxHeight {
					currentMaxHeight = grid[x][y].Height
					gridSpace := TreeSpace{Visible: true, Height: grid[x][y].Height}
					grid[x][y] = gridSpace
				}
			}
		}
		break
	}
}

// from our starting tree, check up,down,right,left
// 1. Up
// 2. down
// 3. right
// 4. left
func CheckVisibleTreesFromTree(grid map[int]map[int]TreeSpace, startX, startY, direction, maxX, maxY int) int {
	numberOfTreesVisible := 0
	switch direction {
	case 0:
		currentMaxHeight := grid[startX][startY].Height
		for y := startY; y < maxY; y++ {
			if y == startY {
				continue
			}
			if grid[startX][y].Height < currentMaxHeight {
				numberOfTreesVisible++
			} else {
				numberOfTreesVisible++
				break
			}
		}
		break
	case 1:
		currentMaxHeight := grid[startX][startY].Height
		for y := startY; y >= 0; y-- {
			if y == startY {
				continue
			}
			if grid[startX][y].Height < currentMaxHeight {
				numberOfTreesVisible++
			} else {
				numberOfTreesVisible++
				break
			}
		}
		break
	case 2:
		currentMaxHeight := grid[startX][startY].Height
		for x := startX; x < maxX; x++ {
			if x == startX {
				continue
			}
			if grid[x][startY].Height < currentMaxHeight {
				numberOfTreesVisible++
			} else {
				numberOfTreesVisible++
				break
			}
		}
		break
	case 3:
		currentMaxHeight := grid[startX][startY].Height
		for x := startX; x >= 0; x-- {
			if x == startX {
				continue
			}
			if grid[x][startY].Height < currentMaxHeight {
				numberOfTreesVisible++
			} else {
				numberOfTreesVisible++
				break
			}
		}
		break
	}
	return numberOfTreesVisible
}

func countVisible(grid map[int]map[int]TreeSpace, maxX, maxY int) int {
	totalVisible := 0
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if grid[x][y].Visible {
				totalVisible++
			}
		}
	}
	return totalVisible
}

func printGrid(grid map[int]map[int]TreeSpace, maxX, maxY int) string {
	print := ""
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			print += strconv.Itoa(grid[x][y].Height)
		}
		print += "\n"
	}
	return print
}
