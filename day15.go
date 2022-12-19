package main

import (
	"bufio"
	"fmt"
	"os"
)

func day15() {
	do15("data/day15_0.txt", 10, 26)
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
		fmt.Println(sx, sy, bx, by)
		mDist := abs(sx-bx) + abs(sy-by)
		yDist := abs(sy - targetY)
		xDist := mDist - yDist
		for x := sx - xDist; x <= sx+xDist; x++ {
			foo[x] = true
		}
		if by == targetY {
			beaconsInLine[bx] = true
		}
		fmt.Println("now", len(foo)-len(beaconsInLine))
	}
	answer := len(foo) - len(beaconsInLine)
	fmt.Printf("answer: %d, expected %d, correct:%t\n", answer, expected, answer == expected)
}
