package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
	timeStart := time.Now().UnixNano()
	sumTable := GenerateSumTable(grid)
	maxValue = 0
	maxX = -1
	maxY = -1
	maxSize := -1
	for j := 0; j < 300; j++ {
		for i := 0; i < 300; i++ {
			A := sumTable[i][j]
			for size := 1; i + size < 300 && j + size < 300; size++ {
				D := sumTable[i+size][j+size]
				C := sumTable[i][j+size]
				B := sumTable[i+size][j]
				sum := D+A-C-B
				if sum > maxValue {
					maxValue = sum
					maxX = j + 2 // these are +2 because of the bounding boxes for the sum tables
					maxY = i + 2
					maxSize = size
				}
			}
		}
	}
	fmt.Println(maxX, ",", maxY, ",", maxSize)
	fmt.Println("Took ", float64(time.Now().UnixNano()-timeStart)/1000000, " milliseconds")
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

func GenerateSumTable(grid [][]int) (output [][]int){
	output = make([][]int, XY)
	for row := range output {
		output[row] = make([]int, XY)
	}
	for j := range grid {
		for i := range grid[j] {
			total := 0
			total += grid[i][j]
			if i > 0 {
				total += output[i-1][j]
			}
			if j > 0 {
				total += output[i][j-1]
			}
			if i > 0 && j > 0 {
				total -= output[i-1][j-1]
			}
			output[i][j] = total
		}
	}
	return
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