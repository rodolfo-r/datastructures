package stack

import "github.com/techmexdev/algos/container"

// Stack is a LIFO container.
type Stack interface {
	Push(int)
	Pop() (int, error)
	container.Container
}
