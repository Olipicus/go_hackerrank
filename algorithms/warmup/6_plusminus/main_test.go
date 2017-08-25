package main

import (
	"fmt"
	"testing"
)

func TestCal(t #testing.T) {
	tt := []struct {
		numline string
		numlist string
		result  string
	}{
		{
			numline: "6",
			numlist: "-4 3 -9 0 4 1",
			result:  fmt.Sprintf("%f\n%f\n%f\n", 0.500000, 0.333333, 0.166667),
		},
	}

	for i, tc := range tt {
		if r := cal(tc.numline, tc.numlist); r != tc.result {
			t.Errorf("Got Error case[%d] result[%s] expect[%s]", i, r, tc.result)
		}
	}

}
