package day1

import (
	"testing"

	"github.com/mikelikesbikes/advent-of-code-2024/util"
)

func TestPart1Example(t *testing.T) {
	l1, l2, err := parseInput(
		`3   4
4   3
2   5
1   3
3   9
3   3`)
	if err != nil {
		t.Errorf("%v", err)
	}
	result := Part1(l1, l2)
	if result != 11 {
		t.Errorf(`Expected distance of 11, got %v`, result)
	}
}

func TestPart1ActulInput(t *testing.T) {
	input, err := util.ReadFile(`day_1_input.txt`)
	if err != nil {
		t.Errorf("%v", err)
	}

	l1, l2, err := parseInput(input)
	if err != nil {
		t.Errorf("%v", err)
	}

	result := Part1(l1, l2)
	if result != 1882714 {
		t.Errorf(`Expected distance of 11, got %v`, result)
	}
}

func TestPart2SampleInput(t *testing.T) {
	l1, l2, err := parseInput(
		`3   4
4   3
2   5
1   3
3   9
3   3`)
	if err != nil {
		t.Errorf("%v", err)
	}
	result := Part2(l1, l2)
	if result != 31 {
		t.Errorf(`Expected distance of 31, got %v`, result)
	}
}

func TestPart2ActulInput(t *testing.T) {
	input, err := util.ReadFile(`day_1_input.txt`)
	if err != nil {
		t.Errorf("%v", err)
	}

	l1, l2, err := parseInput(input)
	if err != nil {
		t.Errorf("%v", err)
	}

	result := Part2(l1, l2)
	if result != 19437052 {
		t.Errorf(`Expected distance of 11, got %v`, result)
	}
}
