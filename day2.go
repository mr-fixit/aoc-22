package main

import (
	"bufio"
	"fmt"
	"os"
)

func day2() {
	// file, err := os.Open("data/day2_0.txt")
	file, err := os.Open("data/day2_1.txt")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}
	outcomes := make(map[string]string)
	outcomes["A X"] = "d"
	outcomes["A Y"] = "w"
	outcomes["A Z"] = "l"
	outcomes["B X"] = "l"
	outcomes["B Y"] = "d"
	outcomes["B Z"] = "w"
	outcomes["C X"] = "w"
	outcomes["C Y"] = "l"
	outcomes["C Z"] = "d"
	shapeScore := make(map[string]int)
	shapeScore["X"] = 1
	shapeScore["Y"] = 2
	shapeScore["Z"] = 3
	outcomeScore := make(map[string]int)
	outcomeScore["w"] = 6
	outcomeScore["l"] = 0
	outcomeScore["d"] = 3
	outcomeScore["X"] = 0
	outcomeScore["Y"] = 3
	outcomeScore["Z"] = 6
	move := make(map[string]string)
	move["A X"] = "Z"
	move["A Y"] = "X"
	move["A Z"] = "Y"
	move["B X"] = "X"
	move["B Y"] = "Y"
	move["B Z"] = "Z"
	move["C X"] = "Y"
	move["C Y"] = "Z"
	move["C Z"] = "X"

	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanWords)
	var total1, total2 int
	for scanner.Scan() {
		var m1, m2 string
		fmt.Sscan(scanner.Text(), &m1, &m2)
		outcome := outcomes[m1+" "+m2]
		move := move[m1+" "+m2]
		score1 := shapeScore[m2] + outcomeScore[outcome]
		total1 += score1
		score2 := shapeScore[move] + outcomeScore[m2]
		fmt.Println(m1, m2, move, shapeScore[move], outcomeScore[m2], score2, total2)
		total2 += score2
	}
	fmt.Println("part 1: ", total1)
	fmt.Println("part 2: ", total2)

}
