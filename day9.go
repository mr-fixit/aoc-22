package main

import (
	"bufio"
	"fmt"
	"os"
)

const nKnots = 2

var knotX = make([]int, 2)
var knotY = make([]int, 2)

func day9(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	dxdy := map[string][]int{
		"U": {0, -1},
		"D": {0, 1},
		"L": {-1, 0},
		"R": {1, 0},
	}
	for i := 0; i < nKnots; i++ {
		knotX[i] = 0
		knotY[i] = 0
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var direction string
		var nTimes int
		fmt.Sscanf(scanner.Text(), "%s %d", &direction, &nTimes)
		fmt.Println(direction, nTimes, dxdy[direction])
		for move := 0; move < nTimes; move++ {
			moveKnot(0, dxdy[direction][0], dxdy[direction][1])
		}
	}
}

func moveKnot(knotIndex int, dx int, dy int) {
	knotX[knotIndex] += dx
	knotY[knotIndex] += dy
	fmt.Println(knotX[knotIndex], knotY[knotIndex])
}
