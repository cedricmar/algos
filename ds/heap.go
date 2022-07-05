package ds

type Heap []int

// Arr[(i-1)/2] 	Returns the parent node
// Arr[(2*i)+1] 	Returns the left child node
// Arr[(2*i)+2] 	Returns the right child node

// Max Heap Property: The value of each node is less than or
// equal to the value of its parent, with the maximum value at the root

// Pop removes and returns the smallest element of the slice
func (h *Heap) Pop() (int, bool) {
	if len(*h) <= 0 {
		return 0, false
	}
	e := (*h)[0]
	*h = (*h)[1:]
	return e, true
}

// InsertMin add an element to a min heap
func (h *Heap) InsertMin(v int) {
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
func (h *Heap) Parent(i int) (int, int) {
	if i == 0 {
		return -1, 0
	}
	pid := (i - 1) / 2
	return pid, (*h)[pid]
}
