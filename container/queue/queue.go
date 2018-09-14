package queue

import "github.com/rodolfo-r/algos/container"

// Queue is a FIFO container.
type Queue interface {
	Enqueue(value interface{})
	Dequeue() (value interface{}, ok bool)
	container.Container
}
