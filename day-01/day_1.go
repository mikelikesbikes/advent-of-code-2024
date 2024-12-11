package day1

import (
	"sort"
	"strconv"
	"strings"
)

func parseInput(input string) (l1 []int, l2 []int, err error) {
	for _, line := range strings.Split(input, "\n") {
		results := strings.Split(line, `   `)

		v1, _ := strconv.Atoi(results[0])
		l1 = append(l1, v1)

		v2, _ := strconv.Atoi(results[1])
		l2 = append(l2, v2)
	}

	sort.Ints(l1)
	sort.Ints(l2)
	return l1, l2, nil
}

func Part1(l1 []int, l2 []int) int {
	distSum := 0
	for i := 0; i < len(l1); i++ {
		dist := l1[i] - l2[i]
		if dist < 0 {
			dist = -dist
		}
		distSum += dist
	}
	return distSum
}

func Part2(l1 []int, l2 []int) int {
	counts := make(map[int]int)
	for _, v := range l2 {
		counts[v] += 1
	}

	similarity := 0
	for _, v := range l1 {
		similarity += v * counts[v]
	}
	return similarity
}
