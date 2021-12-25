package main

import (
	"fmt"
)

func sumSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i*i
	}
	return sum
}

func squareSum(n int) int {
	sum := n*(n+1)/2
	return sum*sum
}

func main() {
	fmt.Println(squareSum(100) - sumSquares(100))
}
