package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day16() {
	day16_1("data/day16_0.txt")
}

type Node16 struct {
	name  string
	rate  int
	open  bool
	exits []string
}

type NodeMap map[string]Node16

var nPaths int = 0

func day16_1(fileName string) {
	var nodes NodeMap = make(NodeMap, 0)
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
		var name string
		var rate int
		fmt.Sscanf(line, "Valve %s has a flow rate=", &name)
		fmt.Sscanf(line[strings.Index(line, "rate=")+5:], "%d", &rate)
		tailIdx := strings.Index(line, "to valve")
		var valves []string

		if line[tailIdx+8:tailIdx+9] == "s" {
			valvesStr := line[tailIdx+10:]
			valves = strings.Split(valvesStr, ", ")
		} else {
			valves = append(valves, line[tailIdx+9:])
		}
		// fmt.Printf("%s %d %d valves: '%s' \n", name, rate, len(valves), valves)
		nodes[name] = Node16{name, rate, false, valves}
	}
	for _, v := range nodes {
		fmt.Printf("%s: %d %s\n", v.name, v.rate, v.exits)
	}

	path := []string{"AA"}
	move16(nodes, path)
	fmt.Println("nPaths:", nPaths)
}

// path: [ aa, dd, *, cc, bb, *, aa, ii... ]
func move16(nodes NodeMap, path []string) {
	if len(path) == 31 || CountofClosed(nodes) == 0 {
		// fmt.Println("path:", path)
		nPaths += 1
		return
	}
	lastMove := path[len(path)-1]
	if lastMove == "*" {
		lastMove = path[len(path)-2]
	}
	cur := nodes[lastMove]
	var moves []string
	if cur.rate > 0 && !cur.open {
		moves = append(moves, "*")
	}
	moves = append(moves, cur.exits...)

	for _, move := range moves {
		// check if move will go back to where we just came from
		if move != "*" && len(path) > 2 && move == path[len(path)-2] {
			continue
		}
		path = append(path, move)
		if move == "*" {
			cur.open = true
			nodes[cur.name] = cur
		}

		move16(nodes, path)

		path = path[:len(path)-1] // remove last move
		if move == "*" {
			cur.open = false
			nodes[cur.name] = cur
		}
	}
}

func CountofClosed(nodes NodeMap) int {
	result := 0
	// fmt.Printf("   n closed: ")
	for _, val := range nodes {
		if !val.open && val.rate > 0 {
			// fmt.Printf("%s ", val.name)
			result += 1
		}
	}
	// fmt.Printf(" = %d\n", result)
	return result
}
