package sort

// Selection applies selection sort to nums.
func Selection(nums []int) {
	selection(nums, 0)
}

func selection(nums []int, start int) {
	if start < len(nums)-1 {
		minIndex := findMinIndex(nums, start)
		nums[start], nums[minIndex] = nums[minIndex], nums[start]
		selection(nums, start+1)
	}
}

func findMinIndex(nums []int, start int) int {
	minIndex := start

	for i := start; i < len(nums); i++ {
		if nums[i] < nums[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}
