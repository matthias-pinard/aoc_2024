package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input/2.input
var input string

func main() {
	l := make([][]int, 0)
	for _, s := range strings.Split(input, "\n") {
		v := strings.Split(s, " ")
		if len(v) == 1 {
			continue
		}
		l1 := make([]int, 0)
		for _, n := range v {
			p, _ := strconv.Atoi(n)
			l1 = append(l1, p)
		}
		l = append(l, l1)
	}

	safe := 0
	for _, r := range l {
		fmt.Printf("v: %+v\n", r)
		if checkSafe(r) {
			safe += 1
		} else {
			for i := range len(r) {
				sr := append([]int{}, r[:i]...)
				sr = append(sr, r[i+1:]...)
				if checkSafe(sr) {
					safe += 1
					break
				}
			}
		}
	}
	print(safe)
}

func checkSafe(r []int) bool {
	inc := r[0]-r[1] < 0
	for i := range len(r) - 1 {
		e := r[i+1] - r[i]
		if abs(e) < 1 || abs(e) > 3 {
			return false
		}
		if (inc && e < 0) || (!inc && e > 0) {
			return false
		}
	}
	return true
}

func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}

func Part1() {
	l := make([][]int, 0)
	for _, s := range strings.Split(input, "\n") {
		v := strings.Split(s, " ")
		if len(v) == 1 {
			continue
		}
		l1 := make([]int, 0)
		for _, n := range v {
			p, _ := strconv.Atoi(n)
			l1 = append(l1, p)
		}
		l = append(l, l1)
	}

	safe := 0
	for _, r := range l {
		fmt.Printf("v: %+v\n", r)
		if checkSafe(r) {
			println("safe")
			safe += 1
		} else {
			println("unsafe")
		}
	}
	print(safe)
}
