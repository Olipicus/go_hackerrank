package main

import "testing"

func TestSolution(t *testing.T) {
	tt := []struct {
		find    int
		numlist []int
		result  int
	}{
		{find: 4, numlist: []int{1, 4, 5, 7, 9, 12}, result: 1},
	}

	for index, tc := range tt {
		if r := solution(tc.find, tc.numlist); r != tc.result {
			t.Errorf("Got Error on case[%d] result[%d] expect[%d]", index, r, tc.result)
		}
	}

}
