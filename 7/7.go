package main

import (
	"fmt"

	"github.com/schmittydog/projecteuler/bitsieve"
)

const (
	max = 1000000
)

func main() {
	bs := bitsieve.NewBitSieve(max)
	primeCount := 0
	for n := 2; n < max; n++ {
		if bs.IsIn(n) {
			primeCount += 1
		}
		if primeCount == 10001 {
			fmt.Println(n)
			break
		}
	}
}
