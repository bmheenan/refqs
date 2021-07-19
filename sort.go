// refqs, or reflective quicksort, is a quicksort-family algorithm that attempts to provide n*log(n) average time
// performance, but performs better than most common implementations of quicksort for common edge cases (e.g already
// sorted, almost already sorted, reverse sorted)
package refqs

import "sort"

// Sort sorts the given int slice in place, using reflective quicksort
func Sort(s sort.Interface) {
	sortSection(s, 0, s.Len())
}

func sortSection(s sort.Interface, min, max int) {
	if max-min <= 1 {
		return
	}
	var (
		low   = min
		pivot = min + ((max - min) / 2)
		high  = max - 1
	)
	for low < pivot || high > pivot {
		switch {
		case low < pivot && !s.Less(pivot, low):
			low++
		case pivot < high && !s.Less(high, pivot):
			high--
		case low < pivot && pivot < high:
			s.Swap(low, high)
			low++
			high--
		case low == pivot:
			p := pivot + min1((high-pivot)/2)
			s.Swap(p, pivot)
			pivot = p
		case high == pivot:
			p := pivot - min1((pivot-low)/2)
			s.Swap(p, pivot)
			pivot = p
		}
	}
	sortSection(s, min, pivot)
	sortSection(s, pivot+1, max)
}

// min1 returns the given int if it's at least 1. Otherwise, it returns 1
func min1(a int) int {
	if a < 1 {
		return 1
	}
	return a
}
