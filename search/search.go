package search

import (
	"github.com/cedricmar/algos/constraints"
)

// Binary returns the index of the found item or -1
func Binary[T constraints.Ordered](s []T, el T) int {
	i, j := 0, len(s)-1
	mid := func() int {
		return ((j - i) / 2) + i
	}

	for {
		if s[mid()] == el {
			return mid()
		} else if s[mid()] > el {
			j = mid() - 1
		} else {
			i = mid() + 1
		}
		if j < i {
			break
		}
	}

	return -1
}
