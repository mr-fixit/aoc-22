package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput() ([][]int, int, int) {
	out := make([][]int, 0)
	file, err := os.Open("day8_1.txt")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowText := scanner.Text()
		row := make([]int, len(rowText))
		for i, c := range rowText {
			row[i] = int(c - '0')
		}
		out = append(out, row)
	}
	return out, len(out), len(out[0])
}

func makeVis(nrows, ncols int) (vis [][]bool) {
	vis = make([][]bool, nrows)
	for i := 0; i < nrows; i++ {
		vis[i] = make([]bool, ncols)
	}
	for x := 0; x < ncols; x++ {
		for y := 0; y < nrows; y++ {
			if x == 0 || x == ncols-1 || y == 0 || y == nrows-1 {
				vis[y][x] = true
			}
		}
	}
	return
}

func print2D(m [][]int) {
	for i, val := range m {
		fmt.Printf("%d: %v\n", i, val)
	}
}

func main() {
	forest, nrows, ncols := readInput()
	print2D(forest)

	vis := makeVis(nrows, ncols)
	nVisible := 2 * (nrows + ncols - 2)

	for x := 1; x < ncols-1; x++ {
		for y := 1; y < nrows-1; y++ {

			// for each tree in the interior
			h := forest[y][x]

			for dx := -1; dx <= 1; dx += 2 {

				for nx := x + dx; !vis[y][x] && nx >= 0 && nx < nrows; nx += dx {
					if h <= forest[y][nx] {
						break // tree is hidden
					}
					if nx == 0 || nx == nrows-1 {
						// we made it to the edge, we must be visible
						vis[y][x] = true
						nVisible++
						break
					}
				}
			}
			for dy := -1; dy <= 1; dy += 2 {
				for ny := y + dy; !vis[y][x] && ny >= 0 && ny < ncols; ny += dy {
					if h <= forest[ny][x] {
						break
					}
					if ny == 0 || ny == ncols-1 {
						// we made it to the edge, we must be visible
						vis[y][x] = true
						nVisible++
						break
					}
				}
			}
		}
	}
	var status string
	if nVisible == 1798 {
		status = "RIGHT "
	} else {
		status = "WRONG"
	}
	fmt.Printf("part1: %d %s", nVisible, status)

	// 582 too low
	// 1642 too low
	// part1: 1798
}
