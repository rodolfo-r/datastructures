package binaryheap_test

import (
	"testing"

	"github.com/rodolfo-r/datastructures"
	"github.com/rodolfo-r/datastructures/container/priorityqueue/binaryheap"
)

func TestMaxEnqueue(t *testing.T) {
	var pq datastructures.PriorityQueue = binaryheap.NewMax()

	nn := []int{5, 2, 1, 3, 4}
	for _, n := range nn {
		pq.Enqueue(n)
	}
}

func TestMaxPeek(t *testing.T) {
	var pq datastructures.PriorityQueue = binaryheap.NewMax()

	nn := []int{5, 2, 1, 3, 4}
	for _, n := range nn {
		pq.Enqueue(n)
	}

	if n, ok := pq.Peek(); !ok {
		t.Fatalf("expected to be ok")
	} else if n != 5 {
		t.Fatalf("have n %v, want %v", n, 1)
	}
}

func TestMaxDequeue(t *testing.T) {
	var pq datastructures.PriorityQueue = binaryheap.NewMax()

	nn := []int{5, 2, 1, 3, 4}
	for _, n := range nn {
		pq.Enqueue(n)
	}

	maxes := []int{5, 4, 3, 2, 1}
	for i := 0; i < len(maxes); i++ {
		if n, ok := pq.Dequeue(); !ok {
			t.Fatalf("expected to be ok")
		} else if n != maxes[i] {
			t.Fatalf("have n: %v, want %v", n, maxes[i])
		}
	}
}

func TestBuildMax(t *testing.T) {
	nn := []int{5, 2, 1, 3, 4}
	var pq datastructures.PriorityQueue = binaryheap.BuildMax(nn)

	maxes := []int{5, 4, 3, 2, 1}
	for i := 0; i < len(maxes); i++ {
		if n, ok := pq.Dequeue(); !ok {
			t.Fatalf("expected to be ok")
		} else if n != maxes[i] {
			t.Fatalf("have n: %v, want %v", n, maxes[i])
		}
	}
}
