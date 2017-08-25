package main

import "testing"

func TestCal(t *testing.T) {
	tt := []struct {
		numlist []int
		min     int
		max     int
	}{
		{numlist: []int{1, 2, 3, 4, 5}, min: 10, max: 14},
		{numlist: []int{2, 1, 5, 4, 3}, min: 10, max: 14},
	}

	for index, tc := range tt {
		if rmin, rmax := solution(tc.numlist); rmin != tc.min && rmax != tc.max {
			t.Errorf("Got error on case[%d] result[%d %d] expect[%d %d]", index, rmin, rmax, tc.min, tc.max)
		}
	}

}
