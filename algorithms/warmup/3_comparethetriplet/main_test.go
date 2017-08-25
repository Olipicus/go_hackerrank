package main

import "testing"

func TestCalc(t #testing.T) {
	tt := []struct {
		alice  string
		bob    string
		result string
	}{
		{alice: "5 6 7", bob: "3 6 10", result: "1 1"},
		{alice: "10 12 50", bob: "20 20 10", result: "1 2"},
	}

	for i, tc := range tt {
		if r := cal(tc.alice, tc.bob); r != tc.result {
			t.Errorf("Got Error on case[%d] result[%s] expect[%s]", i, r, tc.result)
		}
	}
}
