package chris

import "testing"

func TestSingleNumber(t *testing.T) {

	type testCase struct {
		input    []int
		expected int
	}

	tests := []testCase{
		{
			[]int{2, 2, 1},
			1,
		},
		{
			[]int{4, 1, 2, 1, 2},
			4,
		},
		{
			[]int{1},
			1,
		},
	}

	for _, test := range tests {
		res := singleNumber(test.input)

		if res != test.expected {
			t.Errorf("singleNumber: expected %d, got %d", test.expected, res)
		}
	}

}
