package main

import (
	"fmt"
	"testing"
)

func TestMonitoringStationSmallCheck(t *testing.T) {

	input := []string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##",
	}
	expected := 8

	asteroidMap := createAsteroidMap(input)

	highest := 0

	for i, value := range asteroidMap {
		if value {
			num := numberOfAsteroidsSeenFromSpace(asteroidMap, i)
			fmt.Println(i, num)
			if num > highest {
				highest = num
			}
		}
	}

	if highest != expected {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", highest))
	}

}

func TestMonitoringStationVapeSmallCheck(t *testing.T) {

	input := []string{
		".#....#####...#..",
		"##...##.#####..##",
		"##...#...#.#####.",
		"..#.....#...###..",
		"..#.#.....#....##",
	}
	expected := spaceWithCenter{x: 15, y: 1, center: space{x: 8, y: 3}}

	asteroidMap := createAsteroidMap(input)

	highestSpace := space{x: 8, y: 3}

	// Best space is highestSpace
	the200thAsteroid := vapeAsteroids(asteroidMap, highestSpace, 9)

	fmt.Println("The 9th space to be destroyed is", the200thAsteroid)

	if the200thAsteroid != expected {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", the200thAsteroid))
	}

}
