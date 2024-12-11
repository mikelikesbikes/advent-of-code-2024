package day2

import (
	"testing"

	"github.com/mikelikesbikes/advent-of-code-2024/util"
)

func TestReportSafetyRules(t *testing.T) {
	reports := parseInput(
		`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)

	if reports[0].IsSafe() != true {
		t.Errorf("report 1 should be safe")
	}
	if reports[1].IsSafe() != false {
		t.Errorf("report 2 should be unsafe")
	}
	if reports[2].IsSafe() != false {
		t.Errorf("report 3 should beunsafe")
	}
	if reports[3].IsSafe() != false {
		t.Errorf("report 4 should be unsafe")
	}
	if reports[4].IsSafe() != false {
		t.Errorf("report 5 should be unsafe")
	}
	if reports[5].IsSafe() != true {
		t.Errorf("report 6 should be safe")
	}
}

func TestPart1Sample(t *testing.T) {
	reports := parseInput(
		`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)
	result := Part1(reports)
	if result != 2 {
		t.Errorf("expected 2 safe reports, got %v", result)
	}
}

func TestPart1Input(t *testing.T) {
	input, err := util.ReadFile(`day_2_input.txt`)
	if err != nil {
		t.Errorf("%v", err)
	}
	reports := parseInput(input)

	result := Part1(reports)
	if result != 472 {
		t.Errorf("expected 472 safe reports, got %v", result)
	}
}

func TestReportSafetyRulesWithDampener(t *testing.T) {
	reports := parseInput(
		`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)

	if reports[0].IsSafeWithDampener() != true {
		t.Errorf("report 1 should be safe")
	}
	if reports[1].IsSafeWithDampener() != false {
		t.Errorf("report 2 should be unsafe")
	}
	if reports[2].IsSafeWithDampener() != false {
		t.Errorf("report 3 should beunsafe")
	}
	if reports[3].IsSafeWithDampener() != true {
		t.Errorf("report 4 should be safe")
	}
	if reports[4].IsSafeWithDampener() != true {
		t.Errorf("report 5 should be safe")
	}
	if reports[5].IsSafeWithDampener() != true {
		t.Errorf("report 6 should be safe")
	}
}

func TestPart2Sample(t *testing.T) {
	reports := parseInput(
		`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)
	result := Part2(reports)
	if result != 4 {
		t.Errorf("expected 4 safe reports, got %v", result)
	}
}

func TestPart2Input(t *testing.T) {
	input, err := util.ReadFile(`day_2_input.txt`)
	if err != nil {
		t.Errorf("%v", err)
	}
	reports := parseInput(input)

	result := Part2(reports)
	if result != 520 {
		t.Errorf("expected 520 safe reports, got %v", result)
	}
}
