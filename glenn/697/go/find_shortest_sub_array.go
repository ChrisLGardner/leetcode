package glenn

type summaryEntry struct {
	firstIndex int
	degree int
}

func findShortestSubArray(nums []int) int {

	summary := make(map[int]*summaryEntry)

	//degree := make(map[int]int, 0)
  maxDegree := -1
  minSpan := -1

	for idx, val := range nums {
		if _, ok := summary[val]; !ok {
			summary[val] = &summaryEntry{
				firstIndex: idx,
				degree: 0,
			}
		}
		(summary[val]).degree = summary[val].degree + 1
		span := idx - summary[val].firstIndex + 1

		if summary[val].degree > maxDegree || (summary[val].degree == maxDegree && span < minSpan) {
			maxDegree = summary[val].degree
			minSpan = span
		}
	}

	return minSpan
}

// Runtime: 24 ms, faster than 98.41% of Go online submissions for Degree of an Array.
// Memory Usage: 6.7 MB, less than 14.29% of Go online submissions for Degree of an Array.
// Fast but a tad memory hungry
