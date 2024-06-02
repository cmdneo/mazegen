package maze

import (
	"math/rand/v2"
	"mazegen/ds"
)

// A variation of Krushkal's algorithm
func AlgoKruskal(m *Maze) {
	// Union find keeps components as integers [0, max)
	// So we map every vertex to a unique integer as such.
	vertexID := func(pt Point) int { return m.grid.cols*pt.Y + pt.X }

	// There are |V| - 1 edges in a tree.
	edges := genFullEdgeList(&m.grid)
	components := ds.MakeUnionFind(edges.Len())

	for edges.Len() > 0 {
		// Extract a random edge and try adding it
		edges.Swap(rand.IntN(edges.Len()), edges.Len()-1)
		e := edges.Pop()
		v1_id, v2_id := vertexID(e.v1), vertexID(e.v2)

		if components.InSameSet(v1_id, v2_id) {
			continue
		}

		components.Merge(v1_id, v2_id)
		m.grid.SetEdgeWeight(e.v1, e.v2, 1)
	}

	m.solve()
}

type edge struct {
	v1, v2 Point
}

// Generates a list of every possible edge in the grid.
func genFullEdgeList(g *Grid) ds.Stack[edge] {
	edge_cnt := g.Cols()*g.rows*4 - 2*(g.Cols()+g.rows)
	edges := ds.MakeStack[edge](edge_cnt)

	for y := range g.rows {
		for x := range g.Cols() - 1 {
			edges.Push(edge{Point{x, y}, Point{x + 1, y}})
		}
	}

	for y := range g.rows - 1 {
		for x := range g.Cols() {
			edges.Push(edge{Point{x, y}, Point{x, y + 1}})
		}
	}

	return edges
}
