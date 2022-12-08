package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput() ([][]byte, int, int) {
	out := make([][]byte, 0)
	file, err := os.Open("day8_1.txt")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out = append(out, scanner.Bytes())
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
				vis[x][y] = true
			}
		}
	}
	return
}

func main() {
	forest, nrows, ncols := readInput()
	fmt.Println(nrows, ncols)

	vis := makeVis(nrows, ncols)
	nVisible := 2 * (nrows + ncols - 2)

	for x := 1; x < ncols-1; x++ {
		for y := 1; y < nrows-1; y++ {
			h := forest[x][y]
			if (vis[x-1][y] && forest[x-1][y] < h) ||
				(vis[x+1][y] && forest[x+1][y] < h) ||
				(vis[x][y-1] && forest[x][y-1] < h) ||
				(vis[x][y+1] && forest[x][y+1] < h) {
				vis[x][y] = true
				nVisible++
			}
		}
	}
	fmt.Println("part1: ", nVisible)

}
