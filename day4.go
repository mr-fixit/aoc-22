package main

import (
	"bufio"
	"fmt"
	"os"
)

func day4() {
	//file, err := os.Open("day4_0.txt")
	file, err := os.Open("day4_1.txt")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	sumPart1, sumPart2 := 0, 0
	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		wholeLine := scanner.Text()
		// 2-4,6-8
		// 2-3,4-5
		// 5-7,7-9
		// 2-8,3-7
		// 6-6,4-6
		// 2-6,4-8
		var s0, f0, s1, f1 int
		fmt.Sscanf(wholeLine, "%d-%d,%d-%d", &s0, &f0, &s1, &f1)
		s := maxInt(s0, s1)
		f := minInt(f0, f1)
		if (s == s0 && f == f0) || (s == s1 && f == f1) {
			// fmt.Println(s, f)
			sumPart1 += 1
		}
		if (s1 <= f0 && f1 >= s0) || (s0 <= f1 && f0 >= s1) {
			sumPart2 += 1
		}
	}
	fmt.Println("part1: ", sumPart1, "part2: ", sumPart2)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
