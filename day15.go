package main

import (
	"bufio"
	"fmt"
	"os"
)

type Line15 struct {
	sx, sy, bx, by int // sensor & beacon: x & y
	mDist          int // manhattan distance from sensor to beacon
}

func day15() {
	// do15("data/day15_0.txt", 10, 26)
	// do15("data/day15_1.txt", 2000000, 5832528)

	maxCoord := 4000000

	file, err := os.Open("data/day15_1.txt")
	check(err)
	defer file.Close()

	lines := make([]Line15, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := Line15{}
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &line.sx, &line.sy, &line.bx, &line.by)
		line.mDist = ManhattanDist(line)
		if line.sy-line.mDist > maxCoord {
			fmt.Println("ignoring line ", line)
		} else {
			lines = append(lines, line)
		}
	}

	for targetY := 0; targetY < maxCoord; targetY++ {
		if targetY%10 == 0 {
			fmt.Println("line ", targetY)
		}
		do15_2(lines, targetY, maxCoord)
	}
}

func do15_2(lines []Line15, targetY int, maxCoord int) {
	foo := make(map[int]bool, 1000)
	for _, line := range lines {
		yDist := abs(line.sy - targetY)
		xDist := line.mDist - yDist
		for x := line.sx - xDist; x <= line.sx+xDist; x++ {
			if x >= 0 && x <= maxCoord {
				foo[x] = true
			}
		}
	}
	nPossible := len(foo) + 1 - maxCoord
	if nPossible > 0 {
		for i := 0; i < maxCoord+1; i++ {
			if !foo[i] {
				fmt.Printf("found x,y %d,%d, tuning: %d\n", i, targetY, 4000000*i+targetY)
			}
		}
	}
}

func ManhattanDist(line Line15) int {
	return abs(line.sx-line.bx) + abs(line.sy-line.by)
}

func do15(fileName string, targetY int, expected int) {
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)
	foo := make(map[int]bool, 1000)
	beaconsInLine := make(map[int]bool, 0) // true if there's a beacon at X in the target line
	for scanner.Scan() {
		var sx, sy, bx, by int
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		// fmt.Println(sx, sy, bx, by)
		mDist := abs(sx-bx) + abs(sy-by)
		yDist := abs(sy - targetY)
		xDist := mDist - yDist
		for x := sx - xDist; x <= sx+xDist; x++ {
			foo[x] = true
		}
		if by == targetY {
			beaconsInLine[bx] = true
		}
	}
	answer := len(foo) - len(beaconsInLine)
	fmt.Printf("answer: %d, expected %d, correct:%t\n", answer, expected, answer == expected)
}
