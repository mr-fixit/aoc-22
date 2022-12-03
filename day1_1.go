package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var max int
	for {
		var sum int
		for {
			if !scanner.Scan() {
				fmt.Println("max: ", max)
				return
			}
			text := scanner.Text()
			if len(text) != 0 {
				var i int
				fmt.Sscanf(text, "%d", &i)
				sum += i
			} else {
				if sum > max {
					max = sum
					fmt.Println("max: ", max)
				}
				break
			}
		}
	}
	fmt.Println("max: ", max)
}
