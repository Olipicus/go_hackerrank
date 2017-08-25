package main

import "testing"

func TestSolution(t *testing.T) {
	tt := []struct {
		numlist []int
		result  int
	}{
		{numlist: []int{1, 1, 1, 1, 1}, result: 5},
		{numlist: []int{3, 2, 1, 3}, result: 2},
	}

	for index, tc := range tt {
		if r := solution(tc.numlist); r != tc.result {
			t.Errorf("Got error on case[%d] result[%d] expect[%d]", index, r, tc.result)
		}
	}

}
