package rbt

import "github.com/rodolfo-r/datastructures/container/stack/doublystack"

// Tree is a red black tree.
// Reference: https://en.wikipedia.org/wiki/Red-black_tree
type Tree struct {
	root *Node
	size int
}

// Node is a red black tree node.
type Node struct {
	red                 bool
	Val                 int
	right, left, parent *Node
}

// New creates a red black tree.
func New() *Tree {
	return &Tree{}
}

// Insert adds a node as a leaf to t,
// then balances the resulting tree.
func (t *Tree) Insert(val int) {
	if t.root == nil {
		t.root = &Node{Val: val}
		return
	}

	n, parent := t.root, &Node{}
	for n != nil {
		parent = n
		if val < n.Val {
			n = n.left
		} else {
			n = n.right
		}
	}

	leaf := &Node{Val: val, red: true, parent: parent}

	if leaf.Val < parent.Val {
		parent.left = leaf
	} else {
		parent.right = leaf
	}

	t.fixColor(leaf)
}

// InOrder applies fn to each node,
// in in-order traversal.
func (t *Tree) InOrder(fn func(*Node)) {
	stack := doublystack.New()

	n := t.root

	for n != nil && stack.Size() != 0 {
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

func (t *Tree) fixColor(n *Node) {
	if n == t.root || n == nil {
		return
	}
	if n.red && n.parent.red {
		t.correctNode(n)
	}
	t.fixColor(n.parent)
}

func (t *Tree) correctNode(n *Node) {
	if un, ok := uncle(n); ok && un.red {
		n.parent.parent.red, n.parent.red, un.red = true, false, false
	} else {
		t.rotate(n)
		n.parent.red, n.red = false, true
		sib, ok := sibling(n)
		if ok {
			sib.red = true
		}
	}
}

func (t *Tree) rotate(n *Node) {
	grandpa := n.parent.parent
	if n == n.parent.left {
		if n.parent == grandpa.left {
			t.rightRotate(grandpa)
		} else {
			t.rightRotate(n.parent)
			t.leftRotate(grandpa)
		}
	} else if n.parent == grandpa.right {
		t.leftRotate(grandpa)
	} else {
		t.leftRotate(n.parent)
		t.rightRotate(grandpa)
	}
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

func uncle(n *Node) (*Node, bool) {
	grandpa := n.parent.parent
	if n == nil || grandpa == nil || grandpa.right == nil || grandpa.left == nil {
		return &Node{}, false
	}
	if grandpa.right == n.parent {
		return grandpa.left, true
	}

	return grandpa.right, true
}

func sibling(n *Node) (*Node, bool) {
	if n == nil || n.parent == nil || n.parent.left == nil || n.parent.right == nil {
		return &Node{}, false
	}

	if n == n.parent.right {
		return n.parent.left, true
	}
	return n.parent.right, true
}
