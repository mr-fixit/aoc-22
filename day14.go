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
	b := ByteMap2{maxY: 0}
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
			//fmt.Printf("moveto(%s) x,y: %d %d\n", isMoveTo, x, y)
			commands = append(commands, Command14{isMoveTo: isMoveTo, x: x - 500, y: y})
			isMoveTo = true
			ExpandMap14(&b, y)
		}
	}
	AllocateMap14(&b)
	DrawCommands(b, commands)
	drawMap14(b)
	nGrains := 0
	for ; !DropOne(b); nGrains++ {
		// drawMap14(b)
	}
	fmt.Println("part 1:", nGrains)
}

func DropOne(b ByteMap2) (done bool) {
	x := 0 + b.nCols/2
	y := 0
	if b.bytes[y][x] != '.' {
		return true
	}
	for stopped := false; !stopped; {
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
	return false
}

type ByteMap2 struct {
	bytes        [][]byte
	maxY         int
	nRows, nCols int
}

type Command14 struct {
	isMoveTo bool
	x, y     int
}

func MakeMap14(maxY int) (b ByteMap2) {
	b = ByteMap2{maxY: maxY}
	b.maxY = maxY
	b.nRows = maxY + 2
	b.nCols = b.nRows*2 + 1
	return
}

func ExpandMap14(b *ByteMap2, y int) {
	if b.maxY < y {
		b.maxY = y
	}
}

func AllocateMap14(b *ByteMap2) {
	b.nRows = b.maxY + 3
	b.nCols = b.nRows*2 + 1
	b.bytes = make([][]byte, 0)
	for y := 0; y < b.nRows; y++ {
		fillChar := '.'
		if y == b.nRows-1 { // bottom row
			fillChar = '#'
		}
		row := make([]byte, b.nCols)
		for x := 0; x < b.nCols; x++ {
			row[x] = byte(fillChar)
		}
		b.bytes = append(b.bytes, row)
	}
}

func drawMap14(m ByteMap2) {
	for y := 0; y < m.nRows; y++ {
		fmt.Printf("%3d: ", y)
		for x := 0; x < m.nCols; x++ {
			fmt.Printf("%s", string(m.bytes[y][x]))
		}
		fmt.Println()
	}
}

func DrawCommands(b ByteMap2, cmds []Command14) {
	var curX, curY int
	for _, cmd := range cmds {
		cmdX := cmd.x + b.nCols/2
		if cmd.isMoveTo {
			curX = cmdX
			curY = cmd.y
		} else { // line to
			dx, dy := 0, 0
			if curX != cmdX {
				if curX < cmdX {
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
			b.bytes[curY][curX] = '#'
			for curX != cmdX || curY != cmd.y {
				curX += dx
				curY += dy
				b.bytes[curY][curX] = '#'
			}
		}
	}
}
