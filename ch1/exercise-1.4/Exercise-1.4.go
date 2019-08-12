package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Printf("no file set\n")
		countLines(os.Stdin, counts )
	} else {
		for _,arg :=range files {
			f,err :=os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err )
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
    for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}
		
	}
	
}

func countLines (f *os.File, counts map [string]int){
	input := bufio.NewScanner(f)
	for input.Scan() {
		     
		     count_str := f.Name() + "\t" + input.Text() 
         counts[count_str]++
	}
}


//invocation
//either go run Exercise-1.4.go and list multiple strings of text with duplicates
//of go run Exercise-1.4.go Exercise-1.4-files.txt