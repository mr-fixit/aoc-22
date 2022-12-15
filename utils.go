package main

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
