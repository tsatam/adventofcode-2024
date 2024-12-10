package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsatam/adventofcode-2024/common/cartesian"
)

func TestHandlePart1(t *testing.T) {
	input := `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
	`
	want := 36
	got := handlePart1(input)

	assert.Equal(t, want, got)
}

func TestHandlePart2(t *testing.T) {
	input := `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
	`
	want := 81
	got := handlePart2(input)

	assert.Equal(t, want, got)
}

func TestParseInput(t *testing.T) {
	input := `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
	`
	want := heightMap{
		{8, 9, 0, 1, 0, 1, 2, 3},
		{7, 8, 1, 2, 1, 8, 7, 4},
		{8, 7, 4, 3, 0, 9, 6, 5},
		{9, 6, 5, 4, 9, 8, 7, 4},
		{4, 5, 6, 7, 8, 9, 0, 3},
		{3, 2, 0, 1, 9, 0, 1, 2},
		{0, 1, 3, 2, 9, 8, 0, 1},
		{1, 0, 4, 5, 6, 7, 3, 2},
	}
	got := parseInput(input)
	assert.Equal(t, want, got)
}

func TestFindTrailheads(t *testing.T) {
	heightMap := heightMap{
		{8, 9, 0, 1, 0, 1, 2, 3},
		{7, 8, 1, 2, 1, 8, 7, 4},
		{8, 7, 4, 3, 0, 9, 6, 5},
		{9, 6, 5, 4, 9, 8, 7, 4},
		{4, 5, 6, 7, 8, 9, 0, 3},
		{3, 2, 0, 1, 9, 0, 1, 2},
		{0, 1, 3, 2, 9, 8, 0, 1},
		{1, 0, 4, 5, 6, 7, 3, 2},
	}

	want := []cartesian.Point{
		{X: 2, Y: 0}, {X: 4, Y: 0},
		{X: 4, Y: 2},
		{X: 6, Y: 4},
		{X: 2, Y: 5}, {X: 5, Y: 5},
		{X: 0, Y: 6}, {X: 6, Y: 6},
		{X: 1, Y: 7},
	}

	got := findTrailheads(heightMap)

	assert.ElementsMatch(t, want, got)
}

func TestScore(t *testing.T) {
	example := heightMap{
		{8, 9, 0, 1, 0, 1, 2, 3},
		{7, 8, 1, 2, 1, 8, 7, 4},
		{8, 7, 4, 3, 0, 9, 6, 5},
		{9, 6, 5, 4, 9, 8, 7, 4},
		{4, 5, 6, 7, 8, 9, 0, 3},
		{3, 2, 0, 1, 9, 0, 1, 2},
		{0, 1, 3, 2, 9, 8, 0, 1},
		{1, 0, 4, 5, 6, 7, 3, 2},
	}

	for _, tt := range []struct {
		name string
		hm   heightMap
		p    cartesian.Point
		want int
	}{
		{
			name: "height 9 returns 1",
			hm:   heightMap{{9}},
			p:    cartesian.Point{X: 0, Y: 0},
			want: 1,
		},
		{
			name: "height 8 with 3 surrounding 9s returns 3",
			hm: heightMap{
				{9, 9, 9},
				{9, 8, 9},
				{7, 7, 7},
			},
			p:    cartesian.Point{X: 1, Y: 1},
			want: 3,
		},
		{
			name: "simple trailhead",
			hm: heightMap{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 2, 0, 0, 0},
				{6, 5, 4, 3, 4, 5, 6},
				{7, 0, 0, 0, 0, 0, 7},
				{8, 0, 0, 0, 0, 0, 8},
				{9, 0, 0, 0, 0, 0, 9},
			},
			p:    cartesian.Point{X: 3, Y: 0},
			want: 2,
		},
		{
			name: "trailhead 0",
			hm:   example,
			p:    cartesian.Point{X: 2, Y: 0},
			want: 5,
		},
		{
			name: "trailhead 1",
			hm:   example,
			p:    cartesian.Point{X: 4, Y: 0},
			want: 6,
		},
		{
			name: "trailhead 2",
			hm:   example,
			p:    cartesian.Point{X: 4, Y: 2},
			want: 5,
		},
		{
			name: "trailhead 3",
			hm:   example,
			p:    cartesian.Point{X: 6, Y: 4},
			want: 3,
		},
		{
			name: "trailhead 4",
			hm:   example,
			p:    cartesian.Point{X: 2, Y: 5},
			want: 1,
		},
		{
			name: "trailhead 5",
			hm:   example,
			p:    cartesian.Point{X: 5, Y: 5},
			want: 3,
		},
		{
			name: "trailhead 6",
			hm:   example,
			p:    cartesian.Point{X: 0, Y: 6},
			want: 5,
		},
		{
			name: "trailhead 7",
			hm:   example,
			p:    cartesian.Point{X: 6, Y: 6},
			want: 3,
		},
		{
			name: "trailhead 8",
			hm:   example,
			p:    cartesian.Point{X: 1, Y: 7},
			want: 5,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			for p := range foundTerminatingPoints {
				delete(foundTerminatingPoints, p)
			}

			got := tt.hm.score(tt.p)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRating(t *testing.T) {
	example := heightMap{
		{8, 9, 0, 1, 0, 1, 2, 3},
		{7, 8, 1, 2, 1, 8, 7, 4},
		{8, 7, 4, 3, 0, 9, 6, 5},
		{9, 6, 5, 4, 9, 8, 7, 4},
		{4, 5, 6, 7, 8, 9, 0, 3},
		{3, 2, 0, 1, 9, 0, 1, 2},
		{0, 1, 3, 2, 9, 8, 0, 1},
		{1, 0, 4, 5, 6, 7, 3, 2},
	}

	for _, tt := range []struct {
		name string
		hm   heightMap
		p    cartesian.Point
		want int
	}{
		{
			name: "height 9 returns 1",
			hm:   heightMap{{9}},
			p:    cartesian.Point{X: 0, Y: 0},
			want: 1,
		},
		{
			name: "height 8 with 3 surrounding 9s returns 3",
			hm: heightMap{
				{9, 9, 9},
				{9, 8, 9},
				{7, 7, 7},
			},
			p:    cartesian.Point{X: 1, Y: 1},
			want: 3,
		},
		{
			name: "simple trailhead",
			hm: heightMap{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 2, 0, 0, 0},
				{6, 5, 4, 3, 4, 5, 6},
				{7, 0, 0, 0, 0, 0, 7},
				{8, 0, 0, 0, 0, 0, 8},
				{9, 0, 0, 0, 0, 0, 9},
			},
			p:    cartesian.Point{X: 3, Y: 0},
			want: 2,
		},
		{
			name: "trailhead 0",
			hm:   example,
			p:    cartesian.Point{X: 2, Y: 0},
			want: 20,
		},
		{
			name: "trailhead 1",
			hm:   example,
			p:    cartesian.Point{X: 4, Y: 0},
			want: 24,
		},
		{
			name: "trailhead 2",
			hm:   example,
			p:    cartesian.Point{X: 4, Y: 2},
			want: 10,
		},
		{
			name: "trailhead 3",
			hm:   example,
			p:    cartesian.Point{X: 6, Y: 4},
			want: 4,
		},
		{
			name: "trailhead 4",
			hm:   example,
			p:    cartesian.Point{X: 2, Y: 5},
			want: 1,
		},
		{
			name: "trailhead 5",
			hm:   example,
			p:    cartesian.Point{X: 5, Y: 5},
			want: 4,
		},
		{
			name: "trailhead 6",
			hm:   example,
			p:    cartesian.Point{X: 0, Y: 6},
			want: 5,
		},
		{
			name: "trailhead 7",
			hm:   example,
			p:    cartesian.Point{X: 6, Y: 6},
			want: 8,
		},
		{
			name: "trailhead 8",
			hm:   example,
			p:    cartesian.Point{X: 1, Y: 7},
			want: 5,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			for p := range foundRating {
				delete(foundRating, p)
			}

			got := tt.hm.rating(tt.p)
			assert.Equal(t, tt.want, got)
		})
	}
}
