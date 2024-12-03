package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"

	"github.com/tsatam/adventofcode-2024/common/fp"
)

const (
	enabledInstruction  = "do()"
	disabledInstruction = "don't()"
)

var (
	//go:embed input
	input string

	mulRegex = regexp.MustCompile(`mul\(\d+,\d+\)`)
)

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	allMuls := mulRegex.FindAllString(input, -1)
	calculated := fp.Map(allMuls, calcMul)
	return fp.Sum(calculated)
}

func handlePart2(input string) int {
	enabledInstructions := getEnabledInstructions(input)
	filteredInput := strings.Join(enabledInstructions, "")
	return handlePart1(filteredInput)
}

func calcMul(instruction string) int {
	var a, b int
	if _, err := fmt.Sscanf(instruction, "mul(%d,%d)", &a, &b); err != nil {
		panic(err)
	}

	return a * b
}

func getEnabledInstructions(input string) []string {
	firstDontIndex := strings.Index(input, disabledInstruction)
	if firstDontIndex == -1 {
		return []string{input}
	}

	lines := []string{input[:firstDontIndex]}
	rest := input[firstDontIndex+len(disabledInstruction):]

	firstDoIndex := strings.Index(rest, enabledInstruction)
	if firstDoIndex == -1 {
		return lines
	}
	return append(lines, getEnabledInstructions(rest[firstDoIndex+len(enabledInstruction):])...)
}
