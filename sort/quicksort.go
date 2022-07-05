package sort

import (
	"math/rand"
	"time"
)

// Quicksort implementation
func Quicksort(s []int) []int {
	if len(s) > 0 {
		p := partition(s)
		Quicksort(s[:p])
		Quicksort(s[p+1:])
	}
	return s
}

func partition(s []int) int {
	// Choose high pivot
	i := 0
	hi := len(s) - 1
	pivot := s[hi]
	// Move
	for j := i; j <= hi-1; j++ {
		if s[j] <= pivot {
			// Swap
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	// Swap pivot (high) with i
	s[i], s[hi] = s[hi], s[i]
	return i
}

func partitionRand(s []int) int {
	hi := len(s) - 1
	randIdx := selectRand(hi)
	s[randIdx], s[hi] = s[hi], s[randIdx]
	return partition(s)
}

func selectRand(l int) int {
	if l == 0 {
		return 0
	}
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	return rand.Intn(l)
}
