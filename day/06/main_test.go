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
		name      string
		lab       lab
		wantRoute []cartesian.Point
		wantLoop  bool
	}{
		{
			name: "1x1 lab",
			lab: lab{
				bounds:    cartesian.Point{X: 1, Y: 1},
				obstacles: []cartesian.Point{},
				guard:     cartesian.Point{X: 1, Y: 1},
			},
			wantRoute: []cartesian.Point{{X: 1, Y: 1}},
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
			wantRoute: []cartesian.Point{
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
		{
			name: "example lab obstacle 1",
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
					{X: 3, Y: 6},
				},
				guard: cartesian.Point{X: 4, Y: 6},
			},
			wantLoop: true,
		},
		{
			name: "example lab obstacle 2",
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
					{X: 6, Y: 7},
				},
				guard: cartesian.Point{X: 4, Y: 6},
			},
			wantLoop: true,
		},
		{
			name: "example lab obstacle 3",
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
					{X: 7, Y: 7},
				},
				guard: cartesian.Point{X: 4, Y: 6},
			},
			wantLoop: true,
		},
		{
			name: "example lab obstacle 4",
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
					{X: 1, Y: 8},
				},
				guard: cartesian.Point{X: 4, Y: 6},
			},
			wantLoop: true,
		},
		{
			name: "example lab obstacle 5",
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
					{X: 3, Y: 8},
				},
				guard: cartesian.Point{X: 4, Y: 6},
			},
			wantLoop: true,
		},
		{
			name: "example lab obstacle 6",
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
					{X: 7, Y: 9},
				},
				guard: cartesian.Point{X: 4, Y: 6},
			},
			wantLoop: true,
		},
		{
			name: "check 1d loop",
			lab: readInput(`
.#.#.#.
#.....#
.#.^.#.
			`),
			wantLoop: true,
		},
		{
			name: "check trapped",
			lab: readInput(`
.#.
#^#
.#.
			`),
			wantLoop: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			gotRoute, gotLoop := simulateGuardRoute(tt.lab)
			if tt.wantRoute != nil {
				assert.ElementsMatch(t, tt.wantRoute, gotRoute)
			}
			assert.Equal(t, tt.wantLoop, gotLoop)
		})
	}
}
