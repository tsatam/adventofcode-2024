package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlePart1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	want := 11
	got := handlePart1(input)

	assert.Equal(t, want, got)
}

func TestHandlePart2(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	want := 31
	got := handlePart2(input)

	assert.Equal(t, want, got)
}

func TestParseLists(t *testing.T) {
	for i, tt := range []struct {
		lines                []string
		wantList1, wantList2 []int
	}{
		{},
		{
			lines:     []string{"1   2"},
			wantList1: []int{1},
			wantList2: []int{2},
		},
		{
			lines: []string{
				"3   4",
				"4   3",
				"2   5",
				"1   3",
				"3   9",
				"3   3",
			},
			wantList1: []int{3, 4, 2, 1, 3, 3},
			wantList2: []int{4, 3, 5, 3, 9, 3},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			gotList1, gotList2 := parseLists(tt.lines)
			assert.Equal(t, gotList1, tt.wantList1)
			assert.Equal(t, gotList2, tt.wantList2)
		})
	}
}

func TestGetDistances(t *testing.T) {
	for i, tt := range []struct {
		list1, list2 []int
		want         []int
	}{
		{
			list1: []int{1},
			list2: []int{2},
			want:  []int{1},
		},
		{
			list1: []int{3, 4, 2, 1, 3, 3},
			list2: []int{4, 3, 5, 3, 9, 3},
			want:  []int{2, 1, 0, 1, 2, 5},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := getDistances(tt.list1, tt.list2)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestSimilarity(t *testing.T) {
	for i, tt := range []struct {
		list1, list2 []int
		want         []int
	}{
		{
			list1: []int{1},
			list2: []int{2},
			want:  []int{0},
		},
		{
			list1: []int{3, 4, 2, 1, 3, 3},
			list2: []int{4, 3, 5, 3, 9, 3},
			want:  []int{9, 4, 0, 0, 9, 9},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := similarity(tt.list1, tt.list2)
			assert.Equal(t, got, tt.want)
		})
	}
}
