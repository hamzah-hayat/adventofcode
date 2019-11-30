package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Use Flags to run a part
	methodP := flag.String("method", "p1", "The method/part that should be run, valid are p1,p2 and test")
	flag.Parse()

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

func PartOne() {
	input := readInput()

	i := 0
	for {
		h := md5.New()
		io.WriteString(h, input)
		io.WriteString(h, strconv.Itoa(i))
		str := fmt.Sprintf("%x", h.Sum(nil))

		if strings.HasPrefix(str, "00000") {
			fmt.Println("The correct number is", i)
			break
		}
		i++
	}
}

func PartTwo() {
	input := readInput()

	i := 0
	for {
		h := md5.New()
		io.WriteString(h, input)
		io.WriteString(h, strconv.Itoa(i))
		str := fmt.Sprintf("%x", h.Sum(nil))

		if strings.HasPrefix(str, "000000") {
			fmt.Println("The correct number is", i)
			break
		}
		i++
	}
}

// Read data from input.txt
// Return the string, so that we can deal with it however
func readInput() string {

	var input string

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input += scanner.Text()
	}
	return input
}
