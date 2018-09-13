package singly

import (
	"fmt"
	"strings"
)

// Node is a singly linked List node.
type Node struct {
	Value interface{}
	next  *Node
	list  *List
}

// Next returns the previous list node or nil.
func (n *Node) Next() *Node {
	if next := n.next; n.list != nil && next != nil {
		return next
	}
	return nil
}

// List is an unorderd singly Linked list.
type List struct {
	head, tail *Node
}

// New creates a Linked List
func New() *List {
	return &List{}
}

// Append adds a node to the tail of a linked List.
// O(1)
func (l *List) Append(value interface{}) *Node {
	n := &Node{Value: value, list: l}
	if l.tail != nil {
		l.tail.next = n
	}
	l.tail = n

	if l.head == nil {
		l.head = n
	}
	return n
}

// Search looks for a node that has a value, and returns false if not found.
// O(n) time
func (l *List) Search(value interface{}) (*Node, bool) {
	for n := l.head; n != nil; n = n.next {
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
	if l.head == n {
		l.head = n.next
		return
	}

	for curr := l.head; curr != nil; curr = curr.next {
		if curr.next == n {
			if l.tail == n {
				curr.next, l.tail = nil, curr
				return
			}
			curr.next = curr.next.next
			n.next = nil
			break
		}
	}
}

// Reverse reverses all nodes following node head.
// O(n) time, O(n) space
func (l *List) Reverse() {
	var prev *Node
	n := l.head
	for n != nil {
		next := n.next
		n.next = prev
		prev = n
		n = next
	}
	l.tail, l.head = l.head, l.tail
}

// Each applies a function to every node.
func (l *List) Each(fun func(*Node)) {
	for n := l.head; n != nil; n = n.next {
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
	return l.head
}

// Back returns the last node in the list or nil if there is only root.
func (l *List) Back() *Node {
	return l.tail
}

// Equals compares every value in both linked Lists.
// O(n)
func (l *List) Equals(l2 *List) bool {
	for n, n2 := l.head, l2.head; n != nil || n2 != nil; n, n2 = n.next, n2.next {
		if n == nil || n2 == nil || n.Value != n2.Value {
			return false
		}
	}
	return true
}

// String creates a string representing the linked List.
func (l *List) String() string {
	var strs []string
	for n := l.head; n != nil; n = n.next {
		strs = append(strs, fmt.Sprintf("%+v", n.Value))
	}

	return strings.Join(strs, " <--> ")
}
