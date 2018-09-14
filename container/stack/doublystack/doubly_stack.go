package doublystack

import (
	"github.com/rodolfo-r/datastructures"
	"github.com/rodolfo-r/datastructures/list/linked/doubly"
)

// Stack represents a stack.
type Stack struct {
	list *doubly.List
}

func assertStackImplementation() {
	var _ datastructures.Stack = (*Stack)(nil)
}

// New creates a stack from a doubly linked list.
func New() *Stack {
	return &Stack{list: doubly.New()}
}

// Push adds an element to the stack.
func (s *Stack) Push(value interface{}) {
	s.list.Append(value)
}

// Pop removes the last inserted element.
func (s *Stack) Pop() (value interface{}, ok bool) {
	b := s.list.Back()
	if b == nil {
		return nil, false
	}

	val := b.Value
	s.list.Delete(b)
	return val, true
}
