package fp

import (
	"reflect"
	"testing"
)

func TestZipFuncEmptyReturnsEmpty(t *testing.T) {
	in1 := []int{}
	in2 := []int{}
	combine := func(elem1, elem2 int) int { return 0 }

	want := []int{}
	got := ZipFunc(in1, in2, combine)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestZipFuncIntsSumReturnsSum(t *testing.T) {
	in1 := []int{1, 2, 3, 4, 5}
	in2 := []int{5, 4, 3, 2, 1}
	combine := func(elem1, elem2 int) int { return elem1 + elem2 }

	want := []int{6, 6, 6, 6, 6}
	got := ZipFunc(in1, in2, combine)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestZipReturnsTuple(t *testing.T) {
	in1 := []int{1, 2, 3, 4, 5}
	in2 := []int{5, 4, 3, 2, 1}

	want := [][2]int{{1, 5}, {2, 4}, {3, 3}, {4, 2}, {5, 1}}
	got := Zip(in1, in2)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
