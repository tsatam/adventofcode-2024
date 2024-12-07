package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/tsatam/adventofcode-2024/common/fp"
)

var (
	//go:embed input
	input string
)

type Equation struct {
	testValue uint64
	numbers   []uint64
}

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) uint64 {
	equations := parseInput(input)
	validEquations := fp.Filter(equations, func(e Equation) bool { return e.isValid(false) })
	validTestValues := fp.Map(validEquations, func(e Equation) uint64 { return e.testValue })
	return fp.Sum(validTestValues)
}

func handlePart2(input string) uint64 {
	equations := parseInput(input)
	validEquations := fp.Filter(equations, func(e Equation) bool { return e.isValid(true) })
	validTestValues := fp.Map(validEquations, func(e Equation) uint64 { return e.testValue })
	return fp.Sum(validTestValues)
}

func parseInput(input string) []Equation {
	return fp.Map(strings.Split(strings.TrimSpace(input), "\n"), func(line string) Equation {
		split := strings.Split(line, ":")

		return Equation{
			testValue: parseInt(split[0]),
			numbers:   fp.Map(strings.Split(strings.TrimSpace(split[1]), " "), parseInt),
		}
	})
}

func parseInt(s string) uint64 {
	res, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func (e Equation) isValid(shouldConcat bool) bool {
	return isValid(e.testValue, e.numbers, shouldConcat)
}

func isValid(testValue uint64, numbers []uint64, shouldConcat bool) bool {
	if len(numbers) == 1 {
		return numbers[0] == testValue
	}

	if numbers[0] > testValue {
		return false
	}

	addResult := numbers[0] + numbers[1]
	mulResult := numbers[0] * numbers[1]
	concatResult := concat(numbers[0], numbers[1])

	return isValid(testValue, append([]uint64{addResult}, numbers[2:]...), shouldConcat) ||
		isValid(testValue, append([]uint64{mulResult}, numbers[2:]...), shouldConcat) ||
		(shouldConcat && isValid(testValue, append([]uint64{concatResult}, numbers[2:]...), shouldConcat))
}

func concat(a, b uint64) uint64 {
	return parseInt(strconv.FormatUint(a, 10) + strconv.FormatUint(b, 10))
}
