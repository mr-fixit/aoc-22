package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//file, err := os.Open("day3_0.txt")
	file, err := os.Open("day3_1.txt")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanBytes)
	sum := 0
	for scanner.Scan() {
		wholeLine := scanner.Text()
		halfLen := len(wholeLine) / 2
		s1 := []byte(wholeLine[0:halfLen])
		s2 := []byte(wholeLine[halfLen:])
		only := intersection(s1, s2)
		if len(only) != 1 {
			fmt.Println("wrong", only)
		}
		b := only[0]
		bval := value(b)

		fmt.Println(string(only), b, int(b), bval)
		sum += bval
	}
	fmt.Println("part 1:", sum)
}

func value(b byte) int {
	if b > 'Z' {
		return int(b - 'a' + 1)
	} else {
		return int(b - 'A' + 27)
	}
}

func intersection(s1, s2 []byte) (inter []byte) {
	hash := make(map[byte]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			inter = append(inter, e)
			hash[e] = false
		}
	}
	return
}
