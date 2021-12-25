package main

import (
	"fmt"
)

func largePrime(n int) int {
	lp := 0
	for n%2 == 0 {
		lp = 2
		n /= 2
	}
	div := 3
	for n >= div*div {
		for n%div == 0 {
			lp = div
			n /= div
		}
		div += 2
	}
	if n > 1 {
		return n
	}
	return lp
}

func main() {
	fmt.Println(largePrime(600851475143))
}
