package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlePart1(t *testing.T) {
	input := `
2333133121414131402
	`

	want := 1928
	got := handlePart1(input)

	assert.Equal(t, want, got)
}

func TestHandlePart2(t *testing.T) {
	input := `
2333133121414131402
	`

	want := 2858
	got := handlePart2(input)

	assert.Equal(t, want, got)
}

func TestParseDiskMap(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  []File
	}{
		{

			input: "12345",
			want: []File{
				{b: 0, s: 1}, {b: empty, s: 2},
				{b: 1, s: 3}, {b: empty, s: 4},
				{b: 2, s: 5},
			},
		},
		{
			input: "2333133121414131402",
			want: []File{
				{b: 0, s: 2}, {b: empty, s: 3},
				{b: 1, s: 3}, {b: empty, s: 3},
				{b: 2, s: 1}, {b: empty, s: 3},
				{b: 3, s: 3}, {b: empty, s: 1},
				{b: 4, s: 2}, {b: empty, s: 1},
				{b: 5, s: 4}, {b: empty, s: 1},
				{b: 6, s: 4}, {b: empty, s: 1},
				{b: 7, s: 3}, {b: empty, s: 1},
				{b: 8, s: 4},
				{b: 9, s: 2},
			},
		},
	} {
		t.Run(tt.input, func(t *testing.T) {
			got := parseDiskMap(tt.input)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFilesToBlocks(t *testing.T) {
	for _, tt := range []struct {
		name  string
		files []File
		want  []Block
	}{
		{

			name: "12345",
			files: []File{
				{b: 0, s: 1}, {b: empty, s: 2},
				{b: 1, s: 3}, {b: empty, s: 4},
				{b: 2, s: 5},
			},
			want: []Block{0, empty, empty, 1, 1, 1, empty, empty, empty, empty, 2, 2, 2, 2, 2},
		},
		{
			name: "2333133121414131402",
			files: []File{
				{b: 0, s: 2}, {b: empty, s: 3},
				{b: 1, s: 3}, {b: empty, s: 3},
				{b: 2, s: 1}, {b: empty, s: 3},
				{b: 3, s: 3}, {b: empty, s: 1},
				{b: 4, s: 2}, {b: empty, s: 1},
				{b: 5, s: 4}, {b: empty, s: 1},
				{b: 6, s: 4}, {b: empty, s: 1},
				{b: 7, s: 3}, {b: empty, s: 1},
				{b: 8, s: 4},
				{b: 9, s: 2},
			},
			want: []Block{0, 0, empty, empty, empty, 1, 1, 1, empty, empty, empty, 2, empty, empty, empty, 3, 3, 3, empty, 4, 4, empty, 5, 5, 5, 5, empty, 6, 6, 6, 6, empty, 7, 7, 7, empty, 8, 8, 8, 8, 9, 9},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := fileToBlocks(tt.files)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCompactByBlock(t *testing.T) {
	for _, tt := range []struct {
		name         string
		blocks, want []Block
	}{
		{
			name:   "0..111....22222 -> 022111222",
			blocks: []Block{0, empty, empty, 1, 1, 1, empty, empty, empty, empty, 2, 2, 2, 2, 2},
			want:   []Block{0, 2, 2, 1, 1, 1, 2, 2, 2},
		},
		{
			name:   "00...111...2...333.44.5555.6666.777.888899 -> 0099811188827773336446555566",
			blocks: []Block{0, 0, empty, empty, empty, 1, 1, 1, empty, empty, empty, 2, empty, empty, empty, 3, 3, 3, empty, 4, 4, empty, 5, 5, 5, 5, empty, 6, 6, 6, 6, empty, 7, 7, 7, empty, 8, 8, 8, 8, 9, 9},
			want:   []Block{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := compactByBlock(tt.blocks)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCompactByFile(t *testing.T) {
	for _, tt := range []struct {
		name        string
		files, want []File
	}{
		{
			name:  "0..111....22222 -> 0..111....22222",
			files: []File{{b: 0, s: 1}, {b: empty, s: 2}, {b: 1, s: 3}, {b: empty, s: 4}, {b: 2, s: 5}},
			want:  []File{{b: 0, s: 1}, {b: empty, s: 2}, {b: 1, s: 3}, {b: empty, s: 4}, {b: 2, s: 5}},
		},
		{
			name: "00...111...2...333.44.5555.6666.777.888899 -> 00992111777.44.333....5555.6666.....8888..",
			files: []File{
				{b: 0, s: 2}, {b: empty, s: 3},
				{b: 1, s: 3}, {b: empty, s: 3},
				{b: 2, s: 1}, {b: empty, s: 3},
				{b: 3, s: 3}, {b: empty, s: 1},
				{b: 4, s: 2}, {b: empty, s: 1},
				{b: 5, s: 4}, {b: empty, s: 1},
				{b: 6, s: 4}, {b: empty, s: 1},
				{b: 7, s: 3}, {b: empty, s: 1},
				{b: 8, s: 4},
				{b: 9, s: 2},
			},
			want: []File{
				{b: 0, s: 2}, {b: 9, s: 2}, {b: 2, s: 1},
				{b: 1, s: 3}, {b: 7, s: 3}, {b: empty, s: 1},
				{b: 4, s: 2}, {b: empty, s: 1},
				{b: 3, s: 3}, {b: empty, s: 4},
				{b: 5, s: 4}, {b: empty, s: 1},
				{b: 6, s: 4}, {b: empty, s: 5},
				{b: 8, s: 4}, {b: empty, s: 2},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := compactByFile(tt.files)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestChecksum(t *testing.T) {
	for _, tt := range []struct {
		name   string
		blocks []Block
		want   int
	}{
		{
			name:   "022111222 -> 60",
			blocks: []Block{0, 2, 2, 1, 1, 1, 2, 2, 2},
			want:   60,
		},
		{
			name:   "0099811188827773336446555566 -> 1928",
			blocks: []Block{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6},
			want:   1928,
		},
		{
			name:   "00992111777.44.333....5555.6666.....8888.. -> 2858",
			blocks: []Block{0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, empty, 4, 4, empty, 3, 3, 3, empty, empty, empty, empty, 5, 5, 5, 5, empty, 6, 6, 6, 6, empty, empty, empty, empty, empty, 8, 8, 8, 8, empty, empty},
			want:   2858,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := checksum(tt.blocks)
			assert.Equal(t, tt.want, got)
		})
	}
}
