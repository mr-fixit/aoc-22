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
	verb := "move to"
	b := MakeMap(500, 500, 0, 0)
	for scanner.Scan() {
		text := scanner.Text()
		var x, y int
		// 498,4 -> 498,6 -> 496,6
		// 503,4 -> 502,4 -> 502,9 -> 494,9
		if text == "->" {
			verb = "line to"
		} else {
			_, err := fmt.Sscanf(text, "%d,%d", &x, &y)
			check(err)
			fmt.Printf("%s x,y: %d %d\n", verb, x, y)
			verb = "move to"
			ExpandMap(&b, x, y)
		}
	}
}

type ByteMap2 struct {
	bytes                  [][]byte
	minX, minY, maxX, maxY int
	nRows, nCols           int
}

func MakeMap(minX, maxX, minY, maxY int) (b ByteMap2) {
	b = ByteMap2{minX: minX, maxX: maxX, minY: minY, maxY: maxY}
	b.bytes = make([][]byte, 0)
	b.nRows = maxY - minY + 1
	b.nCols = maxX - minX + 1
	for y := 0; y < b.nRows; y++ {
		row := make([]byte, b.nCols)
		for x := 0; x < b.nCols; x++ {
			row[x] = '.'
		}
		b.bytes = append(b.bytes, row)
	}
	return
}

func ExpandMap(b *ByteMap2, x, y int) {
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

func drawMap14(m ByteMap2) {
	for y := 0; y < m.nRows; y++ {
		fmt.Printf("%d: ", y+m.minY)
		for x := 0; x < m.nCols; x++ {
			fmt.Printf("%s", string(m.bytes[y][x]))
		}
		fmt.Println()
	}
}
