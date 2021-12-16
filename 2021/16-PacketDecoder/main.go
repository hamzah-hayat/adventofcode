package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	methodP *string
)

func parseFlags() {
	methodP = flag.String("method", "p2", "The method/part that should be run, valid are p1,p2 and test")
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
		break
	case "p2":
		fmt.Println("Gold:" + PartTwo("input"))
		break
	}
}

func PartOne(filename string) string {
	input := readInput(filename)

	packet := ParsePacket(0, convertHexToBinary(input[0]))

	versionSum := 0
	versionSum += getVersionsRecursive(packet)

	versionSumStr := strconv.Itoa(versionSum)

	return versionSumStr
}

func PartTwo(filename string) string {
	input := readInput(filename)

	packet := ParsePacket(0, convertHexToBinary(input[0]))

	value := EvaluatePacketRecursive(packet)

	valueSumStr := strconv.FormatInt(value, 10)

	return valueSumStr
}

func ParsePacket(currentIndex int64, inputBinary string) Packet {

	// Create a packet
	// First three bits are version
	version, _ := strconv.ParseInt(inputBinary[currentIndex:currentIndex+3], 2, 64)
	currentIndex = currentIndex + 3
	// Then type id
	id, _ := strconv.ParseInt(inputBinary[currentIndex:currentIndex+3], 2, 64)
	currentIndex = currentIndex + 3

	// value
	var num int64
	num = 0

	// length
	var length int64
	length = 6

	// length
	var lengthId int64
	lengthId = 0

	// Packets
	var packets []Packet
	packets = make([]Packet, 0)

	// Check if we have an operator packet or literal packet
	if id == 4 {
		// Literal
		numbinary := ""
		keepReading := true
		for keepReading {
			keepReading = false
			if inputBinary[currentIndex] == '1' {
				keepReading = true
			}
			currentIndex++
			// Read this value
			numbinary += inputBinary[currentIndex : currentIndex+4]
			currentIndex = currentIndex + 4
			length = length + 5
		}
		// now convert to num
		num, _ = strconv.ParseInt(numbinary, 2, 64)
		currentIndex = length
	} else if id != 4 {
		// Operator
		lengthIdInt, _ := strconv.Atoi(string(inputBinary[currentIndex]))
		lengthId = int64(lengthIdInt)
		currentIndex++
		length++
		if lengthId == 0 {
			lengthNum, _ := strconv.ParseInt(inputBinary[currentIndex:currentIndex+15], 2, 64)
			var currentParsedBits int64
			currentParsedBits = 0
			totalBits := lengthNum
			currentIndex = currentIndex + 15
			length += 15

			// Now parse subpackets
			notFinished := true
			for notFinished {
				if currentParsedBits != totalBits {
					notFinished = true
				} else {
					break
				}
				p := ParsePacket(0, inputBinary[currentIndex:])
				packets = append(packets, p)
				currentIndex += p.length
				currentParsedBits += p.length
				length += p.length
			}
		} else {
			packetsNum, _ := strconv.ParseInt(inputBinary[currentIndex:currentIndex+11], 2, 64)

			var currentParsedPackets int64
			currentParsedPackets = 0
			totalPackets := packetsNum
			currentIndex = currentIndex + 11
			length += 11

			// Now parse subpackets
			notFinished := true
			for notFinished {
				if currentParsedPackets != totalPackets {
					notFinished = true
				} else {
					break
				}
				p := ParsePacket(0, inputBinary[currentIndex:])
				packets = append(packets, p)
				currentIndex += p.length
				currentParsedPackets++
				length += p.length
			}
		}
	}

	return Packet{version: version, id: id, length: length, lengthId: lengthId, value: num, subPackets: packets}
}

func EvaluatePacketRecursive(packet Packet) int64 {
	var value int64
	value = 0

	switch packet.id {
	case 0:
		//sum
		for _, sp := range packet.subPackets {
			value += EvaluatePacketRecursive(sp)
		}
	case 1:
		//product
		value = EvaluatePacketRecursive(packet.subPackets[0])
		for i := 1; i < len(packet.subPackets); i++ {
			value *= EvaluatePacketRecursive(packet.subPackets[i])
		}
	case 2:
		//min
		var min int64
		min = math.MaxInt
		for _, sp := range packet.subPackets {
			if sp.value < int64(min) {
				min = EvaluatePacketRecursive(sp)
			}
		}
		value = min
	case 3:
		//max
		var max int64
		max = 0
		for _, sp := range packet.subPackets {
			if sp.value > int64(max) {
				max = EvaluatePacketRecursive(sp)
			}
		}
		value = max
	case 4:
		value = packet.value
	case 5:
		// Greater than
		if EvaluatePacketRecursive(packet.subPackets[0]) > EvaluatePacketRecursive(packet.subPackets[1]) {
			value = 1
		} else {
			value = 0
		}
	case 6:
		// Less than
		if EvaluatePacketRecursive(packet.subPackets[0]) < EvaluatePacketRecursive(packet.subPackets[1]) {
			value = 1
		} else {
			value = 0
		}
	case 7:
		// Equal
		if EvaluatePacketRecursive(packet.subPackets[0]) == EvaluatePacketRecursive(packet.subPackets[1]) {
			value = 1
		} else {
			value = 0
		}
	default:

	}

	return value
}

type Packet struct {
	version    int64
	id         int64
	length     int64
	lengthId   int64
	value      int64
	subPackets []Packet
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

// Turn hex into binary
func convertHexToBinary(inputHex string) string {
	binary := ""

	for _, v := range inputHex {
		switch v {
		case '0':
			binary += "0000"
		case '1':
			binary += "0001"
		case '2':
			binary += "0010"
		case '3':
			binary += "0011"
		case '4':
			binary += "0100"
		case '5':
			binary += "0101"
		case '6':
			binary += "0110"
		case '7':
			binary += "0111"
		case '8':
			binary += "1000"
		case '9':
			binary += "1001"
		case 'A':
			binary += "1010"
		case 'B':
			binary += "1011"
		case 'C':
			binary += "1100"
		case 'D':
			binary += "1101"
		case 'E':
			binary += "1110"
		case 'F':
			binary += "1111"
		}

	}

	return binary
}

func getVersionsRecursive(p Packet) int {
	versionSum := 0
	versionSum += int(p.version)
	for _, sp := range p.subPackets {
		versionSum += getVersionsRecursive(sp)
	}
	return versionSum
}
