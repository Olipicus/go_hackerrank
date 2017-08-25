package main

import "testing"

func TestCal(t #testing.T) {
	tt := []struct {
		listnum string
		result  string
	}{
		{listnum: "1000000001 1000000002 1000000003 1000000004 1000000005", result: "5000000015"},
	}

	for i, tc := range tt {
		if r := cal(tc.listnum); r != tc.result {
			t.Errorf("Got Error case[%d] result[%s] expect[%s]", i, r, tc.result)
		}
	}

}
