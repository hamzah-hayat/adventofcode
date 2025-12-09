package main

import (
	"bufio"
	"cmp"
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
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
		fmt.Println("Silver:" + PartOne("input", 1000))
		fmt.Println("Gold:" + PartTwo("input"))
	case "p1":
		fmt.Println("Silver:" + PartOne("input", 1000))
	case "p2":
		fmt.Println("Gold:" + PartTwo("input"))
	}
}

func PartOne(filename string, numberOfConnections int) string {
	input := readInput(filename)

	points := []Point{}
	distances := []DistanceBetweenPoints{}

	for _, lines := range input {
		splitPoint := strings.Split(lines, ",")
		xValue, _ := strconv.Atoi(splitPoint[0])
		yValue, _ := strconv.Atoi(splitPoint[1])
		zValue, _ := strconv.Atoi(splitPoint[2])
		point := Point{xValue, yValue, zValue}

		points = append(points, point)
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			distance := DistanceBetweenPoints{points[i], points[j], 0}
			xdistanceFloat := math.Abs(float64(points[i].x - points[j].x))
			ydistanceFloat := math.Abs(float64(points[i].y - points[j].y))
			zdistanceFloat := math.Abs(float64(points[i].z - points[j].z))
			distance.distance = math.Sqrt(math.Pow(xdistanceFloat, 2) + math.Pow(ydistanceFloat, 2) + math.Pow(zdistanceFloat, 2))
			distances = append(distances, distance)
		}
	}

	slices.SortFunc(distances,
		func(a, b DistanceBetweenPoints) int {
			return cmp.Compare(a.distance, b.distance)
		})

	// Now we have distances, we can start connecting "points" (aka junctions)
	circuits := []Circuit{}
	// Each junction starts as a circuit
	for _, point := range points {
		circuit := Circuit{points: []Point{point}}
		circuits = append(circuits, circuit)
	}

	// Now start merging circuits based on distances
	// We only want to connect the closest points, so we limit the number of connections
	for i, distance := range distances {
		if i == numberOfConnections {
			break
		}
		circuits = addDistanceToCurcuit(distance, circuits)
	}

	slices.SortFunc(circuits,
		func(a, b Circuit) int {
			return cmp.Compare(len(a.points), len(b.points))
		})

	biggestThreeCircuits := circuits[len(circuits)-3:]
	totalSize := 1
	for _, v := range biggestThreeCircuits {
		totalSize *= len(v.points)
	}

	return strconv.Itoa(totalSize)
}

func addDistanceToCurcuit(distance DistanceBetweenPoints, circuits []Circuit) []Circuit {

	firstCircuit := Circuit{}
	secondCircuit := Circuit{}
	firstCircuitIndex := -1
	secondCircuitIndex := -1
	for i, circuit := range circuits {
		if slices.Contains(circuit.points, distance.p1) && slices.Contains(circuit.points, distance.p2) {
			// Both points are already in the circuit, do nothing
			return circuits
		}
		if slices.Contains(circuit.points, distance.p1) {
			firstCircuitIndex = i
			firstCircuit = circuits[i]
		}
	}
	circuits = append(circuits[:firstCircuitIndex], circuits[firstCircuitIndex+1:]...)

	for i, circuit := range circuits {
		if slices.Contains(circuit.points, distance.p2) {
			secondCircuitIndex = i
			secondCircuit = circuits[i]
		}
	}

	circuits = append(circuits[:secondCircuitIndex], circuits[secondCircuitIndex+1:]...)

	newCircuit := Circuit{}
	newCircuit.points = append(firstCircuit.points, secondCircuit.points...)

	circuits = append(circuits, newCircuit)

	return circuits
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func PartTwo(filename string) string {
	input := readInput(filename)

	points := []Point{}
	distances := []DistanceBetweenPoints{}

	for _, lines := range input {
		splitPoint := strings.Split(lines, ",")
		xValue, _ := strconv.Atoi(splitPoint[0])
		yValue, _ := strconv.Atoi(splitPoint[1])
		zValue, _ := strconv.Atoi(splitPoint[2])
		point := Point{xValue, yValue, zValue}

		points = append(points, point)
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			distance := DistanceBetweenPoints{points[i], points[j], 0}
			xdistanceFloat := math.Abs(float64(points[i].x - points[j].x))
			ydistanceFloat := math.Abs(float64(points[i].y - points[j].y))
			zdistanceFloat := math.Abs(float64(points[i].z - points[j].z))
			distance.distance = math.Sqrt(math.Pow(xdistanceFloat, 2) + math.Pow(ydistanceFloat, 2) + math.Pow(zdistanceFloat, 2))
			distances = append(distances, distance)
		}
	}

	slices.SortFunc(distances,
		func(a, b DistanceBetweenPoints) int {
			return cmp.Compare(a.distance, b.distance)
		})

	// Now we have distances, we can start connecting "points" (aka junctions)
	circuits := []Circuit{}
	// Each junction starts as a circuit
	for _, point := range points {
		circuit := Circuit{points: []Point{point}}
		circuits = append(circuits, circuit)
	}

	// Now start merging circuits based on distances
	// Merge until one circuit left
	xValue := 0
	for _, distance := range distances {
		circuits = addDistanceToCurcuit(distance, circuits)
		if len(circuits) == 1 {
			xValue = distance.p1.x *distance.p2.x
			break
		}
	}

	return strconv.Itoa(xValue)
}

type Point struct {
	x int
	y int
	z int
}

type DistanceBetweenPoints struct {
	p1       Point
	p2       Point
	distance float64
}

type Circuit struct {
	points []Point
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
