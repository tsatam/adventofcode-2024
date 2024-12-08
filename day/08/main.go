package main

import (
	_ "embed"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/tsatam/adventofcode-2024/common/cartesian"
	"github.com/tsatam/adventofcode-2024/common/fp"
)

var (
	//go:embed input
	input string
)

type CityMap struct {
	bounds   cartesian.Point
	antennas map[rune][]cartesian.Point
}

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	cityMap := parseInput(input)
	allAntennas := slices.Collect(maps.Values(cityMap.antennas))
	allAntinodesForAntennas := fp.Map(allAntennas, findAntinodes)
	antinodes := fp.Flatten(allAntinodesForAntennas)
	antinodes = distinct(antinodes)
	antinodes = fp.Filter(antinodes, cityMap.inBounds)

	return len(antinodes)
}

func handlePart2(input string) int {
	cityMap := parseInput(input)
	allAntennas := slices.Collect(maps.Values(cityMap.antennas))
	allAntinodesForAntennas := fp.Map(allAntennas, cityMap.findAntinodesPart2)
	antinodes := fp.Flatten(allAntinodesForAntennas)
	antinodes = distinct(antinodes)
	antinodes = fp.Filter(antinodes, cityMap.inBounds)

	return len(antinodes)
}

func parseInput(input string) CityMap {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	bounds := cartesian.Point{X: len(lines[0]), Y: len(lines)}

	antennas := map[rune][]cartesian.Point{}

	for y, line := range lines {
		for x, c := range []rune(line) {
			switch c {
			case '.':
				continue
			default:
				if _, ok := antennas[c]; !ok {
					antennas[c] = []cartesian.Point{}
				}
				antennas[c] = append(antennas[c], cartesian.Point{X: x, Y: y})
			}
		}
	}

	return CityMap{bounds: bounds, antennas: antennas}
}

func findAntinodes(antennas []cartesian.Point) []cartesian.Point {
	result := []cartesian.Point{}

	for aidx, a := range antennas[:len(antennas)-1] {
		for _, b := range antennas[aidx+1:] {
			result = append(result, getAntinodeForPoints(a, b), getAntinodeForPoints(b, a))
		}
	}

	return result
}

func (m CityMap) findAntinodesPart2(antennas []cartesian.Point) []cartesian.Point {
	result := []cartesian.Point{}

	for aidx, a := range antennas[:len(antennas)-1] {
		for _, b := range antennas[aidx+1:] {
			diffA := a.Diff(b)
			for antinode := a.Add(diffA); m.inBounds(antinode); antinode = antinode.Add(diffA) {
				result = append(result, antinode)
			}

			diffB := b.Diff(a)
			for antinode := b.Add(diffB); m.inBounds(antinode); antinode = antinode.Add(diffB) {
				result = append(result, antinode)
			}
		}
	}

	return distinct(result)

}

func (m CityMap) inBounds(p cartesian.Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < m.bounds.X && p.Y < m.bounds.Y
}

func getAntinodeForPoints(a, b cartesian.Point) cartesian.Point {
	return cartesian.Point{
		X: a.X - (b.X - a.X),
		Y: a.Y - (b.Y - a.Y),
	}
}

func distinct[T comparable](in []T) []T {
	set := map[T]interface{}{}
	for _, t := range in {
		set[t] = nil
	}

	return slices.Collect(maps.Keys(set))
}
