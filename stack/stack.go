package stack

type Stack struct {
	nodes []any
}

func New() *Stack {
	return &Stack{}
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
	return st.nodes[0], true
}

func (st *Stack) Push(node any) {
	st.nodes = append([]any{node}, st.nodes...)
}

func (st *Stack) Pop() (top any, ok bool) {
	top, ok = st.Top()
	st.nodes = st.nodes[1:]
	return
}
