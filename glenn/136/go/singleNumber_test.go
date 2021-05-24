package glenn

import "testing"

func runTest(t *testing.T, input []int, expected int) {
	result := singleNumber(input)

	if result != expected {
		t.Errorf("runTest: expected %v, got %v", expected, result)
	}
}

func TestCase1(t *testing.T) {
	runTest(t, []int{2,2,1}, 1)
}

func TestCase2(t *testing.T) {
	runTest(t, []int{4,1,2,1,2}, 4)
}

func TestCase3(t *testing.T) {
	runTest(t, []int{4}, 4)
}

