// day13.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"golang.org/x/exp/slices"
)

func day13(fileName string, expectedValue int) {
	doTests()
	//sum := 0
	fmt.Println(fileName)
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)
	// for i := 0; scanner.Scan(); i++ {
	// 	left := scanner.Text()
	// 	scanner.Scan()
	// 	right := scanner.Text()
	// 	scanner.Scan()
	// 	result := Compare(left, right, 0)
	// 	fmt.Printf("pair #: %d result: %s\n", i+1, result)
	// 	if result == LT {
	// 		fmt.Println("Found ", i+1)
	// 		sum += i + 1
	// 	}
	// }
	// fmt.Println("part1: ", sum) // 5606 too high, 5555 is right
	var strings = []string{"[[2]]", "[[6]]"}
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) != 0 {
			strings = append(strings, scanner.Text())
		}
	}
	sort.Slice(strings, func(i, j int) bool {
		return Compare(strings[i], strings[j], 0) == LT
	})
	i1 := slices.Index(strings, "[[2]]") + 1
	i2 := slices.Index(strings, "[[6]]") + 1
	fmt.Println("part2: ", i1, "*", i2, "=", i1*i2)
}

func doTests() {

	type CompareTest struct {
		left, right string
		expected    CompareResult
	}
	compareTests := []CompareTest{
		{"[]", "[]", EQ},
		{"[]", "[1]", LT},
		{"1", "1", EQ},
		{"3", "5", LT},
		{"[]", "1", LT},
		{"[]", "[1]", LT},
		{"[]", "[[],1]", LT},
		{"4", "[]", GT},
		{"[4]", "[]", GT},
		{"[[],4]", "[]", GT},
		{"[1,2,3]", "[1,2,3]", EQ},
		{"-13", "13", LT},
		{"13", "-13", GT},
		{"[1,1,3,1,1]", "[1,1,5,1,1]", LT},
		{"[[1],[2,3,4]]", "[[1],4]", LT},
		{"[9]", "[[8,7,6]]", GT},
		{"[[4,4],4,4]", "[[4,4],4,4,4]", LT},
		{"[7,7,7,7]", "[7,7,7]", GT},
		{"[]", "[3]", LT},
		{"[[[]]]", "[[]]", GT},
		{"[1,[2,[3,[4,[5,6,7]]]],8,9]", "[1,[2,[3,[4,[5,6,0]]]],8,9]", GT},
	}
	for _, test := range compareTests {
		result := Compare(test.left, test.right, 0)
		if result != test.expected {
			fmt.Println("failed test: ", test, result)
			Compare(test.left, test.right, 0)
		}
	}
}

func Elements(in string) (result []string) {
	result = make([]string, 0)
	for i := 1; i < len(in); i++ {
		char1 := in[i : i+1]
		if char1 == "[" {
			nestLevel := 1
			j := i + 1
			for ; nestLevel > 0; j++ {
				if in[j:j+1] == "[" {
					nestLevel++
				} else if in[j:j+1] == "]" {
					nestLevel--
				}
			}
			e := in[i:j]
			i = j
			result = append(result, e)
		} else if (char1 >= "0" && char1 <= "9") || char1 == "-" {
			j := i + 1
			for ; j < len(in) && in[j:j+1] >= "0" && in[j:j+1] <= "9"; j++ {
			}
			result = append(result, in[i:j])
		}
	}
	return
}

func Compare(left, right string, depth int) (result CompareResult) {
	typeL := TypeOf(left)
	typeR := TypeOf(right)
	// defer func() {
	// 	fmt.Printf("%s result: %d", pad(depth), result)
	// 	fmt.Println()
	// }()
	// fmt.Printf("%sL: %s\n", pad(depth), left)
	// fmt.Printf("%sR: %s\n", pad(depth), right)
	if typeL == null {
		if typeR == null {
			result = EQ
		}
		return LT
	} else if typeR == null {
		return GT
	} else if typeL == typeR {
		if typeL == number {
			return CompareInts(left, right)
		} else { // both are lists
			elemsL := Elements(left)
			elemsR := Elements(right)
			for i, left := range elemsL {
				if i >= len(elemsR) {
					return GT
				}
				innerResult := Compare(left, elemsR[i], depth+1)
				if innerResult != EQ {
					return innerResult
				}
			}
			if len(elemsL) < len(elemsR) {
				return LT
			} else {
				return EQ
			}
		}
	} else {
		// left and right are different: 1 int, 1 list
		if typeL == number {
			return Compare("["+left+"]", right, depth+1)
		} else {
			result = Compare(left, "["+right+"]", depth+1)
		}
	}
	return
}

func pad(depth int) string {
	return "                                           "[:depth*2]
}

type ElementType byte

const (
	list ElementType = iota
	listEnd
	number
	null
)

func TypeOf(s string) ElementType {
	// if len(s) == 0 {
	// 	return null
	// } else
	if s[0:1] == "[" {
		return list
		// } else if s[0:1] == "]" {
		// 	return listEnd
	} else if (s[0:1] >= "0" && s[0:1] <= "9") || s[0:1] == "-" {
		return number
	} else {
		panic("wtf")
	}
}

type CompareResult byte

const (
	LT CompareResult = iota
	EQ
	GT
)

func CompareInts(lStr, rStr string) CompareResult {
	var l, r int
	_, err := fmt.Sscanf(lStr, "%d", &l)
	check(err)
	_, err = fmt.Sscanf(rStr, "%d", &r)
	check(err)
	if l < r {
		return LT
	} else if l > r {
		return GT
	}
	return EQ
}
