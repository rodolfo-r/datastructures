package bst

import "github.com/rodolfo-r/datastructures/container/stack/doublystack"

// Tree is a binary search tree.
// Reference: https://en.wikipedia.org/wiki/Binary_search_tree
type Tree struct {
	Root *Node
}

// Node is a binary search sub-tree
type Node struct {
	Val         int
	left, right *Node
}

// New creates a binary search tree.
func New() *Tree {
	return &Tree{}
}

// Insert creates a node with val, and
// adds it as a leaf to t.
func (t *Tree) Insert(val int) *Node {
	if t.Root == nil {
		t.Root = &Node{Val: val}
		return t.Root
	}
	n, par := t.Root, &Node{}

	for n != nil {
		par = n
		if val < n.Val {
			n = n.left
		} else {
			n = n.right
		}
	}

	if val < par.Val {
		par.left = &Node{Val: val}
		return par.left
	}

	par.right = &Node{Val: val}
	return par.right
}

// InOrder applies fn to each node in in-order traversal recursively.
func InOrder(n *Node, fn func(*Node)) {
	if n != nil {
		InOrder(n.left, fn)
		fn(n)
		InOrder(n.right, fn)
	}
}

// PreOrder applies fn to each node in pre-order traversal recursively.
func PreOrder(n *Node, fn func(*Node)) {
	if n != nil {
		fn(n)
		PreOrder(n.left, fn)
		PreOrder(n.right, fn)
	}
}

// PostOrder applies fn to each node in post-order traversal recursively.
func PostOrder(n *Node, fn func(*Node)) {
	if n != nil {
		PostOrder(n.left, fn)
		PostOrder(n.right, fn)
		fn(n)
	}
}

// InOrderIterative applies fn to each node in in-order traversal iteratively.
func InOrderIterative(root *Node, fn func(*Node)) {
	stack := doublystack.New()

	n := root
	for n != nil || !stack.Empty() {
		if n != nil {
			stack.Push(n)
			n = n.left
		} else {
			popped, _ := stack.Pop()
			n = popped.(*Node)
			fn(n)
			n = n.right
		}
	}

}

// PreOrderIterative applies fn to each node in pre-order traversal recursively.
func PreOrderIterative(root *Node, fn func(*Node)) {
	stack := doublystack.New()
	stack.Push(root)

	for !stack.Empty() {
		popped, _ := stack.Pop()
		n := popped.(*Node)
		fn(n)
		if n.right != nil {
			stack.Push(n.right)
		}

		if n.left != nil {
			stack.Push(n.left)
		}
	}
}

// PostOrderIterative applies fn to each node in post-order traversal recursively.
func PostOrderIterative(root *Node, fn func(*Node)) {
	stack1, stack2 := doublystack.New(), doublystack.New()

	stack1.Push(root)
	for !stack1.Empty() {
		popped, _ := stack1.Pop()
		n := popped.(*Node)
		stack2.Push(n)

		if n.left != nil {
			stack1.Push(n.left)
		}

		if n.right != nil {
			stack1.Push(n.right)
		}
	}

	for !stack2.Empty() {
		popped, _ := stack2.Pop()
		fn(popped.(*Node))
	}
}
