package queue

import "github.com/techmexdev/algos/container"

// Queue is a FIFO container.
type Queue interface {
	Enqueue(value interface{})
	Dequeue() (value interface{}, ok bool)
	container.Container
}
