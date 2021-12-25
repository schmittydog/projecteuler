package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := os.Open("file")
	defer f.Close()
	i, _ := ioutil.ReadAll(f)
	for _, x := range i { fmt.Println(int(x) - 48) }
}
