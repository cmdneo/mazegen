package maze

import (
	"mazegen/ds"
)

// In case an algorithm cannot generate the solution along with the maze.
// This solver should be used to fill in the solution.
// If no solution is find the the maze generated in not solvable.
// Returns true and fills in the solution if found, otherwise returns false.
func (m *Maze) solve() bool {
	visited := ds.MakeMatrix[bool](m.grid.cols, m.grid.rows)
	stack := ds.MakeStack[Point](0)

	at := m.start
	neighs := make([]Point, 0, 4)
	m.solution = []Point{}

	for {
		stack.Push(at)
		visited.Set(at.X, at.Y, true)

		if at == m.end {
			m.solution = stack.ToNewSlice()
			return true
		}

		for {
			GetAdjacentCells(&m.grid, at, &neighs, &visited)
			neighs = filter_disconnected(&m.grid, at, neighs)

			if len(neighs) > 0 {
				break
			}

			stack.Pop() // Backtrack
			if stack.Len() == 0 {
				return false
			}
			at = stack.Top()
		}

		at = neighs[0]
		stack.Push(at)
		visited.Set(at.X, at.Y, true)
	}
}

// Remove cells in input which have no edge with v.
func filter_disconnected(g *Grid, v Point, input []Point) []Point {
	// Partition and then cut off the end.
	end := 0
	for i, w := range input {
		if g.GetEdgeWeight(v, w) == 0 {
			continue
		}
		input[i], input[end] = input[end], input[i]
		end++
	}

	return input[:end]
}
