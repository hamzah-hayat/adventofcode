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

func parseFlags() {
	methodP = flag.String("method", "all", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()
}

func main() {

	parseFlags()

	switch *methodP {
	case "all":
		fmt.Println("Silver:" + PartOne("input"))
		fmt.Println("Gold:" + PartTwo("input", 2000))
	case "p1":
		fmt.Println("Silver:" + PartOne("input"))
	case "p2":
		fmt.Println("Gold:" + PartTwo("input", 2000))
	}
}

func PartOne(filename string) string {
	input := readInputInt(filename)

	// Work out the 2000th secret
	secrets := make(map[int]int)
	for _, secret := range input {
		// Calculate 2000 times
		secrets[CalculateSecretAfterLoop(secret, 2000)] += 1
	}

	total := 0
	for calcSecret := range secrets {
		total += calcSecret
	}

	return strconv.Itoa(total)
}

// Calculate the result of multiplying the secret number by 64. Then, mix this result into the secret number. Finally, prune the secret number.
// Calculate the result of dividing the secret number by 32. Round the result down to the nearest integer. Then, mix this result into the secret number. Finally, prune the secret number.
// Calculate the result of multiplying the secret number by 2048. Then, mix this result into the secret number. Finally, prune the secret number.
func CalculateSecretAfterLoop(secret, loops int) int {
	secretNumber := secret
	for i := 0; i < loops; i++ {

		firstCalc := 64 * secretNumber
		secretNumber = secretNumber ^ firstCalc
		secretNumber = secretNumber % 16777216

		secondCalc := secretNumber / 32
		secretNumber = secretNumber ^ secondCalc
		secretNumber = secretNumber % 16777216

		thirdCalc := secretNumber * 2048
		secretNumber = secretNumber ^ thirdCalc
		secretNumber = secretNumber % 16777216
	}
	return secretNumber
}

func PartTwo(filename string, loops int) string {
	input := readInputInt(filename)

	// Work out all prices for each secret
	// Then add all price changes to map
	// Find best price sequences
	// Then return the best price sequence total
	bestPriceChangeMap := make(map[string]int)
	for _, secret := range input {
		prices := CalculatePricesAfterLoop(secret, loops)

		// This map stores each sequence and its total price at the end
		priceChangeMap := make(map[string]int)
		for i := 0; i < len(prices)-4; i++ {
			price0 := prices[i]
			price1 := prices[i+1]
			price2 := prices[i+2]
			price3 := prices[i+3]
			price4 := prices[i+4]

			priceChange1 := strconv.Itoa(price1 - price0)
			priceChange2 := strconv.Itoa(price2 - price1)
			priceChange3 := strconv.Itoa(price3 - price2)
			priceChange4 := strconv.Itoa(price4 - price3)

			priceChangeArray := []string{priceChange1, priceChange2, priceChange3, priceChange4}
			priceChangeString := strings.Join(priceChangeArray, ",")
			currentPrice := prices[i+4]

			_, exists := priceChangeMap[priceChangeString]
			if !exists {
				priceChangeMap[priceChangeString] = currentPrice
			}
		}

		for priceChangeSequence, price := range priceChangeMap {
			bestPriceChangeMap[priceChangeSequence] += price
		}
	}

	bestSequence := ""
	highest := 0
	for sequence, bestPrice := range bestPriceChangeMap {
		if bestPrice > highest {
			bestSequence = sequence
			highest = bestPrice
		}
	}
	fmt.Println("The best sequence is", bestSequence)

	return strconv.Itoa(highest)
}

func CalculatePricesAfterLoop(secret, loops int) []int {
	prices := []int{}
	strnum := strconv.Itoa(secret)
	price, _ := strconv.Atoi(strnum[len(strnum)-1:])
	prices = append(prices, price)

	secretNumber := secret
	for i := 0; i < loops; i++ {

		firstCalc := 64 * secretNumber
		secretNumber = secretNumber ^ firstCalc
		secretNumber = secretNumber % 16777216

		secondCalc := secretNumber / 32
		secretNumber = secretNumber ^ secondCalc
		secretNumber = secretNumber % 16777216

		thirdCalc := secretNumber * 2048
		secretNumber = secretNumber ^ thirdCalc
		secretNumber = secretNumber % 16777216

		strnum := strconv.Itoa(secretNumber)
		price, _ := strconv.Atoi(strnum[len(strnum)-1:])
		prices = append(prices, price)
	}
	return prices
}

// Read data from input.txt
// Return the string as int
func readInputInt(filename string) []int {

	var input []int

	f, _ := os.Open(filename + ".txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		input = append(input, num)
	}
	return input
}
