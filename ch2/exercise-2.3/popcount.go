package main

import (
	"fmt"
	"strconv"

	"gopl.io/ch2/popcount"
)

func main() {

	if pop_v, err := strconv.ParseUint(arg, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", pop_v, pop_v)
		fmt.Println(pop_v)
		pop := popcount.PopCount(pop_v)
		fmt.Println(pop)
	}

}
