package queue

import "github.com/techmexdev/algos/container"

// Queue is a FIFO container.
type Queue interface {
	Enqueue(int)
	Dequeue() (int, error)
	container.Container
}
