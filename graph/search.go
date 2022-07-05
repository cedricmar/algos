package graph

import (
	"github.com/cedricmar/algos/ds"
)

func Dfs(edges [][]int, root int) []int {
	stack := ds.Stack{}
	visited := make(map[int]struct{})

	// Push
	stack.Push(root)
	visited[root] = struct{}{}
	res := []int{root}

	for stack.Length() > 0 {
		edge := stack.Pop()

		if _, ok := visited[edge]; !ok {
			res = append(res, edge)
			visited[edge] = struct{}{}
		}

		// Get connected edges
		prox := edges[edge]

		for i := len(prox) - 1; i >= 0; i-- {
			// Seen?
			if _, ok := visited[prox[i]]; !ok {
				// Push
				stack.Push(prox[i])
			}
		}
	}

	return res
}

func Bfs(edges [][]int, root int) []int {
	queue := ds.Queue{}
	visited := make(map[int]struct{})

	queue.Enqueue(root)
	visited[root] = struct{}{}
	res := []int{root}

	for len(queue) > 0 {
		v := queue.Dequeue()
		prox := edges[v]

		for _, v := range prox {
			// Seen?
			if _, ok := visited[v]; !ok {
				queue.Enqueue(v)
				visited[v] = struct{}{}
				res = append(res, v)
			}
		}
	}
	return res
}
