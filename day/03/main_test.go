package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlePart1(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	want := 161
	got := handlePart1(input)

	assert.Equal(t, want, got)
}

func TestHandlePart2(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

	want := 48
	got := handlePart2(input)

	assert.Equal(t, want, got)
}

func TestCalcMul(t *testing.T) {
	for _, tt := range []struct {
		instruction string
		want        int
	}{
		{
			instruction: "mul(44,46)",
			want:        2024,
		},
		{
			instruction: "mul(123,4)",
			want:        492,
		},
	} {
		t.Run(fmt.Sprintf("%s -> %d", tt.instruction, tt.want), func(t *testing.T) {
			got := calcMul(tt.instruction)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetEnabledInstructions(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  []string
	}{
		{},
		{
			input: "asdfdon't()ghjkl",
			want:  []string{"asdf"},
		},
		{
			input: "asdfdon't()ghjkldo()zxcvbdon't()nm,./",
			want:  []string{"asdf", "zxcvb"},
		},
	} {
		t.Run(tt.input, func(t *testing.T) {
			got := getEnabledInstructions(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
