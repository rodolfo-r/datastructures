package stack

import (
	"testing"

	"github.com/techmexdev/algos/stack"
)

func TestPush(t *testing.T) {
	s := stack.New()
	s.Push(0)
	s.Push(1)
	s.Push(2)
}

func TestPop(t *testing.T) {
	s := stack.New()
	s.Push(0)
	s.Push(1)
	s.Push(2)

	nn := []int{2, 1, 0}
	for i := 0; i < len(nn); i++ {
		n, err := s.Pop()
		if err != nil {
			t.Fatalf("have error %s, want nil", err)
		} else if n != nn[i] {
			t.Fatalf("have popped number %v, want %v", n, nn[i])
		}
	}

	_, err := s.Pop()
	if err == nil {
		t.Fatal("have nil error, want non-nil")
	}

}
