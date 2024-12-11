package day3

import (
	"testing"

	"github.com/mikelikesbikes/advent-of-code-2024/util"
)

func TestPart1(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	res := Part1(input)
	if res != 161 {
		t.Errorf("expected 161, got %v", res)
	}

	input = `mul(100,123)`
	res = Part1(input)
	if res != 12300 {
		t.Errorf("expected 12300, got %v", res)
	}
}
func TestPart1Input(t *testing.T) {
	input, err := util.ReadFile(`input.txt`)
	if err != nil {
		t.Errorf("%v", err)
	}

	result := Part1(input)
	if result != 183669043 {
		t.Errorf("expected 183669043, got %v", result)
	}
}

func TestPart2(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	res := Part2(input)
	if res != 48 {
		t.Errorf("expected 48, got %v", res)
	}
}

func TestPart2Input(t *testing.T) {
	input, err := util.ReadFile(`input.txt`)
	if err != nil {
		t.Errorf("%v", err)
	}

	result := Part2(input)
	if result != 59097164 {
		t.Errorf("expected 59097164, got %v", result)
	}
}
