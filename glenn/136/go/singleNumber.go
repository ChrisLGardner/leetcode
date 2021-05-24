package glenn

func singleNumber(nums []int) int {
	h := make(map[int]int, 0)
	for _, num := range nums {
		if _, p := h[num]; p {
			delete(h, num)
		} else {
			h[num] = 1
		}
	}

	for k := range h { return k }
	return -1
}

// Runtime: 24 ms, faster than 31.91% of Go online submissions for Single Number.
// Memory Usage: 6.3 MB, less than 13.76% of Go online submissions for Single Number.
