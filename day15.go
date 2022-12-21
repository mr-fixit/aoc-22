package main

import (
	"bufio"
	"fmt"
	"os"
)

type Line15 struct {
	sx, sy, bx, by int   // sensor & beacon: x & y
	mDist          int   // manhattan distance from sensor to beacon
	pt             Point // sensor coordinates
}

func day15() {
	// do15("data/day15_0.txt", 10, 26)
	// do15("data/day15_1.txt", 2000000, 5832528)

	var maxCoord int
	var fileName string
	if false {
		maxCoord = 20
		fileName = "data/day15_0.txt"
	} else {
		maxCoord = 4000000
		fileName = "data/day15_1.txt"
	}

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	lines := make([]Line15, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := Line15{}
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &line.sx, &line.sy, &line.bx, &line.by)
		line.mDist = ManhattanDist(line)
		line.pt = Point{line.sx, line.sy}
		if line.sy-line.mDist > maxCoord {
			fmt.Println("ignoring line ", line)
		} else {
			lines = append(lines, line)
		}
	}

	do15_2(lines, maxCoord)
}

func do15_2(sensors []Line15, maxCoord int) {
	// look for sensors that have a gap between them
	gappers := make([]Line15, 0)
	for i := 0; i < len(sensors); i++ {
		s1 := sensors[i]
		for j := i + 1; j < len(sensors); j++ {
			s2 := sensors[j]
			sensorDistance := abs(s1.sx-s2.sx) + abs(s1.sy-s2.sy)
			combinedRange := s1.mDist + s2.mDist
			delta := sensorDistance - combinedRange
			if delta == 2 {
				fmt.Println(i, j, delta)
				gappers = append(append(gappers, s1), s2)
			}
		}
	}
	s1 := gappers[0]
	x := s1.sx - s1.mDist - 1
	y := s1.sy
	dx, dy := 1, -1
	nextTurn := Point{x: s1.sx, y: s1.sy - s1.mDist - 1}
	for y != nextTurn.y || x != nextTurn.x {
		here := Point{x, y}
		man1 := ManDistPoint(here, s1.pt) - s1.mDist
		if man1 != 1 {
			panic("bad")
		}
		man2 := ManDistPoint(here, gappers[1].pt) - gappers[1].mDist
		if man2 == 1 {
			man3 := ManDistPoint(here, gappers[2].pt) - gappers[2].mDist
			if man3 == 1 {
				fmt.Println(here, man1, man2, man3, ManDistPoint(here, gappers[3].pt)-gappers[3].mDist, 4000000*x+y)
			}

		}
		x += dx
		y += dy
	}

}

func ManDistPoint(p1, p2 Point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
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
