package main

import (
	"fmt"
	"testing"
)

func TestManyWorlds_MazeSmall1(t *testing.T) {

	input := []string{
		"#########",
		"#b.A.@.a#",
		"#########",
	}
	expected := 8

	maze := createMaze(input)

	lowestSteps := solveMaze(maze)

	if expected != lowestSteps {
		t.Error(fmt.Sprint("Exected output ", expected, " but got ", lowestSteps))
	}

}

func TestManyWorlds_MazeSmall2(t *testing.T) {

	input := []string{
		"########################",
		"#f.D.E.e.C.b.A.@.a.B.c.#",
		"######################.#",
		"#d.....................#",
		"########################",
	}
	expected := 86

	maze := createMaze(input)

	lowestSteps := solveMaze(maze)

	if expected != lowestSteps {
		t.Error(fmt.Sprint("Exected output ", expected, " but got ", lowestSteps))
	}

}
