package main
// performs the task in 4.5-4.75 seconds
import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	cum := 0
	var d int
	inputList := []int{}
	seen := []int{}
	seen = append(seen, cum)
	// read in from STDIN
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
	length := len(inputList)
	var i int
	for {
		if i == length {
			i = 0
		}
		cum += inputList[i]
		i++
		for _, previous := range seen {
			if cum == previous {
				fmt.Println(cum)
				os.Exit(0)
			}
		}
		seen = append(seen, cum)
	}
}
