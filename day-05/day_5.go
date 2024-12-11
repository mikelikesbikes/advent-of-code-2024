package day5

import (
	"slices"
	"strconv"
	"strings"
)

type Manual struct {
	rules map[int]map[int]bool
}

func NewManual(input string) Manual {
	rules := make(map[int]map[int]bool)
	for _, s := range strings.Split(input, "\n") {
		strs := strings.Split(s, "|")
		a, _ := strconv.Atoi(strs[0])
		b, _ := strconv.Atoi(strs[1])
		if rules[a] != nil {
			rules[a][b] = true
		} else {
			rules[a] = map[int]bool{b: true}
		}
	}
	return Manual{rules}
}

func (m *Manual) IsSorted(updates []int) bool {
	return slices.IsSortedFunc(updates, m.compare)
}

func (m *Manual) compare(x int, y int) int {
	if m.rules[x][y] {
		return -1
	} else if m.rules[y][x] {
		return 1
	} else {
		return 0
	}
}

func (m *Manual) Sort(updates []int) []int {
	temp := make([]int, len(updates))
	copy(temp, updates)
	slices.SortFunc(temp, m.compare)
	return temp
}

func parseInput(input string) (Manual, [][]int) {
	s := strings.Split(input, "\n\n")
	manual := NewManual(s[0])

	lines := strings.Split(s[1], "\n")

	updates := make([][]int, len(lines))
	for i, line := range lines {
		strs := strings.Split(line, ",")
		for _, str := range strs {
			x, _ := strconv.Atoi(str)
			updates[i] = append(updates[i], x)
		}

	}

	return manual, updates
}

func Part1(manual Manual, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		if manual.IsSorted(update) {
			sum += update[len(update)/2]
		}
	}
	return sum
}

func Part2(manual Manual, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		if !manual.IsSorted(update) {
			tempUpdate := manual.Sort(update)
			sum += tempUpdate[len(tempUpdate)/2]
		}
	}
	return sum
}
