package doublyqueue

import (
	"github.com/techmexdev/algos/container/queue"
	"github.com/techmexdev/algos/list/linked/doubly"
)

// Queue represents a queue.
type Queue struct {
	list *doubly.List
	len  int
}

func assertQueueImplementation() {
	var _ queue.Queue = (*Queue)(nil)
}

// New creates a queue from a doubly linked list.
func New() *Queue {
	return &Queue{list: doubly.New()}
}

// Enqueue adds an element to the queue.
func (q *Queue) Enqueue(value interface{}) {
	q.list.Append(value)
	q.len++
}

// Dequeue removes the element from the front of the queue.
func (q *Queue) Dequeue() (value interface{}, ok bool) {
	f := q.list.Front()
	if f == nil {
		return 0, false
	}

	val := f.Value
	q.list.Delete(f)
	q.len--
	return val, true
}

// Size returns the length of the list.
func (q *Queue) Size() int {
	return q.len
}

// Front returns the value of the first element.
func (q *Queue) Front() interface{} {
	return q.list.Front().Value
}
