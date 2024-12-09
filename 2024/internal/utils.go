package internal

import (
	"os"
	"strings"
)

func ReadFile(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.TrimRight(string(input), "\n")
}

func AbsouluteValue(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
