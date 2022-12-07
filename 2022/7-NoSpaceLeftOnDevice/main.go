package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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

	root := makeFileTree(filename)

	// now we have processed tree
	// now traverse tree and find all folders less then 100000 in total size
	sum := root.getFolderSizesWithLessthenSize(100000)

	return strconv.Itoa(sum)
}

func PartTwo(filename string) string {
	root := makeFileTree(filename)

	// we have 70000000 space total on disk
	currentFreeSpace := 70000000 - root.getSize()
	// we need at least 30000000 space
	neededSpace := 30000000 - currentFreeSpace
	deleteFolder := root.findSmallestFolderOverSize(root, neededSpace)

	smallestSize := strconv.Itoa(deleteFolder.getSize())

	return smallestSize
}

func makeFileTree(filename string) *Directory {
	input := readInput(filename)

	root := Directory{name: "/"}

	cdRegex := regexp.MustCompile(`\$ cd (\/|[a-z]+|..)`)
	lsRegex := regexp.MustCompile(`\$ ls`)
	fileRegex := regexp.MustCompile(`([0-9]+) ([a-z]+)\.?([a-z]+)?`)
	dirRegex := regexp.MustCompile(`dir ([a-z]+)`)

	//currentDirName := "/"
	currentFolder := &root
	for _, line := range input {
		// three choices
		// 1. Changing directory
		// 2. Running ls
		// 3. file output
		// 4. folder output

		cdRegexResult := cdRegex.FindStringSubmatch(line)
		lsRegexResult := lsRegex.FindStringSubmatch(line)
		fileRegexResult := fileRegex.FindStringSubmatch(line)
		dirRegexResult := dirRegex.FindStringSubmatch(line)

		if len(cdRegexResult) > 0 {
			// change directory
			if cdRegexResult[1] == ".." {
				// go up one directory
				//lastSlash := strings.LastIndex("/", currentDirName)
				//currentDirName = currentDirName[:lastSlash]
				currentFolder = currentFolder.parent
			} else if cdRegexResult[1] != "/" {
				// we ignore the / choice and assume this is a dir name
				//currentDirName += "/" + cdRegexResult[1]
				currentFolder = currentFolder.getChildFolderWithName(cdRegexResult[1])
			}
		}

		if len(lsRegexResult) > 0 {
			continue
		}

		if len(fileRegexResult) > 0 {
			// add file to current folder
			fileSize, _ := strconv.Atoi(fileRegexResult[1])
			newFile := File{name: fileRegexResult[2], size: fileSize, extenstion: fileRegexResult[3]}
			currentFolder.files = append(currentFolder.files, newFile)
		}

		if len(dirRegexResult) > 0 {
			newFolder := Directory{name: dirRegexResult[1], parent: currentFolder}
			currentFolder.children = append(currentFolder.children, &newFolder)
		}

	}

	return &root
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

type Directory struct {
	parent   *Directory
	children []*Directory
	name     string
	files    []File
}

type File struct {
	name       string
	size       int
	extenstion string
}

// getSize returns the size of this directory by recursion
func (d Directory) getSize() int {
	totalSize := 0
	for _, file := range d.files {
		totalSize += file.size
	}
	for _, dir := range d.children {
		totalSize += dir.getSize()
	}
	return totalSize
}

// Find size of directory by recursion
func (d Directory) getFolderSizesWithLessthenSize(size int) int {
	sum := 0
	if d.getSize() < size {
		sum += d.getSize()
	}
	for _, child := range d.children {
		sum += child.getFolderSizesWithLessthenSize(size)
	}
	return sum
}

// Find smallest folder that is over size
func (d Directory) findSmallestFolderOverSize(smallest *Directory, overSize int) *Directory {
	currentSize := d.getSize()

	// If this folder is smaller then the smallest folder and bigger then oversize make it new smallest
	if currentSize < smallest.getSize() && currentSize > overSize {
		smallest = &d
	}
	// then, check each folder inside
	for _, child := range d.children {
		smallestChild := child.findSmallestFolderOverSize(smallest, overSize)
		if smallestChild.getSize() < smallest.getSize() && smallest.getSize() > overSize {
			smallest = smallestChild
		}
	}
	return smallest
}

func (d Directory) getChildFolderWithName(name string) *Directory {
	for _, childDir := range d.children {
		if childDir.name == name {
			return childDir
		}
	}
	return &Directory{}
}
