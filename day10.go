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

func day10(fileName string) {
	file, err := os.Open(fileName)
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
		// fmt.Printf("%d: code: %s opval: %d x: %d\n", curCycle, opCode, opValue, x)

		for opTime := cycleTimes[opCode]; opTime > 0; opTime-- {
			rayX := (curCycle % 40) - 1
			if x-rayX < 2 && x-rayX > -2 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
			curCycle++
			if (curCycle-1)%40 == 0 {
				fmt.Println()
			}
			if curCycle == cycleToCheck {
				sigStrength := cycleToCheck * x
				sigStrengthSum += sigStrength
				// fmt.Printf("  %d %d %d\n", cycleToCheck, sigStrength, sigStrengthSum)
				cycleToCheck += 40
			}

		}
		// if curCycle+opTime > cycleToCheck {
		// 	sigStrength := cycleToCheck * x
		// 	sigStrengthSum += sigStrength
		// 	fmt.Printf("  %d %d %d\n", cycleToCheck, sigStrength, sigStrengthSum)
		// 	cycleToCheck += 40
		// }

		// curCycle += opTime
		if opCode == "addx" {
			x += opValue
		}
	}
	// sample 13140
}
