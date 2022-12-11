package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Monkey struct {
	items                   []string
	operation               func(int) int
	testDivisor             int
	trueMonkey, falseMonkey int
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
		thisMonkey.items = strings.Split(itemsStr, ",")

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
	fmt.Println(monkies)
	for i, v := range monkies {
		fmt.Println(i, *v)
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
