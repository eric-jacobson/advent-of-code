package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const RED int = 12
const GREEN int = 13
const BLUE int = 14

type result struct {
	sumIds    int
	sumPowers int
}

func main() {
	input, err := os.ReadFile("cmd/day02/input")
	if err != nil {
		panic(err)
	}
	input = bytes.TrimSpace(input)

	ids, powers := solve(input)

	fmt.Printf("part 1: %v\n", ids)
	fmt.Printf("part 2: %v\n", powers)
}

func solve(input []byte) (int, int) {
	sumIds, sumPowers := 0, 0

	for _, line := range strings.Split(string(input), "\n") {
		red, green, blue := 0, 0, 0
		eligible := true
		id, err := strconv.Atoi(line[5:strings.Index(line, ":")])
		if err != nil {
			panic(err)
		}
		colors := strings.Split(line[strings.Index(line, ":")+2:], "; ")
		for _, color := range colors {
			values := strings.Split(color, ", ")
			for _, value := range values {
				count, err := strconv.Atoi(value[:strings.Index(value, " ")])
				if err != nil {
					panic(err)
				}
				if strings.Contains(value, "red") {
					if count > RED {
						eligible = false
					}
					if count > red {
						red = count
					}
				}
				if strings.Contains(value, "green") {
					if count > GREEN {
						eligible = false
					}
					if count > green {
						green = count
					}
				}
				if strings.Contains(value, "blue") {
					if count > BLUE {
						eligible = false
					}
					if count > blue {
						blue = count
					}
				}
			}
		}
		if eligible {
			sumIds += id
		}
		sumPowers += (red * green * blue)
	}
	return sumIds, sumPowers
}
