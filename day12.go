package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"golang.org/x/exp/slices"
)

type ByteMap struct {
	bytes        [][]byte
	nRows, nCols int
}

type Point struct {
	x, y int
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
	var kMaxCost = 9999999
	b := readInput_12(fileName)
	sX, sY, eX, eY = findEnds(b)
	fmt.Println("start: ", sX, sY)
	fmt.Println("end: ", eX, eY)
	redOnes := []Point{{sX, sY}, {eX, eY}}
	drawMap12(b, redOnes)

	unvisited := [][]int{}
	cost := make([][]int, b.nRows)
	for y := 0; y < b.nRows; y++ {
		cost[y] = make([]int, b.nCols)
		for x := 0; x < b.nCols; x++ {
			cost[y][x] = kMaxCost
			unvisited = append(unvisited, []int{x, y})
		}
	}
	cost[sY][sX] = 0

	for len(unvisited) > 0 {
		slices.SortFunc(unvisited, func(a, b []int) bool {
			return cost[a[1]][a[0]] < cost[b[1]][b[0]]
		})
		x := unvisited[0][0]
		y := unvisited[0][1]
		// fmt.Println("doing ", x, y, cost[y][x], abs(eX-x)+abs(eY-y))
		unvisited = unvisited[1:]
		curCost := cost[y][x]
		//drawIntMap(cost, kMaxCost)
		for _, dxdy := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			newX := x + dxdy[0]
			newY := y + dxdy[1]
			if canGo(x, y, newX, newY, b, cost) {
				cost[newY][newX] = curCost + 1
				if newX == eX && newY == eY {
					fmt.Println("end cost", curCost+1)
					os.Exit(0)
				}
			}
		}
	}
	drawIntMap(cost, kMaxCost)
}

func drawMap12(m ByteMap, redOnes []Point) {
	red := color.New(color.FgRed).PrintfFunc()

	for y := 0; y < m.nRows; y++ {
		fmt.Printf("%3d: ", y)
		for x := 0; x < m.nCols; x++ {
			if slices.Contains(redOnes, Point{x, y}) {
				red("%s", string(m.bytes[y][x]))
			} else {
				fmt.Printf("%s", string(m.bytes[y][x]))
			}
		}
		fmt.Println()
	}
}

func drawIntMap(m [][]int, ignore int) {
	for y := 0; y < len(m); y++ {
		fmt.Printf("%4d: ", y)
		for x := 0; x < len(m[y]); x++ {
			v := m[y][x]
			if v == ignore {
				fmt.Printf("    ")
			} else {
				fmt.Printf("%4d", v)
			}
		}
		fmt.Println()
	}
}
