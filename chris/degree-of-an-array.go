package chris

func findShortestSubArray(nums []int) int {

	type indices struct {
		start  int
		finish int
		count  int
	}
	count := map[int]indices{}

	for i, v := range nums {
		if _, ok := count[v]; !ok {
			count[v] = indices{
				start:  i,
				finish: i,
				count:  1,
			}
		} else {
			count[v] = indices{
				start:  count[v].start,
				finish: i,
				count:  count[v].count + 1,
			}
		}
	}

	highestCount := 0
	for k, v := range count {
		if v.count > highestCount {
			highestCount = v.count
			for i, j := range count {
				if j.count < highestCount {
					delete(count, i)
				}
			}
		} else if v.count < highestCount {
			delete(count, k)
		}
	}

	diff := len(nums)

	for _, v := range count {
		if newDiff := v.finish - (v.start - 1); newDiff < diff {
			diff = newDiff
		}
	}

	return diff

}
