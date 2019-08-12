package main

import (
	"fmt"
	"os"
)

func main() {

	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	// fmt.Println(s)
}

//Invoke
//go build echo-1.1.go
//./echo-1.1 Tom Jerry Itch Scratchy
