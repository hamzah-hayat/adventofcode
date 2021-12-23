package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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

	cubeMap := make(map[Point3D]bool)

	cubeMap = turnOnCubes(input, cubeMap)

	on := 0
	for _, c := range cubeMap {
		if c {
			on++
		}
	}

	num := strconv.Itoa(on)

	return num
}

func PartTwo(filename string) string {
	input := readInput(filename)

	cubeMap := make(map[Rectangle3D]bool)

	cubeMap = turnOnCubesRectangle(input, cubeMap)

	num := strconv.Itoa(0)

	return num
}

func turnOnCubes(input []string, cubemap map[Point3D]bool) map[Point3D]bool {

	for _, line := range input {
		//on x=-20..26,y=-36..17,z=-47..7
		lineRegex := regexp.MustCompile(`(on|off) x=(-*\d+)..(-*\d+),y=(-*\d+)..(-*\d+),z=(-*\d+)..(-*\d+)`)
		match := lineRegex.FindAllStringSubmatch(line, 1)

		// figure out vars
		var turnOnCube bool
		if match[0][1] == "on" {
			turnOnCube = true
		} else {
			turnOnCube = false
		}

		xMin, _ := strconv.Atoi(match[0][2])
		xMax, _ := strconv.Atoi(match[0][3])
		yMin, _ := strconv.Atoi(match[0][4])
		yMax, _ := strconv.Atoi(match[0][5])
		zMin, _ := strconv.Atoi(match[0][6])
		zMax, _ := strconv.Atoi(match[0][7])

		if xMin < -50 && xMax < -50 || (xMin > 50 && xMax > 50) {
			continue
		}
		if yMin < -50 && yMax < -50 || (yMin > 50 && yMax > 50) {
			continue
		}
		if (zMin < -50 && zMax < -50) || (zMin > 50 && zMax > 50) {
			continue
		}

		for x := xMin; x < xMax+1; x++ {
			for y := yMin; y < yMax+1; y++ {
				for z := zMin; z < zMax+1; z++ {
					cubemap[Point3D{X: x, Y: y, Z: z}] = turnOnCube
				}
			}
		}

	}

	return cubemap
}

func turnOnCubesRectangle(input []string, cubemap map[Rectangle3D]bool) map[Rectangle3D]bool {

	for _, line := range input {
		//on x=-20..26,y=-36..17,z=-47..7
		lineRegex := regexp.MustCompile(`(on|off) x=(-*\d+)..(-*\d+),y=(-*\d+)..(-*\d+),z=(-*\d+)..(-*\d+)`)
		match := lineRegex.FindAllStringSubmatch(line, 1)

		// figure out vars
		var turnOnCube bool
		if match[0][1] == "on" {
			turnOnCube = true
		} else {
			turnOnCube = false
		}

		xMin, _ := strconv.Atoi(match[0][2])
		xMax, _ := strconv.Atoi(match[0][3])
		yMin, _ := strconv.Atoi(match[0][4])
		yMax, _ := strconv.Atoi(match[0][5])
		zMin, _ := strconv.Atoi(match[0][6])
		zMax, _ := strconv.Atoi(match[0][7])

		newCube := Rectangle3D{xMin: xMin, xMax: xMax, yMin: yMin, yMax: yMax, zMin: zMin, zMax: zMax}

		//cubemap[Rectangle3D{xMin: xMin, xMax: xMax, yMin: yMin, yMax: yMax, zMin: zMin, zMax: zMax}] = turnOnCube

		cubemap = intersectCubes(newCube, turnOnCube, cubemap)
	}

	return cubemap
}

func intersectCubes(newCube Rectangle3D, turnOnCube bool, cubemap map[Rectangle3D]bool) map[Rectangle3D]bool {

	for _, cube := range cubemap {
		if !turnOnCube && cube {
			// we have opposites so build intersection cube and cancel out
		}
	}


	return cubemap
}

type Point3D struct {
	X, Y, Z int
}

type Rectangle3D struct {
	xMin, xMax, yMin, yMax, zMin, zMax int
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
