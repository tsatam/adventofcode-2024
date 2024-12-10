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

	foundTerminatingPoints map[cartesian.Point][]cartesian.Point = map[cartesian.Point][]cartesian.Point{}
	foundRating            map[cartesian.Point]int               = map[cartesian.Point]int{}
)

type heightMap [][]int

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	hm := parseInput(input)
	trailheads := findTrailheads(hm)
	scores := fp.Map(trailheads, hm.score)

	return fp.Sum(scores)
}

func handlePart2(input string) int {
	hm := parseInput(input)
	trailheads := findTrailheads(hm)
	ratings := fp.Map(trailheads, hm.rating)

	return fp.Sum(ratings)
}

func parseInput(input string) heightMap {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	return fp.Map(lines, func(l string) []int {
		return fp.Map([]rune(l), func(r rune) int { return int(r - '0') })
	})
}

func findTrailheads(hm heightMap) []cartesian.Point {
	result := []cartesian.Point{}
	for y, r := range hm {
		for x, h := range r {
			if h == 0 {
				result = append(result, cartesian.Point{X: x, Y: y})
			}
		}
	}

	return result
}

func (hm heightMap) at(p cartesian.Point) int {
	return hm[p.Y][p.X]
}

func (hm heightMap) inBounds(p cartesian.Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.Y < len(hm) && p.X < len(hm[p.Y])
}

func (hm heightMap) score(p cartesian.Point) int {
	points := hm.terminatingPoints(p)
	return len(points)
}

func (hm heightMap) rating(p cartesian.Point) int {
	if ps, ok := foundRating[p]; ok {
		return ps
	}

	if hm.at(p) == 9 {
		foundRating[p] = 1
		return 1
	}

	toCheck := []cartesian.Point{
		p.Move(cartesian.Up),
		p.Move(cartesian.Down),
		p.Move(cartesian.Left),
		p.Move(cartesian.Right),
	}

	toCheck = fp.Filter(toCheck, hm.inBounds)
	toCheck = fp.Filter(toCheck, func(q cartesian.Point) bool { return hm.at(q) == hm.at(p)+1 })
	ratings := fp.Map(toCheck, hm.rating)
	rating := fp.Sum(ratings)

	foundRating[p] = rating
	return rating
}

func (hm heightMap) terminatingPoints(p cartesian.Point) []cartesian.Point {
	if ps, ok := foundTerminatingPoints[p]; ok {
		return ps
	}

	if hm.at(p) == 9 {
		result := []cartesian.Point{p}
		foundTerminatingPoints[p] = result
		return result
	}

	toCheck := []cartesian.Point{
		p.Move(cartesian.Up),
		p.Move(cartesian.Down),
		p.Move(cartesian.Left),
		p.Move(cartesian.Right),
	}

	toCheck = fp.Filter(toCheck, hm.inBounds)
	toCheck = fp.Filter(toCheck, func(q cartesian.Point) bool { return hm.at(q) == hm.at(p)+1 })
	points := fp.Flatten(fp.Map(toCheck, hm.terminatingPoints))
	points = distinct(points)

	foundTerminatingPoints[p] = points
	return points
}

func distinct[T comparable](in []T) []T {
	set := map[T]interface{}{}
	for _, t := range in {
		set[t] = nil
	}

	return slices.Collect(maps.Keys(set))
}
