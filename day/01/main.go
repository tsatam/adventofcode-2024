package main

import (
	_ "embed"
	"fmt"
	"slices"
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
	list1, list2 := parseLists(lines)

	pairs := pairLists(list1, list2)
	distances := getDistances(pairs)

	sum := fp.Sum(distances)

	return sum
}

func handlePart2(input string) int {
	lines := readInput(input)
	list1, list2 := parseLists(lines)

	similarities := similarity(list1, list2)

	sum := fp.Sum(similarities)

	return sum
}

func readInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func parseLists(lines []string) (list1, list2 []int) {
	for _, line := range lines {
		var elem1, elem2 int
		if _, err := fmt.Sscanf(line, "%d   %d", &elem1, &elem2); err != nil {
			panic(err)
		}
		list1 = append(list1, elem1)
		list2 = append(list2, elem2)
	}
	return
}

func pairLists(list1, list2 []int) [][2]int {
	pair := make([][2]int, len(list1))

	slices.Sort(list1)
	slices.Sort(list2)

	for i := range len(list1) {
		pair[i] = [2]int{list1[i], list2[i]}
	}

	return pair
}

func getDistances(pairs [][2]int) []int {
	return fp.Map(pairs, func(in [2]int) int {
		diff := in[1] - in[0]
		if diff < 0 {
			return -diff
		} else {
			return diff
		}
	})
}

func similarity(list1, list2 []int) []int {
	return fp.Map(list1, func(elem int) int {
		return elem * fp.Reduce(list2, 0, func(curr, next int) int {
			switch {
			case next == elem:
				return curr + 1
			default:
				return curr
			}
		})
	})
}
