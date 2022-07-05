package graph

type rec struct {
	gr   Graph
	seen map[int]struct{}
	res  map[int][]int
}

// Kosaraju algorithm
// https://en.wikipedia.org/wiki/Kosaraju%27s_algorithm
func Kosaraju(g Graph) map[int][]int {
	r := rec{
		gr:   g,
		seen: map[int]struct{}{},
		res:  map[int][]int{},
	}

	// Let L be empty
	lst := []int{}

	// For each vertex u of the graph
	for u := range g {
		// do Visit(u)
		//     where Visit(u) is the recursive subroutine:
		r.visit(u, &lst)
	}

	r.seen = map[int]struct{}{}
	// For each element u of L in order,
	for _, u := range lst {
		// do Assign(u,u)
		//     where Assign(u,root) is the recursive subroutine:
		r.assign(u, u)
	}

	return r.res
}

func (r *rec) visit(u int, lst *[]int) {
	// If u is unvisited then:
	if _, ok := r.seen[u]; !ok {
		//     Mark u as visited
		r.seen[u] = struct{}{}
		//     For each out-neighbour v of u,
		for _, v := range r.gr[u].out_nb {
			r.visit(v, lst)
			//     Prepend u to L.
			prependInt(lst, u)
		}
	}
	//     Otherwise do nothing.
}

func (r *rec) assign(u int, root int) {
	//     If u has not been assigned to a component then:
	if _, ok := r.seen[u]; !ok {
		//         Assign u as belonging to the component whose root is root.
		r.seen[u] = struct{}{}
		r.res[root] = append(r.res[root], u)
		//         For each in-neighbour v of u, do Assign(v,root).
		for _, v := range r.gr[u].in_nb {
			r.assign(v, root)
		}
	}
	//     Otherwise do nothing.
}

func prependInt(s *[]int, v int) {
	*s = append(*s, 0)
	copy((*s)[1:], *s)
	(*s)[0] = v
}

func reverseGraph(g [][]int) [][]int {
	r := make([][]int, len(g))
	for i, cons := range g {
		for _, dest := range cons {
			r[dest] = append(r[dest], i)
		}
	}
	return r
}
