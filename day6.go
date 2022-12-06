package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	const dayStr = "6"
	//file, err := os.Open("day" + dayStr + "_0.txt")
	file, err := os.Open("day" + dayStr + "_1.txt")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wholeLine := scanner.Text()
		found := false
		for i := 4; i < len(wholeLine) && !found; i++ {
			c4 := wholeLine[i-4 : i]
			b := []byte(c4)
			sort.Slice(b, func(i2, j int) bool {
				return b[i2] < b[j]
			})
			// fmt.Println(c4, b)
			// look for 4 unique
			for j := 0; j < 13; j++ {
				// fmt.Println(j, b[j], b[j+1], b[j] == b[j+1])
				if b[j] == b[j+1] {
					break
				}
				if j == 12 {
					found = true
					fmt.Println("found", c4, "index", i)
				}
			}
		}
		println()
	}
}
