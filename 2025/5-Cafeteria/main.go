package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

	ranges := make([]Range, 0)
	items := make([]int, 0)

	grabbingRangeInputs := true
	for _, line := range input {
		if grabbingRangeInputs {
			if line == "" {
				grabbingRangeInputs = false
				continue
			} else {
				rangeSplit := strings.Split(line, "-")
				lower, _ := strconv.Atoi(rangeSplit[0])
				upper, _ := strconv.Atoi(rangeSplit[1])
				ranges = append(ranges, Range{lower: lower, upper: upper})
			}
		} else {
			item, _ := strconv.Atoi(line)
			items = append(items, item)
		}
	}

	freshItems := 0
	for _, item := range items {
		for _, r := range ranges {
			if numberInRange(item, r) {
				freshItems++
				break
			}
		}
	}

	return strconv.Itoa(freshItems)
}

func PartTwo(filename string) string {
	input := readInput(filename)

	ranges := make([]Range, 0)
	mergedRanges := make([]Range, 0)

	for _, line := range input {
		if line == "" {
			break
		} else {
			rangeSplit := strings.Split(line, "-")
			lower, _ := strconv.Atoi(rangeSplit[0])
			upper, _ := strconv.Atoi(rangeSplit[1])
			ranges = append(ranges, Range{lower: lower, upper: upper})
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].lower == ranges[j].lower {
			return ranges[i].upper < ranges[j].upper
		} else {
			return ranges[i].lower < ranges[j].lower
		}
	})

	// Add the first range to the merged ranges
	mergedRanges = append(mergedRanges, ranges[0])
	for i := 1; i < len(ranges); i++ {

		lastRange := mergedRanges[len(mergedRanges)-1]

		// Go through our ranges and merge them together
		// Start by always adding the final range in full
		if ranges[i].lower <= lastRange.lower && ranges[i].upper <= lastRange.upper {
			continue
		} else if ranges[i].lower <= lastRange.upper && ranges[i].upper > lastRange.upper {
			mergedRanges = mergedRanges[:len(mergedRanges)-1]
			mergedRanges = append(mergedRanges, Range{lower: lastRange.lower, upper: ranges[i].upper})
			continue
		} else if ranges[i].lower > lastRange.upper {
			mergedRanges = append(mergedRanges, ranges[i])
			continue
		}
	}

	totalFreshIDs := 0
	for _, r := range mergedRanges {
		totalFreshIDs += r.upper - r.lower + 1
	}

	return strconv.Itoa(totalFreshIDs)
}

type Range struct {
	lower int
	upper int
}

func numberInRange(value int, r Range) bool {
	return value >= r.lower && value <= r.upper
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
