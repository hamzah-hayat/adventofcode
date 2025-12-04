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
	totalJoltage := 0

	for _, bank := range input {
		highestJoltage := 0
		for i := 9; i > 0; i-- {
			for index, battery := range bank {
				batteryNum, _ := strconv.Atoi(string(battery))
				// We've found the highest number battery, now check after it
				if batteryNum == i {
					highest := 0
					for j := index + 1; j < len(bank); j++ {
						batteryNum2, _ := strconv.Atoi(string(bank[j]))
						if batteryNum2 > highest {
							highest = batteryNum2
						}
					}

					// Did we find a higher battery after the current one?
					if highest != 0 {
						joltageStr := string(battery) + strconv.Itoa(highest)
						joltage, _ := strconv.Atoi(joltageStr)
						if highestJoltage < joltage {
							highestJoltage = joltage
						}
					}
				}
			}
		}
		totalJoltage += highestJoltage
	}

	return strconv.Itoa(totalJoltage)
}

func PartTwo(filename string) string {
	input := readInput(filename)
	totalJoltage := 0

	for _, bank := range input {
		highestJoltage := findHighestJoltage(bank,12)
		totalJoltage += highestJoltage
	}

	return strconv.Itoa(totalJoltage)
}

func findHighestJoltage(bank string, batteryNumber int) int {
	highestJoltage := 0
	batteries := ""

	bankIndex:=-1
	for i := 0; i < batteryNumber; i++ {
		bankIndex++
		currentBestBattery := bank[bankIndex]
		for j := bankIndex+1; j < len(bank)-(batteryNumber-len(batteries)-1); j++ {
			checkBattery := bank[j]
			
			currentBestBatteryNum, _ := strconv.Atoi(string(currentBestBattery))
			checkBatteryNum, _ := strconv.Atoi(string(checkBattery))
			if checkBatteryNum > currentBestBatteryNum {
				currentBestBattery = checkBattery
				bankIndex = j
			}
		}
		batteries += string(currentBestBattery)
	}

	highestJoltage, _ = strconv.Atoi(batteries)

	return highestJoltage
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
