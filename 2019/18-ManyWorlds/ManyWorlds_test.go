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

	fmt.Println(solveCalls)

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

	fmt.Println(solveCalls)

	if expected != lowestSteps {
		t.Error(fmt.Sprint("Exected output ", expected, " but got ", lowestSteps))
	}

}

func TestManyWorlds_MazeSmall3(t *testing.T) {

	input := []string{
		"########################",
		"#...............b.C.D.f#",
		"#.######################",
		"#.....@.a.B.c.d.A.e.F.g#",
		"########################",
	}
	expected := 132

	maze := createMaze(input)

	lowestSteps := solveMaze(maze)

	fmt.Println(solveCalls)

	if expected != lowestSteps {
		t.Error(fmt.Sprint("Exected output ", expected, " but got ", lowestSteps))
	}

}

func TestManyWorlds_MazeSmall4(t *testing.T) {

	input := []string{
		"#################",
		"#i.G..c...e..H.p#",
		"########.########",
		"#j.A..b...f..D.o#",
		"########@########",
		"#k.E..a...g..B.n#",
		"########.########",
		"#l.F..d...h..C.m#",
		"#################",
	}
	expected := 136

	maze := createMaze(input)

	lowestSteps := solveMaze(maze)

	fmt.Println(solveCalls)

	if expected != lowestSteps {
		t.Error(fmt.Sprint("Exected output ", expected, " but got ", lowestSteps))
	}

}

func TestManyWorlds_MazeSmall5(t *testing.T) {

	input := []string{
		"########################",
		"#@..............ac.GI.b#",
		"###d#e#f################",
		"###A#B#C################",
		"###g#h#i################",
		"########################",
	}
	expected := 81

	maze := createMaze(input)

	lowestSteps := solveMaze(maze)

	fmt.Println(solveCalls)

	if expected != lowestSteps {
		t.Error(fmt.Sprint("Exected output ", expected, " but got ", lowestSteps))
	}

}
