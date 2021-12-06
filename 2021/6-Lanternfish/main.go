package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	methodP *string
)

func init() {
	// Use Flags to run a part
	methodP = flag.String("method", "p2", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()
}

func main() {
	switch *methodP {
	case "p1":
		PartOne()
		break
	case "p2":
		PartTwo()
		break
	case "test":
		break
	}
}

// Horrible array solution, takes time and unusable with days > ~200
func PartOne() {
	input := readInput()

	var fishs []int

	fish := 0
	fishs = append(fishs, fish)

	for i := 0; i < 80; i++ {
		for i := range fishs {
			fishs = processFishDay(i, fishs)
		}
		//fmt.Println("Passed day" + strconv.Itoa(i))
	}

	//fmt.Println(len(fishs))

	//0 = 1421
	//1 = 1401
	//2 = 1191
	//3 = 1154
	//4 = 1034
	//5 = 950

	sum := 0

	for _, v := range strings.Split(input[0], ",") {
		num, _ := strconv.Atoi(v)
		switch num {
		case 0:
			sum += 1421
		case 1:
			sum += 1401
		case 2:
			sum += 1191
		case 3:
			sum += 1154
		case 4:
			sum += 1034
		case 5:
			sum += 950

		}
	}
	fmt.Println(sum)

}

// Much better map solution, runs in constant time pretty much
func PartTwo() {
	input := readInput()

	var fishs map[int]int
	fishs = make(map[int]int)

	for _, v := range strings.Split(input[0], ",") {
		num, _ := strconv.Atoi(v)
		fishs[num]++
	}

	for i := 0; i < 256; i++ {
		fishs = processFishDayMap(fishs)
		//fmt.Println("Passed day" + strconv.Itoa(i))
	}

	sum := 0
	for _, v := range fishs {
		sum += v
	}
	fmt.Println(sum)
}

func processFishDay(fishNum int, fishs []int) []int {
	if fishs[fishNum] == 0 {
		fishs[fishNum] = 6
		// Also create a new Fish
		fishs = append(fishs, 8)
	} else {
		fishs[fishNum]--
	}
	return fishs
}

func processFishDayMap(fishs map[int]int) map[int]int {
	var newfishs map[int]int
	newfishs = make(map[int]int)
	newfishs[8] = fishs[0]
	newfishs[7] = fishs[8]
	newfishs[6] = fishs[0] + fishs[7]
	newfishs[5] = fishs[6]
	newfishs[4] = fishs[5]
	newfishs[3] = fishs[4]
	newfishs[2] = fishs[3]
	newfishs[1] = fishs[2]
	newfishs[0] = fishs[1]
	return newfishs
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
