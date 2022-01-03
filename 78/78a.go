package main

import "fmt"

func parts(n int) int {
	arr := make([]int, n+1)
	arr[0] = 1
	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			if arr[j] > 0 {
				arr[j+i] += arr[j]
			}
		}
	}
	//fmt.Println(arr)
	return arr[n]
}

func main() {
	for i := 55000; i < 56000; i++ {
		fmt.Println(i, parts(i))
	}
}
