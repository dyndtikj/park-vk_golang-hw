package stack

import (
	"reflect"
)

type Stack[T any] struct {
	nodes []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{make([]T, 0)}
}

func (st *Stack[T]) Len() int {
	return len(st.nodes)
}

func (st *Stack[T]) IsEmpty() bool {
	return st.Len() == 0
}

func (st *Stack[T]) Peek() (result T, ok bool) {
	if st.IsEmpty() {
		return
	}
	return st.nodes[st.Len()-1], true
}

func (st *Stack[T]) Push(node T) {
	st.nodes = append(st.nodes, node)
}

func (st *Stack[T]) Pop() (top T, ok bool) {
	top, ok = st.Peek()
	if ok {
		st.nodes = st.nodes[:st.Len()-1]
	}
	return
}

func (st *Stack[T]) Values() []T {
	nodesCopy := make([]T, st.Len())
	copy(nodesCopy, st.nodes)
	size := len(nodesCopy)
	swap := reflect.Swapper(nodesCopy)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	return nodesCopy
}
