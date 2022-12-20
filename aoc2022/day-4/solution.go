package day4

import (
	"fmt"
	"strconv"
	"strings"

	"colin-valentini.com/advent-of-code/aoc2022/challenge"
)

// Solver solves the Day 4 challenge of Advent of Code 2022.
type Solver struct {
	input string
	part  challenge.Part
}

// NewSolver returns a new instance of Solver.
func NewSolver(input string, part challenge.Part) *Solver {
	return &Solver{input: input, part: part}
}

func (s *Solver) Solve() (int, error) {
	assignments, err := s.parseInput()
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}
	count := 0
	for _, assign := range assignments {
		if (s.part == challenge.Part1 && assign.hasCover()) ||
			(s.part == challenge.Part2 && assign.hasOverlap()) {
			count++
		}
	}
	return count, nil
}

func (s *Solver) parseInput() ([]*pair, error) {
	pairStrs := strings.Split(s.input, "\n")
	pairs := make([]*pair, len(pairStrs))
	for i, pstr := range pairStrs {
		assignStrs := strings.Split(pstr, ",")
		if len(assignStrs) < 2 {
			return nil, fmt.Errorf("invalid input %q on line %d", pstr, i)
		}
		first, err := s.parseAssignment(assignStrs[0])
		if err != nil {
			return nil, fmt.Errorf("invalid input %q on line %d: %w", pstr, i, err)
		}
		second, err := s.parseAssignment(assignStrs[1])
		if err != nil {
			return nil, fmt.Errorf("invalid input %q on line %d: %w", pstr, i, err)
		}
		pairs[i] = newPair(first, second)
	}
	return pairs, nil
}

func (s *Solver) parseAssignment(str string) (*assignment, error) {
	rng := strings.Split(str, "-")
	if len(rng) < 2 {
		return nil, fmt.Errorf("invalid pair %s", str)
	}
	start, err := strconv.Atoi(rng[0])
	if err != nil {
		return nil, fmt.Errorf("parsing assignment start: %w", err)
	}
	end, err := strconv.Atoi(rng[1])
	if err != nil {
		return nil, fmt.Errorf("parsing assigment end: %w", err)
	}
	return newAssignment(start, end), nil
}

type pair struct {
	first, second *assignment
}

func newPair(first, second *assignment) *pair {
	return &pair{first: first, second: second}
}

func (p *pair) hasCover() bool {
	return p.first.covers(p.second) || p.second.covers(p.first)
}

func (p *pair) hasOverlap() bool {
	return p.first.overlaps(p.second) || p.second.overlaps(p.first)
}

type assignment struct {
	start, end int
}

func newAssignment(start, end int) *assignment {
	return &assignment{start: start, end: end}
}

func (a *assignment) covers(other *assignment) bool {
	return a.start <= other.start && other.end <= a.end
}

func (a *assignment) overlaps(other *assignment) bool {
	return a.start <= other.start && a.end >= other.start
}
