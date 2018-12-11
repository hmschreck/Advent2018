package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	NumberOfChildren int
	Metadata         []int
	Children         []*Node
	Length           int
}

func CreateNode(nums []int) (node Node) {
	node.NumberOfChildren = nums[0]
	node.Metadata = make([]int, nums[1])
	i := 0
	child := 0
	for {
		if child >= node.NumberOfChildren {
			break
		} else {
			childNode := CreateNode(nums[2+i:])
			i += childNode.Length
			node.Children = append(node.Children, &childNode)
			child += 1
		}
	}
	childrenLength := 0
	for _, child := range node.Children {
		childrenLength += child.Length
	}
	node.Length = 2 + childrenLength + len(node.Metadata)
	metadata := nums[2+childrenLength:]
	for i := range node.Metadata {
		node.Metadata[i] = metadata[i]
	}
	return node
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}
	inputArray := strings.Split(strings.TrimSuffix(input, "\n"), " ")
	inputNums := []int{}
	for _, str := range inputArray {
		num, _ := strconv.Atoi(str)
		inputNums = append(inputNums, num)
	}
	timeStart := time.Now().UnixNano()
	root := CreateNode(inputNums)
	sum := GetSumOfMetaData(&root)
	fmt.Println(sum)
	fmt.Println("Ran in ", float64(time.Now().UnixNano()-timeStart)/1000000, " milliseconds")
	timeStart2 := time.Now().UnixNano()
	part2 := GetPart2(&root)
	fmt.Println(part2)
	fmt.Println("Got part two in ", float64(time.Now().UnixNano()-timeStart2)/1000000.0, "milliseconds.")
	fmt.Println("Overall time: ", float64(time.Now().UnixNano()-timeStart)/1000000.0, " milliseconds")
}

func GetSumOfMetaData(node *Node) (output int) {
	for _, data := range node.Metadata {
		output += data
	}
	for _, child := range node.Children {
		output += GetSumOfMetaData(child)
	}
	return
}

func GetPart2(node *Node) (output int) {
	if node.NumberOfChildren == 0 {
		for _, addend := range node.Metadata {
			output += addend
		}
		return
	}
	for _, metadata := range node.Metadata {
		i := metadata - 1
		if !(i < 0 || i >= node.NumberOfChildren) {
			output += GetPart2(node.Children[i])
		}
	}
	return
}
