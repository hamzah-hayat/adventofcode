package main

import (
	"fmt"
	"testing"
)

func TestNBodyCheck_SamePositon(t *testing.T) {

	input := []string{
		"<x=-1, y=0, z=2>",
		"<x=2, y=-10, z=-7>",
		"<x=4, y=-8, z=8>",
		"<x=3, y=5, z=-1>",
	}
	expected := 2772

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
	result := LCM(foundXCounter, foundYCounter, foundZCounter)

	if result != expected {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}

func TestNBodyCheck_SamePositon_2(t *testing.T) {

	input := []string{
		"<x=-8, y=-10, z=0>",
		"<x=5, y=5, z=10>",
		"<x=2, y=-7, z=3>",
		"<x=9, y=-8, z=-3>",
	}
	expected := 4686774924

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
	result := LCM(foundXCounter, foundYCounter, foundZCounter)

	if result != expected {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", result))
	}

}
