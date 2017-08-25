package main

import "testing"

func TestSolution(t *testing.T) {

	tt := []struct {
		numlist []int
		result  []int
	}{
		{numlist: []int{1, 2, 3, 4, 5, 6}, result: []int{6, 5, 4, 3, 2, 1}},
		{numlist: []int{8, 4, 2, 1, 8}, result: []int{8, 1, 2, 4, 8}},
	}

	for index, tc := range tt {
		reverse(tc.numlist)
		for i, num := range tc.numlist {
			if num != tc.result[i] {
				t.Fatalf("Got Error on [%d] result[%v] expect[%v]", index, tc.numlist, tc.result)
			}
		}
	}

}
