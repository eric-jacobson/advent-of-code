package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := os.ReadFile("cmd/day03/input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes.TrimSpace(input)), "\n")
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	// part1, part2 := solve(matrix)
	fmt.Printf("part 1: %v\n", solvePartOne(matrix))
	fmt.Printf("part 2: %v\n", solvePartTwo(matrix))
}

var adjacentPositions = [][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func isSymbol(c rune) bool {
	return !unicode.IsDigit(c) && c != '.'
}

func solvePartOne(matrix [][]rune) int {
	partNumbers := []int{}
	candidate := false

	for r := 0; r < len(matrix); r++ {
		numStr := ""
		for c := 0; c < len(matrix[r]); c++ {
			if unicode.IsDigit(matrix[r][c]) {
				numStr += string(matrix[r][c])
				for _, pos := range adjacentPositions {
					row := r + pos[0]
					col := c + pos[1]
					if row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[r]) && isSymbol(matrix[row][col]) {
						candidate = true
						break
					}
				}
				if candidate && c == len(matrix[r])-1 && len(numStr) > 0 {
					number, err := strconv.Atoi(numStr)
					if err != nil {
						panic(err)
					}
					partNumbers = append(partNumbers, number)
					candidate = false
				}
			} else {
				if candidate && len(numStr) > 0 {
					number, err := strconv.Atoi(numStr)
					if err != nil {
						panic(err)
					}
					partNumbers = append(partNumbers, number)
					candidate = false
				}
				numStr = ""
			}
		}
	}
	sumPartNumbers := 0

	for _, num := range partNumbers {
		sumPartNumbers += num
	}

	return sumPartNumbers
}

func solvePartTwo(matrix [][]rune) int {
	partNumbers := []int{}
	gearRatios := make(map[string][]int)
	candidate := false

	for r := 0; r < len(matrix); r++ {
		numStr := ""
		key := ""
		for c := 0; c < len(matrix[r]); c++ {
			if unicode.IsDigit(matrix[r][c]) {
				numStr += string(matrix[r][c])
				for _, pos := range adjacentPositions {
					row := r + pos[0]
					col := c + pos[1]
					if row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[r]) && matrix[row][col] == '*' {
						key = fmt.Sprintf("%d,%d", row, col)
						if _, ok := gearRatios[key]; !ok {
							gearRatios[key] = []int{}
						}
						candidate = true
						break
					}
				}
				if candidate && c == len(matrix[r])-1 && len(numStr) > 0 {
					number, err := strconv.Atoi(numStr)
					if err != nil {
						panic(err)
					}
					gearRatios[key] = append(gearRatios[key], number)
					candidate = false
				}
			} else {
				if candidate && len(numStr) > 0 {
					number, err := strconv.Atoi(numStr)
					if err != nil {
						panic(err)
					}
					gearRatios[key] = append(gearRatios[key], number)
					candidate = false
				}
				numStr = ""
			}
		}
	}
	sumGearRatios := 0
	for _, num := range partNumbers {
		partNumbers = append(partNumbers, num)
	}

	for _, nums := range gearRatios {
		if len(nums) == 2 {
			sumGearRatios += nums[0] * nums[1]
		}
	}
	return sumGearRatios
}
