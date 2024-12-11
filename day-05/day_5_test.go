package day5

import (
	"testing"

	tu "github.com/mikelikesbikes/advent-of-code-2024/testutil"
	"github.com/mikelikesbikes/advent-of-code-2024/util"
)

var sample string = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestPart1Sample(t *testing.T) {
	manual, updates := parseInput(sample)

	tu.Assert(t, manual.IsSorted(updates[0]))
	tu.Assert(t, manual.IsSorted(updates[1]))
	tu.Assert(t, manual.IsSorted(updates[2]))
	tu.Assert(t, !manual.IsSorted(updates[3]))
	tu.Assert(t, !manual.IsSorted(updates[4]))
	tu.Assert(t, !manual.IsSorted(updates[5]))

	tu.AssertEqual(t, 143, Part1(manual, updates))
}

func TestPart1Input(t *testing.T) {
	input, _ := util.ReadFile("input.txt")
	manual, updates := parseInput(input)
	tu.AssertEqual(t, 5208, Part1(manual, updates))
}

func TestPart2Sample(t *testing.T) {
	manual, updates := parseInput(sample)

	tu.AssertEqual(t, []int{97, 75, 47, 61, 53}, manual.Sort(updates[3]))
	tu.AssertEqual(t, []int{61, 29, 13}, manual.Sort(updates[4]))
	tu.AssertEqual(t, []int{97, 75, 47, 29, 13}, manual.Sort(updates[5]))

	tu.AssertEqual(t, 123, Part2(manual, updates))
}

func TestPart2Input(t *testing.T) {
	input, _ := util.ReadFile("input.txt")
	manual, updates := parseInput(input)
	tu.AssertEqual(t, 6732, Part2(manual, updates))
}
