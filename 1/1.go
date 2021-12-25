package main

import (
	"fmt"
)

func main() {
	count := 0
	for num := 1; num < 1000; num++ {
		if num%3 == 0 || num%5 == 0 {
			count += num
		}
	}
	fmt.Println(count)
}
