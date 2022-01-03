package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var cubes = map[string][]int{}


func convertSort(n int) string {
	arr := []string{}
	for n > 0 {
		digit := n%10
		arr = append(arr, strconv.Itoa(digit))
		n /= 10
	}
	sort.Strings(arr)
	return strings.Join(arr, "")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(convertSort(132576))
	for x := 1; x < 10000000; x++ {
		n := x*x*x
		c := convertSort(n)
		_, ok := cubes[c]
		if ok {
			cubes[c] = append(cubes[c], x)
		} else {
			cubes[c] = []int{x}
		}
	}
	mini := 1000000000000000000
	for _, v := range cubes {
		if len(v) == 5 {
			mini = min(mini, v[0]*v[0]*v[0])
		}
	}
	fmt.Println(mini)
}

