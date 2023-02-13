package stack

import (
	"reflect"
)

type Stack struct {
	nodes []any
}

func New() *Stack {
	return &Stack{make([]any, 0)}
}

func (st *Stack) Len() int {
	return len(st.nodes)
}

func (st *Stack) IsEmpty() bool {
	return st.Len() == 0
}

func (st *Stack) Top() (any, bool) {
	if st.IsEmpty() {
		return nil, false
	}
	return st.nodes[st.Len()-1], true
}

func (st *Stack) Push(node any) {
	st.nodes = append(st.nodes, node)
}

func (st *Stack) Pop() (top any, ok bool) {
	top, ok = st.Top()
	if ok {
		st.nodes = st.nodes[:st.Len()-1]
	}
	return
}

func (st *Stack) Values() []any {
	nodesCopy := make([]any, st.Len())
	copy(nodesCopy, st.nodes)
	size := len(nodesCopy)
	swap := reflect.Swapper(nodesCopy)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	return nodesCopy
}
