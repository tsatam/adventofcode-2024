package main

import (
	_ "embed"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/tsatam/adventofcode-2024/common/cartesian"
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
	lab := readInput(input)

	visited := simulateGuardRoute(lab)

	return len(distinct(visited))
}

func handlePart2(input string) int {
	lab := readInput(input)

	numObstacleAdditions := 0

	for y := range lab.bounds.Y {
		for x := range lab.bounds.X {
			newObstacle := cartesian.Point{X: x, Y: y}
			if !slices.Contains(lab.obstacles, newObstacle) && newObstacle != lab.guard && checkLoop(lab, newObstacle) {
				numObstacleAdditions++
			}
		}
	}

	return numObstacleAdditions
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

func simulateGuardRoute(lab lab) []cartesian.Point {
	var simulateRecursive func(pos cartesian.Point, dir cartesian.Direction) []cartesian.Point

	simulateRecursive = func(pos cartesian.Point, dir cartesian.Direction) []cartesian.Point {
		next := pos.Move(dir)

		if !lab.inBounds(next) {
			return []cartesian.Point{pos}
		}

		if slices.Contains(lab.obstacles, next) {
			dir = rotate[dir]
			next = pos.Move(dir)
		}

		return append(simulateRecursive(next, dir), pos)
	}

	return simulateRecursive(lab.guard, cartesian.Up)
}

func checkLoop(lab lab, newObstacle cartesian.Point) bool {
	lab.obstacles = append(lab.obstacles, newObstacle)

	type collision struct {
		p cartesian.Point
		d cartesian.Direction
	}

	var simulateRecursive func(pos cartesian.Point, dir cartesian.Direction, collisions map[collision]interface{}) bool

	simulateRecursive = func(pos cartesian.Point, dir cartesian.Direction, collisions map[collision]interface{}) bool {
		next := pos.Move(dir)

		if !lab.inBounds(next) {
			return false
		}

		if slices.Contains(lab.obstacles, next) {
			collision := collision{p: pos, d: dir}
			if _, ok := collisions[collision]; ok {
				return true
			} else {
				collisions[collision] = nil
				dir = rotate[dir]
				next = pos.Move(dir)
			}
		}

		return simulateRecursive(next, dir, collisions)
	}

	return simulateRecursive(lab.guard, cartesian.Up, map[collision]interface{}{})
}

func (l lab) inBounds(p cartesian.Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < l.bounds.X && p.Y < l.bounds.Y
}

func distinct[T comparable](in []T) []T {
	set := map[T]interface{}{}
	for _, t := range in {
		set[t] = nil
	}

	return slices.Collect(maps.Keys(set))
}
