package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var stacks []string

func main() {
	const dayStr = "5"
	//file, err := os.Open("day" + dayStr + "_0.txt")
	file, err := os.Open("day" + dayStr + "_1.txt")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	readingMode := "stacks"
	var stacklines []string
	var stacks []string
	for scanner.Scan() {
		wholeLine := scanner.Text()
		//  [D]
		//  [N] [C]
		//  [Z] [M] [P]
		//   1   2   3

		// move 1 from 2 to 1
		// move 3 from 1 to 3
		// move 2 from 2 to 1
		// move 1 from 1 to 2
		fmt.Println(readingMode, "'", wholeLine, "'")
		if readingMode == "stacks" {
			if strings.Contains(wholeLine, "1") {
				readingMode = "moves"
				stacks = buildStacks(stacklines)
				fmt.Println("stacks: ", stacks)
				scanner.Scan()
			}
			stacklines = append(stacklines, wholeLine)
		} else { // reading moves
			var nTimes, fromStack, toStack int
			fmt.Sscanf(wholeLine, "move %d from %d to %d", &nTimes, &fromStack, &toStack)
			doMove(stacks, nTimes, fromStack, toStack)
			fmt.Println("stacks: ", stacks)
		}
	}
	var p1 = ""
	for i := 0; i < len(stacks); i++ {
		stack := stacks[i]
		lastChar := stack[len(stack)-1]
		p1 += string(lastChar)
	}
	fmt.Println("part1: ", p1)
}

func popChar(str string) (char string, newStr string) {
	lenMinus1 := len(str) - 1
	char = string(str[lenMinus1])
	newStr = str[0:lenMinus1]
	return
}

func doMove(stacks []string, nTimes, iFrom, iTo int) {
	//println(nTimes, iFrom, iTo)
	for i := 0; i < nTimes; i++ {
		fromStack := stacks[iFrom-1]
		char, newFrom := popChar(fromStack)
		stacks[iFrom-1] = newFrom
		newTo := stacks[iTo-1] + char
		stacks[iTo-1] = newTo
	}
}

func buildStacks(lines []string) (stacks []string) {
	lineLen := len(lines[1])
	nStacks := (lineLen + 1) / 4
	// println("build", lineLen, nStacks)
	stacks = make([]string, 0)
	for i := 0; i < nStacks; i++ {
		stacks = append(stacks, "")
	}

	for iLine := len(lines) - 1; iLine > -1; iLine-- {
		line := lines[iLine]
		// fmt.Println("line: ", line)
		for iStack := 0; iStack < nStacks; iStack++ {
			idx := iStack*4 + 1
			char := line[idx : idx+1]
			// fmt.Printf("char %s \n", char)
			if char != " " {
				stacks[iStack] = stacks[iStack] + char
			}
		}
		// fmt.Println()
	}
	return
}
