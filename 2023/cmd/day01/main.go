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
	input, err := os.ReadFile("cmd/day01/input")
	if err != nil {
		panic(err)
	}
	input = bytes.TrimSpace(input) // trim the new line to avoid error later

	fmt.Printf("part 1: %v\n", sum(input, false))
	fmt.Printf("part 2: %v\n", sum(input, true))
}

func sum(input []byte, replaceWords bool) int {
	var sum int

	for _, line := range strings.Split(string(input), "\n") {
		if replaceWords {
			line = replace(line)
		}

		var first, last string

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				first = string(line[i])
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				last = string(line[i])
				break
			}
		}

		concat, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}
		sum += concat
	}

	return sum
}

func replace(line string) string {
	replcaements := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	index := len(line)

	for index == len(line) {
		first := ""
		for word := range replcaements {
			pos := strings.Index(line, word)
			if pos != -1 && pos < index {
				index = pos
				first = word
			}
		}
		index = -1
		if _, ok := replcaements[first]; ok {
			line = strings.Replace(line, first[:len(first)-1], replcaements[first], -1)
			index = len(line)
		}
	}

	return line
}
