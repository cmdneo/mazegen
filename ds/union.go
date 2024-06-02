package ds

type UnionFind struct {
	parents    []int
	tree_sizes []int
}

func MakeUnionFind(element_count int) UnionFind {
	size := element_count
	ret := UnionFind{
		parents:    make([]int, size),
		tree_sizes: make([]int, size),
	}

	for i := range size {
		ret.parents[i] = i
		ret.tree_sizes[i] = 1
	}

	return ret
}

func (u *UnionFind) InSameSet(v1, v2 int) bool {
	return u.find(v1) == u.find(v2)
}

func (u *UnionFind) Merge(s1, s2 int) {
	root1 := u.find(s1)
	root2 := u.find(s2)

	if root1 == root2 {
		return
	}

	large, small := root1, root2
	if u.tree_sizes[root1] < u.tree_sizes[root2] {
		large, small = small, large
	}

	// Merge the smaller tree in larger tree to minimize overall tree height.
	u.parents[small] = large
	u.tree_sizes[large] += u.tree_sizes[small]
}

func (u *UnionFind) find(v int) int {
	for {
		if u.parents[v] == v {
			return v
		} else {
			v = u.parents[v]
		}
	}
}
