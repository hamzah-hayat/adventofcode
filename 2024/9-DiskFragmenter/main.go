package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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
	total := 0

	// Create disk
	disk := make(map[int]int)
	isFile := true
	fileID := 0
	space := 0
	for _, char := range input[0] {
		num, _ := strconv.Atoi(string(char))
		if isFile {
			for i := 0; i < num; i++ {
				disk[space] = fileID
				space++
			}
			fileID++
		} else {
			for i := 0; i < num; i++ {
				disk[space] = -1
				space++
			}
		}
		isFile = !isFile
	}

	// Print
	// fmt.Println(PrintDisk(disk))

	// Sort disk
	for currentFileID := fileID - 1; currentFileID > 0; currentFileID-- {
		for _, fileLoc := range FindFiles(currentFileID, disk) {
			for diskScan := 0; diskScan < fileLoc; diskScan++ {
				if disk[diskScan] == -1 {
					// Move file up
					disk[diskScan] = disk[fileLoc]
					disk[fileLoc] = -1
					continue
				}
			}
		}
		fmt.Println(currentFileID)
	}

	// Print
	// fmt.Println(PrintDisk(disk))

	// Score
	for i, b := range disk {
		if b != -1 {
			total += i * b
		}
	}

	return strconv.Itoa(total)
}

func FindFiles(fileID int, disk map[int]int) []int {
	foundFiles := make([]int, 0)

	for i, v := range disk {
		if v == fileID {
			foundFiles = append(foundFiles, i)
		}
	}

	return foundFiles
}

func PrintDisk(disk map[int]int) string {
	diskStr := ""
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			diskStr += "."
		} else {
			diskStr += strconv.Itoa(disk[i])
		}
	}
	return diskStr
}

func PartTwo(filename string) string {
	input := readInput(filename)
	total := 0

	// Create disk
	disk := make(map[int]int)
	isFile := true
	fileID := 0
	space := 0
	for _, char := range input[0] {
		num, _ := strconv.Atoi(string(char))
		if isFile {
			for i := 0; i < num; i++ {
				disk[space] = fileID
				space++
			}
			fileID++
		} else {
			for i := 0; i < num; i++ {
				disk[space] = -1
				space++
			}
		}
		isFile = !isFile
	}

	// Print
	// fmt.Println(PrintDisk(disk))

	// Sort disk
	for currentFileID := fileID - 1; currentFileID > 0; currentFileID-- {
		currentFile := FindFiles(currentFileID, disk)
		sort.Ints(currentFile)
		startFileLoc := currentFile[0]

		for diskScan := 0; diskScan < startFileLoc; diskScan++ {
			if disk[diskScan] == -1 {


				// Want to try and move entire file up now
				enoughSpace := true
				for checkAhead := 0; checkAhead < len(currentFile); checkAhead++ {
					if disk[diskScan+checkAhead] != -1 {
						enoughSpace = false
					}
				}

				// Move file up
				if enoughSpace {
					for checkAhead := 0; checkAhead < len(currentFile); checkAhead++ {
						disk[diskScan+checkAhead] = disk[startFileLoc+checkAhead]
						disk[startFileLoc+checkAhead] = -1

						// Print
						// fmt.Println(PrintDisk(disk))
					}
					break
				}
			}
		}
		fmt.Println(currentFileID)
	}

	// Print
	// fmt.Println(PrintDisk(disk))

	// Score
	for i, b := range disk {
		if b != -1 {
			total += i * b
		}
	}

	return strconv.Itoa(total)
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

// Read data from input.txt
// Return the string as int
func readInputInt() []int {

	var input []int

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		input = append(input, num)
	}
	return input
}
