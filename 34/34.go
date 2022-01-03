package main

import "fmt"

var factorials = []int{1}


func digitFact(n int) int {
	ret := 0
	for n > 0 {
		digit := n%10
		ret += factorials[digit]
		n /= 10
	}
	return ret
}

func main() {
	//initialize factorials array
	for i := 1; i < 10; i++ {
		num := factorials[len(factorials)-1] * i
		factorials = append(factorials, num)
	}

	sum := 0
	for i := 3; i < 3628800; i++ {
		if i == digitFact(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
