package main

import (
	"math/big"
	"testing"
)

type testcase struct {
	numlist []big.Int
	result  []big.Int
}

func addBigNum(numlist []big.Int, num string) []big.Int {

	bnum := big.NewInt(0)
	bnum, _ = bnum.SetString(num, 10)
	numlist = append(numlist, *bnum)

	return numlist
}

func TestSolution(t *testing.T) {
	tt := []testcase{}

	numlist := []big.Int{}
	numlist = addBigNum(numlist, "31415926535897932384626433832795")
	numlist = addBigNum(numlist, "1")
	numlist = addBigNum(numlist, "3")
	numlist = addBigNum(numlist, "10")
	numlist = addBigNum(numlist, "3")
	numlist = addBigNum(numlist, "5")

	result := []big.Int{}
	result = addBigNum(result, "1")
	result = addBigNum(result, "3")
	result = addBigNum(result, "3")
	result = addBigNum(result, "5")
	result = addBigNum(result, "10")
	result = addBigNum(result, "31415926535897932384626433832795")

	tt = append(tt, testcase{numlist: numlist, result: result})

	for index, tc := range tt {
		selection_sort(&tc.numlist, 0)
		for i, num := range tc.numlist {
			if num.Cmp(&tc.result[i]) != 0 {
				t.Fatalf("Got Error on case[%d] result[%v] expect[%v]", index, tc.numlist, tc.result)
			}
		}
	}
}
