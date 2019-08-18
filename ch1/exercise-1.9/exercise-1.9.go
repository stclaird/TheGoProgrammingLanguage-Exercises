package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"log"
	"strings"
	"bytes"
)

func main() {

	for _, url := range os.Args[1:] {
  
		prefix := strings.HasPrefix(url, "http://")
        fetch_url := ""

		if prefix == false {
		   var buffer bytes.Buffer
		   buffer.WriteString("http://")
		   buffer.WriteString(url)
		   fetch_url = buffer.String()
		} else {
			fetch_url = url
		}

		resp, err := http.Get(fetch_url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		} 
        
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
          log.Fatal(err)
		}

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		
		fmt.Println("\nStatus:" + resp.Status)
	}
}