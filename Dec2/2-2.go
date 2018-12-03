package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
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
	timeStart := time.Now().UnixNano()
	for i, checkString := range inputList {
		solution, runes := FindOneOff(i, checkString, inputList)
		if solution {
			fmt.Println(string(runes))
			fmt.Println(time.Now().UnixNano() - timeStart)
			os.Exit(0)
		}
	}
}

func FindOneOff(start int, input string, inputList []string) (solution bool, output []rune) {
	inputRunes := []rune(input)
	for i := start + 1; i < len(inputList); i++ {
		checkRunes := []rune(inputList[i])
		outputRunes := []rune{}
		for i, checkRune := range inputRunes {
			if checkRune == checkRunes[i] {
				outputRunes = append(outputRunes, checkRune)
			}
		}
		if len(outputRunes) == len(inputRunes)-1 {
			solution = true
			output = outputRunes
			return
		}
	}
	solution = false
	return
}