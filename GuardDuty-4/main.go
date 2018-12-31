package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type shift struct {
	guardID string // The ID of the guard
	sleep   []bool // For each minute of the shift, was the guard asleep or not
}

func main() {
	PartOne()
	//PartTwo()
}

func PartOne() {
	shifts := readInput()

	// Find out who the most sleepy guard is
	guardSleepSum := make(map[string]int)
	for _, shift := range shifts {
		// Build a map and figure out which guard is the most sleep
		for _, sleepValue := range shift.sleep {
			if sleepValue {
				guardSleepSum[shift.guardID]++
			}
		}
	}

	//Find sleepiest guard
	highestSum := 0
	highestGuard := ""
	for guardID, sleepSum := range guardSleepSum {
		if sleepSum > highestSum {
			highestGuard = guardID
			sleepSum = highestSum
		}
	}

	fmt.Println("The sleepiest guard is " + highestGuard)

	// Find out when this guard was the sleepiest
	guardMinuteSleepSum := make([]int, 60)
	for _, shift := range shifts {
		if shift.guardID == highestGuard {
			for min, val := range shift.sleep {
				if val {
					guardMinuteSleepSum[min]++
				}
			}
		}
	}

	fmt.Println(guardMinuteSleepSum)

	highestMinuteSum := 0
	highestMinute := 0
	for minute, sleepMinSum := range guardMinuteSleepSum {
		if sleepMinSum > highestMinuteSum {
			highestMinute = minute
			highestMinuteSum = sleepMinSum
		}
	}

	highestMinuteStr := strconv.Itoa(highestMinute)
	fmt.Println("They were sleepiest during minute " + highestMinuteStr)

}

func PartTwo() {
	//input := readInput()
}

// Read data from input.txt
// Load it into string array
func readInput() []shift {

	var input []string

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() != "" {
			line := scanner.Text()
			input = append(input, line)
		}
	}

	// Now have all the input, sort it by date/time
	sort.Strings(input)

	// Now we have to build a "shift" for each day
	var shifts []shift
	i := 0
	for i < len(input) {
		newShift, nexti := BuildShift(i, input)
		i = nexti
		shifts = append(shifts, newShift)
	}

	return shifts
}

func BuildShift(i int, input []string) (shift, int) {
	// Build a shift using the input
	// Return the new shift and the next i
	var newShift shift

	// Get the GuardID for this shift
	newShift.guardID = strings.TrimLeft(strings.Fields(input[i])[3], "#")
	newShift.sleep = make([]bool, 60)
	i++

	time := 0

	for {
		//Check if this is last shift, aka EOF
		if i >= len(input) {
			break
		}
		// Check if we are looking at a new shift
		if strings.Contains(input[i], "#") {
			break
		}
		// Otherwise, examine the next action and act accordingly
		if strings.Contains(input[i], "sleep") {
			newTimeStr := strings.Fields(input[i])[1]
			newTimeVal, _ := strconv.Atoi(strings.TrimPrefix(strings.TrimSuffix(newTimeStr, "]"), "00:"))
			for j := time; j < newTimeVal; j++ {
				newShift.sleep[j] = true
			}
			time = newTimeVal
		}
		if strings.Contains(input[i], "wakes") {
			newTimeStr := strings.Fields(input[i])[1]
			newTimeVal, _ := strconv.Atoi(strings.TrimPrefix(strings.TrimSuffix(newTimeStr, "]"), "00:"))
			for j := time; j < newTimeVal; j++ {
				newShift.sleep[j] = false
			}
			time = newTimeVal
		}
		i++
	}
	return newShift, i
}
