package sort

// Insertion applies insertion sort to nums.
func Insertion(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}
