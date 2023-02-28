package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	matrix, down, up := make([][]byte, numRows, numRows), 0, numRows-2
	for i := 0; i < len(s); {
		if down != numRows {
			matrix[down] = append(matrix[down], byte(s[i]))
			down++
			i++
		} else if up > 0 {
			matrix[up] = append(matrix[up], byte(s[i]))
			i++
			up--
		} else {
			up = numRows - 2
			down = 0
		}

	}
	solution := make([]byte, 0, len(s))
	for _, i := range matrix {
		for _, j := range i {
			solution = append(solution, j)
		}
	}
	return string(solution)
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 4
	s2 := convert(s, numRows)
	fmt.Println(s2)

}
