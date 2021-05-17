package glenn

import "testing"

func runTest(t *testing.T, input []int, expected []int) {
	result := removeDuplicates(input)

	if result != len(expected) {
		t.Errorf("runTest: expected %v, got %v", len(expected), result)
	}

	for i, v := range expected {
		if input[i] != v {
			t.Errorf("runTest: expected %v, got %v at index %v", v, input[i], i)
		}
	}
}

func TestCase1(t *testing.T) {
	runTest(t, []int{1,2,3}, []int{1,2,3})
}

func TestCase2(t *testing.T) {
	runTest(t, []int{0,0,1,1,1,2,2,3,3,4}, []int{0,1,2,3,4})
}

func TestCase3(t *testing.T) {
	runTest(t, []int{}, []int{})
}

