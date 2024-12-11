package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlePart1(t *testing.T) {
	input := `
125 17
	`

	want := 55312
	got := handlePart1(input)

	assert.Equal(t, want, got)
}

func TestHandlePart2(t *testing.T) {
	input := `
125 17
	`

	want := 65601038650482
	got := handlePart2(input)

	assert.Equal(t, want, got)
}

func TestProcessStone(t *testing.T) {
	for _, tt := range []struct {
		name  string
		stone uint64
		want  []uint64
	}{
		{name: "first", stone: 0, want: []uint64{1}},
		{name: "second", stone: 1, want: []uint64{2024}},
		{name: "third", stone: 10, want: []uint64{1, 0}},
		{name: "fourth", stone: 99, want: []uint64{9, 9}},
		{name: "fifth", stone: 999, want: []uint64{2021976}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := processStone(tt.stone)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIterate(t *testing.T) {
	for _, tt := range []struct {
		name string
		sm   stonesMap
		want stonesMap
	}{
		{
			name: "initial arrangement -> after 1 blink",
			sm:   stonesMap{125: 1, 17: 1},
			want: stonesMap{253000: 1, 1: 1, 7: 1},
		},
		{
			name: "1->2",
			sm:   stonesMap{253000: 1, 1: 1, 7: 1},
			want: stonesMap{253: 1, 0: 1, 2024: 1, 14168: 1},
		},
		{
			name: "2->3",
			sm:   stonesMap{253: 1, 0: 1, 2024: 1, 14168: 1},
			want: stonesMap{512072: 1, 1: 1, 20: 1, 24: 1, 28676032: 1},
		},
		{
			name: "3->4",
			sm:   stonesMap{512072: 1, 1: 1, 20: 1, 24: 1, 28676032: 1},
			want: stonesMap{512: 1, 72: 1, 2024: 1, 2: 2, 0: 1, 4: 1, 2867: 1, 6032: 1},
		},
		{
			name: "4->5",
			sm:   stonesMap{512: 1, 72: 1, 2024: 1, 2: 2, 0: 1, 4: 1, 2867: 1, 6032: 1},
			want: stonesMap{1036288: 1, 7: 1, 2: 1, 20: 1, 24: 1, 4048: 2, 1: 1, 8096: 1, 28: 1, 67: 1, 60: 1, 32: 1},
		},
		{
			name: "5->6",
			sm:   stonesMap{1036288: 1, 7: 1, 2: 1, 20: 1, 24: 1, 4048: 1, 8096: 1, 28: 1, 67: 1, 60: 1, 32: 1},
			want: stonesMap{2097446912: 1, 14168: 1, 4048: 1, 2: 4, 0: 2, 4: 1, 40: 1, 48: 1, 80: 1, 96: 1, 8: 1, 6: 2, 7: 1, 3: 1},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sm.iterate()
			assert.Equal(t, tt.want, got)
		})
	}
}
