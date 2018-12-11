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

const ADD = 0

type CellNucleus struct {
	X, Y     int
	Contents []Cell
}

type Cell struct {
	X, Y int
}

func main() {
	inputList := []CellNucleus{}
	reader := bufio.NewReader(os.Stdin)
	maxX := 0
	maxY := 0
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		newEntry := GenerateCell(strings.TrimSuffix(text, "\n"))
		if newEntry.X > maxX {
			maxX = newEntry.X
		}
		if newEntry.Y > maxY {
			maxY = newEntry.Y
		}
		inputList = append(inputList, newEntry)
	}
	fmt.Println(inputList)
	grid := make([][]int, maxY+ADD/2)
	for i := range grid {
		grid[i] = make([]int, maxX+ADD/2)
	}

	for i := range grid {
		for j := range grid[i] {
			distances := make([]int, len(inputList))
			cell := Cell{
				X: j,
				Y: i,
			}
			for k, nucleus := range inputList {
				distanceX := cell.X - nucleus.X
				distanceY := cell.Y - nucleus.Y
				if distanceX < 0 {
					distanceX = 0 - distanceX
				}
				if distanceY < 0 {
					distanceY = 0 - distanceY
				}
				distances[k] = distanceX + distanceY
			}
			cellToAddTo := GetIDOfClosest(distances)
			if cellToAddTo == -1 {
				continue
			}
			inputList[cellToAddTo].Contents = append(inputList[cellToAddTo].Contents, cell)
		}
	}
	NonInfiniteCells := []CellNucleus{}
CellLoop:
	for _, wholeCell := range inputList {
		for _, cell := range wholeCell.Contents {
			if cell.X == 0 || cell.Y == 0 || cell.X == maxX+ADD/2-1 || cell.Y == maxY+ADD/2-1 {
				continue CellLoop
			}
		}
		NonInfiniteCells = append(NonInfiniteCells, wholeCell)
	}
	maxSize := 0
	for _, nucleus := range NonInfiniteCells {
		if len(nucleus.Contents) > maxSize {
			maxSize = len(nucleus.Contents)
		}
	}
	fmt.Println(maxSize)

	withinRange := 0
	for y, row := range grid {
		for x, _ := range row {
			cell := Cell{
				X: x,
				Y: y,
			}
			sum := 0
			for _, nucleus := range inputList {
				sum += DistanceToNucleus(&cell, &nucleus)
			}
			if sum < 10000 {
				withinRange += 1
			}
		}
	}
	fmt.Println(withinRange)

}

func GenerateCell(input string) (output CellNucleus) {
	nums := strings.Split(input, ",")
	output.X, _ = strconv.Atoi(nums[0])
	output.X += ADD / 2
	output.Y, _ = strconv.Atoi(strings.Replace(nums[1], " ", "", 1))
	output.Y += ADD / 2
	return
}

func GetIDOfClosest(input []int) (output int) {
	min := 99999999999
	output = 0
	for i, value := range input {
		if value < min {
			min = value
			output = i
		}
	}
	count := 0
	for _, value := range input {
		if value == min {
			count += 1
		}
	}
	if count > 1 {
		output = -1
	}
	return
}

func DistanceToNucleus(cell *Cell, nucleus *CellNucleus) (output int) {
	distanceX := cell.X - nucleus.X
	if distanceX < 0 {
		distanceX = 0 - distanceX
	}
	distanceY := cell.Y - nucleus.Y
	if distanceY < 0 {
		distanceY = 0 - distanceY
	}
	return distanceX + distanceY
}
