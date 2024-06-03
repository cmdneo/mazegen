package maze

import (
	"math/rand/v2"
	"mazegen/ds"
)

const (
	oldWalk int8 = iota + 1
	thisWalk
)

// Loop-erased random walk algorithm, it is slow.
// Plain random walk produces mazes having very long and convoluted paths.
func AlgoLoopErasedRandWalk(m *Maze) {
	// vertices := genAllVertices(m.grid.Cols(), m.grid.Rows())
	visited := ds.MakeMatrix[int8](m.grid.Cols(), m.grid.Rows())
	count := m.grid.Cols() * m.grid.Rows()

	extract_new := func(visit_type int8) Point {
		x, y := visited.Find(func(v int8) bool { return v == 0 })
		visited.Set(x, y, visit_type)
		return Point{x, y}
	}

	// Initialize the maze by marking a cell as visited.
	extract_new(oldWalk)
	count--
	neighs := make([]Point, 0, 4)

	for count > 0 {
		// Start a new loop-erased random walk from an unvisited cell until
		// we collide with a visited cell not related to the current walk.
		v := extract_new(0)
		path := ds.MakeStack[Point](4)

		for visited.Get(v.X, v.Y) != oldWalk {
			path.Push(v)
			if visited.Get(v.X, v.Y) == thisWalk {
				eraseLoop(&path, &visited)
			}
			visited.Set(v.X, v.Y, thisWalk)

			GetAdjacentCells(&m.grid, v, &neighs, nil)
			v = neighs[rand.IntN(len(neighs))]
		}
		count -= path.Len()

		// Merge this walk with rest of the maze.
		path.Push(v)
		for i := range len(path) - 1 {
			v, w := path[i], path[i+1]
			visited.Set(v.X, v.Y, oldWalk)
			m.grid.SetEdgeWeight(v, w, 1)
		}
	}

	m.solve()
}

// Random walk until we have visited all the cells
func AlgoRandWalk(m *Maze) {
	visited := ds.MakeMatrix[bool](m.grid.cols, m.grid.rows)
	stack := ds.MakeStack[Point](0)

	count := m.grid.cols * m.grid.rows
	at := m.start
	neighs := make([]Point, 0, 4)

	for {
		count--
		stack.Push(at)
		visited.Set(at.X, at.Y, true)

		if at == m.end {
			m.solution = stack.ToNewSlice()
		}

		if count == 0 {
			break // All cells visited
		}

		for {
			GetAdjacentCells(&m.grid, at, &neighs, &visited)
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

// Erase the loop and mark the removed cells as unvisited.
func eraseLoop(path *ds.Stack[Point], visited *ds.Matrix[int8]) {
	cross_at := path.Pop()

	for path.Top() != cross_at {
		v := path.Pop()
		// vertices.Push(v)
		visited.Set(v.X, v.Y, 0)
	}
}
