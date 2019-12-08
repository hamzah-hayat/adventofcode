package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hamzah-hayat/adventofcode/intcode"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p1", "The method/part that should be run, valid are p1,p2 and test")
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

	programStr := strings.Split(input[0], ",")
	var program []int

	for _, s := range programStr {
		i, _ := strconv.Atoi(s)
		program = append(program, i)
	}

	// Can only use each signal ONCE
	signals := []int{0, 1, 2, 3, 4}

	highest := RunRocketAmplifiers(program, signals)

	fmt.Println("Highest output is", highest)

}

func partTwo() {
	input := readInput()

	programStr := strings.Split(input[0], ",")
	var program []int

	for _, s := range programStr {
		i, _ := strconv.Atoi(s)
		program = append(program, i)
	}

	// Can only use each signal ONCE
	signals := []int{5, 6, 7, 8, 9}

	highest := RunRocketAmplifiers(program, signals)

	fmt.Println("Highest output is", highest)
}

//RunRocketAmplifiers takes a program, runs it through 5 different amplifiers, then outputs the highest result
func RunRocketAmplifiers(program []int, signals []int) int {
	highestResult := 0

	Perm(signals, func(a []int) {

		// For each permuation of the signals, run the amplifiers and get the highest result
		out := runRocketsWithSignals(program, a)

		if out > highestResult {
			highestResult = out
		}
	})

	return highestResult
}

func runRocketsWithSignals(program []int, signalSet []int) int {

	// First channel is for input
	input := make(chan int)
	// Second channel is for output
	output := make(chan int)

	go intcode.RunIntCodeProgram(program, input, output)
	input <- signalSet[0]
	input <- 0

	out := <-output

	go intcode.RunIntCodeProgram(program, input, output)
	input <- signalSet[1]
	input <- out

	out = <-output

	go intcode.RunIntCodeProgram(program, input, output)
	input <- signalSet[2]
	input <- out

	out = <-output

	go intcode.RunIntCodeProgram(program, input, output)
	input <- signalSet[3]
	input <- out

	out = <-output

	go intcode.RunIntCodeProgram(program, input, output)
	input <- signalSet[4]
	input <- out

	return <-output
}

// Perm calls f with each permutation of a.
func Perm(a []int, f func([]int)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
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
