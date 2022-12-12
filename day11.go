package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items                   []int
	operation               func(int) int
	testDivisor             int
	trueMonkey, falseMonkey int
	count                   int
}

var monkies []*Monkey

func day11(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var monkeyIdx int
		fmt.Sscanf(scanner.Text(), "Monkey %d:", &monkeyIdx)

		thisMonkey := &Monkey{}
		monkies = append(monkies, thisMonkey)

		scanner.Scan()
		var itemsStr string = strings.Split(scanner.Text(), ":")[1]
		for _, itemStr := range strings.Split(itemsStr, ",") {
			itemInt, err := strconv.Atoi(strings.TrimSpace(itemStr))
			if err != nil {
				panic(err)
			} else {
				thisMonkey.items = append(thisMonkey.items, itemInt)
			}
		}

		scanner.Scan()
		thisMonkey.operation = parseOpStr(scanner.Text())

		scanner.Scan()
		var testDivisor int
		fmt.Sscanf(scanner.Text(), "  Test: divisible by %d", &testDivisor)
		thisMonkey.testDivisor = testDivisor

		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "    If true: throw to monkey %d", &thisMonkey.trueMonkey)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "    If false: throw to monkey %d", &thisMonkey.falseMonkey)

		scanner.Scan()
	}
	PrintMonkies()
	thing := 1
	for _, v := range monkies {
		thing *= v.testDivisor
	}

	for roundIdx := 0; roundIdx < 10000; roundIdx++ {
		for _, monkey := range monkies {
			for _, wl := range monkey.items {
				monkey.count += 1
				wl = monkey.operation(wl)
				// wl /= 3
				wl %= thing
				trueDst := monkies[monkey.trueMonkey]
				falseDst := monkies[monkey.falseMonkey]
				if wl%monkey.testDivisor == 0 {
					trueDst.items = append(trueDst.items, wl)
				} else {
					falseDst.items = append(falseDst.items, wl)
				}
			}
			monkey.items = make([]int, 0)
		}
		fmt.Println("Round ", roundIdx+1)
		PrintMonkies()
	}
	sort.Slice(monkies, func(i, j int) bool {
		return monkies[i].count > monkies[j].count
	})
	fmt.Println("part 1: ", monkies[0].count*monkies[1].count)
}

func PrintMonkies() {
	for _, v := range monkies {
		fmt.Println("  ", v.items)
	}
}

func parseOpStr(opStr string) func(int) int {
	var op string
	var rhs int
	nRead, _ := fmt.Sscanf(opStr, "  Operation: new = old %s %d", &op, &rhs)
	if op == "*" {
		if nRead == 1 {
			return func(old int) int { return old * old }
		} else {
			return func(old int) int { return old * rhs }
		}
	} else {
		return func(old int) int { return old + rhs }
	}
}
