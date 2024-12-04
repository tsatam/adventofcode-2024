package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/tsatam/adventofcode-2024/common/cartesian"
	"github.com/tsatam/adventofcode-2024/common/fp"
)

var (
	//go:embed input
	input    string
	wantWord = []rune("XMAS")
)

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	wordsearch := readInput(input)
	allX := findXs(wordsearch)
	return fp.Sum(fp.Map(allX, checkPoint(wordsearch, 0)))
}

func handlePart2(input string) int {
	wordsearch := readInput(input)
	return findXMASs(wordsearch)
}

func readInput(input string) [][]rune {
	return fp.Map(strings.Split(strings.TrimSpace(input), "\n"), func(s string) []rune { return []rune(s) })
}

func findXs(wordsearch [][]rune) []cartesian.Point {
	result := []cartesian.Point{}

	for y := range wordsearch {
		for x := range wordsearch[y] {
			if wordsearch[y][x] == 'X' {
				result = append(result, cartesian.Point{X: x, Y: y})
			}
		}
	}
	return result
}

func findXMASs(wordsearch [][]rune) int {
	xmases := 0

	for y := 1; y < len(wordsearch)-1; y++ {
		for x := 1; x < len(wordsearch)-1; x++ {
			if wordsearch[y][x] == 'A' {
				if isXMAS(wordsearch, cartesian.Point{X: x, Y: y}) {
					xmases++
				}
			}
		}
	}

	return xmases
}

func isXMAS(wordsearch [][]rune, p cartesian.Point) bool {
	x, y := p.X, p.Y

	corners := []rune{
		wordsearch[y-1][x-1],
		wordsearch[y-1][x+1],
		wordsearch[y+1][x-1],
		wordsearch[y+1][x+1],
	}

	totalM, totalS := 0, 0

	for _, c := range corners {
		switch c {
		case 'M':
			totalM++
		case 'S':
			totalS++
		default:
			return false
		}
	}

	return totalM == 2 && totalS == 2 && corners[0] != corners[3] && corners[1] != corners[2]
}

func checkPoint(wordsearch [][]rune, wordIdx int, direction ...cartesian.Direction) func(cartesian.Point) int {
	return func(p cartesian.Point) int {
		if !isInBounds(wordsearch, p) ||
			wordsearch[p.Y][p.X] != wantWord[wordIdx] {
			return 0
		}

		if wordIdx == len(wantWord)-1 {
			return 1
		}

		if wordIdx == 0 {
			toCheckDirections := [][]cartesian.Direction{
				{cartesian.Left},
				{cartesian.Left, cartesian.Up},
				{cartesian.Up},
				{cartesian.Right, cartesian.Up},
				{cartesian.Right},
				{cartesian.Right, cartesian.Down},
				{cartesian.Down},
				{cartesian.Left, cartesian.Down},
			}

			sum := 0

			for _, direction := range toCheckDirections {
				checker := checkPoint(wordsearch, wordIdx+1, direction...)
				nextPoint := fp.Reduce(direction, p, func(curr cartesian.Point, next cartesian.Direction) cartesian.Point { return curr.Move(next) })

				sum += checker(nextPoint)
			}

			return sum
		}

		checker := checkPoint(wordsearch, wordIdx+1, direction...)
		nextPoint := fp.Reduce(direction, p, func(curr cartesian.Point, next cartesian.Direction) cartesian.Point { return curr.Move(next) })
		return checker(nextPoint)
	}
}

func isInBounds(wordsearch [][]rune, p cartesian.Point) bool {
	return p.X >= 0 &&
		p.Y >= 0 &&
		p.X < len(wordsearch[0]) &&
		p.Y < len(wordsearch)
}
