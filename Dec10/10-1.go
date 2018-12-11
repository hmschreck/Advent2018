package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type RocketShip struct {
	PositionX int
	PositionY int
	VelocityX int
	VelocityY int
}

type NightSky struct {
	Ships []*RocketShip
}

func (sky *NightSky) Tick() {
	for _, ship := range sky.Ships {
		ship.PositionX += ship.VelocityX
		ship.PositionY += ship.VelocityY
	}
}

func (sky *NightSky) Print() {
	fmt.Println("In Print")
	width, height, ymin, xmin := sky.BoundingBox()
	grid := make([][]int, height+1)
	for i := range grid {
		grid[i] = make([]int, width+1)
	}
	for _, ship := range sky.Ships {
		grid[ship.PositionY-ymin][ship.PositionX-xmin] += 1
	}
	for i := range grid {
		fmt.Println(grid[i])
	}
}

func (sky *NightSky) BoundingBox() (x, y, ymin, xmin int) {
	var xmax, ymax int
	ymin = 99999999999
	xmin = 99999999999
	for _, ship := range sky.Ships {
		if ship.PositionX < xmin {
			xmin = ship.PositionX
		}
		if ship.PositionX > xmax {
			xmax = ship.PositionX
		}
		if ship.PositionY < ymin {
			ymin = ship.PositionY
		}
		if ship.PositionY > ymax {
			ymax = ship.PositionY
		}
	}
	x = xmax - xmin
	y = ymax - ymin
	return
}

var regexShip = regexp.MustCompile("position=<(.*), (.*)> velocity=<(.*), (.*)>")

func GenerateShip(input string) (ship RocketShip) {
	output := regexShip.FindStringSubmatch(input)
	for i := 1; i < 5; i++ {
		output[i] = strings.TrimPrefix(output[i], " ")
	}
	ship.PositionX, _ = strconv.Atoi(output[1])
	ship.PositionY, _ = strconv.Atoi(output[2])
	ship.VelocityX, _ = strconv.Atoi(output[3])
	ship.VelocityY, _ = strconv.Atoi(output[4])
	return
}

func main() {
	sky := new(NightSky)
	sky.Ships = make([]*RocketShip, 0)
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		newEntry := GenerateShip(strings.TrimSuffix(text, "\n"))
		sky.Ships = append(sky.Ships, &newEntry)
	}
	timeStart := time.Now().UnixNano()
	area := 9999999999999999
	cycles := 0
	for {
		cycles += 1
		sky.Tick()
		width, height, _, _ := sky.BoundingBox()
		checkArea := width * height
		if area > checkArea {
			area = checkArea
		} else {
			fmt.Println(cycles - 1)
			fmt.Println(height, width)
			for _, ship := range sky.Ships {
				ship.PositionX -= ship.VelocityX
				ship.PositionY -= ship.VelocityY
			}
			sky.Print()
			fmt.Println(float64(time.Now().UnixNano()-timeStart) / 1000000.0)
			break
		}
	}
}
