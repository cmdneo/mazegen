package ds

type Stack[T any] []T

func MakeStack[T any](init_cap int) Stack[T] {
	return make(Stack[T], 0, init_cap)
}

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() T {
	n := len(*s)
	ret := (*s)[n-1]

	*s = (*s)[:n-1]
	return ret
}

func (s Stack[T]) Top() T {
	return s[len(s)-1]
}

func (s Stack[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Stack[T]) ToNewSlice() []T {
	ret := make([]T, len(s))
	copy(ret, s)
	return ret
}

func (s Stack[T]) Len() int {
	return len(s)
}
