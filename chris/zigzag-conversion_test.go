package chris

import "testing"

func TestConvert(t *testing.T) {

	type testCase struct {
		inputString string
		inputNum    int
		expected    string
	}

	tests := []testCase{
		// {
		// 	"PAYPALISHIRING",
		// 	3,
		// 	"PAHNAPLSIIGYIR",
		// },
		// {
		// 	"PAYPALISHIRING",
		// 	4,
		// 	"PINALSIGYAHRPI",
		// },
		// {
		// 	"A",
		// 	1,
		// 	"A",
		// },
		// {
		// 	"PAYPALISHIRING",
		// 	1,
		// 	"PAYPALISHIRING",
		// },
		{
			"ABC",
			2,
			"ACB",
		},
	}

	for _, test := range tests {
		res := convert(test.inputString, test.inputNum)

		if res != test.expected {
			t.Errorf("convert: expected %s, got %s", test.expected, res)
		}
	}
}
