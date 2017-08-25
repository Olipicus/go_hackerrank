package main

import "testing"

func TestCal(t *testing.T) {

	tt := []struct {
		listnum []string
		result  int
	}{
		{listnum: []string{"11 2 4", "4 5 6", "10 8 -12"}, result: 15},
	}

	for i, tc := range tt {
		if r := cal(tc.listnum); r != tc.result {
			t.Errorf("Got Error case[%d] result[%d] expect[%d]", i, r, tc.result)
		}
	}

}
