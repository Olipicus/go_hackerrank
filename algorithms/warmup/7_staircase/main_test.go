package main

import (
	"fmt"
	"testing"
)

func TestCal(t *testing.T) {
	tt := []struct {
		num    int
		result string
	}{
		{num: 3, result: fmt.Sprintf("  #\n ##\n###\n")},
		{num: 4, result: fmt.Sprintf("   #\n  ##\n ###\n####\n")},
	}

	for index, tc := range tt {
		if r := cal(tc.num); r != tc.result {
			t.Errorf("Got Error case[%d] expect[\n%s\n] got[\n%s\n]", index, tc.result, r)
		}
	}
}
