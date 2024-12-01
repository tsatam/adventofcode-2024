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

	distances := getDistances(list1, list2)

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

func getDistances(list1, list2 []int) []int {
	slices.Sort(list1)
	slices.Sort(list2)

	return fp.ZipFunc(list1, list2, func(elem1, elem2 int) int {
		return max(elem2-elem1, elem1-elem2)
	})
}

func similarity(list1, list2 []int) []int {
	return fp.Map(list1, func(elem int) int {
		return elem * len(fp.Filter(list2, func(i int) bool { return i == elem }))
	})
}
