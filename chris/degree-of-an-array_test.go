package chris

import "testing"

func TestFindShortestSubArray(t *testing.T) {

	type testCase struct {
		input    []int
		expected int
	}

	tests := []testCase{
		{
			[]int{1, 2, 2, 3, 1},
			2,
		},
		{
			[]int{1, 2, 2, 3, 1, 4, 2},
			6,
		},
		{
			[]int{8, 8, 7, 7, 7},
			3,
		},
		{
			[]int{47, 47, 72, 47, 72, 47, 79, 47, 12, 92, 13, 47, 47, 83, 33, 15, 18, 47, 47, 47, 47, 64, 47, 65, 47, 47, 47, 47, 70, 47, 47, 55, 47, 15, 60, 47, 47, 47, 47, 47, 46, 30, 58, 59, 47, 47, 47, 47, 47, 90, 64, 37, 20, 47, 100, 84, 47, 47, 47, 47, 47, 89, 47, 36, 47, 60, 47, 18, 47, 34, 47, 47, 47, 47, 47, 22, 47, 54, 30, 11, 47, 47, 86, 47, 55, 40, 49, 34, 19, 67, 16, 47, 36, 47, 41, 19, 80, 47, 47, 27},
			99,
		},
	}

	for _, test := range tests {
		res := findShortestSubArray(test.input)

		if res != test.expected {
			t.Errorf("findShortestSubArray: expected %d, got %d for input: %v", test.expected, res, test.input)
		}
	}
}
