package doubly

import (
	"fmt"
	"strconv"
	"strings"
)

// List is a doubly Linked list.
type List struct {
	root Node // sentinel list element
}

// Node is linked List node.
type Node struct {
	Val        int
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

// New creates a Linked List
func New() *List {
	l := &List{}
	l.root.prev, l.root.next = &l.root, &l.root

	return l
}

// Insert inserts n after at, and returns n.
func (l *List) Insert(n, at *Node) *Node {
	next := at.next
	at.next = n
	n.list, n.prev, n.next = l, at, next
	next.prev = n

	return n
}

// Append adds a node to the tail of a linked List.
// O(1)
func (l *List) Append(val int) *Node {
	return l.Insert(&Node{Val: val}, l.root.prev)
}

// Prepend adds a node to the front of a Linked List.
// O(1)
func (l *List) Prepend(val int) *Node {
	return l.Insert(&Node{Val: val}, &l.root)
}

// Search looks for a node that has a value.
// O(n) time
func (l *List) Search(val int) (*Node, error) {
	for n := l.root.next; n != &l.root; n = n.next {
		if n.Val == val {
			return n, nil
		}
	}
	return &Node{}, fmt.Errorf("Node with value %v not in list", val)
}

// Delete removes a node from its Linked List.
// O(1)
func (l *List) Delete(n *Node) {
	n.prev.next, n.next.prev = n.next, n.prev
	n.prev, n.next, n.list = nil, nil, nil
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
func (l *List) Map(fun func(int) int) *List {
	mapped := New()
	l.Each(func(n *Node) {
		mapped.Append(fun(n.Val))
	})
	return mapped
}

// Copy returns a copy of a linked List.
func (l *List) Copy() *List {
	return l.Map(func(num int) int {
		return num
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
		if n == &l.root || n2 == &l2.root || n.Val != n2.Val {
			return false
		}
	}
	return true
}

// String creates a string representing the linked List.
func (l *List) String() string {
	var strs []string
	for n := l.root.next; n != &l.root; n = n.next {
		strs = append(strs, strconv.Itoa(n.Val))
	}

	return strings.Join(strs, " <--> ")
}
