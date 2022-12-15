package main

import (
	"fmt"
	"os"
)

type ByteMap struct {
	bytes        [][]byte
	nRows, nCols int
}

func inBounds(x, y int, board ByteMap) bool {
	return x >= 0 && y >= 0 && x < board.nCols && y < board.nRows
}

func readInput_12(fileName string) (out ByteMap) {
	dat, err := os.ReadFile(fileName)
	check(err)
	nBytes := len(dat)
	nCols := 0
	for i, v := range dat {
		if v == '\n' {
			nCols = i
			break
		}
	}
	nRows := (nBytes + 1) / nCols
	bytes := make([][]byte, 0)
	for i := 0; i < nRows; i++ {
		idx0 := i * (nCols + 1)
		row := dat[idx0 : idx0+nCols]
		bytes = append(bytes, row)
	}
	return ByteMap{nRows: nRows, nCols: nCols, bytes: bytes}
}

func drawMap(m ByteMap) {
	for y := 0; y < m.nRows; y++ {
		fmt.Printf("%2d: ", y)
		for x := 0; x < m.nCols; x++ {
			fmt.Printf("%s", string(m.bytes[y][x]))
		}
		fmt.Println()
	}
}

func findEnds(m ByteMap) (startX, startY, endX, endY int) {
	for y := 0; y < m.nRows; y++ {
		for x := 0; x < m.nCols; x++ {
			if m.bytes[y][x] == 'S' {
				startX = x
				startY = y
				m.bytes[y][x] = 'a'
			}
			if m.bytes[y][x] == 'E' {
				endX = x
				endY = y
				m.bytes[y][x] = 'z'
			}
		}
	}
	return
}

func canGo(x, y, nx, ny int, b ByteMap, cost [][]int) (canGo bool) {
	canGo = inBounds(nx, ny, b)
	canGo = canGo && b.bytes[ny][nx]-b.bytes[y][x] <= 1
	canGo = canGo && cost[ny][nx] > cost[y][x]+1
	return
}

var sX, sY, eX, eY int

func day12(fileName string) {

	b := readInput_12(fileName)
	drawMap(b)
	sX, sY, eX, eY = findEnds(b)
	fmt.Println("start: ", sX, sY)
	fmt.Println("end: ", eX, eY)

	todo := [][]int{{sX, sY}}
	cost := make([][]int, b.nRows)
	for y := 0; y < b.nRows; y++ {
		//cost = append(cost, make([]int, b.nRows))
		cost[y] = make([]int, b.nCols)
		for x := 0; x < b.nCols; x++ {
			cost[y][x] = 1000000
		}
	}
	cost[sY][sX] = 0

	for len(todo) > 0 {
		x := todo[0][0]
		y := todo[0][1]
		fmt.Println("doing ", x, y, cost[y][x], abs(eX-x)+abs(eY-y))
		todo = todo[1:]
		curCost := cost[y][x]
		for _, dxdy := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			newX := x + dxdy[0]
			newY := y + dxdy[1]
			if canGo(x, y, newX, newY, b, cost) {
				cost[newY][newX] = curCost + 1
				if newX == eX && newY == eY {
					fmt.Println("end cost", curCost+1)
					os.Exit(0)
				}
				todo = append(todo, []int{newX, newY})
			}
		}
	}

}
