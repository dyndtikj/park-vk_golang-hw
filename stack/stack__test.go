package stack

import "testing"

func TestStack_Push(t *testing.T) {
	stack := New()
	if val := stack.IsEmpty(); val != true {
		t.Errorf("Expected empty stack")
	}
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if val := stack.Values(); val[0] != 3 || val[1] != 2 || val[2] != 1 || val[3] != 0 {
		t.Errorf("Got %v expected %v", val, "[3 2 1 0]")
	}
	if val := stack.IsEmpty(); val != false {
		t.Errorf("Got %v expected %v", val, false)
	}
	if val := stack.Len(); val != 4 {
		t.Errorf("Got %v expected %v", val, 3)
	}
	if val, ok := stack.Top(); val != 3 || !ok {
		t.Errorf("Got %v expected %v", val, 3)
	}
}

func TestStack_Pop(t *testing.T) {
	stack := New()
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if val, ok := stack.Pop(); val != 3 || !ok {
		t.Errorf("Got %v expected %v", val, 3)
	}
	if val, ok := stack.Pop(); val != 2 || !ok {
		t.Errorf("Got %v expected %v", val, 2)
	}
	if val, ok := stack.Pop(); val != 1 || !ok {
		t.Errorf("Got %v expected %v", val, 1)
	}
	if val, ok := stack.Pop(); val != 0 || !ok {
		t.Errorf("Got %v expected %v", val, 0)
	}
	if val, ok := stack.Pop(); val != nil || ok {
		t.Errorf("Got %v expected %v", val, nil)
	}
}

func TestStack_Top(t *testing.T) {
	stack := New()
	if val, ok := stack.Top(); val != nil || ok {
		t.Errorf("Expected empty stack")
	}
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if val, ok := stack.Top(); val != 3 || !ok {
		t.Errorf("Got %v expected %v", val, 3)
	}
	_, _ = stack.Pop()
	if val, ok := stack.Top(); val != 2 || !ok {
		t.Errorf("Got %v expected %v", val, 2)
	}
	_, _ = stack.Pop()
	if val, ok := stack.Top(); val != 1 || !ok {
		t.Errorf("Got %v expected %v", val, 1)
	}
	_, _ = stack.Pop()
	if val, ok := stack.Top(); val != 0 || !ok {
		t.Errorf("Got %v expected %v", val, 0)
	}
	_, _ = stack.Pop()
	if val, ok := stack.Top(); val != nil || ok {
		t.Errorf("Got %v expected %v", val, nil)
	}
}
