package main

import (
	"math/big"
)

type Fib struct {
	Input  int       `json:"input"`
	Output *big.Int  `json:"fibonacci"`
}

type Fibs []Fib
