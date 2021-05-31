package glenn

import "testing"

func runTest(t *testing.T, input []int, expected int) {
	result := maxProfit(input)

	if result != expected {
		t.Errorf("runTest: expected %v, got %v", expected, result)
	}
}

func TestCase1(t *testing.T) {
	runTest(t, []int{1}, 0)
}

func TestCase2(t *testing.T) {
	runTest(t, []int{7,1,5,3,6,4}, 5)
}

func TestCase4(t *testing.T) {
	runTest(t, []int{7,6,4,3,1}, 0)
}

func TestCase5(t *testing.T) {
	runTest(t, []int{17,12,15,13,16,14}, 4)
}

func TestCase6(t *testing.T) {
	runTest(t, []int{17,12,15,13,16,14,11}, 4)
}

func TestCase7(t *testing.T) {
	runTest(t, []int{17,12,15,13,16,14,11,18}, 7)
}

