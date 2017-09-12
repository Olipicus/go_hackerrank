package main

import "testing"

func TestSolution(t *testing.T) {

	tt := []struct {
		input  []float64
		output float64
	}{
		{input: []float64{3, -7, 0}, output: 3},
	}

	for index, tc := range tt {
		result := Solution(tc.input)

		if result != tc.output {
			t.Fatalf("Got Error on Case[%d] input[%v] result[%v] expect[%v] ", index, tc.input, result, tc.output)
		}
	}

}
