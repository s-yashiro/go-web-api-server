package main

import (
	"math/big"
)

func fibonacci(n int) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)
	for i := 0; i < n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return a
}
