package day11

import (
	"testing"

	"github.com/mikelikesbikes/advent-of-code-2024/testutil"
	"github.com/mikelikesbikes/advent-of-code-2024/util"
)

func TestPart1SampleInput(t *testing.T) {
	input := []int{125, 17}
	gen1 := BlinkAll(input)
	testutil.AssertEqual(t, []int{253000, 1, 7}, gen1)
	gen2 := BlinkAll(gen1)
	testutil.AssertEqual(t, []int{253, 0, 2024, 14168}, gen2)
	gen3 := BlinkAll(gen2)
	testutil.AssertEqual(t, []int{512072, 1, 20, 24, 28676032}, gen3)
	gen4 := BlinkAll(gen3)
	testutil.AssertEqual(t, []int{512, 72, 2024, 2, 0, 2, 4, 2867, 6032}, gen4)
	gen5 := BlinkAll(gen4)
	testutil.AssertEqual(t, []int{1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32}, gen5)
	gen6 := BlinkAll(gen5)
	testutil.AssertEqual(t, []int{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2}, gen6)

	stones := BlinkAllN(input, 25)
	testutil.AssertEqual(t, 55312, stones)
}

func TestPart1Input(t *testing.T) {
	input, _ := util.ReadFile("input.txt")
	stones := BlinkAllN(parseInput(input), 25)
	testutil.AssertEqual(t, 218079, stones)
}

func TestPart2Input(t *testing.T) {
	input, _ := util.ReadFile("input.txt")
	stones := BlinkAllN(parseInput(input), 75)
	testutil.AssertEqual(t, 259755538429618, stones)
}
