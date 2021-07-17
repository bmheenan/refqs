package refqs

import "math"

// Sort sorts the given int slice in place, using reflective quicksort
func Sort(s []int) {
	if len(s) <= 1 {
		return
	}
	var (
		low   = 0
		pivot = len(s) / 2
		high  = len(s) - 1
	)
	for low < pivot || high > pivot {
		switch {
		case low < pivot && s[low] <= s[pivot]:
			low++
		case pivot < high && s[pivot] <= s[high]:
			high--
		case low < pivot && pivot < high:
			s[low], s[high] = s[high], s[low]
			low++
			high--
		case low == pivot:
			jump := pivot + int(math.Max(1, float64(high-pivot)/2))
			s[pivot], s[jump] = s[jump], s[pivot]
			pivot = jump
		case high == pivot:
			jump := pivot - int(math.Max(1, float64(pivot-low)/2))
			s[pivot], s[jump] = s[jump], s[pivot]
			pivot = jump
		}
	}
	Sort(s[:pivot])
	Sort(s[pivot+1:])
}
