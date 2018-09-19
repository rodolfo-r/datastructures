package bst_test

import (
	"testing"

	"github.com/rodolfo-r/datastructures/tree/bst"
)

func TestNew(t *testing.T) {
	_ = bst.New()
}

func TestInsert(t *testing.T) {
	tree := bst.New()
	for n := 0; n < 12; n++ {
		_ = tree.Insert(n)
	}
}

func TestInOrderRecursive(t *testing.T) {
	ordered := []int{1, 2, 3, 4, 5}
	tt := []struct {
		name string
		in   []int
	}{
		{name: "Sorted", in: []int{1, 2, 3, 4, 5}},
		{name: "Reveresed", in: []int{5, 4, 3, 2, 1}},
		{name: "Mixed", in: []int{3, 4, 2, 5, 1}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tree := bst.New()
			for _, n := range tc.in {
				tree.Insert(n)
			}

			nn := []int{}
			fn := func(n *bst.Node) {
				nn = append(nn, n.Val)
			}

			bst.InOrder(tree.Root, fn)
			for i := range ordered {
				if nn[i] != ordered[i] {
					t.Fatalf("have %v, want %v", nn[i], ordered[i])
				}
			}
		})
	}
}
func TestPreOrderRecursive(t *testing.T) {
	tt := []struct {
		name     string
		in, want []int
	}{
		{name: "Sorted", in: []int{1, 2, 3, 4, 5}, want: []int{1, 2, 3, 4, 5}},
		{name: "Reveresed", in: []int{5, 4, 3, 2, 1}, want: []int{5, 4, 3, 2, 1}},
		{name: "Mixed", in: []int{3, 4, 2, 5, 1}, want: []int{3, 2, 1, 4, 5}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tree := bst.New()
			for _, n := range tc.in {
				tree.Insert(n)
			}

			nn := []int{}
			fn := func(n *bst.Node) {
				nn = append(nn, n.Val)
			}

			bst.PreOrder(tree.Root, fn)
			for i := range tc.want {
				if nn[i] != tc.want[i] {
					t.Fatalf("have %v, want %v", nn[i], tc.want[i])
				}
			}
		})
	}
}

func TestPostOrderRecursive(t *testing.T) {
	tt := []struct {
		name     string
		in, want []int
	}{
		{name: "Sorted", in: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
		{name: "Reveresed", in: []int{5, 4, 3, 2, 1}, want: []int{1, 2, 3, 4, 5}},
		{name: "Mixed", in: []int{3, 4, 2, 5, 1}, want: []int{1, 2, 5, 4, 3}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tree := bst.New()
			for _, n := range tc.in {
				tree.Insert(n)
			}

			nn := []int{}
			fn := func(n *bst.Node) {
				nn = append(nn, n.Val)
			}

			bst.PostOrder(tree.Root, fn)
			for i := range tc.want {
				if nn[i] != tc.want[i] {
					t.Fatalf("have %v, want %v", nn[i], tc.want[i])
				}
			}
		})
	}
}

func TestInOrderIterative(t *testing.T) {
	ordered := []int{1, 2, 3, 4, 5}
	tt := []struct {
		name string
		in   []int
	}{
		{name: "Sorted", in: []int{1, 2, 3, 4, 5}},
		{name: "Reveresed", in: []int{5, 4, 3, 2, 1}},
		{name: "Mixed", in: []int{3, 4, 2, 5, 1}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tree := bst.New()
			for _, n := range tc.in {
				tree.Insert(n)
			}

			nn := []int{}
			fn := func(n *bst.Node) {
				nn = append(nn, n.Val)
			}

			bst.InOrderIterative(tree.Root, fn)

			for i := range ordered {
				if nn[i] != ordered[i] {
					t.Fatalf("have %v, want %v", nn[i], ordered[i])
				}
			}
		})
	}
}

func TestPreOrderIterative(t *testing.T) {
	tt := []struct {
		name     string
		in, want []int
	}{
		{name: "Sorted", in: []int{1, 2, 3, 4, 5}, want: []int{1, 2, 3, 4, 5}},
		{name: "Reveresed", in: []int{5, 4, 3, 2, 1}, want: []int{5, 4, 3, 2, 1}},
		{name: "Mixed", in: []int{3, 4, 2, 5, 1}, want: []int{3, 2, 1, 4, 5}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tree := bst.New()
			for _, n := range tc.in {
				tree.Insert(n)
			}

			nn := []int{}
			fn := func(n *bst.Node) {
				nn = append(nn, n.Val)
			}

			bst.PreOrderIterative(tree.Root, fn)
			for i := range tc.want {
				if nn[i] != tc.want[i] {
					t.Fatalf("have %v, want %v", nn[i], tc.want[i])
				}
			}
		})
	}
}

func TestPostOrderIterative(t *testing.T) {
	tt := []struct {
		name     string
		in, want []int
	}{
		{name: "Sorted", in: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
		{name: "Reveresed", in: []int{5, 4, 3, 2, 1}, want: []int{1, 2, 3, 4, 5}},
		{name: "Mixed", in: []int{3, 4, 2, 5, 1}, want: []int{1, 2, 5, 4, 3}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tree := bst.New()
			for _, n := range tc.in {
				tree.Insert(n)
			}

			nn := []int{}
			fn := func(n *bst.Node) {
				nn = append(nn, n.Val)
			}

			bst.PostOrderIterative(tree.Root, fn)
			for i := range tc.want {
				if nn[i] != tc.want[i] {
					t.Fatalf("have %v, want %v", nn[i], tc.want[i])
				}
			}
		})
	}
}
