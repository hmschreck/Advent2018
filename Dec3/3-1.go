package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Claim struct {
	ClaimNumber int
	Down        int
	Right       int
	Width       int
	Height      int
}

func main() {
	inputList := []Claim{}
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		claim := new(Claim)
		breakString := strings.Split(text, " ")
		claim.ClaimNumber, _ = strconv.Atoi(strings.Replace(breakString[0], "#", "", 1))
		position := strings.Split(strings.Replace(breakString[2], ":", "", 1), ",")
		claim.Down, _ = strconv.Atoi(position[1])
		claim.Right, _ = strconv.Atoi(position[0])
		size := strings.Split(breakString[3], "x")
		claim.Width, _ = strconv.Atoi(size[0])
		claim.Height, _ = strconv.Atoi(strings.TrimSuffix(size[1], "\n"))
		inputList = append(inputList, *claim)
	}

	var grid [2000][2000]int
	for _, newClaim := range inputList {
		for x := 0; x < newClaim.Height; x++ {
			for y := 0; y < newClaim.Width; y++ {
				grid[newClaim.Down+x][newClaim.Right+y] += 1
			}
		}
	}
	total := 0
	for _, row := range grid {
		for _, column := range row {
			if column > 1 {
				total += 1
			}
		}
	}
	fmt.Println(total)

}
