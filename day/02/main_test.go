package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlePart1(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	want := 2
	got := handlePart1(input)

	assert.Equal(t, want, got)
}

func TestHandlePart2(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	want := 4
	got := handlePart2(input)

	assert.Equal(t, want, got)
}

func TestParseReports(t *testing.T) {
	for i, tt := range []struct {
		lines []string
		want  [][]int
	}{
		{
			lines: []string{},
			want:  [][]int{},
		},
		{
			lines: []string{"1 2"},
			want:  [][]int{{1, 2}},
		},
		{
			lines: []string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			want: [][]int{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := parseReports(tt.lines)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReportSafe(t *testing.T) {
	for i, tt := range []struct {
		report []int
		want   bool
	}{
		{
			report: []int{7, 6, 4, 2, 1},
			want:   true,
		},
		{
			report: []int{1, 2, 7, 8, 9},
			want:   false,
		},
		{
			report: []int{9, 7, 6, 2, 1},
			want:   false,
		},
		{
			report: []int{1, 3, 2, 4, 5},
			want:   false,
		},
		{
			report: []int{8, 6, 4, 4, 1},
			want:   false,
		},
		{
			report: []int{1, 3, 6, 7, 9},
			want:   true,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := reportSafe(tt.report)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReportSafeDampened(t *testing.T) {
	for i, tt := range []struct {
		report []int
		want   bool
	}{
		{
			report: []int{7, 6, 4, 2, 1},
			want:   true,
		},
		{
			report: []int{1, 2, 7, 8, 9},
			want:   false,
		},
		{
			report: []int{9, 7, 6, 2, 1},
			want:   false,
		},
		{
			report: []int{1, 3, 2, 4, 5},
			want:   true,
		},
		{
			report: []int{8, 6, 4, 4, 1},
			want:   true,
		},
		{
			report: []int{1, 3, 6, 7, 9},
			want:   true,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := reportSafeDampened(tt.report)
			assert.Equal(t, tt.want, got)
		})
	}
}
