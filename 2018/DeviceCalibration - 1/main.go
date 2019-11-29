package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readDevice()

	fmt.Println("Total of all frequencies is " + strconv.Itoa(SumAllFrequencies(input)))
	fmt.Println("First time a frequency is repeated is " + strconv.Itoa(FindFirstRepeatedFrequencey(input)))

}

//Read data from device.txt
//Load this data into a int array
func readDevice() []int {

	var input []int

	f, _ := os.Open("device.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() != "" {
			// Take input and put into int array
			number, _ := strconv.Atoi(scanner.Text())
			input = append(input, number)
		}
	}
	return input
}

// Sum all the frequencies together
// Simple for loop
func SumAllFrequencies(input []int) int {
	total := 0

	for _, num := range input {
		total += num
	}
	return total
}

func FindFirstRepeatedFrequencey(input []int) int {
	currentTotal := 0
	var usedFrequencies []int
	for {
		for _, num := range input {
			usedFrequencies = append(usedFrequencies, currentTotal)
			currentTotal += num

			for _, used := range usedFrequencies {
				//fmt.Println("Checking " + strconv.Itoa(used) + " and " + strconv.Itoa(currentTotal))
				if currentTotal == used {
					return used
				}
			}
		}
	}
	//fmt.Println(usedFrequencies)
}
