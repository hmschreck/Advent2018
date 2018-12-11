package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var alphabet = strings.Split("abcdefghijklmnopqrstuvwxyz", "")

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}
	input = strings.TrimSuffix(input, "\n")

	characters := []rune(input)
	cycle := 0
	for {
		newCharacter, changes := React(characters)
		cycle += 1
		characters = newCharacter
		if changes == 0 {
			break
		}
	}
	fmt.Println("Cycles: ", cycle)
	fmt.Println(len(characters))
	minLength := 9999999999
	outputChan := make(chan int, 26)
	for _, letter := range alphabet {
		go func(inputX string, letter string, outputChannel chan int) {
			output := React2(inputX, letter)
			input := []rune(output)
			for {
				newCharacter, changes := React(input)
				input = newCharacter
				if changes == 0 {
					break
				}
			}
			outputChan <- len(input)
		}(input, letter, outputChan)
	}
	numbers := []int{}
	for i := 0; i < 26; i++ {
		numbers = append(numbers, <-outputChan)
	}
	for _, num := range numbers {
		if num < minLength {
			minLength = num
		}
	}
	fmt.Println(minLength)
}

func React(input []rune) (output []rune, changes int) {
	for i := 0; i < len(input)-1; i++ {
		if distance := int(input[i]) - int(input[i+1]); distance == 32 || distance == -32 {
			changes += 1
			output = append(output, input[i+2:]...)
			break
		} else {
			output = append(output, input[i])
		}
		if i >= len(input)-2 {
			output = input
			break
		}
	}
	return
}

func React2(input string, letter string) (output string) {
	replaceString := strings.ToLower(letter) + "|" + strings.ToUpper(letter)
	replaceRegex := regexp.MustCompile(replaceString)
	oldOutput := " "
	output = input
	for {
		if output == oldOutput {
			return
		}
		oldOutput = output
		output = replaceRegex.ReplaceAllString(oldOutput, "")
	}
	return

}
