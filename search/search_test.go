package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchInt(t *testing.T) {
	s := []int{1, 2, 3, 4, 7, 8, 90, 567, 890}

	type test struct {
		input int
		idx   int
	}

	tests := []test{
		{input: 1, idx: 0},
		{input: 2, idx: 1},
		{input: 3, idx: 2},
		{input: 4, idx: 3},
		{input: 7, idx: 4},
		{input: 8, idx: 5},
		{input: 90, idx: 6},
		{input: 567, idx: 7},
		{input: 890, idx: 8},
		{input: 9, idx: -1},
		{input: -1, idx: -1},
		{input: 900, idx: -1},
		{input: 300, idx: -1},
		{input: 0, idx: -1},
		{input: 6, idx: -1},
		{input: -567, idx: -1},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.idx, Binary(s, tc.input))
	}
}

func TestBinarySearchString(t *testing.T) {
	s := []string{"a", "aa", "aab", "bc", "kk", "tyu", "yyy", "z", "zaa"}

	type test struct {
		input string
		idx   int
	}

	tests := []test{
		{input: "a", idx: 0},
		{input: "aa", idx: 1},
		{input: "aab", idx: 2},
		{input: "bc", idx: 3},
		{input: "kk", idx: 4},
		{input: "tyu", idx: 5},
		{input: "yyy", idx: 6},
		{input: "z", idx: 7},
		{input: "zaa", idx: 8},
		{input: "", idx: -1},
		{input: "t", idx: -1},
		{input: "zzz", idx: -1},
		{input: "0", idx: -1},
		{input: "_", idx: -1},
		{input: "aaa", idx: -1},
		{input: "cc", idx: -1},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.idx, Binary(s, tc.input))
	}
}
