package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getArray(fname string) [][]int {
	f, _ := os.Open(fname)
	defer f.Close()
	b, _ := ioutil.ReadAll(f)
	str := string(b[:len(b)-1])
	arr := [][]int{}
	strs := strings.Split(str, "\n")
	for _, row := range strs {
		narr := []int{}
		for _, col := range strings.Split(row, " ") {
			num, _ := strconv.Atoi(col)
			narr = append(narr, num)
		}
		arr = append(arr, narr)
	}
	return arr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxProd(arr [][]int, i, j int) int {
	var maxProd int
	if i < 17 {
		maxProd = max(maxProd, arr[i][j]*arr[i+1][j]*arr[i+2][j]+arr[i+3][j])
	}
	if j < 17 {
		maxProd = max(maxProd, arr[i][j]*arr[i][j+1]*arr[i][j+2]*arr[i][j+3])
	}
	if i < 17 && j < 17 {
		maxProd = max(maxProd, arr[i][j]*arr[i+1][j+1]*arr[i+2][j+2]*arr[i+3][j+3])
	}
	if i < 17 && j > 3 {
		maxProd = max(maxProd, arr[i][j]*arr[i+1][j-1]*arr[i+2][j-2]*arr[i+3][j-3])
	}
	return maxProd
}

func main() {
	arr := getArray("file")
	var maxProd int
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			maxProd = max(maxProd, getMaxProd(arr, i, j))
		}
	}
	fmt.Println(maxProd)
}
