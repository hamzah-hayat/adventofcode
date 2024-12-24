package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

	total := 0
	cache := make(map[RobotCache]int)
	for _, code := range input {
		total += RunCode(code, 2, cache)
	}

	return strconv.Itoa(total)
}

func RunCode(startCode string, robots int, cache map[RobotCache]int) int {

	bigKeypadcode := BigKeyPadBestPath(startCode)
	length := RunSmallKeyPadRecursive(bigKeypadcode, cache, 0, robots)

	keyPadNumber, _ := strconv.Atoi(startCode[:len(startCode)-1])

	return keyPadNumber * length
}

type RobotCache struct {
	code  string
	robot int
}

func RunSmallKeyPadRecursive(code string, cache map[RobotCache]int, robotNumBefore, maxRobots int) int {
	codeExtended := SmallKeyPadBestPath(code)

	if robotNumBefore == maxRobots {
		return len(code)
	}

	length := 0
	for _, curr := range SplitOnA(codeExtended) {
		cacheEntry := RobotCache{curr, robotNumBefore + 1}
		if _, ok := cache[cacheEntry]; ok {
			length += cache[cacheEntry]
			continue
		}
		count := RunSmallKeyPadRecursive(curr, cache, robotNumBefore+1, maxRobots)
		length += count
	}

	cache[RobotCache{code, robotNumBefore}] = length
	return length
}

func SplitOnA(code string) []string {
	splits := make([]string, 0)
	start := 0
	for i, c := range code {
		if c == 'A' {
			splits = append(splits, code[start:i+1])
			start = i + 1
		}
	}
	return splits
}

func BigKeyPadBestPath(code string) string {
	// +---+---+---+
	// | 7 | 8 | 9 |
	// +---+---+---+
	// | 4 | 5 | 6 |
	// +---+---+---+
	// | 1 | 2 | 3 |
	// +---+---+---+
	//     | 0 | A |
	//     +---+---+
	// Robot always starts at A

	expanded := ""
	current := "A"
	for i := 0; i < len(code); i++ {
		start := current
		end := string(code[i])
		expanded += BigKeypadPathFind(start, end)
		expanded += "A"
		current = end
	}

	return expanded
}

func BigKeypadPathFind(start, end string) string {
	// +---+---+---+
	// | 7 | 8 | 9 |
	// +---+---+---+
	// | 4 | 5 | 6 |
	// +---+---+---+
	// | 1 | 2 | 3 |
	// +---+---+---+
	//     | 0 | A |
	//     +---+---+
	bigKeypad := make(map[string]Point)
	bigKeypad["7"] = Point{0, 0}
	bigKeypad["8"] = Point{1, 0}
	bigKeypad["9"] = Point{2, 0}
	bigKeypad["4"] = Point{0, 1}
	bigKeypad["5"] = Point{1, 1}
	bigKeypad["6"] = Point{2, 1}
	bigKeypad["1"] = Point{0, 2}
	bigKeypad["2"] = Point{1, 2}
	bigKeypad["3"] = Point{2, 2}
	bigKeypad["0"] = Point{1, 3}
	bigKeypad["A"] = Point{2, 3}

	startPoint := bigKeypad[start]
	endPoint := bigKeypad[end]

	xDiff := endPoint.x - startPoint.x
	yDiff := endPoint.y - startPoint.y

	vertical := ""
	for yDiff < 0 {
		vertical += "^"
		yDiff++
	}
	for yDiff > 0 {
		vertical += "v"
		yDiff--
	}

	horizontal := ""
	for xDiff < 0 {
		horizontal += "<"
		xDiff++
	}
	for xDiff > 0 {
		horizontal += ">"
		xDiff--
	}

	xDiff = endPoint.x - startPoint.x

	// Priority: < over ^ over v over >
	if startPoint.y == 3 && endPoint.x == 0 {
		return vertical + horizontal
	} else if startPoint.x == 0 && endPoint.y == 3 {
		return horizontal + vertical
	} else if xDiff < 0 {
		return horizontal + vertical
	} else {
		return vertical + horizontal
	}
}

func SmallKeyPadBestPath(code string) string {
	// 	   +---+---+
	//     | ^ | A |
	// +---+---+---+
	// | < | v | > |
	// +---+---+---+
	// Again, start from A

	expanded := ""
	current := 'A'
	for i := 0; i < len(code); i++ {
		start := current
		end := rune(code[i])
		expanded += SmallKeypadPathFind(start, end)
		expanded += "A"
		current = end
	}

	return expanded
}

func SmallKeypadPathFind(start, end rune) string {

	// 	   +---+---+
	//     | ^ | A |
	// +---+---+---+
	// | < | v | > |
	// +---+---+---+
	smallKeypad := make(map[rune]Point)
	smallKeypad['^'] = Point{1, 0}
	smallKeypad['A'] = Point{2, 0}
	smallKeypad['<'] = Point{0, 1}
	smallKeypad['v'] = Point{1, 1}
	smallKeypad['>'] = Point{2, 1}

	startPoint := smallKeypad[start]
	endPoint := smallKeypad[end]

	xDiff := endPoint.x - startPoint.x
	yDiff := endPoint.y - startPoint.y

	vertical := ""
	for yDiff < 0 {
		vertical += "^"
		yDiff++
	}
	for yDiff > 0 {
		vertical += "v"
		yDiff--
	}

	horizontal := ""
	for xDiff < 0 {
		horizontal += "<"
		xDiff++
	}
	for xDiff > 0 {
		horizontal += ">"
		xDiff--
	}

	xDiff = endPoint.x - startPoint.x

	// Priority: < over ^ over v over >
	if startPoint.x == 0 && endPoint.y == 0 {
		return horizontal + vertical
	} else if startPoint.y == 0 && endPoint.x == 0 {
		return vertical + horizontal
	} else if xDiff < 0 {
		return horizontal + vertical
	} else {
		return vertical + horizontal
	}
}

type Point struct {
	x int
	y int
}

func PartTwo(filename string) string {
	input := readInput(filename)

	total := 0
	cache := make(map[RobotCache]int)
	for _, code := range input {
		total += RunCode(code, 25, cache)
	}

	return strconv.Itoa(total)
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
