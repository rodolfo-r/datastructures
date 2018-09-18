package avl

import (
	"math"
)

// Tree is an AVL tree.
// Reference: https://en.wikipedia.org/wiki/AVL_tree
type Tree struct {
	root *Node
}

// Node is an AVL node.
type Node struct {
	Value               int
	left, right, parent *Node
}

// New creates an AVL tree.
func New() *Tree {
	return &Tree{}
}

// Insert adds a new node with value
// as a leaf, then balances the resulting new tree.
func (t *Tree) Insert(value int) {
	if t.root == nil {
		t.root = &Node{Value: value}
		return
	}

	leaf := t.insert(t.root, value)
	t.balance(leaf)
}

// Delete removes n from the tree,
// then balances the resulting new tree.
func (t *Tree) Delete(n *Node) {
}

// InOrder applies fn to each node
// in in-order traversal
func (t *Tree) InOrder(fn func(*Node)) {
	t.inOrder(t.root, fn)
}

func (t *Tree) insert(n *Node, value int) *Node {
	if value < n.Value {
		if n.left == nil {
			n.left = &Node{parent: n, Value: value}
			return n.left
		}
		return t.insert(n.left, value)
	}

	if n.right == nil {
		n.right = &Node{parent: n, Value: value}
		return n.right
	}
	return t.insert(n.right, value)
}

func (t *Tree) inOrder(n *Node, fn func(*Node)) {
	if n != nil {
		t.inOrder(n.left, fn)
		fn(n)
		t.inOrder(n.right, fn)
	}
}

func (t *Tree) balance(n *Node) {
	if n == nil {
		return
	}

	if lh, rh := height(n.left), height(n.right); lh-rh > 1 {
		if height(n.left.left) > height(n.left.right) {
			n = t.rightRotate(n)
		} else {
			t.leftRotate(n.left)
			n = t.rightRotate(n)
		}
	} else if rh-lh > 1 {
		if height(n.right.right) > height(n.right.left) {
			n = t.leftRotate(n)
		} else {
			t.rightRotate(n.right)
			n = t.leftRotate(n)
		}
	}

	if t.root == nil {
		t.root = n
		return
	}
	t.balance(n.parent)
}

func (t *Tree) rightRotate(n *Node) *Node {
	l := n.left
	n.left = l.right

	if l.right != nil {
		l.right.parent = n
	}

	l.parent = n.parent
	if n.parent == nil {
		t.root = l
	} else if n == n.parent.left {
		n.parent.left = l
	} else {
		n.parent.right = l
	}

	n.parent, l.right = l, n
	return l
}

func (t *Tree) leftRotate(n *Node) *Node {
	r := n.right
	n.right = r.left

	if r.left != nil {
		r.left.parent = n
	}

	r.parent = n.parent
	if n.parent == nil {
		t.root = r
	} else if n == n.parent.right {
		n.parent.right = r
	} else {
		n.parent.left = r
	}

	n.parent, r.left = r, n
	return r
}

func height(n *Node) int {
	if n == nil {
		return -1
	}
	h := int(math.Max(float64(height(n.left)), float64(height(n.right)))) + 1

	return h
}
