package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("cmd/day04/input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes.TrimSpace(input)), "\n")

	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		numbers := strings.Split(line[strings.Index(line, ":")+1:], "|")
		regex := regexp.MustCompile("[0-9]+")
		left := regex.FindAllString(numbers[0], -1)
		right := regex.FindAllString(numbers[1], -1)
		// fmt.Printf("left: %v - right: %v\n", left, right)
		lineSum := 0
		for _, num := range right {
			if slices.Contains(left, num) {
				if lineSum == 0 {
					lineSum++
				} else {
					lineSum *= 2
				}
			}
		}
		sum += lineSum
	}

	return sum
}

func part2(lines []string) int {
	idMap := make(map[int]int)
	queue := make([]int, 0)
	sum := 0

	for _, line := range lines {
		id, err := strconv.Atoi(strings.TrimSpace(line[4:strings.Index(line, ":")]))
		if err != nil {
			panic(err)
		}

		numbers := strings.Split(line[strings.Index(line, ":")+1:], "|")
		regex := regexp.MustCompile("[0-9]+")
		left := regex.FindAllString(numbers[0], -1)
		right := regex.FindAllString(numbers[1], -1)

		matches := 0
		for _, num := range right {
			if slices.Contains(left, num) {
				matches++
			}
		}

		idMap[id] = matches
		queue = append(queue, id)
	}

	for len(queue) > 0 {
		sum++
		temp := queue[0]
		queue = queue[1:]
		for i := temp + 1; i <= temp+idMap[temp]; i++ {
			queue = append(queue, i)
		}
	}

	return sum
}
