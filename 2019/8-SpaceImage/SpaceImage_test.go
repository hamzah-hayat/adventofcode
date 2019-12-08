package main

import (
	"fmt"
	"testing"
)

func TestSpaceImage_small(t *testing.T) {

	input := []int{1, 1, 1, 0, 1, 1, 1, 2, 0, 0, 1, 2}
	expected := 0

	width := 3
	height := 2

	pixels, maxLayers := buildPixelMap(input, height, width)

	lowestLayer := findLayerWithLowestNumber(pixels, 0, maxLayers)

	num1 := findNumberCountInLayer(pixels, 1, lowestLayer)
	num2 := findNumberCountInLayer(pixels, 2, lowestLayer)

	if num1*num2 != expected {
		t.Error(fmt.Sprint("Expected output ", expected, " but got ", num1*num2))
	}

}
