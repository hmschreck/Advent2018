package main

import (
	"fmt"
	"io"
	"log"
)
// Solution:
// Load in
func main() {
	cum := 0
	var d int
	for {
		_, err := fmt.Scan(&d)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		cum += d
	}
	fmt.Println(cum)
}
