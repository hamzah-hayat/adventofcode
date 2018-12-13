package main

import (
	"bufio"
	"os"
)

func main() {

}

//Read data from device.txt
//Load this data into a int array
func readDevice() {
	f, err := os.Open("device.txt")
	reader := bufio.NewReader(f)
}
