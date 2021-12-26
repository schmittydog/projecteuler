package main

import (
	"fmt"

	"github.com/schmittydog/projecteuler/bitsieve"
)

func main() {
	var counter int
	bs := bitsieve.NewBitSieve(2000000)
	for i := 2; i <= 2000000; i++ {
		if bs.IsIn(i) {
			counter += i
		}
	}
	fmt.Println(counter)
}
