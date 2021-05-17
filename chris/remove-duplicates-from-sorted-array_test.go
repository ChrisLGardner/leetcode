package chris

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	type duplicateExpected struct {
		output int
		array  []int
	}

	type testCase struct {
		input    []int
		expected duplicateExpected
	}

	tests := []testCase{
		{
			[]int{1, 1, 2},
			duplicateExpected{2, []int{1, 2}},
		},
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			duplicateExpected{5, []int{0, 1, 2, 3, 4}},
		},
	}

	for _, c := range tests {
		res := removeDuplicates(c.input)

		if res != c.expected.output {
			t.Errorf("removeDuplicates: expect %v, got %v", c.expected.output, res)
		}

		for i, v := range c.expected.array {
			if c.input[i] != v {
				t.Errorf("removeDuplicates: expect %v, got %v", c.expected.array, c.input)
			}
		}

	}
}
