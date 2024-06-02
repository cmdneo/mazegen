package maze

import (
	"math/rand/v2"
	"mazegen/ds"
)

func AlgoRandWalk(m *Maze) {
	in_tree := ds.MakeMatrix[bool](m.grid.cols, m.grid.rows)
	stack := ds.MakeStack[Point](0)
	// Predicate for if a vertex should be a neighbour.
	is_new_neigh := func(v Point) bool { return !in_tree.Get(v.X, v.Y) }

	count := 0
	at := m.start
	neighs := make([]Point, 0, 4)

	for {
		count++
		stack.Push(at)
		in_tree.Set(at.X, at.Y, true)

		if at == m.end {
			m.solution = stack.ToNewSlice()
		}

		if count == m.grid.cols*m.grid.rows {
			break // All cells visited
		}

		for {
			GetAdjacentCells(&m.grid, at, &neighs, is_new_neigh)
			if len(neighs) > 0 {
				break
			}

			stack.Pop() // Backtrack
			at = stack.Top()
		}

		// Select a random neighbour and descend
		next := neighs[rand.N(len(neighs))]
		m.grid.SetEdgeWeight(at, next, 1)
		at = next
	}
}
