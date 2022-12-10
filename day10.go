package main

import (
	"bufio"
	"fmt"
	"os"
)

var cycleTimes = map[string]int{
	"noop": 1,
	"addx": 2,
}

func main() {
	file, err := os.Open("day10_1.txt")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	curCycle := 1
	x := 1
	cycleToCheck := 20
	sigStrengthSum := 0
	for scanner.Scan() {
		var opCode string
		var opValue int
		fmt.Sscanf(scanner.Text(), "%s %d", &opCode, &opValue)
		opTime := cycleTimes[opCode]
		// fmt.Printf("%d: code: %s opval: %d x: %d\n", curCycle, opCode, opValue, x)

		if curCycle+opTime > cycleToCheck {
			sigStrength := cycleToCheck * x
			sigStrengthSum += sigStrength
			fmt.Printf("  %d %d %d\n", cycleToCheck, sigStrength, sigStrengthSum)
			cycleToCheck += 40
		}

		curCycle += opTime
		if opCode == "addx" {
			x += opValue
		}
	}
	// sample 13140
}
