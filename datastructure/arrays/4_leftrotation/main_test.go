package main

import "testing"

func TestSolution(t *testing.T) {

	tt := []struct {
		numlist []int
		round   int
		result  []int
	}{
		{numlist: []int{1, 2, 3, 4, 5}, round: 4, result: []int{5, 1, 2, 3, 4}},
	}

	for index, tc := range tt {
		r := rotate(tc.numlist, tc.round)
		for i, num := range r {
			if num != tc.result[i] {
				t.Fatalf("Got Error on case[%d] result[%v] expect[%v]", index, r, tc.result)
			}
		}
	}

}
