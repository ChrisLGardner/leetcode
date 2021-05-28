package glenn

import "testing"

func runTest(t *testing.T, input []int, expected int) {
	result := findShortestSubArray(input)

	if result != expected {
		t.Errorf("runTest: expected %v, got %v", expected, result)
	}
}

func TestCase1(t *testing.T) {
	runTest(t, []int{1}, 1)
}

func TestCase2(t *testing.T) {
	runTest(t, []int{1,2,2,3,1}, 2)
}

func TestCase3(t *testing.T) {
	runTest(t, []int{1,2,2,3,1,4,2}, 6)
}

