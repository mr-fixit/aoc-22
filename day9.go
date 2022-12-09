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
			move head
			calc dist
			move tail
			add tail position
	sort/unique tailPositions
	count tailPositions
*/

func main() {
	fmt.Println("Day 9")

	file, err := os.Open("day9_1.txt")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}
	var headX, headY, tailX, tailY int
	tailPositions := map[string]bool{
		"0 0": true,
	}

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
			headX += headMove[0]
			headY += headMove[1]

			tailMoved := true
			if headX-tailX == 2 {
				tailX += 1
				tailY = headY
			} else if headX-tailX == -2 {
				tailX -= 1
				tailY = headY
			} else if headY-tailY == 2 {
				tailY += 1
				tailX = headX
			} else if headY-tailY == -2 {
				tailY -= 1
				tailX = headX
			} else {
				tailMoved = false
			}
			if tailMoved {
				tailPositionStr := fmt.Sprintf("%d %d", tailX, tailY)
				fmt.Println("moved tailPosition", tailX, tailY)
				tailPositions[tailPositionStr] = true
			}
		} // for nTimes
		fmt.Println("head now at ", headX, headY)
	}
	fmt.Println("part1: ", len(tailPositions))

}
