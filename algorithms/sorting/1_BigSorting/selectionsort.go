package main

import "math/big"

func selection_sort(numlist *[]big.Int, position int) {

	l := *numlist
	min := l[position]
	minPosition := position

	for i := position; i < len(l); i++ {
		if min.Cmp(&l[i]) == 1 {
			min = l[i]
			minPosition = i
		}
	}

	l[position], l[minPosition] = min, l[position]

	if position < len(*numlist)-1 {
		selection_sort(numlist, position+1)
	}

}
