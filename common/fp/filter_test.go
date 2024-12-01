package fp

import (
	"reflect"
	"testing"
)

func TestFilerEmptyReturnsEmpty(t *testing.T) {
	in := []string{}
	predicate := func(t string) bool {
		return true
	}
	want := []string{}

	got := Filter(in, predicate)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFilterString(t *testing.T) {
	type Item struct {
		id    int
		inner bool
	}
	in := []Item{{1, true}, {2, true}, {3, false}, {4, true}, {5, false}}
	predicate := func(t Item) bool {
		return t.inner
	}
	want := []Item{{1, true}, {2, true}, {4, true}}

	got := Filter(in, predicate)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAllMatchEmptyReturnsTrue(t *testing.T) {
	in := []string{}
	predicate := func(t string) bool {
		return true
	}
	want := true

	got := AllMatch(in, predicate)

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestAllMatchTrue(t *testing.T) {
	type Item struct {
		id    int
		inner bool
	}
	in := []Item{{1, true}, {2, true}, {3, true}, {4, true}, {5, true}}
	predicate := func(t Item) bool {
		return t.inner
	}
	want := true
	got := AllMatch(in, predicate)

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestAllMatchFalse(t *testing.T) {
	type Item struct {
		id    int
		inner bool
	}
	in := []Item{{1, true}, {2, true}, {3, false}, {4, true}, {5, false}}
	predicate := func(t Item) bool {
		return t.inner
	}
	want := false
	got := AllMatch(in, predicate)

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}
