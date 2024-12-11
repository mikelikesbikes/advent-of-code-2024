package day4

import (
	"testing"

	"github.com/mikelikesbikes/advent-of-code-2024/util"
)

func TestPart1Sample(t *testing.T) {
	p := NewPuzzle(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)
	res := Part1(p)
	if res != 18 {
		t.Errorf("expected 18, got %v", res)
	}
}

func TestPart1Input(t *testing.T) {
	input, _ := util.ReadFile(`input.txt`)
	p := NewPuzzle(input)
	res := Part1(p)
	if res != 2662 {
		t.Errorf("expected 2662, got %v", res)
	}
}

func TestPart2Sample(t *testing.T) {
	p := NewPuzzle(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)
	res := Part2(p)
	if res != 9 {
		t.Errorf("expected 9, got %v", res)
	}
}

func TestPart2Input(t *testing.T) {
	input, _ := util.ReadFile(`input.txt`)
	p := NewPuzzle(input)
	res := Part2(p)
	if res != 2034 {
		t.Errorf("expected 2034, got %v", res)
	}
}
