package graph

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBfs(t *testing.T) {
	beg := getNano()

	tcs := []struct {
		edges, exp [][]int
	}{
		{
			[][]int{
				0: {1},
				1: {2, 3},
				2: {0},
				3: {3},
			},
			[][]int{
				0: {0, 1, 2, 3},
				1: {1, 2, 3, 0},
				2: {2, 0, 1, 3},
			},
		},
		{
			[][]int{
				0: {1, 2, 4},
				1: {3, 5},
				2: {6},
				3: {},
				4: {5},
				5: {},
				6: {},
			},
			[][]int{
				0: {0, 1, 2, 4, 3, 5, 6},
			},
		},
	}

	for _, tc := range tcs {
		for i, x := range tc.exp {
			sorted := Bfs(tc.edges, i)
			assert.Equal(t, x, sorted)
		}
	}

	end := getNano()
	fmt.Printf("total time: %d\n", end-beg)
}

func TestDfs(t *testing.T) {
	beg := getNano()

	tcs := []struct {
		edges, exp [][]int
	}{
		{
			[][]int{
				0: {1},
				1: {2, 3},
				2: {0},
				3: {3},
			},
			[][]int{
				0: {0, 1, 2, 3},
				1: {1, 2, 0, 3},
				2: {2, 0, 1, 3},
			},
		},
		{
			[][]int{
				0: {1, 2, 4},
				1: {3, 5},
				2: {6},
				3: {},
				4: {5},
				5: {},
				6: {},
			},
			[][]int{
				0: {0, 1, 3, 5, 2, 6, 4},
			},
		},
	}

	for _, tc := range tcs {
		for i, x := range tc.exp {
			sorted := Dfs(tc.edges, i)
			assert.Equal(t, x, sorted)
		}
	}

	end := getNano()
	fmt.Printf("total time: %d\n", end-beg)
}

// func TestDfs2(t *testing.T) {
// 	beg := getNano()

// 	sorted := Dfs2(2)
// 	fmt.Println(sorted)

// 	end := getNano()
// 	fmt.Printf("total time: %d\n", end-beg)
// }

func getNano() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}
