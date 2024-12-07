package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlePart1(t *testing.T) {
	input := `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
	`

	want := uint64(3749)
	got := handlePart1(input)

	assert.Equal(t, want, got)
}

func TestHandlePart2(t *testing.T) {
	input := `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
	`

	want := uint64(11387)
	got := handlePart2(input)

	assert.Equal(t, want, got)
}

func TestIsValid(t *testing.T) {
	for _, tt := range []struct {
		name         string
		testValue    uint64
		numbers      []uint64
		shouldConcat bool
		want         bool
	}{
		{
			name:         "single value returns true if equal to testValue",
			testValue:    42,
			numbers:      []uint64{42},
			shouldConcat: false,
			want:         true,
		},
		{
			name:         "single value returns false if not equal to testValue",
			testValue:    42,
			numbers:      []uint64{12},
			shouldConcat: false,
			want:         false,
		},
		{
			name:         "two values, returns true if multiplication equal to testValue",
			testValue:    190,
			numbers:      []uint64{10, 19},
			shouldConcat: false,
			want:         true,
		},
		{
			name:         "two values, returns false if multiplication not equal to testValue",
			testValue:    190,
			numbers:      []uint64{10, 12},
			shouldConcat: false,
			want:         false,
		},
		{
			name:         "two values, returns true if addition equal to testValue",
			testValue:    19,
			numbers:      []uint64{10, 9},
			shouldConcat: false,
			want:         true,
		},
		{
			name:         "two values, returns false if addition not equal to testValue",
			testValue:    19,
			numbers:      []uint64{10, 2},
			shouldConcat: false,
			want:         false,
		},
		{
			name:         "if value already greater than test value, immediately return false",
			testValue:    1,
			numbers:      []uint64{2, 1, 1, 1},
			shouldConcat: false,
			want:         false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.testValue, tt.numbers, tt.shouldConcat)
			assert.Equal(t, tt.want, got)
		})
	}
}
