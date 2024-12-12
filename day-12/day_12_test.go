package day12

import (
	"testing"

	"github.com/mikelikesbikes/advent-of-code-2024/testutil"
	"github.com/mikelikesbikes/advent-of-code-2024/util"
)

func TestPart1SmallSample(t *testing.T) {
	input := `AAAA
BBCD
BBCC
EEEC`
	garden := parseInput(input)
	testutil.AssertEqual(t, 140, garden.TotalPrice())
}

func TestPart1LargeSample(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	garden := parseInput(input)
	testutil.AssertEqual(t, 1930, garden.TotalPrice())
}

func TestPart1Input(t *testing.T) {
	input, _ := util.ReadFile("input.txt")
	garden := parseInput(input)
	testutil.AssertEqual(t, 1387004, garden.TotalPrice())
}

func TestPart2SmallSample(t *testing.T) {
	input := `AAAA
BBCD
BBCC
EEEC`
	garden := parseInput(input)
	testutil.AssertEqual(t, 80, garden.TotalPriceWithDiscount())
}

func TestPart2LargeSample(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	garden := parseInput(input)
	testutil.AssertEqual(t, 1206, garden.TotalPriceWithDiscount())
}

func TestPart2EShape(t *testing.T) {
	input := `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`
	garden := parseInput(input)
	testutil.AssertEqual(t, 236, garden.TotalPriceWithDiscount())
}

func TestPart2InteriorCorners(t *testing.T) {
	input := `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`
	garden := parseInput(input)
	testutil.AssertEqual(t, 368, garden.TotalPriceWithDiscount())
}

func TestPart2Input(t *testing.T) {
	input, _ := util.ReadFile("input.txt")
	garden := parseInput(input)
	testutil.AssertEqual(t, 844198, garden.TotalPriceWithDiscount())
}
