package main

import (
	"fmt"
	"os"
)

func main() {

	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("%v %v\n", i, os.Args[i])
	}

}

//invoked  go run Exercise-1.2.go dog cat newt frog
