package util

import (
	"testing"
)

func TestBits(t *testing.T) {
	tests := [][]int{
		{-1, 0},
		{-100, 0},
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 3},
		{7, 3},
		{8, 3},
		{9, 4},
		{15, 4},
		{16, 4},
		{17, 5},
		{31, 5},
		{32, 5},
		{33, 6},
		{64, 6},
		{65, 7},
	}
	for _, test := range tests {
		if in, res, exp := test[0], Bits(test[0]), test[1]; res != uint(exp) {
			t.Errorf("Expected %d bits for %d values, but got %d", exp, in, res)
		}
	}
}
