package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsatam/adventofcode-2024/common/cartesian"
)

func TestHandlePart1(t *testing.T) {
	input := `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
	`

	want := 14
	got := handlePart1(input)

	assert.Equal(t, want, got)
}

func TestHandlePart2(t *testing.T) {
	input := `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
	`

	want := 34
	got := handlePart2(input)

	assert.Equal(t, want, got)
}

func TestParseInput(t *testing.T) {
	input := `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
	`

	want := CityMap{
		bounds: cartesian.Point{X: 12, Y: 12},
		antennas: map[rune][]cartesian.Point{
			'0': {{X: 8, Y: 1}, {X: 5, Y: 2}, {X: 7, Y: 3}, {X: 4, Y: 4}},
			'A': {{X: 6, Y: 5}, {X: 8, Y: 8}, {X: 9, Y: 9}},
		},
	}

	got := parseInput(input)

	assert.Equal(t, want, got)
}

func TestFindAntinodes(t *testing.T) {
	for _, tt := range []struct {
		name     string
		antennas []cartesian.Point
		want     []cartesian.Point
	}{
		{
			name:     "simple",
			antennas: []cartesian.Point{{X: 1, Y: 1}, {X: 2, Y: 2}},
			want:     []cartesian.Point{{X: 0, Y: 0}, {X: 3, Y: 3}},
		},
		{
			name:     "0",
			antennas: []cartesian.Point{{X: 8, Y: 1}, {X: 5, Y: 2}, {X: 7, Y: 3}, {X: 4, Y: 4}},
			want: []cartesian.Point{
				{X: 11, Y: 0}, {X: 2, Y: 3},
				{X: 9, Y: -1}, {X: 6, Y: 5},
				{X: 12, Y: -2}, {X: 0, Y: 7},
				{X: 3, Y: 1}, {X: 9, Y: 4},
				{X: 3, Y: 6}, {X: 6, Y: 0},
				{X: 1, Y: 5}, {X: 10, Y: 2},
			},
		},
		{
			name:     "A",
			antennas: []cartesian.Point{{X: 6, Y: 5}, {X: 8, Y: 8}, {X: 9, Y: 9}},
			want: []cartesian.Point{
				{X: 4, Y: 2}, {X: 10, Y: 11},
				{X: 3, Y: 1}, {X: 12, Y: 13},
				{X: 7, Y: 7}, {X: 10, Y: 10},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := findAntinodes(tt.antennas)

			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
func TestFindAntinodesPart2(t *testing.T) {
	for _, tt := range []struct {
		name     string
		bounds   cartesian.Point
		antennas []cartesian.Point
		want     []cartesian.Point
	}{
		{
			name:     "simple",
			bounds:   cartesian.Point{X: 5, Y: 5},
			antennas: []cartesian.Point{{X: 1, Y: 1}, {X: 2, Y: 2}},
			want:     []cartesian.Point{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 3}, {X: 4, Y: 4}},
		},
		{
			name:     "T",
			bounds:   cartesian.Point{X: 10, Y: 10},
			antennas: []cartesian.Point{{X: 0, Y: 0}, {X: 3, Y: 1}, {X: 1, Y: 2}},
			want: []cartesian.Point{
				{X: 0, Y: 0}, {X: 3, Y: 1}, {X: 1, Y: 2},
				{X: 6, Y: 2}, {X: 9, Y: 3},
				{X: 2, Y: 4}, {X: 3, Y: 6}, {X: 4, Y: 8},
				{X: 5, Y: 0},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			cityMap := CityMap{bounds: tt.bounds}
			got := cityMap.findAntinodesPart2(tt.antennas)

			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
