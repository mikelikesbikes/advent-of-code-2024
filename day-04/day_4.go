package day4

import "strings"

type Puzzle struct {
	board []string
}

func NewPuzzle(input string) Puzzle {
	return Puzzle{strings.Split(input, "\n")}
}

func (p *Puzzle) countAllXMAS() int {
	count := 0
	for y, s := range p.board {
		for x := range s {
			count += p.countXMASAt(x, y)
		}
	}
	return count
}

func (p *Puzzle) compAt(x int, y int, s rune) bool {
	if x < 0 || y < 0 || y >= len(p.board) || x >= len(p.board[0]) {
		return false
	}
	return p.board[x][y] == byte(s)
}

func (p *Puzzle) countXMASAt(x int, y int) int {
	count := 0
	if !p.compAt(x, y, 'X') {
		return count
	}

	// diagonal up left
	if p.compAt(x-1, y-1, 'M') && p.compAt(x-2, y-2, 'A') && p.compAt(x-3, y-3, 'S') {
		count += 1
	}
	// up
	if p.compAt(x, y-1, 'M') && p.compAt(x, y-2, 'A') && p.compAt(x, y-3, 'S') {
		count += 1
	}
	// diagonal up right
	if p.compAt(x+1, y-1, 'M') && p.compAt(x+2, y-2, 'A') && p.compAt(x+3, y-3, 'S') {
		count += 1
	}
	// right
	if p.compAt(x+1, y, 'M') && p.compAt(x+2, y, 'A') && p.compAt(x+3, y, 'S') {
		count += 1
	}
	// diagonal down right
	if p.compAt(x+1, y+1, 'M') && p.compAt(x+2, y+2, 'A') && p.compAt(x+3, y+3, 'S') {
		count += 1
	}
	// down
	if p.compAt(x, y+1, 'M') && p.compAt(x, y+2, 'A') && p.compAt(x, y+3, 'S') {
		count += 1
	}
	// diagonal down left
	if p.compAt(x-1, y+1, 'M') && p.compAt(x-2, y+2, 'A') && p.compAt(x-3, y+3, 'S') {
		count += 1
	}
	// left
	if p.compAt(x-1, y, 'M') && p.compAt(x-2, y, 'A') && p.compAt(x-3, y, 'S') {
		count += 1
	}

	return count
}

func (p *Puzzle) countAllX_MAS() int {
	count := 0
	for y, s := range p.board {
		for x := range s {
			count += p.countX_MASAt(x, y)
		}
	}
	return count
}
func (p *Puzzle) countX_MASAt(x int, y int) int {
	count := 0
	if !p.compAt(x, y, 'A') {
		return count
	}

	// top
	if p.compAt(x-1, y-1, 'M') && p.compAt(x+1, y-1, 'M') &&
		p.compAt(x-1, y+1, 'S') && p.compAt(x+1, y+1, 'S') {
		count += 1
	}
	// right
	if p.compAt(x-1, y-1, 'S') && p.compAt(x+1, y-1, 'M') &&
		p.compAt(x-1, y+1, 'S') && p.compAt(x+1, y+1, 'M') {
		count += 1
	}
	// bottom
	if p.compAt(x-1, y-1, 'S') && p.compAt(x+1, y-1, 'S') &&
		p.compAt(x-1, y+1, 'M') && p.compAt(x+1, y+1, 'M') {
		count += 1
	}
	// left
	if p.compAt(x-1, y-1, 'M') && p.compAt(x+1, y-1, 'S') &&
		p.compAt(x-1, y+1, 'M') && p.compAt(x+1, y+1, 'S') {
		count += 1
	}

	return count
}

func Part1(p Puzzle) int {
	return p.countAllXMAS()
}

func Part2(p Puzzle) int {
	return p.countAllX_MAS()
}
