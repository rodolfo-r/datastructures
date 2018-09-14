package binaryheap_test

import (
	"testing"

	"github.com/rodolfo-r/datastructures"
	"github.com/rodolfo-r/datastructures/container/priorityqueue/binaryheap"
)

func TestInsert(t *testing.T) {
	var pq datastructures.PriorityQueue = binaryheap.NewMin()

	nn := []int{5, 2, 1, 3, 4}
	for _, n := range nn {
		pq.Insert(n)
	}
}

func TestMin(t *testing.T) {
	var pq datastructures.PriorityQueue = binaryheap.NewMin()

	nn := []int{5, 2, 1, 3, 4}
	for _, n := range nn {
		pq.Insert(n)
	}

	if n, ok := pq.Min(); !ok {
		t.Fatalf("expected to be ok")
	} else if n != 1 {
		t.Fatalf("have n %v, want %v", n, 1)
	}
}

func TestExtractMin(t *testing.T) {
	var pq datastructures.PriorityQueue = binaryheap.NewMin()

	nn := []int{5, 2, 1, 3, 4}
	for _, n := range nn {
		pq.Insert(n)
	}

	mins := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(mins); i++ {
		if n, ok := pq.ExtractMin(); !ok {
			t.Fatalf("expected to be ok")
		} else if n != mins[i] {
			t.Fatalf("have n: %v, want %v", n, mins[i])
		}
	}
}
