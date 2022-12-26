package day5

import (
	"fmt"
	"strings"

	"colin-valentini.com/advent-of-code/aoc2022/challenge"
)

// Solver solves the Day 4 challenge of Advent of Code 2022.
type Solver struct {
	input string
	part  challenge.Part
	crane *crane
	moves []*move
}

// NewSolver returns a new instance of Solver.
func NewSolver(input string, part challenge.Part) *Solver {
	return &Solver{input: input, part: part}
}

func (s *Solver) Solve() (string, error) {
	err := s.parseInput()
	if err != nil {
		return "", fmt.Errorf("parsing input: %w", err)
	}
	for i, m := range s.moves {
		if err := s.crane.apply(m, s.part); err != nil {
			return "", fmt.Errorf("applying move %d", i)
		}
	}
	var builder strings.Builder
	for i := 0; i < s.crane.numStacks(); i++ {
		builder.WriteRune(rune(s.crane.top(i)))
	}
	return builder.String(), nil
}

func (s *Solver) parseInput() error {
	crane, moves, err := newParser().parse(s.input)
	if err != nil {
		return fmt.Errorf("parsing input: %w", err)
	}
	s.crane = crane
	s.moves = moves
	return nil
}
