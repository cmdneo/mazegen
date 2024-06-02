package maze

import (
	"image"
	"image/color"
)

type Maze struct {
	start, end Point
	grid       Grid
	// Cells which lie in the solution path
	solution []Point
}

// Return image and overlay solution (if requested)
func (maze *Maze) ToImage(
	cellSize, wallSize int, withSolution bool,
	wallColor, cellColor, startColor, endColor, solutionColor color.RGBA,
) *image.RGBA {
	lenCalc := func(blocks int) int {
		return cellSize*blocks + wallSize*(blocks+1)
	}

	// Image is composed of cells and walls.
	img := image.NewRGBA(
		image.Rect(0, 0, lenCalc(maze.grid.Cols()), lenCalc(maze.grid.Rows())),
	)
	fillMazeCell := func(x, y, pad int, color color.RGBA) {
		ix0 := wallSize*(x+1) + x*cellSize + pad
		iy0 := wallSize*(y+1) + y*cellSize + pad
		size := cellSize - pad*2
		fillBlock(img, ix0, iy0, size, size, color)
	}

	xend, yend := maze.grid.edges.GetSize()

	// Image is constructed in several steps:
	// 1. Fill eveything with wall
	fillBlock(img, 0, 0, img.Rect.Dx(), img.Rect.Dy(), wallColor)

	// 2. Fill cells between walls
	for y := range yend {
		for x := range xend {
			fillMazeCell(x, y, 0, cellColor)
		}
	}

	// 3. Check and delete any walls which should not exist.
	// For vertical walls (H edges).
	for y := 0; y < yend; y += 2 {
		for x := 0; x < xend-1; x++ {
			if maze.grid.edges.Get(x, y) == 0 {
				continue
			}

			ix0 := (cellSize + wallSize) * (x + 1)
			iy0 := wallSize*(1+y/2) + cellSize*y/2
			fillBlock(img, ix0, iy0, wallSize, cellSize, cellColor)
		}
	}

	// For horizontal walls (V edges).
	for y := 1; y < yend; y += 2 {
		for x := 0; x < xend; x++ {
			if maze.grid.edges.Get(x, y) == 0 {
				continue
			}

			ix0 := wallSize*(x+1) + cellSize*x
			iy0 := (wallSize + cellSize) * (y/2 + 1)
			fillBlock(img, ix0, iy0, cellSize, wallSize, cellColor)
		}
	}

	// 4. Draw the start and end markers.
	fillMazeCell(maze.start.X, maze.start.Y, 0, startColor)
	fillMazeCell(maze.end.X, maze.end.Y, 0, endColor)

	// 5. If solution requested then fill the solution path.
	if withSolution {
		for _, cell := range maze.solution {
			fillMazeCell(cell.X, cell.Y, cellSize/4, solutionColor)
		}
	}

	return img
}

func GenerateMaze(algorithm func(*Maze),
	width, height int, start, end Corner) Maze {
	if start == end {
		panic("Start and End corner should be different.")
	}

	maze := makeMaze(width, height)
	maze.start = start.toPosition(width, height)
	maze.end = end.toPosition(width, height)

	algorithm(&maze)
	return maze
}

func makeMaze(width, height int) Maze {
	return Maze{
		grid: MakeGrid(width, height),
	}
}

func fillBlock(img *image.RGBA, x0, y0, w, h int, color color.RGBA) {
	for y := range h {
		for x := range w {
			img.Set(x0+x, y0+y, color)
		}
	}
}
