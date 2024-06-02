package maze

import "image"

type Corner uint8
type weightInt uint8
type Point = image.Point

const (
	TopLeft Corner = iota
	TopRight
	BottomLeft
	BottomRight
)

var (
	leftSide   = Point{-1, 0}
	rightSide  = Point{1, 0}
	topSide    = Point{0, -1}
	bottomSide = Point{0, 1}

	directions = []Point{leftSide, rightSide, topSide, bottomSide}
)

var _ = directions // Supress `directions` unused warning.

func (c Corner) toPosition(width, height int) Point {
	switch c {
	case TopLeft:
		return Point{0, 0}
	case TopRight:
		return Point{width - 1, 0}
	case BottomLeft:
		return Point{0, height - 1}
	case BottomRight:
		return Point{width - 1, height - 1}

	default:
		panic("Invalid corner value.")
	}
}
