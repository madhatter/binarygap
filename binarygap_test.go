package main

import (
	"strconv"
	"testing"
)

func TestFindGap(t *testing.T) {
	tests := []struct {
		num, expectedGap int
	}{
		{12, 0},
		{6, 0},
		{328, 2},
		{5, 1},
		{16, 0},
		{1024, 0},
		{1162, 3},
		{5, 1},
		{51712, 2},
		{20, 1},
		{561892, 3},
		{9, 2},
		{66561, 9},
		{6291457, 20},
		{74901729, 4},
		{805306373, 25},
		{1376796946, 5},
		{1073741825, 29},
		{1610612737, 28},
	}

	for _, test := range tests {
		if observed := findGap(strconv.FormatInt(int64(test.num), 2)); observed != test.expectedGap {
			t.Fatalf("findGap(%v) = %v, want %v", test.num, observed, test.expectedGap)
		}
	}
}
