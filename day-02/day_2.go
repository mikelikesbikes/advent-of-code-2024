package day2

import (
	"strconv"
	"strings"
)

func parseInput(input string) (reports []Report) {
	for _, line := range strings.Split(input, "\n") {
		svals := strings.Split(line, ` `)

		vals := make([]int, len(svals))
		for i, s := range svals {
			v, _ := strconv.Atoi(s)
			vals[i] = v
		}
		reports = append(reports, Report{vals})
	}
	return reports
}

type Report struct {
	vals []int
}

func (r *Report) IsSafe() bool {
	increasing := eachConsMatch(r.vals, func(i int, j int) bool {
		return i < j
	})

	decreasing := eachConsMatch(r.vals, func(i int, j int) bool {
		return i > j
	})

	gaps := eachConsMatch(r.vals, func(i int, j int) bool {
		gap := i - j
		return (gap >= 1 && gap <= 3) || (gap <= -1 && gap >= -3)
	})

	return (increasing || decreasing) && gaps
}

func (r *Report) IsSafeWithDampener() bool {
	if r.IsSafe() {
		return true
	}
	for i := 0; i < len(r.vals); i++ {
		tempVals := make([]int, len(r.vals))
		copy(tempVals, r.vals)
		tempReport := Report{append(tempVals[:i], tempVals[i+1:]...)}
		if tempReport.IsSafe() {
			return true
		}
	}
	return false
}

func Part1(reports []Report) int {
	count := 0
	for _, report := range reports {
		if report.IsSafe() {
			count += 1
		}
	}
	return count
}

func Part2(reports []Report) int {
	count := 0
	for _, report := range reports {
		if report.IsSafeWithDampener() {
			count += 1
		}
	}
	return count
}

type consMatcher func(i int, j int) bool

func eachConsMatch(l []int, fn consMatcher) bool {
	for i := 0; i < len(l)-1; i++ {
		if !fn(l[i], l[i+1]) {
			return false
		}
	}
	return true
}
