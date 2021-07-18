package refqs

import (
	"math/rand"
	"testing"
	"time"
)

// TestSortBasics tests small slices with basic edge cases (e.g. no elements, all the same, already sorted)
func TestSortBasics(t *testing.T) {
	for _, test := range [][]int{
		{},
		{10},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 2, 3, 4, 5, 6},
		{6, 5, 4, 3, 2, 1},
		{3, 2, 1},
		{1, 9, 2, 8, 3, 7, 4, 6, 5},
		{1, 1, 1, 1, 2, 2, 2, 2, 3, 7, 7, 7, 7, 7, 7, 7, 5, 5, 5, 4, 3, 2, 2, 9},
	} {
		Sort(test)
		if !sorted(test) {
			t.Errorf("%v is not sorted", test)
		}
	}
}

// TestSortBig tests sorting large int slices (sized 10,000 - 100,000,000) of random ints.
func TestSortBig(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for _, n := range []int{1e4, 1e5, 1e6, 1e7, 1e8} {
		test := randSlice(n)
		Sort(test)
		if !sorted(test) {
			t.Errorf("Could not sort []int of size %d", n)
		}
	}
}

// randSlice returns a []int of size n with random ints
func randSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = rand.Int()
	}
	return s
}

// sorted checks if the given slice is sorted
func sorted(s []int) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}
