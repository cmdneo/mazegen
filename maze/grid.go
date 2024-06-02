package maze

import (
	"math/rand/v2"
	"mazegen/ds"
)

// Grid with weighted edges among cells.
// If an has edge weight zero, then it means that no such edge exists.
type Grid struct {
	cols, rows int
	// Each edge is stored only once.
	// Even numbered rows store the horizontal edge weights (=widths-1 edges).
	// Odd numbered rows store the vertical edge weights    (=width edges).
	// Total rows = 2*height - 1.
	// We store edge weights in the same order as they appear in the grid
	// left-to-right and top-to-bottom.
	//
	// It can be visualized as follows for a 3x3 grid:
	// H and V are edges among cells.
	//
	// +-----+-----+-----+    Weights     Row-Index
	// |     |     |     |
	// |     H     H     |    [w, w, -] : 0
	// |     |     |     |
	// +--V--+--V--+--V--+    [w, w, w] : 1
	// |     |     |     |
	// |     H     H     |    [w, w, -] : 2
	// |     |     |     |
	// +--V--+--V--+--V--+    [w, w, w] : 3
	// |     |     |     |
	// |     H     H     |    [w, w, -] : 4
	// |     |     |     |
	// +-----+-----+-----+
	edges ds.Matrix[weightInt]
}

func MakeGrid(columns, rows int) Grid {
	return Grid{
		cols:  columns,
		rows:  rows,
		edges: ds.MakeMatrix[weightInt](columns, 2*rows-1),
	}
}

func (g *Grid) RandomizeWeights(min, max weightInt) {
	xend, yend := g.edges.GetSize()

	for y := range yend {
		for x := range xend {
			// In range [min, max]
			v := rand.N(max-min+1) + min
			g.edges.Set(x, y, v)
		}
	}

	// Clear unused entries in case of H edges.
	for y := 0; y < yend; y += 2 {
		g.edges.Set(xend-1, y, 0)
	}

}

// Returns edge weight, 0 if no such edge exists.
func (g *Grid) GetEdgeWeight(v1, v2 Point) weightInt {
	x, y := calcEdgeIndex(v1, v2)

	if g.edges.IsIndexValid(x, y) {
		return g.edges.Get(x, y)
	} else {
		return 0
	}
}

func (g *Grid) SetEdgeWeight(v1, v2 Point, w weightInt) {
	x, y := calcEdgeIndex(v1, v2)
	g.edges.Set(x, y, w)
}

func (g *Grid) Cols() int { return g.cols }

func (g *Grid) Rows() int { return g.rows }

func GetAdjacentCells(g *Grid, v Point, result *[]Point, pred func(Point) bool) {
	*result = (*result)[:0]

	appendIf := func(cond bool, n Point) {
		if cond && pred(n) {
			*result = append(*result, n)
		}
	}

	appendIf(v.X > 0, Point{v.X - 1, v.Y})
	appendIf(v.Y > 0, Point{v.X, v.Y - 1})
	appendIf(v.X < g.cols-1, Point{v.X + 1, v.Y})
	appendIf(v.Y < g.rows-1, Point{v.X, v.Y + 1})
}

func calcEdgeIndex(v1, v2 Point) (int, int) {
	d := v2.Sub(v1)
	x, y := v1.X, v1.Y

	switch d {
	case leftSide, topSide, bottomSide:
		return x + d.X, 2*y + d.Y
	case rightSide:
		return x, 2*y + d.Y

	default:
		return -1, -1
	}
}
