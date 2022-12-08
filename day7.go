package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	nodeType string
	name     string
	parent   *Node
	subNodes map[string]*Node
	size     int
}

func readInput() (topNode *Node) {
	const dayStr = "7"
	//file, err := os.Open("day" + dayStr + "_0.txt")
	file, err := os.Open("day" + dayStr + "_1.txt")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	var curNode *Node

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wholeLine := scanner.Text()
		if wholeLine[0] == '$' {
			var cmd, dirName string
			fmt.Sscanf(wholeLine, "$ %s %s", &cmd, &dirName)
			//fmt.Println("cmd: ", cmd, dirName)
			if cmd == "cd" {
				if dirName == "/" {
					topNode = &Node{name: "/", subNodes: make(map[string]*Node)}
					curNode = topNode
				} else if dirName == ".." {
					curNode.parent.size += curNode.size
					curNode = curNode.parent
				} else {
					curNode = curNode.subNodes[dirName]
				}
			} else if cmd == "ls" {
				// don't do anything
			}
		} else {
			var sizeOrDir, nodeName string
			fmt.Sscanf(wholeLine, "%s %s", &sizeOrDir, &nodeName)
			//fmt.Println(nodeName, sizeOrDir)
			if sizeOrDir == "dir" {
				newNode := Node{name: nodeName, parent: curNode, nodeType: "dir", subNodes: make(map[string]*Node)}
				curNode.subNodes[nodeName] = &newNode
			} else {
				var thisSize int
				fmt.Sscanf(sizeOrDir, "%d", &thisSize)
				newNode := Node{name: nodeName, parent: curNode, size: thisSize, nodeType: "file"}
				curNode.subNodes[nodeName] = &newNode
				curNode.size += thisSize
			}
		}
	}
	for ; curNode != topNode; curNode = curNode.parent {
		curNode.parent.size += curNode.size
	}
	return topNode
}

func Print(node *Node, depth int) {
	fmt.Printf("                                             "[0 : depth*2])
	fmt.Println(node.name, node.nodeType, node.size)
	for _, subNode := range node.subNodes {
		Print(subNode, depth+1)
	}
	return
}

func SumOfThoseDirs(node *Node, sum *int) {
	if node.nodeType == "dir" && node.size <= 100000 {
		*sum += node.size
	}
	for _, subNode := range node.subNodes {
		SumOfThoseDirs(subNode, sum)
	}
}

func main() {
	var topNode = readInput()
	fmt.Println("topNode: ", topNode)
	Print(topNode, 0)

	sum := 0
	SumOfThoseDirs(topNode, &sum)
	fmt.Println("part 1:", sum)
}
