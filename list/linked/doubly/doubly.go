package doubly

import (
	"fmt"
	"strings"
)

// Node is linked List node.
type Node struct {
	Value      interface{}
	prev, next *Node
	list       *List
}

// Prev returns the previous list node or nil.
func (n *Node) Prev() *Node {
	if prev := n.prev; n.list != nil && prev != &n.list.root {
		return prev
	}
	return nil
}

// Next returns the previous list node or nil.
func (n *Node) Next() *Node {
	if next := n.next; n.list != nil && next != &n.list.root {
		return next
	}
	return nil
}

// List is an unorderd doubly Linked list.
type List struct {
	root Node // sentinel list element
	len  int
}

// New creates a Linked List
func New() *List {
	l := &List{}
	l.root.prev, l.root.next = &l.root, &l.root

	return l
}

// Append adds a node to the tail of a linked List.
// O(1)
func (l *List) Append(value interface{}) *Node {
	return l.insertAfter(&Node{Value: value}, l.root.prev)
}

// Prepend adds a node to the front of a Linked List.
// O(1)
func (l *List) Prepend(value interface{}) *Node {
	return l.insertAfter(&Node{Value: value}, &l.root)
}

// Search looks for a node that has a value, and returns false if not found.
// O(n) time
func (l *List) Search(value interface{}) (*Node, bool) {
	for n := l.root.next; n != &l.root; n = n.next {
		if n.Value == value {
			return n, true
		}
	}
	return &Node{}, false
}

// Get retrieves a node's value.
// O(n) time
func (l *List) Get(n *Node) (value interface{}) {
	return n.Value
}

// Delete removes a node from its Linked List.
// O(1)
func (l *List) Delete(n *Node) {
	n.prev.next, n.next.prev = n.next, n.prev
	n.prev, n.next, n.list = nil, nil, nil
	l.len--
}

// Size returns the length of l.
func (l *List) Size() int {
	return l.len
}

// Empty checks wether the length
// of the list is 0
func (l *List) Empty() bool {
	return l.len == 0
}

// Reverse reverses all nodes following node head.
// O(n) time, O(n) space
func (l *List) Reverse() {
	for n := l.root.next; n != &l.root; n = n.prev {
		n.prev, n.next = n.next, n.prev
	}

	l.root.next, l.root.prev = l.root.prev, l.root.next
}

// Each applies a function to every node.
func (l *List) Each(fun func(*Node)) {
	for n := l.root.next; n != &l.root; n = n.next {
		fun(n)
	}
}

// Map returns a copy of every node value applied to a function.
func (l *List) Map(fun func(interface{}) interface{}) *List {
	mapped := New()
	l.Each(func(n *Node) {
		mapped.Append(fun(n.Value))
	})
	return mapped
}

// Copy returns a copy of a linked List.
func (l *List) Copy() *List {
	return l.Map(func(value interface{}) interface{} {
		return value
	})
}

// Front returns the first node in the list or nil if there is only root.
func (l *List) Front() *Node {
	if l.root.next != &l.root {
		return l.root.next
	}
	return nil
}

// Back returns the last node in the list or nil if there is only root.
func (l *List) Back() *Node {
	if l.root.prev != &l.root {
		return l.root.prev
	}
	return nil
}

// Equals compares every value in both linked Lists.
// O(n)
func (l *List) Equals(l2 *List) bool {
	for n, n2 := l.root.next, l2.root.next; n != &l.root || n2 != &l2.root; n, n2 = n.next, n2.next {
		if n == &l.root || n2 == &l2.root || n.Value != n2.Value {
			return false
		}
	}
	return true
}

// String creates a string representing the linked List.
func (l *List) String() string {
	var strs []string
	for n := l.root.next; n != &l.root; n = n.next {
		strs = append(strs, fmt.Sprintf("%+v", n.Value))
	}

	return strings.Join(strs, " <--> ")
}

func (l *List) insertAfter(n, at *Node) *Node {
	next := at.next
	at.next = n
	n.list, n.prev, n.next = l, at, next
	next.prev = n

	l.len++
	return n
}
