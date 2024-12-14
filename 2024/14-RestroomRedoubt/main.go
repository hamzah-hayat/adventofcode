package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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
		fmt.Println("Silver:" + PartOne("input", 101, 103))
		fmt.Println("Gold:" + PartTwo("input", 101, 103))
	case "p1":
		fmt.Println("Silver:" + PartOne("input", 101, 103))
	case "p2":
		fmt.Println("Gold:" + PartTwo("input", 101, 103))
	}
}

func PartOne(filename string, gridMaxX, gridMaxY int) string {
	input := readInput(filename)

	robotRegex := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	robots := make([]Robot, 0)

	// Read Robots
	for _, robot := range input {
		robotStrings := robotRegex.FindAllStringSubmatch(robot, -1)
		posX, _ := strconv.Atoi(robotStrings[0][1])
		posY, _ := strconv.Atoi(robotStrings[0][2])
		dirX, _ := strconv.Atoi(robotStrings[0][3])
		dirY, _ := strconv.Atoi(robotStrings[0][4])
		newRobot := Robot{Vector{posX, posY}, Vector{dirX, dirY}}

		robots = append(robots, newRobot)
	}

	// Now we run simulation 100 times

	//fmt.Println(PrintGrid(robots, gridMaxX, gridMaxY))
	for i := 0; i < 100; i++ {
		SimulateRobots(robots, gridMaxX, gridMaxY)
		//fmt.Println(PrintGrid(robots, gridMaxX, gridMaxY))
	}
	//fmt.Println(PrintGrid(robots, gridMaxX, gridMaxY))

	// Count robots in each quadrant
	NWQuad := 0
	NEQuad := 0
	SEQuad := 0
	SWQuad := 0

	halfX := gridMaxX / 2
	halfY := gridMaxY / 2

	for _, robot := range robots {
		if robot.position.x < halfX && robot.position.y < halfY {
			NWQuad++
		}
		if robot.position.x > halfX && robot.position.y < halfY {
			NEQuad++
		}
		if robot.position.x > halfX && robot.position.y > halfY {
			SEQuad++
		}
		if robot.position.x < halfX && robot.position.y > halfY {
			SWQuad++
		}
	}

	total := NEQuad * SEQuad * NWQuad * SWQuad

	return strconv.Itoa(total)
}

func SimulateRobots(robots []Robot, gridMaxX, gridMaxY int) {
	for i := 0; i < len(robots); i++ {
		// Change position of each Robot
		// Check if we leave map in each direction and overflow
		movedX := false
		movedY := false

		// North
		if robots[i].position.y+robots[i].direction.y < 0 {
			movedY = true
			robots[i].position.y = gridMaxY + robots[i].position.y + robots[i].direction.y
		}

		// East
		if robots[i].position.x+robots[i].direction.x >= gridMaxX {
			movedX = true
			robots[i].position.x = ((robots[i].position.x + robots[i].direction.x) % gridMaxX)
		}

		// South
		if robots[i].position.y+robots[i].direction.y >= gridMaxY {
			movedY = true
			robots[i].position.y = ((robots[i].position.y + robots[i].direction.y) % gridMaxY)
		}

		// West
		if robots[i].position.x+robots[i].direction.x < 0 {
			movedX = true
			robots[i].position.x = gridMaxX + robots[i].position.x + robots[i].direction.x
		}

		// Otherwise just move
		if !movedX {
			robots[i].position.x += robots[i].direction.x
		}

		if !movedY {
			robots[i].position.y += robots[i].direction.y
		}
	}
}

func PrintGrid(robots []Robot, gridMaxX, gridMaxY int) string {
	// Print grid
	gridPrint := ""
	for y := 0; y < gridMaxY; y++ {
		for x := 0; x < gridMaxX; x++ {
			numRobotsHere := 0
			for _, r := range robots {
				if r.position.x == x && r.position.y == y {
					numRobotsHere++
				}
			}

			if numRobotsHere > 0 {
				gridPrint += "⬜"
			} else {
				gridPrint += "⬛"
			}

		}
		gridPrint += "\n"
	}
	return gridPrint
}

type Vector struct {
	x int
	y int
}

type Robot struct {
	position  Vector
	direction Vector
}

// Lol, wut?
// Searching for a christmas tree, done via inspection
func PartTwo(filename string, gridMaxX, gridMaxY int) string {
	input := readInput(filename)

	robotRegex := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	robots := make([]Robot, 0)

	// Read Robots
	for _, robot := range input {
		robotStrings := robotRegex.FindAllStringSubmatch(robot, -1)
		posX, _ := strconv.Atoi(robotStrings[0][1])
		posY, _ := strconv.Atoi(robotStrings[0][2])
		dirX, _ := strconv.Atoi(robotStrings[0][3])
		dirY, _ := strconv.Atoi(robotStrings[0][4])
		newRobot := Robot{Vector{posX, posY}, Vector{dirX, dirY}}

		robots = append(robots, newRobot)
	}

	// Now we run simulation 100 times

	//fmt.Println(PrintGrid(robots, gridMaxX, gridMaxY))
	for i := 0; i < 10000; i++ {
		SimulateRobots(robots, gridMaxX, gridMaxY)
		//fmt.Println("Current second", i+1)
		if strings.Contains(PrintGrid(robots, gridMaxX, gridMaxY), "⬜⬜⬜⬜⬜⬜⬜⬜") {
			fmt.Println("Current second", i+1)
			fmt.Println(PrintGrid(robots, gridMaxX, gridMaxY))
			break
		}
	}
	//fmt.Println(PrintGrid(robots, gridMaxX, gridMaxY))

	// Answer is done via inspection I guess
	return strconv.Itoa(0)
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
