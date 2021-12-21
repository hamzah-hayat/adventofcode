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

	players := GetPlayerStarts(input)

	rolls := PlayDiceDeterministic(players)

	// find loser score * rolls
	loserScore := 0
	for _, p := range players {
		if p.score < 1000 {
			loserScore = p.score
		}
	}

	num := strconv.Itoa(rolls * loserScore)

	return num
}

func PartTwo(filename string) string {
	//input := readInput(filename)

	return "0"
}

func GetPlayerStarts(input []string) []*Player {
	players := make([]*Player, 0)

	playerRegex := regexp.MustCompile("Player [0-9] starting position: ([0-9])")
	for _, line := range input {
		match := playerRegex.FindAllStringSubmatch(line, 1)
		start, _ := strconv.Atoi(match[0][1])
		p := Player{position: start, score: 0}
		players = append(players, &p)
	}
	return players
}

func PlayDiceDeterministic(players []*Player) int {
	dice := 1
	rolls := 0
	playerNotWon := true
	for playerNotWon {
		// Roll and move
		for _, p := range players {
			movement := dice + (dice + 1) + (dice + 2)
			rolls = rolls + 3
			dice = dice + 3
			p.position = (p.position + movement)
			for p.position > 10 {
				p.position = p.position - 10
			}
			p.score += p.position
			playerNotWon = !CheckWin(1000, players)
			if !playerNotWon {
				break
			}
		}
		if !playerNotWon {
			break
		}
	}
	return rolls
}

func GetPlayerStartsQuantum(input []string) []*PlayerQuantum {
	players := make([]*PlayerQuantum, 0)

	playerRegex := regexp.MustCompile("Player [0-9] starting position: ([0-9])")
	for _, line := range input {
		match := playerRegex.FindAllStringSubmatch(line, 1)
		start, _ := strconv.Atoi(match[0][1])
		startMap := make(map[int]int)
		startMap[start] = 1
		p := PlayerQuantum{position: startMap}
		players = append(players, &p)
	}
	return players
}


func CheckWin(win int, players []*Player) bool {
	for _, p := range players {
		if p.score >= win {
			return true
		}
	}
	return false
}

type Player struct {
	position int
	score    int
}

type PlayerQuantum struct {
	position map[int]int
	score    map[int]int
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
