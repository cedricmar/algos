package graph

type direction struct {
	out_nb []int
	in_nb  []int
}

type Graph []direction

func NewGraphFromAdjList(adj [][]int) Graph {
	g := make(Graph, len(adj))

	for vtx, nbs := range adj {
		for _, nb := range nbs {
			g[vtx].out_nb = append(g[vtx].out_nb, nb)
			g[nb].in_nb = append(g[nb].in_nb, vtx)
		}
	}

	return g
}
