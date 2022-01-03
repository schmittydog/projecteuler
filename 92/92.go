package main

import "fmt"

var (
	eight9s = map[int]bool{89: true}
	ones = map[int]bool{1: true}
)

func next(n int) int {
	ret := 0
	for n > 0 {
		d := n%10
		ret += d*d
		n /= 10
	}
	return ret
}

func chain(n int) bool {
	arr := []int{}
	for {
		_, ok := ones[n]
		if ok {
                        for _, i := range arr {
                                ones[i] = true
                        }
			return false
		}
		_, ok = eight9s[n]
		if ok {
                        for _, i := range arr {
                                eight9s[i] = true
                        }
			return true
		}
		arr = append(arr, n)
		n = next(n)
	}
}

func main() {
	count := 0
	for x := 1; x < 10000000; x++ {
		if chain(x) {
			count += 1
		}
	}
	fmt.Println(count)
}
