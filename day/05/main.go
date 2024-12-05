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
	rules, updates := readInput(input)

	validUpdates := fp.Filter(updates, updateIsValid(rules))
	middlePages := fp.Map(validUpdates, middlePageNumber)
	return fp.Sum(middlePages)
}

func handlePart2(input string) int {
	rules, updates := readInput(input)

	invalidUpdates := fp.Filter(updates, func(update []int) bool {
		return !updateIsValid(rules)(update)
	})
	invalidUpdatesSorted := fp.Map(invalidUpdates, func(update []int) []int {
		slices.SortStableFunc(update, sortCmp(rules))
		return update
	})
	middlePages := fp.Map(invalidUpdatesSorted, middlePageNumber)
	return fp.Sum(middlePages)
}

func readInput(input string) (rules map[int][]int, updates [][]int) {
	split := strings.Split(strings.TrimSpace(input), "\n\n")
	rawRules, rawPages := split[0], split[1]

	return readRules(rawRules), readPages(rawPages)
}

func readRules(input string) map[int][]int {
	return fp.Reduce(strings.Split(input, "\n"), map[int][]int{}, func(curr map[int][]int, next string) map[int][]int {
		var before, after int
		if _, err := fmt.Sscanf(next, "%d|%d", &before, &after); err != nil {
			panic(err)
		}
		if _, ok := curr[before]; !ok {
			curr[before] = []int{after}
		} else {
			curr[before] = append(curr[before], after)
		}
		return curr
	})
}

func readPages(input string) [][]int {
	return fp.Map(strings.Split(input, "\n"), func(line string) []int {
		return fp.Map(strings.Split(line, ","), func(elem string) int {
			res, err := strconv.ParseInt(elem, 10, 0)
			if err != nil {
				panic(err)
			}
			return int(res)
		})
	})
}

func updateIsValid(rules map[int][]int) func(update []int) bool {
	cmp := sortCmp(rules)
	return func(update []int) bool {
		return slices.IsSortedFunc(update, cmp)
	}
}

func middlePageNumber(update []int) int {
	return update[len(update)/2]
}

func sortCmp(rules map[int][]int) func(a, b int) int {
	return func(a, b int) int {
		if elems, ok := rules[a]; ok {
			if slices.Contains(elems, b) {
				return -1
			}
		}

		if elems, ok := rules[b]; ok {
			if slices.Contains(elems, a) {
				return 1
			}
		}

		return 0
	}
}
