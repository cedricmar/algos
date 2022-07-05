package ds

import (
	"github.com/cedricmar/algos/constraints"
)

type Heap[T constraints.Ordered] []T

// Arr[(i-1)/2] 	Returns the parent node
// Arr[(2*i)+1] 	Returns the left child node
// Arr[(2*i)+2] 	Returns the right child node

// Max Heap Property: The value of each node is less than or
// equal to the value of its parent, with the maximum value at the root

// Pop removes and returns the smallest element of the slice
func (h *Heap[T]) Pop() (T, bool) {
	var e T
	if len(*h) <= 0 {
		return e, false
	}
	e = (*h)[0]
	*h = (*h)[1:]
	return e, true
}

// InsertMin add an element to a min heap
func (h *Heap[T]) InsertMin(v T) {
	*h = append(*h, v)
	l := len(*h) - 1
	for i := l; i > 0; i-- {
		if (*h)[i] < (*h)[i-1] {
			(*h)[i], (*h)[i-1] = (*h)[i-1], (*h)[i]
		}
	}
}

// Parent returns the index and value of the parent
// Returns -1 (index) for the root element
func (h *Heap[T]) Parent(i int) (int, T) {
	if i == 0 {
		var v T
		return -1, v
	}
	pid := (i - 1) / 2
	return pid, (*h)[pid]
}
