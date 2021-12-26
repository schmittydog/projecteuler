package main

import (
	"fmt"
)

func collatz(n int) int {
	count := 1
	for n > 1 {
		count += 1
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxCollatz := 0
	maxNum := 0
	for x := 1; x < 1000000; x++ {
		c := collatz(x)
		if c > maxCollatz {
			maxNum = x
			maxCollatz = c
		}
	}
	fmt.Println(maxNum)
}
