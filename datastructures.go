package datastructures

// Queue is a FIFO container.
type Queue interface {
	Enqueue(value interface{})
	Dequeue() (value interface{}, ok bool)
}

// Stack is a LIFO container.
type Stack interface {
	Push(value interface{})
	Pop() (value interface{}, ok bool)
}

// PriorityQueue is a priority queue.
// Reference: https://en.wikipedia.org/wiki/Priority_queue
type PriorityQueue interface {
	Enqueue(int)
	Peek() (int, bool)
	Dequeue() (int, bool)
}
