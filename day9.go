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

const nKnots = 10

var knotX = make([]int, nKnots)
var knotY = make([]int, nKnots)
var tailPositions = map[string]bool{"0 0": true}
var xMax, xMin, yMax, yMin int

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
		fmt.Println("  headMove", headMove, "times", nTimes)
		// drawKnots("head")
		for i := 0; i < nTimes; i++ {
			moveKnot(0, headMove[0], headMove[1])
			// drawKnots(fmt.Sprintf("dir:%s n:%d i:%d", direction, nTimes, i))
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
	updateMinMax(knotIndex, headX, headY)
	if knotIndex == nKnots-1 {
		tailPositionStr := fmt.Sprintf("%d %d", headX, headY)
		fmt.Println("moved tailPosition", headX, headY)
		tailPositions[tailPositionStr] = true
		return
	}
	tailX := knotX[knotIndex+1]
	tailY := knotY[knotIndex+1]
	moves := [][][]int{
		{{1, 1}, {1, 1}, {0, 1}, {-1, 1}, {-1, 1}},
		{{1, 1}, {0, 0}, {0, 0}, {0, 0}, {-1, 1}},
		{{1, 0}, {0, 0}, {0, 0}, {0, 0}, {-1, 0}},
		{{1, -1}, {0, 0}, {0, 0}, {0, 0}, {-1, -1}},
		{{1, -1}, {1, -1}, {0, -1}, {-1, -1}, {-1, -1}},
	}
	dxdy := moves[tailY-headY+2][tailX-headX+2]
	if dxdy[0] == 0 && dxdy[1] == 0 {
		return
	} else if dxdy[0] == 9 || dxdy[1] == 9 {
		panic("wtf?")
	}
	moveKnot(knotIndex+1, dxdy[0], dxdy[1])
}

func drawKnots(title string) {
	fmt.Println("---", title, "---")
	for row := yMin; row <= yMax; row++ {
		for col := xMin; col <= xMax; col++ {
			noKnot := true
			for knotIdx := 0; knotIdx < nKnots-1; knotIdx++ {
				if col == knotX[knotIdx] && row == knotY[knotIdx] {
					if knotIdx == 0 {
						fmt.Print("H")
					} else if knotIdx == nKnots-1 {
						fmt.Printf("T")
					} else {
						fmt.Print(knotIdx)
					}
					noKnot = false
					break
				}
			}
			if noKnot {
				if row == 0 && col == 0 {
					fmt.Printf("s")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}

}

func updateMinMax(knotIdx, x, y int) {
	if knotIdx != 0 {
		return
	}
	if x > xMax {
		xMax = x
	}
	if y > yMax {
		yMax = y
	}
	if x < xMin {
		xMin = x
	}
	if y < yMin {
		yMin = y
	}
}
