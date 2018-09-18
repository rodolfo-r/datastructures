package binaryheap

import (
	"github.com/rodolfo-r/datastructures"
)

// MaxHeap is a max binary heap.
type MaxHeap struct {
	List    []int
	lastPos int
}

// MaxNode is a node in the binary heap.
type MaxNode struct {
	Value int
}

func assertPriorityQueueMaxHeapImplementation() {
	var _ datastructures.PriorityQueue = (*MaxHeap)(nil)
}

// NewMax creates a MaxHeap.
func NewMax() *MaxHeap {
	return &MaxHeap{lastPos: -1}
}

// BuildMax creates a MaxHeap from nums.
func BuildMax(nums []int) *MaxHeap {
	h := &MaxHeap{List: nums, lastPos: len(nums) - 1}
	for i := h.lastPos; i >= 0; i-- {
		h.trickleDown(i)
	}
	return h
}

// Enqueue inserts an element to MaxHeap.
// O(log n)
func (h *MaxHeap) Enqueue(value int) {
	h.List = append(h.List, value)
	h.lastPos++
	h.trickleUp(h.lastPos)
}

func (h *MaxHeap) trickleUp(index int) {
	if index == 0 {
		return
	}

	if par, ok := h.parent(index); ok && h.List[index] > h.List[par] {
		h.List[index], h.List[par] = h.List[par], h.List[index]
		h.trickleUp(par)
	}
}

// Peek retrieves the smallest number
// O(1)
func (h *MaxHeap) Peek() (int, bool) {
	if h.lastPos > -1 {
		return h.List[0], true
	}
	return 0, false
}

// Dequeue retrieves the smallest number, and deletes it
// O(log n)
func (h *MaxHeap) Dequeue() (int, bool) {
	if h.lastPos < 0 {
		return 0, false
	}
	if h.lastPos == 0 {
		h.lastPos = -1
		return h.List[0], true
	}

	h.List[0], h.List[h.lastPos] = h.List[h.lastPos], h.List[0]
	h.lastPos--
	h.trickleDown(0)

	return h.List[h.lastPos+1], true
}

func (h *MaxHeap) trickleDown(index int) {
	if index == h.lastPos {
		return
	}

	if ch, ok := h.maxChild(index); ok && h.List[index] < h.List[ch] {
		h.List[index], h.List[ch] = h.List[ch], h.List[index]
		h.trickleDown(ch)
	}
}

func (h *MaxHeap) parent(index int) (int, bool) {
	if index == 0 {
		return 0, false
	}

	return (index - 1) / 2, true
}

func (h *MaxHeap) maxChild(index int) (int, bool) {
	ch1, ch2 := (2*index)+1, (2*index)+2
	if ch1 > h.lastPos && ch2 > h.lastPos {
		return 0, false
	}

	if ch2 > h.lastPos || h.List[ch1] > h.List[ch2] {
		return ch1, true
	}
	return ch2, true
}
