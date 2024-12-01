package fp

import (
	"reflect"
	"testing"
)

func TestReduceEmptyReturnsIdentity(t *testing.T) {
	in := []int{}
	identity := 42
	combine := func(curr int, next int) int {
		return -12
	}

	want := identity
	got := Reduce(in, identity, combine)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReduceIntsSumReturnsSum(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	identity := 0
	combine := func(curr int, next int) int {
		return curr + next
	}

	want := 15
	got := Reduce(in, identity, combine)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumInts(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	want := 15
	got := Sum(in)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSumFloats(t *testing.T) {
	in := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	want := 16.5
	got := Sum(in)

	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

func TestSumFrom(t *testing.T) {
	type Item struct {
		n int
	}
	in := []Item{{1}, {2}, {3}, {4}, {5}}
	toNumber := func(i Item) int {
		return i.n
	}
	want := 15
	got := SumFrom(in, toNumber)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
