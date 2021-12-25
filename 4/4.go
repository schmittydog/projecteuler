package main

import (
	"fmt"
)

func isPali(n int) bool {
	x := n
	revNum := 0
	for x > 0 {
		revNum *= 10
		revNum += x % 10
		x /= 10
	}
	return revNum == n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxNum := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			if isPali(i * j) {
				maxNum = max(i*j, maxNum)
			}
		}
	}
	fmt.Println(maxNum)
}
