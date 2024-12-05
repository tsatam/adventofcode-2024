package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlePart1(t *testing.T) {
	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	want := 143
	got := handlePart1(input)
	assert.Equal(t, want, got)
}

func TestHandlePart2(t *testing.T) {
	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	want := 123
	got := handlePart2(input)
	assert.Equal(t, want, got)
}

func TestUpdateIsValid(t *testing.T) {
	rules := map[int][]int{
		1: {2, 3, 4},
		2: {3, 4},
		3: {4},
	}

	for _, tt := range []struct {
		update []int
		want   bool
	}{
		{
			update: []int{1, 2, 3, 4, 5},
			want:   true,
		},
		{
			update: []int{5, 4, 3, 2, 1},
			want:   false,
		},
	} {
		t.Run(fmt.Sprintf("%v->%v", tt.update, tt.want), func(t *testing.T) {
			got := updateIsValid(rules)(tt.update)
			assert.Equal(t, tt.want, got)
		})
	}
}
