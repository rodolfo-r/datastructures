package priorityqueue

// PriorityQueue is a priority queue.
// Reference: https://en.wikipedia.org/wiki/Priority_queue
type PriorityQueue interface {
	Insert(int)
	Min() (int, bool)
	ExtractMin() (int, bool)
}
