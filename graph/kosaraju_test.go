package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKosaraju(t *testing.T) {
	assert.Equal(t,
		map[int][]int{1: {1, 0, 2}, 3: {3}},
		Kosaraju(
			NewGraphFromAdjList(graph),
		),
	)
}

func TestKosarajuComplex(t *testing.T) {
	exp := map[int][]int{
		1:  {1, 0, 3, 4, 2, 5, 6},
		7:  {7, 9, 8},
		10: {10},
		11: {11},
		12: {12},
	}
	assert.Equal(t, exp, Kosaraju(
		NewGraphFromAdjList(graph_complex),
	))
}

func TestReverse(t *testing.T) {
	assert.Equal(t, reversed, reverseGraph(graph))
}

var graph = [][]int{
	0: {1},
	1: {2, 3},
	2: {0},
	3: {3},
}

var reversed = [][]int{
	0: {2},
	1: {0},
	2: {1},
	3: {1, 3},
}

var graph_complex = [][]int{
	0:  {1},
	1:  {2},
	2:  {4, 6},
	3:  {0, 5},
	4:  {3},
	5:  {4, 7},
	6:  {4, 7},
	7:  {8, 10},
	8:  {9, 11},
	9:  {7},
	10: {},
	11: {12},
	12: {},
}
