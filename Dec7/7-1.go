package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type Step struct {
	Letter string
	Prev []Step
	Next []Step
}


func main() {
	stepsRegex := regexp.MustCompile("Step (.) must be finished before step (.) can begin.")
	StepsList := make(map[string]Step, 0)
	for _, value := range []string(strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")) {
		StepsList[value] = Step{Letter: value}
	}
	inputList := []string{}
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		inputList = append(inputList, text)
	}
	for _, line := range inputList {
		steps := stepsRegex.FindStringSubmatch(line)
		StepsList[steps[1]].Next = append(StepsList[steps[1]].Next, StepsList[steps[2]])
	}


}