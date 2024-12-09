package main

import (
	"2024/internal"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := internal.ReadFile("cmd/day_03/input")

	fmt.Println("Part 1, productSum: ", solve(input, false))
	fmt.Println("Part 2, productSum: ", solve(input, true))
}

func solve(input string, filter bool) int {
	var potentialInstructions []string
	addInstruction := true
	for i := 0; i < len(input)-3; i++ {
		if filter {
			if len(input)-i >= 7 && input[i:i+7] == "don't()" {
				// fmt.Println(input[i : i+7])
				addInstruction = false
				continue
			}
			if len(input)-i >= 7 && input[i:i+4] == "do()" {
				// fmt.Println(input[i : i+4])
				addInstruction = true
				continue
			}
		}

		if addInstruction {
			if input[i:i+4] == "mul(" {
				// fmt.Println("found mul instruction opening paren index: ", i)
				for j := i + 4; j < len(input); j++ {
					if input[j] == ')' {
						// fmt.Println("found closing paren index: ", j)
						potentialInstructions = append(potentialInstructions, input[i+4:j+1])
						break
					}
				}

			}
		}
	}

	productSum := 0
	for _, instruction := range potentialInstructions {
		substrings := strings.Split(instruction, ",")
		if len(substrings) != 2 {
			continue
		}
		num1, _ := strconv.Atoi(substrings[0])
		num2, _ := strconv.Atoi(strings.TrimRight(substrings[1], ")"))
		productSum += num1 * num2
	}
	return productSum
}
