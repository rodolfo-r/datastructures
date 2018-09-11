package queue

import (
	"errors"

	"github.com/techmexdev/algos/doubly"
)

// Queue represents a queue.
type Queue struct {
	list *doubly.List
	len  int
}

// New creates a queue from a doubly linked list.
func New() *Queue {
	return &Queue{list: doubly.New()}
}

// Enqueue adds an element to the queue.
func (q *Queue) Enqueue(val int) {
	q.list.Append(val)
	q.len++
}

// Dequeue removes the element from the front of the queue.
func (q *Queue) Dequeue() (int, error) {
	f := q.list.Front()
	if f == nil {
		return 0, errors.New("no more elements in queue")
	}

	val := f.Val
	q.list.Delete(f)
	q.len--
	return val, nil
}

// Size returns the length of the list.
func (q *Queue) Size() int {
	return q.len
}
