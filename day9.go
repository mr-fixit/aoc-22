package main

import (
	"bufio"
	"fmt"
	"os"
)

/*

  tail position 0,0
	each line
		read line
		move N moves
			move k1
			if k2 moves
			  move k2
				if k3 moves... etc
					... if kN == tail
						add tail position
	sort/unique tailPositions
	count tailPositions
*/

const nKnots = 2

var knotX = make([]int, nKnots)
var knotY = make([]int, nKnots)
var tailPositions = map[string]bool{"0 0": true}

func day9(fileName string) {
	fmt.Println("Day 9")

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}
	var headX, headY int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var direction string
		var nTimes int
		_, err := fmt.Sscanf(scanner.Text(), "%s %d", &direction, &nTimes)
		if err != nil {
			fmt.Println("err: ", err)
			os.Exit(1)
		}
		fmt.Println(direction, nTimes)
		headMoveMap := map[string][]int{
			"R": {1, 0}, "U": {0, -1}, "L": {-1, 0}, "D": {0, 1},
		}
		headMove := headMoveMap[direction]
		fmt.Println("  headMove", headMove)
		for i := 0; i < nTimes; i++ {
			moveKnot(0, headMove[0], headMove[1])
		} // for nTimes
		fmt.Println("head now at ", headX, headY)
	}
	fmt.Println("part1: ", len(tailPositions))

}

func moveKnot(knotIndex int, dx int, dy int) {
	knotX[knotIndex] += dx
	knotY[knotIndex] += dy

	headX := knotX[knotIndex]
	headY := knotY[knotIndex]
	if knotIndex == nKnots-1 {
		tailPositionStr := fmt.Sprintf("%d %d", headX, headY)
		fmt.Println("moved tailPosition", headX, headY)
		tailPositions[tailPositionStr] = true
		return
	}
	tailX := knotX[knotIndex+1]
	tailY := knotY[knotIndex+1]
	var dx1, dy1 int
	deltaX := headX - tailX
	deltaY := headY - tailY
	if deltaX == 2 {
		dx1 = 1
		dy1 = deltaY
	} else if deltaX == -2 {
		dx1 = -1
		dy1 = deltaY
	} else if deltaY == 2 {
		dy1 = 1
		dx1 = deltaX
	} else if deltaY == -2 {
		dy1 = -1
		dx1 = deltaX
	} else {
		return
	}
	moveKnot(knotIndex+1, dx1, dy1)
}
