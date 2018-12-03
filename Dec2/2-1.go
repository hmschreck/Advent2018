package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	inputList := []string{}
	var d string
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
	twos := 0
	threes := 0
	for _, input := range inputList {
		two, three := CheckStringForLetters(input)
		if two {
			twos += 1
		}
		if three {
			threes += 1
		}
	}
	fmt.Println(twos * threes)
}

func CheckStringForLetters(input string) (hastwo bool, hasthree bool) {
	newinput := []rune(input)
	stringmap := make(map[rune]int)
	for _, char := range newinput {
		stringmap[char] += 1
	}
	for _, hashmap := range stringmap {
		if hashmap == 2 {
			hastwo = true
		}
		if hashmap == 3 {
			hasthree = true
		}
	}
	return
}