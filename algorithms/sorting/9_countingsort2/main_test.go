package main

import "testing"

func TestSolution(t *testing.T) {
	tt := []struct {
		unsort []int
		sort   []int
	}{
		{unsort: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, sort: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{unsort: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, sort: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{unsort: []int{9, 5, 1, 4, 7, 8, 3, 2, 6}, sort: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for index, tc := range tt {
		QuickSort(tc.unsort)
		for i, num := range tc.unsort {
			if num != tc.sort[i] {
				t.Fatalf("Got Error on case[%d] result[%v] expect[%v]", index, tc.unsort, tc.sort)
			}
		}
	}
}
