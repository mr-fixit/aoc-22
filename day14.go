package main

import (
	"bufio"
	"fmt"
	"os"
)

func day14(fileName string) {
	fmt.Println(fileName)
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	isMoveTo := true
	commands := make([]Command14, 0)
	b := ByteMap2{minX: 500, maxX: 500, minY: 0, maxY: 0}
	for scanner.Scan() {
		text := scanner.Text()
		var x, y int
		// 498,4 -> 498,6 -> 496,6
		// 503,4 -> 502,4 -> 502,9 -> 494,9
		if text == "->" {
			isMoveTo = false
		} else {
			_, err := fmt.Sscanf(text, "%d,%d", &x, &y)
			check(err)
			fmt.Printf("moveto(%s) x,y: %d %d\n", isMoveTo, x, y)
			commands = append(commands, Command14{isMoveTo: isMoveTo, x: x, y: y})
			isMoveTo = true
			ExpandMap14(&b, x, y)
		}
	}
	AllocateMap14(&b)
	DrawCommands(b, commands)
	nGrains := 0
	for ; DropOne(b); nGrains++ {
		//		drawMap14(b)
	}
	fmt.Println("part 1:", nGrains)
}

func DropOne(b ByteMap2) (stopped bool) {
	x := 500 - b.minX + 1
	y := 0 - b.minY
	stopped = false
	for !stopped {
		if y > b.maxY-1 {
			return false
		}
		if b.bytes[y+1][x] == '.' {
			y += 1
		} else if b.bytes[y+1][x-1] == '.' {
			y += 1
			x -= 1
		} else if b.bytes[y+1][x+1] == '.' {
			y += 1
			x += 1
		} else {
			stopped = true
		}
	}
	b.bytes[y][x] = 'o'
	return stopped
}

type ByteMap2 struct {
	bytes                  [][]byte
	minX, minY, maxX, maxY int
	nRows, nCols           int
}

type Command14 struct {
	isMoveTo bool
	x, y     int
}

func MakeMap14(minX, maxX, minY, maxY int) (b ByteMap2) {
	b = ByteMap2{minX: minX, maxX: maxX, minY: minY, maxY: maxY}
	b.nRows = maxY - minY + 2
	b.nCols = maxX - minX + 2
	return
}

func ExpandMap14(b *ByteMap2, x, y int) {
	if b.minX > x {
		b.minX = x
	}
	if b.maxX < x {
		b.maxX = x
	}
	if b.minY > y {
		b.minY = y
	}
	if b.maxY < y {
		b.maxY = y
	}
}

func AllocateMap14(b *ByteMap2) {
	b.nRows = b.maxY - b.minY + 3
	b.nCols = b.maxX - b.minX + 3
	b.bytes = make([][]byte, 0)
	for y := 0; y < b.nRows; y++ {
		row := make([]byte, b.nCols)
		for x := 0; x < b.nCols; x++ {
			row[x] = '.'
		}
		b.bytes = append(b.bytes, row)
	}
}

func drawMap14(m ByteMap2) {
	for y := 0; y < m.nRows; y++ {
		fmt.Printf("%3d: ", y+m.minY)
		for x := 0; x < m.nCols; x++ {
			fmt.Printf("%s", string(m.bytes[y][x]))
		}
		fmt.Println()
	}
}

func DrawCommands(b ByteMap2, cmds []Command14) {
	var curX, curY int
	for _, cmd := range cmds {
		if cmd.isMoveTo {
			curX = cmd.x
			curY = cmd.y
		} else { // line to
			dx, dy := 0, 0
			if curX != cmd.x {
				if curX < cmd.x {
					dx = 1
				} else {
					dx = -1
				}
			} else if curY != cmd.y {
				if curY < cmd.y {
					dy = 1
				} else {
					dy = -1
				}
			} else {
				panic("why a zero-length line?")
			}
			b.bytes[curY-b.minY][curX-b.minX+1] = '#'
			for curX != cmd.x || curY != cmd.y {
				curX += dx
				curY += dy
				b.bytes[curY-b.minY][curX-b.minX+1] = '#'
			}
		}
		drawMap14(b)
	}
}
