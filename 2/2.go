package main

import (
	"fmt"
)

func fibGen(n int) chan int {
	c := make(chan int)
	go func() {
		x, y := 0, 1
		for {
			if x > n {
				close(c)
				break
			}
			c <- x
			x, y = y, x+y
		}
	}()
	return c
}

func main() {
	count := 0
	a := fibGen(4000000)
	for n := range a {
		if n%2 == 0 {
			count += n
		}
	}
	fmt.Println(count)
}
