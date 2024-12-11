package day3

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) int {
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := r.FindAllString(input, -1)
	res := 0
	for _, s := range matches {
		parts := strings.Split(s[4:len(s)-1], `,`)
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		res += x * y
	}
	return res
}

func Part2(input string) int {
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\)`)
	matches := r.FindAllString(input, -1)
	enabled := true
	res := 0
	for _, s := range matches {
		if s == "do()" {
			enabled = true
		} else if s == "don't()" {
			enabled = false
		} else {
			parts := strings.Split(s[4:len(s)-1], `,`)
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			if enabled {
				res += x * y
			}
		}
	}

	return res
}
