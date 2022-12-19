package main

import (
	"bufio"
	"fmt"
	"os"
)

func day3() {
	//file, err := os.Open("data/day3_0.txt")
	file, err := os.Open("data/day3_1.txt")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanBytes)
	groupOf3 := make([]string, 0)
	var sum, sumPart2 int
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
		sum += bval

		groupOf3 = append(groupOf3, wholeLine)
		if len(groupOf3) == 3 {
			only := intersection(intersection([]byte(groupOf3[0]), []byte(groupOf3[1])), []byte(groupOf3[2]))
			val := value(only[0])
			sumPart2 += val
			groupOf3 = make([]string, 0)
		}

	}
	fmt.Println("part 1:", sum)
	fmt.Println("part 2:", sumPart2)
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
