package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var stacksP1 []string
var stacksP2 []string

func day5() {
	const dayStr = "5"
	//file, err := os.Open("data/day" + dayStr + "_0.txt")
	file, err := os.Open("data/day" + dayStr + "_1.txt")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	readingMode := "stacks"
	var stacklines []string
	for scanner.Scan() {
		wholeLine := scanner.Text()
		//fmt.Println(readingMode, "'", wholeLine, "'")
		if readingMode == "stacks" {
			//  [D]
			//  [N] [C]
			//  [Z] [M] [P]
			//   1   2   3
			if strings.Contains(wholeLine, "1") {
				readingMode = "moves"
				//stacksP1 = buildStacks(stacklines)
				stacksP1 = buildStacks(stacklines)
				stacksP2 = buildStacks(stacklines)
				// fmt.Println("starting stack: ", stacksP1)
				scanner.Scan()
			}
			stacklines = append(stacklines, wholeLine)
		} else { // reading moves
			// move 1 from 2 to 1
			// move 3 from 1 to 3
			// move 2 from 2 to 1
			// move 1 from 1 to 2
			var nTimes, fromStack, toStack int
			fmt.Sscanf(wholeLine, "move %d from %d to %d", &nTimes, &fromStack, &toStack)
			stacksP1 = doMove1(stacksP1, nTimes, fromStack, toStack)
			stacksP2 = doMove2(stacksP2, nTimes, fromStack, toStack)
			//fmt.Println("stacks: ", stackP2)
		}
	}
	var p1, p2 = getAnswer(stacksP1), getAnswer(stacksP2)
	fmt.Println("part1: ", p1, "part2: ", p2)
	if p1 != "WSFTMRHPP" {
		println("fail part 1")
	}
	if p2 != "GSLCMFBRP" {
		println("fail part 2")
	}
}

func getAnswer(stacks []string) (answer string) {
	answer = ""
	for i := 0; i < len(stacks); i++ {
		stack := stacks[i]
		answer += string(stack[len(stack)-1])
	}
	return
}

func popChar(str string) (char string, newStr string) {
	lenMinus1 := len(str) - 1
	char = string(str[lenMinus1])
	newStr = str[0:lenMinus1]
	return
}

func doMove1(stacks []string, nTimes, iFrom, iTo int) (after []string) {
	after = stacks
	for i := 0; i < nTimes; i++ {
		fromStack := stacks[iFrom-1]
		char, newFrom := popChar(fromStack)
		after[iFrom-1] = newFrom
		newTo := stacks[iTo-1] + char
		after[iTo-1] = newTo
	}
	return
}

func doMove2(stacks []string, nTimes, iFrom, iTo int) (after []string) {
	//println(nTimes, iFrom, iTo)
	after = stacks
	fromStack := stacks[iFrom-1]
	fromLen := len(fromStack)
	moving := string(fromStack[fromLen-nTimes:])
	after[iFrom-1] = fromStack[0 : fromLen-len(moving)]
	newTo := stacks[iTo-1] + moving
	after[iTo-1] = newTo
	return
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
