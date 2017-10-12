package main

import (
	"strconv"
	"testing"
)

func testfindGap(t *testing.T) {
	tests := []struct {
		num, expectedGap int
	}{
		{12, 0},
		{6, 0},
		{328, 2},
		{5, 1},
		{16, 0},
		{1024, 0},
	}

	for _, test := range tests {
		if observed := findGap(strconv.FormatInt(int64(test.num), 2)); observed != test.expectedGap {
			t.Fatalf("findGap(%v) = %v, want %v", test.num, observed, test.expectedGap)
		}
	}
}
