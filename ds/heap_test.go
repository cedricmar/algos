package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetParent(t *testing.T) {

	//      1
	//    /   \
	//   2     3
	//  / \
	// 4   7
	heap := Heap([]int{1, 2, 3, 4, 7})

	i, e := heap.Parent(1)
	assert.Equal(t, i, 0)
	assert.Equal(t, e, 1)
}

func TestPop(t *testing.T) {

	//      2
	//    /   \
	//   3     4
	//  /
	// 7
	heap := Heap([]int{1, 2, 3, 4, 7})

	e, ok := heap.Pop()
	assert.Equal(t, e, 1)
	assert.Equal(t, ok, true)
	assert.Equal(t, Heap{2, 3, 4, 7}, heap)
}

func TestInsertMin(t *testing.T) {

	//      1
	//    /   \
	//   2     3
	//  / \
	// 4   7
	heap := Heap([]int{1, 2, 3, 4, 7})

	//      1
	//    /   \
	//   2     3
	//  / \   /
	// 4   5 7
	// []int{1 2 3 4 5 7}
	heap.InsertMin(5)
	assert.Equal(t, Heap{1, 2, 3, 4, 5, 7}, heap)

	//      0
	//    /   \
	//   1     2
	//  / \   / \
	// 3   4 5   7
	heap.InsertMin(0)
	assert.Equal(t, Heap{0, 1, 2, 3, 4, 5, 7}, heap)

	//       -1
	//      /   \
	//     0     1
	//    / \   / \
	//   2   3 4   5
	//  /
	// 7
	heap.InsertMin(-1)
	assert.Equal(t, Heap{-1, 0, 1, 2, 3, 4, 5, 7}, heap)
}
