package day11

import (
	"math"
	"strconv"
	"strings"
)

func Blink(stone int) []int {
	var res []int
	if stone == 0 {
		res = append(res, 1)
	} else if len(strconv.Itoa(stone))%2 == 0 {
		length := int(math.Log10(float64(stone)) + 1)
		half := int(math.Pow10(length / 2))
		i1 := stone / half
		i2 := stone % half
		res = append(res, i1, i2)
	} else {
		res = append(res, stone*2024)
	}
	return res
}

func BlinkAll(stones []int) []int {
	var res []int
	for _, stone := range stones {
		res = append(res, Blink(stone)...)
	}
	return res
}

func parseInput(input string) []int {
	strs := strings.Split(input, " ")
	var res []int
	for _, str := range strs {
		i, _ := strconv.Atoi(str)
		res = append(res, i)
	}
	return res
}

type Args struct {
	stone, n int
}

var cache = make(map[Args]int)

func BlinkN(stone int, n int) int {
	args := Args{stone, n}
	if v, ok := cache[args]; ok {
		return v
	}
	if n == 0 {
		cache[args] = 1
	} else {
		sum := 0
		for _, newStone := range Blink(stone) {
			sum += BlinkN(newStone, n-1)
		}
		cache[args] = sum
	}
	return cache[args]
}

func BlinkAllN(stones []int, n int) int {
	sum := 0
	for _, stone := range stones {
		sum += BlinkN(stone, n)
	}
	return sum
}
