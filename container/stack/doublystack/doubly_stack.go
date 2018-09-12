package doublystack

import (
	"errors"

	"github.com/techmexdev/algos/dictionary/doubly"
)

// Stack represents a stack.
type Stack struct {
	list *doubly.List
}

// New creates a stack from a doubly linked list.
func New() *Stack {
	return &Stack{list: doubly.New()}
}

// Push adds an element to the stack.
func (s *Stack) Push(val int) {
	s.list.Append(val)
}

// Pop removes the last inserted element.
func (s *Stack) Pop() (int, error) {
	b := s.list.Back()
	if b == nil {
		return 0, errors.New("no more elements in stack")
	}

	val := b.Val
	s.list.Delete(b)
	return val, nil
}
