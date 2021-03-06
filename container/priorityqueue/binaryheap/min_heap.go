package binaryheap

import (
	"log"

	"github.com/rodolfo-r/datastructures"
)

// MinHeap is a min binary heap.
type MinHeap struct {
	list    []int
	lastPos int
}

// MinNode is a node in the binary heap.
type MinNode struct {
	Value int
}

func assertPriorityQueueMinHeapImplementation() {
	var _ datastructures.PriorityQueue = (*MinHeap)(nil)
}

// NewMin creates a MinHeap.
func NewMin() *MinHeap {
	return &MinHeap{lastPos: -1}
}

// Enqueue inserts an element to MinHeap.
// O(log n)
func (h *MinHeap) Enqueue(value int) {
	h.list = append(h.list, value)
	h.lastPos++
	h.trickleUp(h.lastPos)
}

func (h *MinHeap) trickleUp(index int) {
	if index == 0 {
		return
	}

	if par, ok := h.parent(index); ok && h.list[index] < h.list[par] {
		h.list[index], h.list[par] = h.list[par], h.list[index]
		h.trickleUp(par)
	}
}

// Peek retrieves the smallest number
// O(1)
func (h *MinHeap) Peek() (int, bool) {
	if h.lastPos > -1 {
		return h.list[0], true
	}
	return 0, false
}

// Dequeue retrieves the smallest number, and deletes it
// O(log n)
func (h *MinHeap) Dequeue() (int, bool) {
	log.Printf("h.list = %+v\n", h.list)
	if h.lastPos < 0 {
		return 0, false
	}
	if h.lastPos == 0 {
		h.lastPos = -1
		return h.list[0], true
	}

	h.list[0], h.list[h.lastPos] = h.list[h.lastPos], h.list[0]
	h.lastPos--
	h.trickleDown(0)
	return h.list[h.lastPos+1], true

}

func (h *MinHeap) trickleDown(index int) {
	if index == h.lastPos {
		return
	}

	if ch, ok := h.minChild(index); ok && h.list[index] > h.list[ch] {
		h.list[index], h.list[ch] = h.list[ch], h.list[index]
		h.trickleDown(ch)
	}
}

func (h *MinHeap) parent(index int) (int, bool) {
	if index == 0 {
		return 0, false
	}

	return (index - 1) / 2, true
}

func (h *MinHeap) minChild(index int) (int, bool) {
	ch1, ch2 := (2*index)+1, (2*index)+2
	if ch1 > h.lastPos && ch2 > h.lastPos {
		return 0, false
	}

	if ch2 > h.lastPos || h.list[ch1] < h.list[ch2] {
		return ch1, true
	}
	return ch2, true
}
