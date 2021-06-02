package chris

import "testing"

func TestMaxProfit(t *testing.T) {

	type testCase struct {
		input    []int
		expected int
	}

	tests := []testCase{
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
		{
			[]int{1, 2},
			1,
		},
	}

	for _, test := range tests {
		res := maxProfit(test.input)

		if res != test.expected {
			t.Errorf("maxProfit: expected %d, got %d for input: %v", test.expected, res, test.input)
		}
	}
}
