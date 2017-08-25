package main

import "math/big"

type BigInt []big.Int

func (a BigInt) Len() int      { return len(a) }
func (a BigInt) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BigInt) Less(i, j int) bool {
	switch a[i].Cmp(&a[j]) {
	case -1:
		return true
	default:
		return false
	}
}
