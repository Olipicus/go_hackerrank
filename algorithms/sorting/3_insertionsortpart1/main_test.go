package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {

	tt := []struct {
		input  []int
		result []int
	}{
		{input: []int{1, 2, 3, 4, 5}, result: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 4, 3, 2, 1}, result: []int{1, 2, 3, 4, 5}},
		{input: []int{3, 2, 1, 4, 5}, result: []int{1, 2, 3, 4, 5}},
	}

	for index, tc := range tt {
		if result := InsertionSort(tc.input); fmt.Sprintf("%v", tc.result) != fmt.Sprintf("%v", result) {
			t.Errorf("Error on case[%d] input[%v] result[%v] expect[%v]", index, tc.input, result, tc.result)
		}
	}

}
