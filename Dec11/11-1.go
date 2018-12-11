package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var serial = 0
const XY = 300
var grid = [][]int{}

func main() {
	serial, _ = strconv.Atoi(os.Args[1])
	// make our grid
	grid = make([][]int, XY)
	for i := range grid {
		grid[i] = make([]int, XY)
	}
	for y, row := range grid {
		for x, _ := range row {
			grid[y][x] = CalculateFuelCell(x+1, y+1)
		}
	}
	maxValue := 0
	maxX := -1
	maxY := -1
	for y := 0; y < XY-2; y++ {
		for x := 0; x < XY-2; x++ {
			value := Calculate3x3(x, y)
			if value > maxValue {
				maxX = x + 1
				maxY = y + 1
				maxValue = value
			}
		}
	}
	fmt.Println(maxX,",",maxY)

	maxChannel := make(chan *MaxCell, 90000)
	for y := 0; y < XY; y++ {
		for x := 0; x < XY; x++ {
			go CalculateMax(x, y, maxChannel)
		}
	}
	maxResult := new(MaxCell)
	for i := 0; i < 90000; i++ {
		input := <-maxChannel
		if input.Max > maxResult.Max {
			maxResult = input
		}
	}
	fmt.Println(maxResult.X,",",maxResult.Y,",",maxResult.Size)
}

func Calculate3x3(x, y int) (output int) {
	sum := 0
	for yplus := 0; yplus < 3; yplus++ {
		for xplus := 0; xplus <3; xplus++ {
			sum += grid[y+yplus][x+xplus]
		}
	}
	output = sum
	return
}

type MaxCell struct {
	X int
	Y int
	Size int
	Max int
}

func CalculateMax(x, y int, outputChannel chan *MaxCell) {
	output := new(MaxCell)
	max := 0
	maxSize := -1
	output.X = x + 1
	output.Y = y + 1
	maxFails := 5
	fails := 0
	for i := 1; i + x < 300 && i + y < 300; i++ {
		sum := 0
		for yplus := 0; yplus < i; yplus++ {
			for xplus := 0; xplus < i; xplus++ {
				sum += grid[y+yplus][x+xplus]
			}
		}
		if sum > max {
			max = sum
			maxSize = i
			fails = 0
		} else {
			fails += 1
			if fails == maxFails {
				break
			}
		}
	}
	output.Max = max
	output.Size = maxSize
	outputChannel <- output
}

func CalculateFuelCell(x, y int) (output int) {
	rackID := x + 10 // since it's 0 indexed
	powerlevel := rackID * (y)
	powerlevel += serial
	powerlevel *= rackID
	formattedNumber := strings.Split(fmt.Sprintf("%03d", powerlevel), "")
	hundredsDigit, _ := strconv.Atoi(formattedNumber[len(formattedNumber)-3])
	powerlevel = hundredsDigit
	powerlevel -= 5
	output = powerlevel
	return
}