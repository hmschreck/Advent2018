package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Tree struct {
	Value bool
	True *Tree
	False *Tree
}

const offset = 250

func (tree *Tree) Populate(path []bool, outcome bool) {
	currentNode := tree
	for _, value := range path {
		if value {
			if currentNode.True == nil {
				currentNode.True = new(Tree)
			}
			currentNode = currentNode.True
		} else {
			if currentNode.False == nil {
				currentNode.False = new(Tree)
			}
			currentNode = currentNode.False
		}
		currentNode.Value = outcome
	}
}

func ParseInitialState(state string) (output []bool) {
	split := strings.Split(state, "")
	for i := 0; i < offset; i++ {
		output = append(output, false)
	}
	for _, value := range split {
		if value == "#" {
			output = append(output, true)
		} else {
			output = append(output, false)
		}
	}
	for i := 0; i < offset; i++ {
		output = append(output, false)
	}
	return
}

func main() {
	lines := []string{}
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		lines = append(lines, strings.TrimSuffix(text, "\n"))
	}
	tree := new(Tree)
	state := []bool{}
	for _, line := range lines {
		if strings.HasPrefix(line, "initial state") {
			state = ParseInitialState(strings.Replace(line, "initial state: ", "", 1))
		} else if strings.Contains(line, " => ") {
			path := []bool{}
			splitInput := strings.Split(line, " => ")
			for _, value := range strings.Split(splitInput[0], "") {
				if value == "#" {
					path = append(path, true)
				} else {
					path = append(path, false)
				}
			}
			outcome := (splitInput[1] == "#")
			tree.Populate(path, outcome)
		}
	}
	patterns := make([][]bool, 0)
	outerloop:
	for times := 0; times < 50000000000; times++ {
		nextState := make([]bool, len(state))
		for i := range state {
			currentNode := tree
			for j := i - 2; j <= i+2; j++ {
				if j < 0 || j >= len(state) {
					currentNode = currentNode.False
					continue
				}
				if state[j] == true {
					currentNode = currentNode.True
				} else {
					currentNode = currentNode.False
				}
			}
			nextState[i] = currentNode.Value
		}
		for i2 := range patterns {
			match := true
			for x, pattern := range patterns[i2] {
				if nextState[x] != pattern {
					match = false
				}
			}
			if match == true {
				fmt.Println(times-i2)
				state=nextState
				remainder := (50000000000-times+1)%(times-i2)-1
				state = patterns[i2+remainder]
				break outerloop
			}
		}
		patterns = append(patterns, nextState)
		state = nextState
		if times % 1000000 == 0 {
			fmt.Println(times)
		}
	}

	sum := 0
	for i, value := range state {
		if value  == true{
			sum += i - offset
		}
	}
	fmt.Println(sum)
}