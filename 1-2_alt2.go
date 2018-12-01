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

func main() {
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
	var freqlist [180000]bool
	length := len(inputList)
	var i int
	init := make([]int, length)
	cum := 0

	for {
		if i == length {
			break
		}
		cum += inputList[i]
		init[i] = cum
		i++
	}
	net := init[len(init) - 1]
	i = 0
	timeStart := time.Now().UnixNano()

	for {
		if i == length {
			i = 0
		}
		num := init[i]
		if freqlist[num] {
			fmt.Println(time.Now().UnixNano() - timeStart)
			fmt.Println(num)
			os.Exit(0)
		} else {
			freqlist[num] = true
		}
		init[i] += net
		i++
	}
}