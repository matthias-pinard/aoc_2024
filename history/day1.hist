package main

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"
)

//go:embed input/1.input
var input string

func main() {
	l1 := make([]int, 0)
	l2 := make([]int, 0)
	for _, s := range strings.Split(input, "\n") {
		v := strings.Split(s, "   ")
		v1, _ := strconv.Atoi(v[0])
		v2, _ := strconv.Atoi(v[1])
		l1 = append(l1, v1)
		l2 = append(l2, v2)
	}

	mapCount := make(map[int]int)

	for _, v := range l2 {
		mapCount[v] += 1
	}
	acc := 0
	for _, v := range l1 {
		acc += v * mapCount[v]
		// println(acc)
	}
	println(acc)
}

func Part1() {
	l1 := make([]int, 0)
	l2 := make([]int, 0)
	for _, s := range strings.Split(input, "\n") {
		v := strings.Split(s, "   ")
		v1, _ := strconv.Atoi(v[0])
		v2, _ := strconv.Atoi(v[1])
		l1 = append(l1, v1)
		l2 = append(l2, v2)
	}
	slices.Sort(l1)
	slices.Sort(l2)

	acc := 0
	for i, v1 := range l1 {
		v2 := l2[i]
		acc += abs(v1 - v2)
		// println(acc)
	}
	println(acc)
}

func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}
