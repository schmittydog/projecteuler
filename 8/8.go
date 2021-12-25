package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func getIntArray(fname string) []int {
	f, _ := os.Open(fname)
	bytes, _ := ioutil.ReadAll(f)
	arr := []int{}
	for _, b := range bytes {
		n := int(b) - 48
		if n < 0 || n > 9 {
			continue
		}
		arr = append(arr, n)
	}
	return arr
}

func arrayProd(arr []int) int {
	ret := 1
	for _, n := range arr {
		ret *= n
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := getIntArray("file")
	maxSum := 0
	for i := 13; i <= 1000; i++ {
		maxSum = max(maxSum, arrayProd(a[i-13:i]))
	}
	fmt.Println(maxSum)
}
