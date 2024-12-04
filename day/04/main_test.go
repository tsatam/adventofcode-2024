package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsatam/adventofcode-2024/common/cartesian"
)

func TestHandlePart1(t *testing.T) {
	input := `....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX`

	want := 18
	got := handlePart1(input)
	assert.Equal(t, want, got)
}

func TestHandlePart2(t *testing.T) {
	input := `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`

	want := 9
	got := handlePart2(input)
	assert.Equal(t, want, got)
}

func TestFindXs(t *testing.T) {
	wordsearch := readInput(`..X...
.SAMX.
.A..A.
XMAS.S
.X....`)
	want := []cartesian.Point{
		{X: 2, Y: 0},
		{X: 4, Y: 1},
		{X: 0, Y: 3},
		{X: 1, Y: 4},
	}
	got := findXs(wordsearch)
	assert.ElementsMatch(t, want, got)
}
