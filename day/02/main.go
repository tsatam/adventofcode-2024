package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/tsatam/adventofcode-2024/common/fp"
)

var (
	//go:embed input
	input string
)

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	lines := readInput(input)
	reports := parseReports(lines)
	safeReports := fp.Filter(reports, reportSafe)

	return len(safeReports)
}

func handlePart2(input string) int {
	lines := readInput(input)
	reports := parseReports(lines)
	safeReports := fp.Filter(reports, reportSafeDampened)

	return len(safeReports)
}

func readInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func parseReports(lines []string) [][]int {
	result := make([][]int, len(lines))

	for i, line := range lines {
		result[i] = fp.Map(strings.Split(line, " "), func(s string) int { r, _ := strconv.Atoi(s); return r })
	}

	return result
}

func reportSafe(report []int) bool {
	dir := +1
	if report[len(report)-1] < report[0] {
		dir = -1
	}

	for i := range len(report) - 1 {
		diff := (report[i+1] - report[i]) * dir
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func reportSafeDampened(report []int) bool {
	safeAsIs := reportSafe(report)

	if safeAsIs {
		return true
	}

	for i := range report {
		removed := slices.Clone(report)
		removed = slices.Delete(removed, i, i+1)

		if reportSafe(removed) {
			return true
		}
	}

	return false
}
