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
	//PartOne()
	PartTwo()
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
			highestSum = sleepSum
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

	highestGuardStr, _ := strconv.Atoi(highestGuard)
	checksum := highestMinute * highestGuardStr
	fmt.Println("Checksum is " + strconv.Itoa(checksum))

}

func PartTwo() {
	shifts := readInput()

	// Find the person who had the sleepiest minute
	// For each person, find their sleepiest minute and add to a map
	sleepiestMinute := make(map[string]int)

	var guardIDs []string

	// Find all Guard IDs
	for _, shift := range shifts {
		found := false
		for _, guardID := range guardIDs {
			if shift.guardID == guardID {
				found = true
				break
			}
		}
		if !found {
			guardIDs = append(guardIDs, shift.guardID)
		}
	}

	//fmt.Println(guardIDs)

	guardSleepSumList := make(map[string][]int)
	for _, guardID := range guardIDs {
		//Find the sleepiest minute for each Guard
		guardMinuteSleepSum := make([]int, 60)
		for _, shift := range shifts {
			if shift.guardID == guardID {
				for min, val := range shift.sleep {
					if val {
						guardMinuteSleepSum[min]++
					}
				}
			}
		}
		guardSleepSumList[guardID] = guardMinuteSleepSum

		// Find their sleepiest minute and add the value to our sleepiestMinute map
		highestSleepValue := 0
		for _, sleepValue := range guardMinuteSleepSum {
			if sleepValue > highestSleepValue {
				highestSleepValue = sleepValue
			}
		}
		sleepiestMinute[guardID] = highestSleepValue
	}

	fmt.Println(sleepiestMinute)

	// Now find the sleepiest minute guard
	highestSleepMinuteValue := 0
	highestSleepGuard := ""
	for guardID, sleepValue := range sleepiestMinute {
		if sleepValue > highestSleepMinuteValue {
			highestSleepGuard = guardID
			highestSleepMinuteValue = sleepValue
		}
	}

	highestSleepGuardID, _ := strconv.Atoi(highestSleepGuard)

	highestSleepMinute := 0
	// Now that we know which guard has the sleepiest minute, we should find out what minute they were the sleepiest on
	for guardID, sleepValueList := range guardSleepSumList {
		if guardID == highestSleepGuard {
			highestSleepValue := 0
			for minute, sleepValue := range sleepValueList {
				if sleepValue > highestSleepValue {
					highestSleepValue = sleepValue
					highestSleepMinute = minute
				}
			}
		}

	}

	fmt.Println("The highest guard is " + highestSleepGuard)
	fmt.Println("The highest slept minute is " + strconv.Itoa(highestSleepMinute))
	fmt.Printf("The checksum is %v\n", highestSleepGuardID*highestSleepMinute)

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
				newShift.sleep[j] = false
			}
			time = newTimeVal
		}
		if strings.Contains(input[i], "wakes") {
			newTimeStr := strings.Fields(input[i])[1]
			newTimeVal, _ := strconv.Atoi(strings.TrimPrefix(strings.TrimSuffix(newTimeStr, "]"), "00:"))
			for j := time; j < newTimeVal; j++ {
				newShift.sleep[j] = true
			}
			time = newTimeVal
		}
		i++
	}
	return newShift, i
}
