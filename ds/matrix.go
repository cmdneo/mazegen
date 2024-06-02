package ds

import (
	"fmt"
	"strings"
)

type Matrix[T any] struct {
	rows, cols int
	// Data stored in row major form
	data []T
}

func MakeMatrix[T any](columns, rows int) Matrix[T] {
	return Matrix[T]{
		rows: rows,
		cols: columns,
		data: make([]T, rows*columns),
	}
}

func MakeMatrixLike[T, V any](example Matrix[V]) Matrix[T] {
	return MakeMatrix[T](example.cols, example.rows)
}

func (m *Matrix[T]) Slice() []T {
	return m.data
}

func (m *Matrix[T]) Get(x, y int) T {
	return m.data[m.SliceIndex(x, y)]
}

func (m *Matrix[T]) GetPtr(x, y int) *T {
	return &m.data[m.SliceIndex(x, y)]
}

func (m *Matrix[T]) Set(x, y int, value T) {
	m.data[m.SliceIndex(x, y)] = value
}

func (m *Matrix[T]) IsIndexValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < m.cols && y < m.rows
}

func (m *Matrix[T]) SliceIndex(x, y int) int {
	return x + y*m.cols
}

func (m *Matrix[T]) GetSize() (int, int) {
	return m.cols, m.rows
}

func (m Matrix[T]) String() string {
	sb := strings.Builder{}

	for y := range m.rows {
		sb.WriteString(fmt.Sprintf("%v\n", m.data[y*m.cols:(y+1)*m.cols-1]))
	}

	return sb.String()
}
