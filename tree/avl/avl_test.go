package avl_test

import (
	"testing"

	"github.com/rodolfo-r/datastructures/tree/avl"
)

func TestNew(t *testing.T) {
	_ = avl.New()
}

func TestInsert(t *testing.T) {
	tree := avl.New()
	for n := 0; n <= 2; n++ {
		tree.Insert(n)
	}
}

func TestInOrder(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	tt := []struct {
		name string
		in   []int
	}{
		{name: "Sorted", in: []int{1, 2, 3, 4, 5}},
		{name: "Reversed", in: []int{5, 4, 3, 2, 1}},
		{name: "Mixed", in: []int{3, 4, 2, 5, 1}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tree := avl.New()
			for _, n := range tc.in {
				tree.Insert(n)
			}

			nn := []int{}
			fn := func(n *avl.Node) {
				nn = append(nn, n.Value)
			}

			tree.InOrder(fn)
			for i := range nn {
				if nn[i] != sorted[i] {
					t.Fatalf("have %v, want %v", nn[i], sorted[i])
				}
			}
		})
	}
}
