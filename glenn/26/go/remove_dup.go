package glenn

func removeDuplicates(nums []int) int {
	if len(nums) < 2 { return len(nums) }

	lastIndex := 1
	lastNum := -1
	for index, num := range nums {
		if index == 0 {
			lastNum = num
			continue
		}
		if lastNum != num {
			lastNum = num
			if lastIndex != index {	nums[lastIndex] = nums[index] }
			lastIndex++
		}
	}

	nums = nums[0:lastIndex]

	return len(nums)
}
