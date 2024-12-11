package main

import (
	_ "embed"
	"fmt"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/tsatam/adventofcode-2024/common/fp"
)

var (
	//go:embed input
	input string

	stoneCache map[uint64][]uint64 = map[uint64][]uint64{}
)

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	stones := parseInput(input)
	sm := toMap(stones)
	for range 25 {
		sm = sm.iterate()
	}
	return fp.Sum(slices.Collect(maps.Values(sm)))
}

func handlePart2(input string) int {
	stones := parseInput(input)
	sm := toMap(stones)
	for range 75 {
		sm = sm.iterate()
	}
	return fp.Sum(slices.Collect(maps.Values(sm)))
}

func parseInput(input string) []uint64 {
	return fp.Map(strings.Split(strings.TrimSpace(input), " "), func(s string) uint64 {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return uint64(n)
	})
}

type stonesMap map[uint64]int

func toMap(stones []uint64) stonesMap {
	s := stonesMap{}
	for _, stone := range stones {
		s.add(stone, 1)
	}
	return s
}

func (s stonesMap) add(stone uint64, count int) {
	if _, ok := s[stone]; !ok {
		s[stone] = 0
	}
	s[stone] += count
}

func (s stonesMap) iterate() stonesMap {
	result := stonesMap{}
	for stone, count := range s {
		newStones := processStoneCached(stone)
		for _, newStone := range newStones {
			result.add(newStone, count)
		}
	}
	return result
}

func processStoneCached(stone uint64) []uint64 {
	if result, ok := stoneCache[stone]; ok {
		return result
	} else {
		result := processStone(stone)
		stoneCache[stone] = result
		return result
	}
}

func processStone(stone uint64) []uint64 {
	digits := numDigits(stone)
	switch {
	case stone == 0:
		return []uint64{1}

	case digits%2 == 0:
		split := uint64(math.Pow10(digits / 2))
		right := stone % split
		left := stone / split

		return []uint64{left, right}
	default:
		return []uint64{2024 * stone}

	}
}

func numDigits(n uint64) int {
	switch {
	case n == 0:
		return 0
	default:
		return int(math.Log10(float64(n))) + 1
	}
}
