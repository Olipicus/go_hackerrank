package main

import "testing"

func TestSolution(t *testing.T) {

	tt := []struct {
		input  []int64
		result []int64
		peek   int64
	}{
		{input: []int64{5, 4, 3, 2, 1}, result: []int64{1, 2, 3, 4, 5}, peek: 1},
		{input: []int64{1, 2, 3, 4, 5}, result: []int64{1, 2, 3, 4, 5}, peek: 1},
		{input: []int64{1, 2, 0, 4, 5}, result: []int64{0, 1, 2, 4, 5}, peek: 0},
		{input: []int64{9, -439859749}, result: []int64{-439859749, 9}, peek: -439859749},
	}

	for index, tc := range tt {
		h := NewHeap()

		for _, inputnum := range tc.input {
			h.Add(inputnum)
		}

		if h.Peek() != tc.peek {
			t.Fatalf("Got Error on case[%d] check Peek result[%d] expect[%d]", index, h.Peek(), tc.peek)
		}

		for i, _ := range tc.input {
			num := h.Poll()

			if num != tc.result[i] {
				t.Fatalf("Got Error on case[%d] check Poll result[%d] expect[%d]", index, num, tc.result[i])
			}
		}

	}

}
