package chris

func singleNumber(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	result := map[int]int{}
	for _, v := range nums {
		if val, ok := result[v]; val == 1 {
			delete(result, v)
		} else if !ok {
			result[v] = 1
		}
	}

	for k := range result {
		return k
	}

	return nums[0]
}
