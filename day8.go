package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput_8() ([][]int, int, int) {
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

func day8() {

	forest, nrows, ncols := readInput_8()
	// print2D(forest)

	vis := makeVis(nrows, ncols)
	nVisible := 2 * (nrows + ncols - 2)

	for x := 1; x < ncols-1; x++ {
		for y := 1; y < nrows-1; y++ {

			// for each tree in the interior
			h := forest[y][x]
			// fmt.Println("doing ", x, y, h)
			for _, dxdy := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				// fmt.Println("  dxdy:", dxdy)
				dx := dxdy[0]
				dy := dxdy[1]
				for nx, ny := x+dx, y+dy; !vis[y][x] && nx >= 0 && nx < ncols && ny >= 0 && ny < nrows; nx, ny = nx+dx, ny+dy {
					// fmt.Println("     nx ny height[nx,ny]", nx, ny, forest[ny][nx])
					if h <= forest[ny][nx] {
						// fmt.Println("      hidden")
						break // tree is hidden
					}
					if nx == 0 || nx == nrows-1 || ny == 0 || ny == ncols-1 {
						// we made it to the edge, we must be visible
						vis[y][x] = true
						nVisible++
						// fmt.Println("visible! now", nVisible)
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
		status = "WRONG, should be 1798"
	}
	fmt.Printf("part1: %d %s\n\n", nVisible, status)

	// 582 too low
	// 1642 too low
	// part1: 1798

	maxSceneScore := 0
	for x := 1; x < ncols-1; x++ {
		for y := 1; y < nrows-1; y++ {

			sceneScore := 1
			// for each tree in the interior
			h := forest[y][x]
			// fmt.Println("doing ", x, y, h)
			for _, dxdy := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				// fmt.Println("  dxdy:", dxdy)
				dx := dxdy[0]
				dy := dxdy[1]
				directionScore := 0
				for nx, ny := x+dx, y+dy; nx >= 0 && nx < ncols && ny >= 0 && ny < nrows; nx, ny = nx+dx, ny+dy {
					directionScore += 1
					// fmt.Println("     nx ny height[nx,ny]", nx, ny, forest[ny][nx])
					if h <= forest[ny][nx] {
						// fmt.Println("      hidden")
						break // tree is hidden
					}
					// if nx == 0 || nx == nrows-1 || ny == 0 || ny == ncols-1 {
					// 	// we made it to the edge, we must be visible
					// 	// fmt.Println("visible! now", nVisible)
					// 	break
					// }
				}
				sceneScore *= directionScore
				// fmt.Println(" dirScore", directionScore, sceneScore)
			} // each direction
			if sceneScore > maxSceneScore {
				maxSceneScore = sceneScore
			}
			// fmt.Println("sceneScore, max", sceneScore, maxSceneScore)
		}
	}
	fmt.Println("part 2: ", maxSceneScore) // 259308 is right

}
