package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// performs the task in --
// Sweet jesus, 0.025 seconds
// Sweeter Jesus, the actual algorithm
// (lines 35 to 48) take 200k NANOSECONDS.
// That's 0.0002 seconds.
// ... I have a problem

func main() {
	// Read in from STDIN
	inputList := []int{}
	var d int
	for {
		_, err := fmt.Scan(&d)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		inputList = append(inputList, d)
	}
	// End Read in from STDIN
	// Setup variables
	var freqlist [180000]bool
	length := len(inputList)
	var i int
	cum := 0
	// End Setup variables
	timeStart := time.Now().UnixNano()

	for {
		// Loop over the input list and add it
		if i == length {
			i = 0
		}
		cum += inputList[i]
		i++
		// End loop over the input list
		// If we've seen this number already,
		// this is our second time.
		if freqlist[cum] {
			fmt.Println(time.Now().UnixNano() - timeStart)
			fmt.Println(cum)
			os.Exit(0)
			// if we haven't seen this number before
			// we have now
		} else {
			freqlist[cum] = true
		}
	}
}
