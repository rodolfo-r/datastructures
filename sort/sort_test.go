package sort_test

import (
	"testing"

	"github.com/techmexdev/algos/sort"
)

func TestQuick(t *testing.T) {
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
			nn := sort.Quick(tc.nums)
			if !equals(nn, tc.want) {
				t.Fatalf("have %v, want %v", nn, tc.want)
			}
		})
	}
}

func equals(nums, nums2 []int) bool {
	if len(nums) != len(nums2) {
		return false
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums2[i] {
			return false
		}
	}
	return true
}
