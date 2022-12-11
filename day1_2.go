package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func day1_2() {
	scanner := bufio.NewScanner(os.Stdin)
	arr := make([]int, 0)
	for {
		var sum int
		for {
			if !scanner.Scan() {
				sort.Slice(arr, func(i, j int) bool {
					return arr[i] > arr[j]
				})

				fmt.Println("arr: ", arr)
				fmt.Println("top 3: ", arr[0]+arr[1]+arr[2])
				return
			}
			text := scanner.Text()
			if len(text) != 0 {
				var i int
				fmt.Sscanf(text, "%d", &i)
				sum += i
			} else {
				fmt.Println("sum", sum)
				arr = append(arr, sum)
				break
			}
		}
	}
}
