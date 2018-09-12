package sort_test

import (
	"testing"

	"github.com/techmexdev/algos/sort"
)

func TestHeap(t *testing.T) {
	tt := []struct {
		name       string
		nums, want []int
	}{
		{
			name: "Ascending",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Descending",
			nums: []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Mixed",
			nums: []int{3, 1, 5, 4, 2},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "One num",
			nums: []int{1},
			want: []int{1},
		},
		{
			name: "Two nums",
			nums: []int{2, 1},
			want: []int{1, 2},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sort.Heap(tc.nums)
			if !equals(tc.nums, tc.want) {
				t.Fatalf("have %v, want %v", tc.nums, tc.want)
			}
		})
	}
}
