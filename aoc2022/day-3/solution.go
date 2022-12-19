package day3

import (
	"strings"

	"colin-valentini.com/advent-of-code/aoc2022/challenge"
)

type Solver struct {
	input string
	part  challenge.Part
}

func NewSolver(input string, part challenge.Part) *Solver {
	return &Solver{input: input, part: part}
}

func (s *Solver) Solve() int {
	if s.part == challenge.Part1 {
		return s.solvePart1()
	}
	return s.solvePart2()
}

func (s *Solver) solvePart1() int {
	total := 0
	for _, ruck := range s.parseInput() {
		total += priority(ruck.misplacedItem())
	}
	return total
}

func (s *Solver) solvePart2() int {
	rucks := s.parseInput()
	groups := make([]*group, 0, len(rucks)/3)
	for i := 0; i < len(rucks)-2; i += 3 {
		groups = append(groups, newGroup([3]*rucksack{
			rucks[i],
			rucks[i+1],
			rucks[i+2],
		}))
	}
	total := 0
	for _, group := range groups {
		total += priority(group.badgeItem())
	}
	return total
}

func (s *Solver) parseInput() []*rucksack {
	ruckStrs := strings.Split(s.input, "\n")
	rucks := make([]*rucksack, len(ruckStrs))
	for i, rs := range ruckStrs {
		rucks[i] = newRucksack(rs)
	}
	return rucks
}

type group struct {
	rucks [3]*rucksack
}

func newGroup(rucks [3]*rucksack) *group {
	return &group{rucks: rucks}
}

func (g *group) badgeItem() rune {
	common := g.rucks[0].toRuneSet()
	for i := 1; i < 3; i++ {
		common = common.intersect(g.rucks[i].toRuneSet())
	}
	// Challenge constraints dictate this is a single valued set.
	for char := range common {
		return char
	}
	return rune(0)
}

type rucksack struct {
	items string
}

func newRucksack(items string) *rucksack {
	return &rucksack{items: items}
}

func (r *rucksack) misplacedItem() rune {
	mid := len(r.items) / 2
	first, second := make(runeSet), make(runeSet)
	for i, char := range r.items {
		if i < mid {
			first[char] = nothing{}
			continue
		}
		if _, ok := first[char]; ok {
			return char
		} else {
			second[char] = nothing{}
		}
	}
	return rune(0)
}

func (r *rucksack) toRuneSet() runeSet {
	set := make(runeSet)
	for _, char := range r.items {
		set[char] = nothing{}
	}
	return set
}

type runeSet map[rune]nothing

func (r runeSet) intersect(other runeSet) runeSet {
	intersection := make(runeSet)
	a, b := r, other
	if len(other) < len(a) {
		a, b = other, r
	}
	for char := range a {
		if _, ok := b[char]; ok {
			intersection[char] = nothing{}
		}
	}
	return intersection
}

type nothing struct{}

func priority(item rune) int {
	if item >= 'a' {
		return int(item-'a') + 1
	}
	return int(item-'A') + 27
}
