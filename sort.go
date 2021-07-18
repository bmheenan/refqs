package refqs

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
			p := pivot + min1((high-pivot)/2)
			s[pivot], s[p] = s[p], s[pivot]
			pivot = p
		case high == pivot:
			p := pivot - min1((pivot-low)/2)
			s[pivot], s[p] = s[p], s[pivot]
			pivot = p
		}
	}
	Sort(s[:pivot])
	Sort(s[pivot+1:])
}

// min1 returns the given int if it's at least 1. Otherwise, it returns 1
func min1(a int) int {
	if a < 1 {
		return 1
	}
	return a
}
