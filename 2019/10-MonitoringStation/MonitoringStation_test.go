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
	expected := spaceWithCenter{x: 8, y: 1, center: space{x: 8, y: 3}}

	asteroidMap := createAsteroidMap(input)

	highestSpace := space{x: 8, y: 3}

	asteroidMap[highestSpace] = false
	// Best space is highestSpace
	the200thAsteroid := vapeAsteroids(asteroidMap, highestSpace, 1)

	fmt.Println("The 1st space to be destroyed is", the200thAsteroid)

	if the200thAsteroid != expected {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", the200thAsteroid))
	}

}

func TestMonitoringStationVapeSmallCheck2(t *testing.T) {

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

	asteroidMap[highestSpace] = false
	// Best space is highestSpace
	the200thAsteroid := vapeAsteroids(asteroidMap, highestSpace, 9)

	if the200thAsteroid != expected {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", the200thAsteroid))
	}

}
