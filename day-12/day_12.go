package day12

import (
	"maps"
	"slices"
	"strings"
)

type Pos struct {
	x, y int
}

func (p Pos) Neighbors() []Pos {
	return []Pos{
		{p.x, p.y - 1},
		{p.x + 1, p.y},
		{p.x, p.y + 1},
		{p.x - 1, p.y},
	}
}

func (p Pos) CornerNeighbors() [][]Pos {
	return [][]Pos{
		{{p.x, p.y - 1}, {p.x + 1, p.y - 1}, {p.x + 1, p.y}},
		{{p.x, p.y + 1}, {p.x + 1, p.y + 1}, {p.x + 1, p.y}},
		{{p.x, p.y + 1}, {p.x - 1, p.y + 1}, {p.x - 1, p.y}},
		{{p.x, p.y - 1}, {p.x - 1, p.y - 1}, {p.x - 1, p.y}},
	}
}

type Garden struct {
	plot map[Pos]rune
}

func (g *Garden) TotalPrice() int {
	sum := 0
	for _, region := range g.regions() {
		price := region.Price()
		sum += price
	}
	return sum
}

func (g *Garden) TotalPriceWithDiscount() int {
	sum := 0
	for _, region := range g.regions() {
		price := region.PriceWithDiscount()
		sum += price
	}
	return sum
}

func (g *Garden) regions() []Region {
	tempPlots := maps.Clone(g.plot)
	return g.findRegions(tempPlots, []Region{})
}

func (g *Garden) findRegions(plots map[Pos]rune, regions []Region) []Region {
	if len(plots) == 0 {
		return regions
	}

	// use the next plot from the available plots
	var startingPlot Pos
	var t rune
	for startingPlot, t = range plots {
		break
	}

	region := Region{make([]Pos, 0), string(t)}
	availablePlots := []Pos{startingPlot}
	for len(availablePlots) > 0 {
		// pop next plot
		n := len(availablePlots) - 1
		plot := availablePlots[n]
		availablePlots = availablePlots[:n]
		if !region.Contains(plot) {
			region.AddPlot(plot)
		}

		// remove plot from plots list
		delete(plots, plot)

		// add neighbors with matching type to available
		for _, cPlot := range plot.Neighbors() {
			if v, ok := plots[cPlot]; ok && string(v) == region.t && !region.Contains(cPlot) {
				availablePlots = append(availablePlots, cPlot)
			}
		}
	}

	regions = append(regions, region)
	return g.findRegions(plots, regions)
}

type Region struct {
	plots []Pos
	t     string
}

func (r *Region) AddPlot(p Pos) {
	r.plots = append(r.plots, p)
}

func (r *Region) Contains(p Pos) bool {
	return slices.Index(r.plots, p) >= 0
}

func (r *Region) Price() int {
	perimeter := 0
	for _, p := range r.plots {
		neighborCount := 4
		for _, n := range p.Neighbors() {
			if r.Contains(n) {
				neighborCount -= 1
			}
		}
		perimeter += neighborCount
	}
	return perimeter * len(r.plots)
}

func (r *Region) PriceWithDiscount() int {
	corners := 0
	for _, p := range r.plots {
		for _, corner := range p.CornerNeighbors() {
			// outside corner
			if !r.Contains(corner[0]) && !r.Contains(corner[2]) {
				corners += 1
			}
			// inside corner
			if r.Contains(corner[0]) && !r.Contains(corner[1]) && r.Contains(corner[2]) {
				corners += 1
			}
		}
	}
	return corners * len(r.plots)
}

func parseInput(input string) *Garden {
	garden := Garden{make(map[Pos]rune)}
	for y, line := range strings.Split(input, "\n") {
		for x, r := range line {
			garden.plot[Pos{x, y}] = r
		}
	}
	return &garden
}
