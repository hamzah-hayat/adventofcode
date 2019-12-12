package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p2", "The method/part that should be run, valid are p1,p2 and test")
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

	moons, velos := makeMoons(input)

	// Run 1000 times
	for i := 0; i < 1000; i++ {
		moons, velos = moveMoons(moons, velos)
	}

	// Now print total energy
	fmt.Println("The total energy in the system is", totalEnergy(moons, velos))
}

func partTwo() {
	input := readInput()

	moons, velos := makeMoons(input)

	xTargets := []int{}
	yTargets := []int{}
	zTargets := []int{}

	for _, moon := range moons {
		xTargets = append(xTargets, moon.x)
		yTargets = append(yTargets, moon.y)
		zTargets = append(zTargets, moon.z)
	}

	// Run until the x,y,z of each planet aligns
	foundX := false
	foundXCounter := 0
	foundY := false
	foundYCounter := 0
	foundZ := false
	foundZCounter := 0
	counter := 1
	for {
		counter++
		moons, velos = moveMoons(moons, velos)

		if !foundX {
			match := true
			for i, moon := range moons {
				if moon.x != xTargets[i] {
					match = false
				}
			}
			if match {
				foundX = true
				foundXCounter = counter
			}
		}

		if !foundY {
			match := true
			for i, moon := range moons {
				if moon.y != yTargets[i] {
					match = false
				}
			}
			if match {
				foundY = true
				foundYCounter = counter
			}
		}

		if !foundZ {
			match := true
			for i, moon := range moons {
				if moon.z != zTargets[i] {
					match = false
				}
			}
			if match {
				foundZ = true
				foundZCounter = counter
			}
		}

		if foundX && foundY && foundZ {
			break
		}
	}

	fmt.Println("X,Y and Z were matching at ", foundXCounter, ",", foundYCounter, ",", foundZCounter, " respectively")
	fmt.Println(LCM(foundXCounter, foundYCounter, foundZCounter))

}

func moveMoons(moons []space, velos []velocity) ([]space, []velocity) {
	newMoons := []space{}
	newVelos := []velocity{}

	// For each moon, compute the velocity gain by pairing it with the other moons
	for i := 0; i < len(moons); i++ {

		veloChange := velocity{x: velos[i].x, y: velos[i].y, z: velos[i].z}

		for j, moon := range moons {
			if j == i {
				continue
			}
			// otherwise calc veloChange
			// for X,Y,Z
			if moons[i].x < moon.x {
				veloChange.x++
			} else if moons[i].x > moon.x {
				veloChange.x--
			}

			if moons[i].y < moon.y {
				veloChange.y++
			} else if moons[i].y > moon.y {
				veloChange.y--
			}

			if moons[i].z < moon.z {
				veloChange.z++
			} else if moons[i].z > moon.z {
				veloChange.z--
			}
		}

		// Now calc the position change
		posChange := space{x: moons[i].x + veloChange.x, y: moons[i].y + veloChange.y, z: moons[i].z + veloChange.z}

		newMoons = append(newMoons, posChange)
		newVelos = append(newVelos, veloChange)
	}

	return newMoons, newVelos
}

func totalEnergy(moons []space, velos []velocity) int {
	energy := 0

	// Potential energy is sum of absoloute values of all moons
	for i, moonSpace := range moons {
		kinetic := Abs(velos[i].x) + Abs(velos[i].y) + Abs(velos[i].z)
		potential := Abs(moonSpace.x) + Abs(moonSpace.y) + Abs(moonSpace.z)
		energy += potential * kinetic
	}

	return energy
}

func makeMoons(input []string) ([]space, []velocity) {
	moons := []space{}
	velos := []velocity{}

	for _, moon := range input {

		values := strings.Split(moon, ",")

		x, _ := strconv.Atoi(strings.Trim(values[0], "<x="))
		y, _ := strconv.Atoi(strings.Trim(values[1], " y="))
		z, _ := strconv.Atoi(strings.TrimRight(strings.Trim(values[2], " z="), ">"))

		moons = append(moons, space{x: x, y: y, z: z})
		velos = append(velos, velocity{x: 0, y: 0, z: 0})
	}
	return moons, velos
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// GCD greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

type space struct {
	x, y, z int
}

type velocity struct {
	x, y, z int
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
