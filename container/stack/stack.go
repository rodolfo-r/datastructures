package stack

import "github.com/techmexdev/algos/container"

// Stack is a LIFO container.
type Stack interface {
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	container.Container
}
