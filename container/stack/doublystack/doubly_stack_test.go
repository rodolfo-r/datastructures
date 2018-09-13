package doublystack_test

import (
	"testing"

	"github.com/techmexdev/algos/container/stack/doublystack"
)

func TestPush(t *testing.T) {
	s := doublystack.New()
	s.Push(0)
	s.Push(1)
	s.Push(2)
}

func TestPop(t *testing.T) {
	s := doublystack.New()
	s.Push(0)
	s.Push(1)
	s.Push(2)

	nn := []int{2, 1, 0}
	for i := 0; i < len(nn); i++ {
		if n, ok := s.Pop(); !ok {
			t.Fatal("expected successful pop")
		} else if n != nn[i] {
			t.Fatalf("have popped number %v, want %v", n, nn[i])
		}
	}

	if _, ok := s.Pop(); ok {
		t.Fatal("expected unsuccessful pop")
	}
}
