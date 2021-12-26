package main

import "fmt"

func S(n, d int) int {
	ret := 1
	for i := 0; i < d; i++ {
		ret *= (n-i)
		ret /= (i+1)
	}
	return ret
}

func main() {
	fmt.Println(S(40,20))
}
