package sort

// Quick applies quicksort to nums.
func Quick(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	pivot := nums[mid]
	left, right := []int{}, []int{}

	for _, n := range nums {
		if n == pivot {
			continue
		} else if n < pivot {
			left = append(left, n)
		} else {
			right = append(right, n)
		}
	}
	return append(append(Quick(left), pivot), Quick(right)...)
}
