package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("cmd/day_01/input")
	if err != nil {
		panic(err)
	}

	var leftColumn []int
	var rightColumn []int
	rightColumnFrequency := make(map[int]int)

	for _, line := range strings.Split(string(input), "\n") {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			left, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			right, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			rightColumnFrequency[right]++
			leftColumn = append(leftColumn, left)
			rightColumn = append(rightColumn, right)
		}
	}
	sort.Ints(leftColumn)
	sort.Ints(rightColumn)

	var totalDistance, adjustedTotalDistance int
	// an assumption is made here that the length of the two columns are the same
	for i, _ := range leftColumn {
		totalDistance += absouluteValue(leftColumn[i] - rightColumn[i])
		adjustedTotalDistance += leftColumn[i] * rightColumnFrequency[leftColumn[i]]
	}

	fmt.Println("Part 1, total distance: ", totalDistance)
	fmt.Println("Part 2, adjusted total distance: ", adjustedTotalDistance)
}

func absouluteValue(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
