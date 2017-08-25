package main

import "testing"

func TestSolution(t *testing.T) {

	tt := []struct {
		time   string
		result string
	}{
		{time: "07:05:45PM", result: "19:05:45"},
		{time: "07:05:45AM", result: "07:05:45"},
		{time: "12:45:54PM", result: "12:45:54"},
		{time: "12:05:39AM", result: "00:05:39"},
	}

	for index, tc := range tt {
		if r := solution(tc.time); r != tc.result {
			t.Errorf("Got Error on Case[%d] result[%s] expect[%s]", index, r, tc.result)
		}
	}

}
