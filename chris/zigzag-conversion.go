package chris

import (
	"strings"
)

func convert(s string, numRows int) string {

	if len(s) == 1 || numRows == 1 {
		return s
	}

	type grid struct {
		row int
		col int
	}

	zig := make([][]rune, numRows)
	for i := range zig {
		zig[i] = make([]rune, (len(s)/2)+1)
	}

	lastRow := 0
	cursor := grid{0, 0}

	for _, c := range s {
		if cursor.col%(numRows-1) == 0 {
			zig[cursor.row][cursor.col] = c
			lastRow = cursor.row
			cursor.row++
			if cursor.row >= numRows {
				cursor.col++
				cursor.row = 0
			}
		} else {
			zig[lastRow-1][cursor.col] = c
			lastRow--
			cursor.col++
			cursor.row = 0
		}
	}

	var sb strings.Builder
	for _, str := range zig {
		for _, r := range str {
			if r > 0 {
				sb.WriteRune(r)
			}
		}
	}
	res := sb.String()
	return res
}
