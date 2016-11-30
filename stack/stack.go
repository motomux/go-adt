package stack

import "container/list"

// Stack represents stack data structre
type Stack struct {
	list *list.List
}

// NewStack initiates stack
func NewStack() *Stack {
	return &Stack{
		list: list.New(),
	}
}

// Push add a value on top of a stack
func (s *Stack) Push(v interface{}) {
	s.list.PushFront(v)
}

// Pop remove a value on top of stack
func (s *Stack) Pop() interface{} {
	front := s.list.Front()
	s.list.Remove(front)
	return front.Value
}
