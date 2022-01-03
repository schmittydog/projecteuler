package main

import (
	"fmt"
)

var (
	cache = map[string]int{"1 1":1, "2 2":2}
	hits = 0
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func partitions(n, l int) int {
	l = min(n, l)
	if l == 1 || n == 0 {
		return 1
	}
	key := fmt.Sprintf("%d %d", n, l)
	val, ok := cache[key]
	if ok {
		fmt.Printf("hit: %d, %d\n", n, l)
		return val
	}
	sum := 0
	for i := min(n,l); i > 0; i-- {
		sum += partitions(n-i, i)
	}
	sum %= 1000000
	cache[key] = sum
	return sum
}

func main() {
	for x := 1; x < 100; x++ {
		a := partitions(x, x)
		key := fmt.Sprintf("%d %d", x, x)
		fmt.Println(x, a)
		if cache[key] == 0 {
			fmt.Println(x, hits)
			break
		}
	}
	fmt.Println(hits, len(cache))
}
