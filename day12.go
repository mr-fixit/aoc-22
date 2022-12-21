package main

import (
	"fmt"
	"math"
	"os"

	"github.com/fatih/color"
	"golang.org/x/exp/slices"
)

type ByteMap struct {
	bytes        [][]byte
	nRows, nCols int
}

var kMaxCost = 9999999

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

var sX, sY, eX, eY int

func SolveOne(b ByteMap, sX, sY int) (minCost int) {

	unvisited := []Point{}
	cost := make([][]int, b.nRows)
	for y := 0; y < b.nRows; y++ {
		cost[y] = make([]int, b.nCols)
		for x := 0; x < b.nCols; x++ {
			cost[y][x] = kMaxCost
			unvisited = append(unvisited, Point{x, y})
		}
	}
	cost[sY][sX] = 0
	done := false
	for len(unvisited) > 0 && !done {
		slices.SortFunc(unvisited, func(a, b Point) bool {
			return cost[a.y][a.x] < cost[b.y][b.x]
		})
		x := unvisited[0].x
		y := unvisited[0].y
		// fmt.Println("doing ", x, y, cost[y][x], abs(eX-x)+abs(eY-y))
		unvisited = unvisited[1:]
		curCost := cost[y][x]
		// drawIntMap(cost, kMaxCost)
		// if x == 23 && y == 3 {
		// 	println()
		// }
		for _, dxdy := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			newX := x + dxdy[0]
			newY := y + dxdy[1]

			canGo := inBounds(newX, newY, b)
			canGo = canGo && slices.Contains(unvisited, Point{newX, newY})
			canGo = canGo && int(b.bytes[newY][newX])-int(b.bytes[y][x]) <= 1
			canGo = canGo && cost[newY][newX] > cost[y][x]+1

			if canGo {
				// fmt.Println("		can go", newX, newY)
				cost[newY][newX] = curCost + 1
				if newX == eX && newY == eY {
					done = true
					break
				}
			}
		}
	}
	minCost = cost[eY][eX]
	// drawIntMap(cost, kMaxCost)
	fmt.Println("cost ", minCost)
	return
}

func day12(fileName string, expectedValue int) {
	b := readInput_12(fileName)
	sX, sY, eX, eY = findEnds(b)
	fmt.Println("end: ", eX, eY)
	// drawMap12(b, Point{sX, sY}, Point{eX, eY})

	var minCost = math.MaxInt32
	for y := 0; y < b.nRows; y++ {
		fmt.Println("doing row", y)
		if b.bytes[y][0] == 'a' {
			thisCost := SolveOne(b, 0, y)
			if thisCost < minCost {
				minCost = thisCost
				fmt.Printf("y %d cost %d\n", y, minCost)
			}
		}
	}
	fmt.Println("mincost", minCost)
}

func drawMap12(m ByteMap, start Point, end Point) {
	return
	printRed := color.New(color.FgRed).PrintfFunc()

	for y := 0; y < m.nRows; y++ {
		fmt.Printf("%3d: ", y)
		for x := 0; x < m.nCols; x++ {
			pt := Point{x, y}
			if pt == start {
				printRed("S")
			} else if pt == end {
				printRed("E")
			} else {
				fmt.Printf("%s", string(m.bytes[y][x]))
			}
		}
		fmt.Println()
	}
}

func drawIntMap(m [][]int, ignore int) {
	fmt.Printf("   ")
	for x := 0; x < len(m[0]); x++ {
		fmt.Printf("%4d", x)
	}
	fmt.Println()
	for y := 0; y < len(m); y++ {
		fmt.Printf("%3d", y)
		for x := 0; x < len(m[y]); x++ {
			v := m[y][x]
			if v == ignore {
				fmt.Printf("    ")
			} else {
				fmt.Printf("%4d", v)
			}
		}
		fmt.Println()
		fmt.Println()
	}
}
