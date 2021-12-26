package main

import (
	"fmt"
	"strings"
	"strconv"

	"github.com/schmittydog/projecteuler/peuler"
)

func main() {
	a := peuler.GetFile("file")
	numStrings := strings.Split(a, "\n")
	count := 0
	for _, numString := range numStrings {
		front := numString[:13]
		num, _ := strconv.Atoi(front)
		count += num
	}
	for count > 10000000000 {
		count /= 10
	}
	fmt.Println(count)
}
