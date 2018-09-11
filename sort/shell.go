package sort

// Shell applies shell sort to nums.
func Shell(nums []int) {
	shell(nums, len(nums)/2)
}

func shell(nums []int, gap int) {
	if gap < 1 {
		return
	}

	for i := 0; i+gap < len(nums); i += gap {
		for j, k := i+gap, i; k >= 0 && nums[j] < nums[k]; j, k = j-gap, k-gap {
			nums[j], nums[k] = nums[k], nums[j]
		}
	}
	shell(nums, gap/2)
}
