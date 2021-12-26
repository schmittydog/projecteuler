package main

import (
	"fmt"
)

// a = m**2 - n**2
// c = 2*m*n
// b = m**2 + n**2
func main() {
	var a, b, c int
	for m := 2; m < 200; m++ {
		for n := 1; n < m; n++ {
			a = m*m - n*n
			b = 2 * m * n
			c = m*m + n*n
			if a+b+c == 1000 {
				fmt.Println(a * b * c)
			}
		}
	}
}
