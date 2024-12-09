package main

import (
	"2024/internal"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := internal.ReadFile("cmd/day_02/input")

	fmt.Println("Part 1, totalSafeReports: ", solve(input, false))
	fmt.Println("Part 2, totalSafeReportsDampened: ", solve(input, true))
}

func solve(input string, dampen bool) int {
	totalSafeReports := 0

	for _, line := range strings.Split(input, "\n") {
		values := strings.Fields(line)
		if validateSafety(values) { // part 1
			totalSafeReports++
		} else {
			if dampen { // part 2
				for i := 0; i < len(values); i++ {
					dampenedValues := append([]string{}, values[:i]...)
					dampenedValues = append(dampenedValues, values[i+1:]...)
					if validateSafety(dampenedValues) {
						totalSafeReports++
						break
					}
				}
			}
		}
	}
	return totalSafeReports
}

func validateSafety(values []string) bool {
	increasing, decreasing := false, false
	for i := 1; i < len(values); i++ {
		current, _ := strconv.Atoi(values[i])
		previous, _ := strconv.Atoi(values[i-1])

		delta := absouluteValue(current - previous)
		if (current > previous && decreasing) || (current < previous && increasing) || current == previous || !(delta >= 1 && delta <= 3) {
			return false
		}

		if current > previous {
			increasing = true
		} else if current < previous {
			decreasing = true
		}
	}
	return true
}

func absouluteValue(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
