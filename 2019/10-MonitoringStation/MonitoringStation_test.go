package main

import (
	"fmt"
	"testing"
)

func MonitoringStationSmallCheck(t *testing.T) {

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
			if num > highest {
				highest = num
			}
		}
	}

	if highest != expected {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", highest))
	}

}
