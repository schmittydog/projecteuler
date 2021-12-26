package main

import (
	"fmt"
	"github.com/schmittydog/projecteuler/peuler"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	peuler.PrimesTo(1000000)
	for i := 2; i < 1000000; i++ {
		divCount := 1
		triNum := (i*i+i)/2
		for _, tup := range peuler.GetPrimeTups(triNum) {
			divCount *= tup[1] + 1
		}
		if divCount > 500 {
			fmt.Println(triNum)
			break
		}
	}
}
