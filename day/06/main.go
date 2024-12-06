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

	rotate = map[cartesian.Direction]cartesian.Direction{
		cartesian.Up:    cartesian.Right,
		cartesian.Right: cartesian.Down,
		cartesian.Down:  cartesian.Left,
		cartesian.Left:  cartesian.Up,
	}
)

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	l := readInput(input)

	visited, _ := simulateGuardRoute(l)

	return len(distinct(visited))
}

func handlePart2(input string) int {
	l := readInput(input)

	a, _ := simulateGuardRoute(l)
	a = distinct(a)
	a = slices.DeleteFunc(a, l.isGuard)
	a = fp.Filter(a, l.doesCauseLoop)

	return len(a)
}

type lab struct {
	bounds    cartesian.Point
	obstacles []cartesian.Point
	guard     cartesian.Point
}

func readInput(input string) lab {
	split := strings.Split(strings.TrimSpace(input), "\n")
	if len(split) == 0 {
		return lab{}
	}

	bounds := cartesian.Point{X: len(split[0]), Y: len(split)}
	obstacles := []cartesian.Point{}
	var guard cartesian.Point

	for y, row := range split {
		for x, char := range []rune(row) {
			switch char {
			case '#':
				obstacles = append(obstacles, cartesian.Point{X: x, Y: y})
			case '^':
				guard = cartesian.Point{X: x, Y: y}
			}
		}
	}

	return lab{bounds: bounds, obstacles: obstacles, guard: guard}
}

func simulateGuardRoute(l lab) ([]cartesian.Point, bool) {
	type collision struct {
		p cartesian.Point
		d cartesian.Direction
	}
	collisions := map[collision]interface{}{}

	var simulateRecursive func(pos cartesian.Point, dir cartesian.Direction) ([]cartesian.Point, bool)

	simulateRecursive = func(pos cartesian.Point, dir cartesian.Direction) ([]cartesian.Point, bool) {
		next := pos.Move(dir)

		if !l.inBounds(next) {
			return []cartesian.Point{pos}, false
		}

		if slices.Contains(l.obstacles, next) {
			collision := collision{p: pos, d: dir}
			if _, ok := collisions[collision]; ok {
				return []cartesian.Point{pos}, true
			} else {
				collisions[collision] = nil
				return simulateRecursive(pos, rotate[dir])
			}
		}

		resultPoints, isLoop := simulateRecursive(next, dir)
		return append(resultPoints, pos), isLoop
	}

	return simulateRecursive(l.guard, cartesian.Up)
}

func (l lab) inBounds(p cartesian.Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < l.bounds.X && p.Y < l.bounds.Y
}

func (l lab) isObstacle(p cartesian.Point) bool {
	return slices.Contains(l.obstacles, p)
}

func (l lab) isGuard(p cartesian.Point) bool {
	return p == l.guard
}

func (l lab) doesCauseLoop(p cartesian.Point) bool {
	_, isLoop := simulateGuardRoute(lab{bounds: l.bounds, guard: l.guard, obstacles: append(l.obstacles, p)})
	return isLoop
}

func distinct[T comparable](in []T) []T {
	set := map[T]interface{}{}
	for _, t := range in {
		set[t] = nil
	}

	return slices.Collect(maps.Keys(set))
}
