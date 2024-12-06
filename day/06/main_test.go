package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsatam/adventofcode-2024/common/cartesian"
)

func TestHandlePart1(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	want := 41
	got := handlePart1(input)
	assert.Equal(t, want, got)
}
func TestHandlePart2(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	want := 6
	got := handlePart2(input)
	assert.Equal(t, want, got)
}

func TestReadInput(t *testing.T) {
	input := `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

	want := lab{
		bounds: cartesian.Point{X: 10, Y: 10},
		obstacles: []cartesian.Point{
			{X: 4, Y: 0},
			{X: 9, Y: 1},
			{X: 2, Y: 3},
			{X: 7, Y: 4},
			{X: 1, Y: 6},
			{X: 8, Y: 7},
			{X: 0, Y: 8},
			{X: 6, Y: 9},
		},
		guard: cartesian.Point{X: 4, Y: 6},
	}
	got := readInput(input)
	assert.Equal(t, want, got)
}

func TestSimulateGuardRoute(t *testing.T) {
	for _, tt := range []struct {
		name string
		lab  lab
		want []cartesian.Point
	}{
		{
			name: "1x1 lab",
			lab: lab{
				bounds:    cartesian.Point{X: 1, Y: 1},
				obstacles: []cartesian.Point{},
				guard:     cartesian.Point{X: 1, Y: 1},
			},
			want: []cartesian.Point{{X: 1, Y: 1}},
		},
		{
			name: "example lab",
			lab: lab{
				bounds: cartesian.Point{X: 10, Y: 10},
				obstacles: []cartesian.Point{
					{X: 4, Y: 0},
					{X: 9, Y: 1},
					{X: 2, Y: 3},
					{X: 7, Y: 4},
					{X: 1, Y: 6},
					{X: 8, Y: 7},
					{X: 0, Y: 8},
					{X: 6, Y: 9},
				},
				guard: cartesian.Point{X: 4, Y: 6},
			},
			want: []cartesian.Point{
				{X: 4, Y: 6}, {X: 4, Y: 5}, {X: 4, Y: 4}, {X: 4, Y: 3}, {X: 4, Y: 2}, {X: 4, Y: 1},
				{X: 5, Y: 1}, {X: 6, Y: 1}, {X: 7, Y: 1}, {X: 8, Y: 1},
				{X: 8, Y: 2}, {X: 8, Y: 3}, {X: 8, Y: 4}, {X: 8, Y: 5}, {X: 8, Y: 6},
				{X: 7, Y: 6}, {X: 6, Y: 6}, {X: 5, Y: 6}, {X: 4, Y: 6}, {X: 3, Y: 6}, {X: 2, Y: 6},
				{X: 2, Y: 5}, {X: 2, Y: 4},
				{X: 3, Y: 4}, {X: 4, Y: 4}, {X: 5, Y: 4}, {X: 6, Y: 4},
				{X: 6, Y: 5}, {X: 6, Y: 6}, {X: 6, Y: 7}, {X: 6, Y: 8},
				{X: 5, Y: 8}, {X: 4, Y: 8}, {X: 3, Y: 8}, {X: 2, Y: 8}, {X: 1, Y: 8},
				{X: 1, Y: 7},
				{X: 2, Y: 7}, {X: 3, Y: 7}, {X: 4, Y: 7}, {X: 5, Y: 7}, {X: 6, Y: 7}, {X: 7, Y: 7},
				{X: 7, Y: 8}, {X: 7, Y: 9},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := simulateGuardRoute(tt.lab)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
